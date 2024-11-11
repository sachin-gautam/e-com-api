package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/sachin-gautam/go-crud-api/internal/dtypes"
	"github.com/sachin-gautam/go-crud-api/internal/utils/response"
)

func Create(w http.ResponseWriter, r *http.Request) {
	slog.Info("Creating a Student")

	var student dtypes.Sudent

	err := json.NewDecoder(r.Body).Decode(&student)

	if errors.Is(err, io.EOF) {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
		return
	}

	if err != nil {
		response.WriteJson(w, http.StatusBadGateway, response.GeneralError(err))
		return
	}

	//Request Validation
	if err := validator.New().Struct(student); err != nil {
		validateErrs := err.(validator.ValidationErrors)
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
		return
	}

	response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
}
