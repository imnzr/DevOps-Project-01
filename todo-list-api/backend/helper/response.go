package helper

import (
	"encoding/json"
	"net/http"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
)

func WriteErrorResponse(writter http.ResponseWriter, statusCode int, message string) {
	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(statusCode)
	json.NewEncoder(writter).Encode(response.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   message,
	})
}
