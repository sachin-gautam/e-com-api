package student

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator" // Adjust path as needed
	jwtutil "github.com/sachin-gautam/go-crud-api/internal/utils/jwt"
	"github.com/sachin-gautam/go-crud-api/internal/utils/response"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h StudentHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	if err := validator.New().Struct(loginReq); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
		return
	}

	if loginReq.Username != "admin" || loginReq.Password != "password" {
		response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(errors.New("invalid credentials")))
		return
	}

	token, err := jwtutil.GenerateToken(loginReq.Username)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	// Send token in response
	response.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}
