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

// Login handles the login request and generates a JWT token if the credentials are valid
func (h StudentHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	// Decode JSON request body into the login request struct
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}

	// Validate request body fields
	if err := validator.New().Struct(loginReq); err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
		return
	}

	// Validate credentials (replace with your actual authentication logic)
	if loginReq.Username != "admin" || loginReq.Password != "password" { // Replace with real validation logic
		response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(errors.New("invalid credentials")))
		return
	}

	// Generate JWT token
	token, err := jwtutil.GenerateToken(loginReq.Username)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}

	// Send token in response
	response.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}
