package validate

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
_:
	_ = v.RegisterValidation("password", password)
	return &CustomValidator{
		Validator: v,
	}
}

func (v *CustomValidator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

func password(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return isValidPassword(password)
}

func isValidPassword(password string) bool {
	// 半角英字チェック
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	// 数字チェック
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	// 記号チェック
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	return hasLetter && hasDigit && hasSpecial
}
