package errors

import (
	"net/http"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func ServerErrorResponse(w http.ResponseWriter, status int, msg string) {
	jsonHelper.WriteJSON(w, status, msg)
}
