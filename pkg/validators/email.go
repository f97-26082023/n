package validators

import (
	"github.com/go-playground/validator/v10"

	"github.com/f97/n/pkg/utils"
)

// ValidEmail returns whether the given email is valid
func ValidEmail(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		if utils.IsValidEmail(value) {
			return true
		}
	}

	return false
}
