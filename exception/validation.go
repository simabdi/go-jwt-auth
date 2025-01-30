package exception

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
)

var (
	Validate = validator.New()
)

func Validation(data interface{}) (dataError string) {
	var errors []string
	validate := validator.New()
	_ = validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())
		return err == nil
	})

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" : "+errorValidation(err.Tag(), err.Type().String(), err.Param()))
		}
	}

	dataError = strings.Join(errors, ", ")
	return dataError
}

func errorValidation(tag interface{}, typeName string, param string) string {
	switch tag {
	case "min":
		return "This field must is min " + param + " character"
	case "max":
		return "This field must is max " + param + " character"
	case "number":
		return "This field must is number"
	case "string":
		return "This field must is string"
	case "required":
		return "This field is required " + typeName
	case "email":
		return "Invalid email"
	case "oneof":
		return "Failed on the 'oneof'"
	case "eqfield":
		return "Password confirmation doesn't match"
	case "date":
		return "Format date must be Y-m-d"
	}

	return ""
}
