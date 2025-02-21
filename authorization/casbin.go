package authorization

import (
	"encoding/csv"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/pkg/errors"
	"log/slog"
	"strings"
)

// CasbinSlogLogger is a wrapper to use slog.Logger with Casbin.
type CasbinSlogLogger struct {
	logger  *slog.Logger
	enabled bool
}

func (l *CasbinSlogLogger) EnableLog(enable bool) {
	l.enabled = enable
}

func (l *CasbinSlogLogger) IsEnabled() bool {
	return l.enabled
}

func (l *CasbinSlogLogger) LogModel(model [][]string) {
	if !l.enabled {
		return
	}

	l.logger.Info("Model", slog.Any("model", model))
}

func (l *CasbinSlogLogger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {
	if !l.enabled {
		return
	}

	l.logger.Info("Enforce",
		slog.String("matcher", matcher),
		slog.Any("request", request),
		slog.Bool("result", result),
		slog.Any("explains", explains),
	)
}

func (l *CasbinSlogLogger) LogRole(roles []string) {
	if !l.enabled {
		return
	}

	l.logger.Info("Roles", slog.Any("roles", roles))
}

func (l *CasbinSlogLogger) LogPolicy(policy map[string][][]string) {
	if !l.enabled {
		return
	}

	l.logger.Info("Policy", slog.Any("policy", policy))
}

func (l *CasbinSlogLogger) LogError(err error, msg ...string) {
	if !l.enabled {
		return
	}

	l.logger.Error("Error", slog.Any("error", err), slog.Any("message", msg))
}

var _ log.Logger = (*CasbinSlogLogger)(nil)

// NewCasbinSlogLogger creates a new adapter for Casbin with slog.Logger.
func NewCasbinSlogLogger(logger *slog.Logger) *CasbinSlogLogger {
	return &CasbinSlogLogger{logger: logger}
}

func addPolicyFromString(enforcer *casbin.Enforcer, policyFileContent string) error {
	reader := csv.NewReader(strings.NewReader(policyFileContent))

	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return errors.Wrap(err, "failed to read policy content")
	}

	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		args := make([]any, len(record))
		for i := range args {
			args[i] = record[i]
		}

		switch args[0] {
		case "p":
			exists, err := enforcer.HasPolicy(args[1:]...)
			if err != nil {
				return errors.Wrap(err, "failed to check policy")
			}

			if exists {
				continue
			}

			_, err = enforcer.AddPolicy(args[1:]...)
			if err != nil {
				return errors.Wrap(err, "failed to add policy")
			}

		case "g":
			exists, err := enforcer.HasGroupingPolicy(args[1:]...)
			if err != nil {
				return errors.Wrap(err, "failed to check policy")
			}

			if exists {
				continue
			}

			_, err = enforcer.AddGroupingPolicy(args[1:]...)
			if err != nil {
				return errors.Wrap(err, "failed to add grouping policy")
			}
		default:
			return errors.Errorf("unknown policy type: %s", args[0])
		}
	}

	return nil
}
