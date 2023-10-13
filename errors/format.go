package errors

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type GormErr struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func IsDuplicateKeyGormError(err error) bool {
	byteErr, _ := json.Marshal(err)
	var newError GormErr
	json.Unmarshal(byteErr, &newError)
	switch newError.Code {
	case "23505":
		return true
	}
	return false
}

var validationErrors = map[string]string{
	"required": " is required, but was not received",
	"min":      "'s value or length is less than allowed",
	"max":      "'s value or length is bigger than allowed",
}

func FormatErrors(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		e := err.(validator.ValidationErrors)
		var errMsg string
		for _, v := range e {
			errMsg = fmt.Sprintf("%s%s", v.Namespace(), getVldErrorMsg(v.ActualTag()))
		}
		return errMsg
	default:
		return err.Error()
	}
}

func getVldErrorMsg(s string) string {
	if v, ok := validationErrors[s]; ok {
		return v
	}
	return " failed on " + s + " validation"
}
