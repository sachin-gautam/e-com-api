package student

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/sachin-gautam/go-crud-api/internal/model"
	"github.com/sachin-gautam/go-crud-api/internal/utils/response"
)

func (h StudentHandler) GetJson(w http.ResponseWriter, r *http.Request) {
	slog.Info("getting from external json")
	var externalList model.ExternalList
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&externalList)
	if err != nil {
		response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
		return
	}
	response.WriteJson(w, http.StatusOK, externalList)
}
