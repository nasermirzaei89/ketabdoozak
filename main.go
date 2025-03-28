package main

import (
	"context"
	_ "embed"
	"encoding/gob"
	goerrors "errors"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nasermirzaei89/env"
	"github.com/nasermirzaei89/ketabdoozak/authentication"
	"github.com/nasermirzaei89/ketabdoozak/db/postgres"
	dbredis "github.com/nasermirzaei89/ketabdoozak/db/redis"
	_ "github.com/nasermirzaei89/ketabdoozak/docs"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/validation"
	"github.com/nasermirzaei89/ketabdoozak/www"
	"github.com/nasermirzaei89/problem"
	problemoutput "github.com/nasermirzaei89/problem/output"
	"github.com/nasermirzaei89/services/api"
	"github.com/nasermirzaei89/services/authorization"
	"github.com/nasermirzaei89/services/swagger"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
	"golang.org/x/oauth2"
	golog "log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//	@title									Ketabdoozak
//	@version								1.0
//	@description							This is ketabdoozak api doc
//	@termsOfService							http://swagger.io/terms/
//
//	@securityDefinitions.oauth2.implicit	OAuth2Implicit
//	@authorizationUrl						https://auth.applicaset.com/realms/ketabdoozak/protocol/openid-connect/auth

//go:embed policy.csv
var casbinPolicyContent string

func main() {
	if err := run(); err != nil {
		golog.Fatalln(errors.Wrap(err, "error in running the program"))
	}
}

const (
	HTTPServerTimeOut = 60 * time.Second
)

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var err error

	// OpenTelemetry Logger
	logExporter, err := autoexport.NewLogExporter(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to initialize log exporter")
	}

	defer func() { err = goerrors.Join(err, logExporter.Shutdown(context.Background())) }()

	global.SetLoggerProvider(log.NewLoggerProvider(log.WithProcessor(log.NewSimpleProcessor(logExporter))))

	logger := otelslog.NewLogger("main")

	// OpenTelemetry Meter
	metricReader, err := autoexport.NewMetricReader(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to initialize metric reader")
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metricReader),
	)

	defer func() { err = goerrors.Join(err, meterProvider.Shutdown(context.Background())) }()

	otel.SetMeterProvider(meterProvider)

	// OpenTelemetry Tracer
	spanExporter, err := autoexport.NewSpanExporter(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to initialize span exporter")
	}

	traceProvider := trace.NewTracerProvider(trace.WithBatcher(spanExporter))

	defer func() { err = goerrors.Join(err, traceProvider.Shutdown(context.Background())) }()

	otel.SetTracerProvider(traceProvider)

	// RFC 7807 Problem
	// TODO: set otel logger
	problem.SetLogger(problemoutput.New())

	// Redis
	rdb := redis.NewClient(&redis.Options{ //nolint:exhaustruct
		Addr:     env.GetString("REDIS_URL", "localhost:6379"),
		Password: env.GetString("REDIS_PASSWORD", ""),
		DB:       env.GetInt("REDIS_DB", 0),
	})

	defer func() { err = goerrors.Join(err, rdb.Close()) }()

	// Database
	dbDSN := env.MustGetString("DATABASE_DSN")

	sqlDB, dbCloseFunc, err := postgres.NewDB(ctx, dbDSN)
	if err != nil {
		return errors.Wrap(err, "error opening db")
	}

	defer func() {
		err = goerrors.Join(err, dbCloseFunc())
	}()

	// Validator
	validate, err := validation.NewValidator()
	if err != nil {
		return errors.Wrap(err, "error initializing validator")
	}

	// Authentication Service
	authOIDCIssuerURL := env.MustGetString("AUTHENTICATION_OIDC_ISSUER_URL")
	authOIDCClientID := env.MustGetString("AUTHENTICATION_OIDC_CLIENT_ID")
	authOIDCUsernameClaim := env.GetString("AUTHENTICATION_OIDC_USERNAME_CLAIM", "sub")

	authSvc, err := authentication.NewService(ctx, authOIDCIssuerURL, authOIDCClientID, authOIDCUsernameClaim)
	if err != nil {
		return errors.Wrap(err, "failed to initialize auth service")
	}

	// Authorization Service
	authzSvc, err := authorization.NewService(sqlDB)
	if err != nil {
		return errors.Wrap(err, "failed to initialize authz service")
	}

	err = authzSvc.AddPolicyFromCSV(ctx, casbinPolicyContent)
	if err != nil {
		return errors.Wrap(err, "failed to add policy")
	}

	// File Manager Service
	fileManagerFilesBucket, err := blob.OpenBucket(ctx, env.MustGetString("FILE_MANAGER_FILES_BUCKET_URL"))
	if err != nil {
		return errors.Wrap(err, "failed to initialize file manager bucket")
	}

	isAccessible, err := fileManagerFilesBucket.IsAccessible(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to check if file manager is accessible")
	}

	if !isAccessible {
		return errors.New("file manager is not accessible")
	}

	defer func() {
		err = goerrors.Join(err, fileManagerFilesBucket.Close())
	}()

	fileRepo := postgres.NewFileRepo(sqlDB)

	var fileManagerSvc filemanager.Service
	fileManagerSvc = filemanager.NewService(fileManagerFilesBucket, fileRepo)
	fileManagerSvc = filemanager.NewAuthorizationMiddleware(fileManagerSvc, authzSvc)

	fileManagerHandler := filemanager.NewHandler(fileManagerSvc)

	// Listing Service
	listingLocation := postgres.NewListingLocationRepo(sqlDB)
	listingItemRepo := postgres.NewListingItemRepo(sqlDB)

	var listingSvc listing.Service
	listingSvc = listing.NewService(listingLocation, listingItemRepo, validate)
	listingSvc = listing.NewAuthorizationMiddleware(listingSvc, authzSvc)

	listingHandler := listing.NewHandler(listingSvc)

	// Cookie Store
	key := env.GetString("WWW_COOKIE_STORE_KEY", "super-secret-key")
	wwwCookieStore := sessions.NewCookieStore([]byte(key))
	wwwCookieStore.Options = &sessions.Options{
		Path:        "/www/",
		Domain:      "",
		MaxAge:      0,
		Secure:      true,
		HttpOnly:    true,
		Partitioned: true,
		SameSite:    http.SameSiteLaxMode,
	}

	// For secure cookie
	gob.Register(new(oauth2.Token))

	// WWW Auth
	wwwOIDCIssuerURL := env.MustGetString("WWW_OIDC_ISSUER_URL")
	wwwOIDCClientID := env.MustGetString("WWW_OIDC_CLIENT_ID")
	wwwOIDCClientSecret := env.MustGetString("WWW_OIDC_CLIENT_SECRET")
	wwwOIDCRedirectURL := env.MustGetString("WWW_OIDC_REDIRECT_URL")
	wwwOIDCLogoutRedirectURL := env.MustGetString("WWW_OIDC_LOGOUT_REDIRECT_URL")

	wwwAuth, err := www.NewAuthenticator(ctx, www.NewAuthenticatorRequest{
		OIDCIssuerURL:    wwwOIDCIssuerURL,
		OIDCClientID:     wwwOIDCClientID,
		OIDCClientSecret: wwwOIDCClientSecret,
		OIDCRedirectURL:  wwwOIDCRedirectURL,
	})
	if err != nil {
		return errors.Wrap(err, "error initializing www authenticator")
	}

	// WWW Handler
	wwwBaseURL := "/www/"

	sessionRepo := dbredis.NewSessionRepo(rdb)

	wwwCSRFAuthKey := []byte(env.MustGetString("WWW_CSRF_AUTH_KEY"))

	wwwHandler, err := www.NewHandler(
		wwwBaseURL,
		string(env.Environment()),
		wwwCookieStore,
		sessionRepo,
		wwwAuth,
		wwwOIDCLogoutRedirectURL,
		wwwCSRFAuthKey,
		authzSvc,
		fileManagerSvc,
		listingSvc,
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize www handler")
	}

	// API Handler
	h, err := api.NewHandler(
		api.RegisterMiddleware(authSvc.AuthenticateMiddleware()),
		api.RegisterEndpoint("/swagger/", http.StripPrefix("/swagger", swagger.NewHandler())),
		api.RegisterEndpoint("/filemanager/", http.StripPrefix("/filemanager", fileManagerHandler)),
		api.RegisterEndpoint("/listing/", http.StripPrefix("/listing", listingHandler)),
		api.RegisterEndpoint("/www/", http.StripPrefix("/www", wwwHandler)),
		api.RedirectRootTo("/www/"),
	)
	if err != nil {
		return errors.Wrap(err, "error creating handler")
	}

	addr := env.GetString("HOST", "0.0.0.0") + ":" + env.GetString("PORT", "8080")

	// HTTP Server
	server := http.Server{
		Addr:                         addr,
		Handler:                      h,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  HTTPServerTimeOut,
		ReadHeaderTimeout:            HTTPServerTimeOut,
		WriteTimeout:                 HTTPServerTimeOut,
		IdleTimeout:                  HTTPServerTimeOut,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  func(_ net.Listener) context.Context { return ctx },
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}
	serverErr := make(chan error, 1)

	go func() {
		logger.InfoContext(ctx, "Listening on http://"+addr)

		serverErr <- server.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-serverErr:
		return err
	case <-ctx.Done():
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = server.Shutdown(context.Background())
	if err != nil {
		return errors.Wrap(err, "error shutting down server")
	}

	return nil
}
