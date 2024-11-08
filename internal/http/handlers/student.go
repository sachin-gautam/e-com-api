package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/sachin-gautam/go-crud-api/internal/dtypes"
	"github.com/sachin-gautam/go-crud-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student dtypes.Sudent

		if err := json.NewDecoder(r.Body).Decode(&student); errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		slog.Info("Creating a Student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
