package errors

import (
	"net/http"
	"whatsapp_clone/internal/helpers"
)

func ServerErrorResponse(w http.ResponseWriter, status int, msg string) {
	helpers.WriteJSON(w, status, msg)
}
