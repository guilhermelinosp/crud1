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
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLocale := en.New()
		uni := ut.New(enLocale, enLocale)
		transl, _ = uni.GetTranslator("en")
		err := entranslation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidatorRequest(validationErr error) *errs.Error {
	var jsonErr *json.UnmarshalTypeError
	var validationErrors validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return errs.NewBadRequest("Invalid field type")
	} else if errors.As(validationErr, &validationErrors) {
		causes := make([]errs.Causes, 0, len(validationErrors))

		for _, err := range validationErrors {
			causes = append(causes, errs.Causes{
				Field:   err.Field(),
				Message: err.Translate(transl),
			})
		}

		return errs.NewValidationError("Invalid field value", causes)
	} else {
		return errs.NewBadRequest("Invalid request")
	}
}
