package api

import (
	"net/http"
	"whatsapp_clone/internal/errors"
	"whatsapp_clone/internal/helpers"
)

func (app *Application) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ChatId      string `json:"chat_id"`
		SenderId    string `json:"sender_id"`
		ReceiverId  string `json:"receiver_id"`
		TextContent string `json:"text_content"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatId, err := helpers.ConvertStringToObjectId(input.ChatId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	senderId, err := helpers.ConvertStringToObjectId(input.ChatId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	receiverId, err := helpers.ConvertStringToObjectId(input.ChatId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.TextContent == "" {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "message is null")
		return
	}

	err = app.Models.Message.InsertMessageInstance(chatId, senderId, receiverId, input.TextContent)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}
