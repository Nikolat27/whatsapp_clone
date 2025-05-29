package api

import (
	"net/http"
	"whatsapp_clone/internal/errors"
	"whatsapp_clone/internal/helpers"
)

func (app *Application) CreateMessage(chatId, senderId string, payload []byte) error {
	chatObjectId, err := helpers.ConvertStringToObjectId(chatId)
	if err != nil {
		return err
	}

	senderObjectId, err := helpers.ConvertStringToObjectId(senderId)
	if err != nil {
		return err
	}

	err = app.Models.Message.InsertMessageInstance(chatObjectId, senderObjectId, payload)
	if err != nil {
		return err
	}
	return nil
}

func (app *Application) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	msgId, err := helpers.ConvertStringToObjectId(id)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = app.Models.Message.DeleteMessageInstance(msgId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "msg deleted successfully")
}
