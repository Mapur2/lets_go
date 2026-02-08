package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

type Response struct {
	Success bool   `json:"success"` //known as struct tags
	Error   string `json:"error"`
}

func GeneralError(err error) Response {
	return Response{
		Success: false,
		Error:   err.Error(),
	}
}

func ValidatorError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return Response{
		Success: false,
		Error:   strings.Join(errMsgs, ", "),
	}
}
