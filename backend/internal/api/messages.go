package api

import (
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	. "whatsapp_clone/internal/errors"
	convertHelper "whatsapp_clone/internal/helpers/convert"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func (app *Application) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	msgId, err := convertHelper.StringToObjectId(id)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = app.Models.Message.DeleteMessageInstance(msgId); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonHelper.WriteJSON(w, http.StatusOK, "msg deleted successfully")
}

func (app *Application) HandleIncomingMessages(chatID, userID string, conn *websocket.Conn) error {
	messageType, payload, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	// parsing, encryption and Insert into the db
	if err = app.saveMessage(chatID, userID, payload); err != nil {
		return err
	}

	return app.broadcastMessage(chatID, userID, messageType, payload)
}

func (app *Application) saveMessage(chatId, senderId string, payload []byte) error {
	chatObjectId, err := convertHelper.StringToObjectId(chatId)
	if err != nil {
		return err
	}

	senderObjectId, err := convertHelper.StringToObjectId(senderId)
	if err != nil {
		return err
	}

	cipherText, err := app.Cipher.Encrypt(payload)
	if err != nil {
		return err
	}
	
	if err = app.Models.Message.InsertMessageInstance(chatObjectId, senderObjectId, cipherText); err != nil {
		return err
	}

	return nil
}

func (app *Application) broadcastMessage(chatID, senderId string, msgType int, payload []byte) error {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	connections, ok := app.ChatConnections[chatID]
	if !ok {
		return errors.New("connection does not exist")
	}

	for userId, conn := range connections {
		if userId == senderId {
			continue // Sender does not receive their own msg
		}

		if err := conn.WriteMessage(msgType, payload); err != nil {
			app.removeWebSocketConn(chatID, userId)
			return err
		}
	}
	return nil
}
