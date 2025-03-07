package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Awais914/go-students-api/internal/types"
	"github.com/Awais914/go-students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.HandleError(fmt.Errorf("payload missing")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.HandleError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"message": "Created"})
	}
}
