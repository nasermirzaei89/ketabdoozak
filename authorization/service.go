package authorization

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

//go:embed model.conf
var casbinModelContent string

//go:embed policy.csv
var casbinPolicyContent string

const ObjectNone = "-"

type Service struct {
	enforcer *casbin.Enforcer
	tracer   trace.Tracer
	logger   *slog.Logger
}

func NewService(sqlDB *sql.DB) (*Service, error) {
	casbinModel, err := model.NewModelFromString(casbinModelContent)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load casbin model")
	}

	adapter, err := sqladapter.NewAdapter(sqlDB, "postgres", "casbin_rules")
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to casbin database")
	}

	enforcer, err := casbin.NewEnforcer(casbinModel, adapter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create casbin enforcer")
	}

	enforcer.EnableAutoSave(true)

	casbinLogger := NewCasbinSlogLogger(logger)

	casbinLogger.EnableLog(true)

	enforcer.SetLogger(casbinLogger)

	err = addPolicyFromString(enforcer, casbinPolicyContent)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load csv policy")
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load db policy")
	}

	return &Service{
		enforcer: enforcer,
		tracer:   tracer,
		logger:   logger,
	}, nil
}

type CheckAccessRequest struct {
	Subject string
	Action  string
	Object  string
}

type AccessDeniedError struct {
	Subject string
	Action  string
	Object  string
}

func (err AccessDeniedError) Error() string {
	if err.Object != "" {
		return fmt.Sprintf("access denied for subject '%s' and object '%s' and action '%s'", err.Subject, err.Object, err.Action)
	}

	return fmt.Sprintf("access denied for subject '%s' and action '%s'", err.Subject, err.Action)
}

func (svc *Service) CheckAccess(ctx context.Context, req CheckAccessRequest) error {
	_, span := svc.tracer.Start(ctx, "CheckAccess")
	defer span.End()

	if req.Object == "" {
		req.Object = ObjectNone
	}

	allowed, err := svc.enforcer.Enforce(req.Subject, req.Object, req.Action)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrap(err, "failed to check permission")
	}

	if !allowed {
		return AccessDeniedError(req)
	}

	return nil
}

type AddPolicyRequest struct {
	Subject string
	Action  string
	Object  string
}

func (svc *Service) AddPolicy(ctx context.Context, reqs ...AddPolicyRequest) error {
	_, span := svc.tracer.Start(ctx, "AddPolicy")
	defer span.End()

	rules := make([][]string, 0, len(reqs))

	for _, req := range reqs {
		if req.Object == "" {
			req.Object = ObjectNone
		}

		rules = append(rules, []string{req.Subject, req.Object, req.Action})
	}

	_, err := svc.enforcer.AddPolicies(rules)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrap(err, "failed to add policies")
	}

	return nil
}

func (svc *Service) AddToGroup(ctx context.Context, sub string, groups ...string) error {
	_, span := svc.tracer.Start(ctx, "AddToGroup")
	defer span.End()

	rules := make([][]string, 0, len(groups))

	for _, group := range groups {
		rules = append(rules, []string{sub, group})
	}

	_, err := svc.enforcer.AddGroupingPolicies(rules)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrap(err, "failed to add grouping policies")
	}

	return nil
}
