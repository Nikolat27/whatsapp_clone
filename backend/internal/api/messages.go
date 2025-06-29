package api

import (
	"errors"
	"net/http"
	. "whatsapp_clone/internal/errors"
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
	
	if err = app.Models.Message.InsertMessageInstance(chatObjectId, senderObjectId, payload); err != nil {
		return err
	}
	
	return nil
}

func (app *Application) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	msgId, err := helpers.ConvertStringToObjectId(id)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	if err = app.Models.Message.DeleteMessageInstance(msgId); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "msg deleted successfully")
}

func (app *Application) PublishMessageToChat(chatID, senderUserID string, messageType int, message []byte) error {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	connections, ok := app.ChatConnections[chatID]
	if !ok {
		return errors.New("connection does not exist")
	}
	
	for userId, conn := range connections {
		if userId == senderUserID {
			continue // Sender does not receive their own msg
		}

		if err := conn.WriteMessage(messageType, message); err != nil {
			app.removeConnection(chatID, userId)
			return err
		}
	}
	return nil
}
