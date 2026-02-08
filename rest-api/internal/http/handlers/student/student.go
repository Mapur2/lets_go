package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mapur2/lets_go/rest-apis/internal/storage"
	"github.com/mapur2/lets_go/rest-apis/internal/types"
	"github.com/mapur2/lets_go/rest-apis/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello students"))
	}
}

func Create(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadGateway, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//request validation
		if er := validator.New().Struct(student); er != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidatorError(er.(validator.ValidationErrors)))
			//er.(validator.ValidationErrors) is type casting
			return
		}

		exists, _, err := storage.GetStudentByEmail(student.Email)

		if exists == true && err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if exists == true {
			existError := errors.New("Email already exists")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(existError))
			return
		}

		fmt.Println(student)
		id, err := storage.CreateStudent(student)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		student.Id = id

		response.WriteJson(w, http.StatusCreated, map[string]any{
			"success": "OK",
			"student": student,
		})
	}
}

func GetStudentByEmail(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.PathValue("email")

		exits, student, err := storage.GetStudentByEmail(email)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		if exits == true {
			response.WriteJson(w, http.StatusAccepted, map[string]any{
				"success": true,
				"student": student,
			})
			return
		}
		response.WriteJson(w,http.StatusBadRequest, response.GeneralError(errors.New("Student not found")))
	}
}
