package main

import (
	"context"
	goerrors "errors"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nasermirzaei89/env"
	"github.com/nasermirzaei89/ketabdoozak/authentication"
	"github.com/nasermirzaei89/ketabdoozak/authorization"
	"github.com/nasermirzaei89/ketabdoozak/db/postgres"
	_ "github.com/nasermirzaei89/ketabdoozak/docs"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/nasermirzaei89/ketabdoozak/validation"
	"github.com/nasermirzaei89/ketabdoozak/www"
	"github.com/nasermirzaei89/problem"
	problemoutput "github.com/nasermirzaei89/problem/output"
	"github.com/nasermirzaei89/services/api"
	"github.com/nasermirzaei89/services/swagger"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/exporters/autoexport"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
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

	// Minio
	minioEndpoint := env.GetString("MINIO_ENDPOINT", "localhost:9000")
	minioUsername := env.MustGetString("MINIO_USERNAME")
	minioPassword := env.MustGetString("MINIO_PASSWORD")
	minioUseSSL := env.GetBool("MINIO_USE_SSL", false)

	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUsername, minioPassword, ""),
		Secure: minioUseSSL,
	})
	if err != nil {
		return errors.Wrap(err, "error creating minio client")
	}

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

	// File Manager Service
	fileManagerMinioBucketName := env.GetString("FILE_MANAGER_MINIO_BUCKET_NAME", "files")

	exists, err := minioClient.BucketExists(ctx, fileManagerMinioBucketName)
	if err != nil {
		return errors.Wrap(err, "error checking if minio bucket exists")
	}

	if !exists {
		if env.GetBool("FILE_MANAGER_MINIO_BUCKET_CREATE_IF_NOT_EXISTS", false) {
			err = minioClient.MakeBucket(ctx, fileManagerMinioBucketName, minio.MakeBucketOptions{})
			if err != nil {
				return errors.Wrap(err, "error creating minio bucket")
			}
		} else {
			return errors.Errorf("minio bucket does not exist: %s", fileManagerMinioBucketName)
		}
	}

	fileRepo := postgres.NewFileRepo(sqlDB)

	fileManagerSvc := filemanager.NewService(authzSvc, minioClient, fileManagerMinioBucketName, fileRepo)

	fileManagerHandler := filemanager.NewHandler(fileManagerSvc)

	// Listing Service
	listingLocation := postgres.NewListingLocationRepo(sqlDB)
	listingItemRepo := postgres.NewListingItemRepo(sqlDB)

	listingSvc := listing.NewService(listingLocation, listingItemRepo, validate)

	// Cookie Store
	key := env.GetString("STORE_KEY", "super-secret-key")
	cookieStore := sessions.NewCookieStore([]byte(key))
	cookieStore.Options = &sessions.Options{
		Path:        "/",
		Domain:      "",
		MaxAge:      0,
		Secure:      false,
		HttpOnly:    false,
		Partitioned: false,
		SameSite:    0,
	}

	// WWW Auth
	wwwOIDCIssuerURL := env.MustGetString("WWW_OIDC_ISSUER_URL")
	wwwOIDCClientID := env.MustGetString("WWW_OIDC_CLIENT_ID")
	wwwOIDCClientSecret := env.MustGetString("WWW_OIDC_CLIENT_SECRET")
	wwwOIDCRedirectURL := env.MustGetString("WWW_OIDC_REDIRECT_URL")
	wwwOIDCLogoutRedirectURL := env.MustGetString("WWW_OIDC_LOGOUT_REDIRECT_URL")

	wwwAuth, err := www.NewAuthenticator(ctx, wwwOIDCIssuerURL, wwwOIDCClientID, wwwOIDCClientSecret, wwwOIDCRedirectURL)
	if err != nil {
		return errors.Wrap(err, "error initializing www authenticator")
	}

	// WWW Handler
	wwwBaseURL := "/www/"

	wwwHandler, err := www.NewHandler(
		wwwBaseURL,
		string(env.Environment()),
		cookieStore,
		wwwAuth,
		wwwOIDCLogoutRedirectURL,
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
		api.RegisterEndpoint("/www/", http.StripPrefix("/www", wwwHandler)),
		api.RedirectRootTo("/www/"),
	)
	if err != nil {
		return errors.Wrap(err, "error creating handler")
	}

	addr := env.GetString("HOST", "0.0.0.0") + ":" + env.GetString("PORT", "8080")

	// HTTP Server
	sever := http.Server{
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
	}
	serverErr := make(chan error, 1)

	go func() {
		logger.InfoContext(ctx, "Listening on http://"+addr)

		serverErr <- sever.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-serverErr:
		return err
	case <-ctx.Done():
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = sever.Shutdown(context.Background())
	if err != nil {
		return errors.Wrap(err, "error shutting down server")
	}

	return nil
}
