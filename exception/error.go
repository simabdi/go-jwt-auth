package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"strings"
)

func Error(err error) (dataError string) {
	switch e := err.(type) {
	case *mysql.MySQLError:
		dataError = "SQL " + e.Error()
	default:
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			var errorArray []string
			for _, e := range err.(validator.ValidationErrors) {
				errorArray = append(errorArray, e.Error())
			}
			dataError = strings.Join(errorArray, ", ")
		} else {
			dataError = err.Error()
		}

	}

	return
}
