package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func Mobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()

	pattern := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`

	ok, _ := regexp.Match(pattern, []byte(mobile))

	return ok
}
