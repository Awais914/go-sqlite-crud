package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Awais914/go-students-api/internal/storage"
	"github.com/Awais914/go-students-api/internal/types"
	"github.com/Awais914/go-students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func Create(storage storage.Storage) http.HandlerFunc {
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

		stdId, err := storage.CreateStudent(student.Email, student.Name, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		slog.Info("student created successfully")
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": stdId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.HandleError(err))
			return
		}

		student, err := storage.GetStudentById(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.HandleError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetAll(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitStr := r.URL.Query().Get("limit")
		var limit int
		if limitStr == "" {
			limit = 10
		} else {
			var err error
			limit, err = strconv.Atoi(limitStr)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.HandleError(fmt.Errorf("invalid limit")))
				return
			}
		}

		pageStr := r.URL.Query().Get("page")
		var page int
		if pageStr == "" {
			page = 1
		} else {
			var err error
			page, err = strconv.Atoi(pageStr)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.HandleError(fmt.Errorf("invalid page")))
				return
			}
		}

		students, err := storage.GetAllStudents(limit, page)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.HandleError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

func UpdateById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.HandleError(err))
			return
		}

		var student types.Student
		err = json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.HandleError(err))
			return
		}

		err = storage.UpdateStudentById(id, student.Email, student.Name, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.HandleError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": "student updated successfully"})
	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.HandleError(err))
			return
		}

		err = storage.DeleteStudentById(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.HandleError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]string{"message": "student deleted successfully"})
	}
}
