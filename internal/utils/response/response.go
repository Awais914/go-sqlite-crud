package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error string `json:"error"`
}

const (
	Status = "OK"
	Error = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func HandleError(err error) Response {
	return Response{
		Status: Error,
		Error: err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
		default:
			errMsgs = append(errMsgs, err.Error())
		}
	}

	return Response{
		Status: Error,
		Error: strings.Join(errMsgs, ", "),
	}
}