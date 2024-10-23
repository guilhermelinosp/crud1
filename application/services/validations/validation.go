package validations

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/guilhermelinosp/crud1/config/errs"
	"regexp"
)

var (
	_      = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		translator := en.New()
		unt := ut.New(translator, translator)
		transl, _ = unt.GetTranslator("en")
		_ = entranslation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateRequest(validationErr error) *errs.Error {
	var jsonErr *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return errs.NewBadRequest("Invalid field type")
	} else if errors.As(validationErr, &validationErrors) {
		var errorCauses []errs.Causes

		for _, e := range validationErrors {
			cause := errs.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return errs.NewValidationError("Some fields are invalid", errorCauses)
	} else {
		return errs.NewBadRequest("Error trying to convert fields")
	}
}

func ValidatePassword(password *string) (bool, *errs.Error) {
	uppercasePattern := `.*[A-Z].*`
	numberPattern := `.*\d.*`
	specialCharPattern := `.*[!@#$%^&*()].*`

	hasUppercase, err := regexp.MatchString(uppercasePattern, *password)
	if err != nil || !hasUppercase {
		return false, errs.NewBadRequest("Password must contain at least one uppercase letter")
	}

	hasNumber, err := regexp.MatchString(numberPattern, *password)
	if err != nil || !hasNumber {
		return false, errs.NewBadRequest("Password must contain at least one number")
	}

	hasSpecialChar, err := regexp.MatchString(specialCharPattern, *password)
	if err != nil || !hasSpecialChar {
		return false, errs.NewBadRequest("Password must contain at least one special character")
	}

	return true, nil
}
