package validator

import (
	"strconv"
	"strings"

	"recommand-chat-bot/domain"

	"github.com/go-playground/validator"
)

func RegisterMovieValidation(v *validator.Validate) error {
	if err := v.RegisterValidation("maxlen", validateMaxLen); err != nil {
		return err
	}
	if err := v.RegisterValidation("validateGenre", validateGenre); err != nil {
		return err
	}
	if err := v.RegisterValidation("validateEmptyStr", validateEmptyStr); err != nil {
		return err
	}

	return nil
}

func validateGenre(fl validator.FieldLevel) bool {
	genre, ok := fl.Field().Interface().(domain.Genre)
	if !ok {
		return false
	}
	return genre.IsValid()
}

func validateEmptyStr(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	rmBlankStr := strings.ReplaceAll(field, " ", "")
	return len(rmBlankStr) > 0
}

func validateMaxLen(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	param := fl.Param()
	maxLen, err := strconv.Atoi(param)
	if err != nil {
		return false
	}
	return len(field) <= maxLen
}
