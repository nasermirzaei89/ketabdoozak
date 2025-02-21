package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"regexp"
)

func NewValidator() (*validator.Validate, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.RegisterValidation("slug", IsValidateSlug)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register slug validation")
	}

	return validate, nil
}

var slugRegex = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]+[a-zA-Z0-9]$")

func IsValidateSlug(fl validator.FieldLevel) bool {
	return slugRegex.MatchString(fl.Field().String())
}
