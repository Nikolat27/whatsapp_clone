package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	. "whatsapp_clone/internal/errors"
	convertHelper "whatsapp_clone/internal/helpers/convert"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

type Response struct {
	Id       string
	MsgText  string
	ChatId   string
	SenderId string
	MsgType  int
}

func (app *Application) HandleIncomingMessages(chatID, userID string, conn *websocket.Conn) error {
	messageType, payload, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	var input struct {
		Text string `json:"msg_text"`
	}

	if err := json.Unmarshal(payload, &input); err != nil {
		return err
	}

	// parsing, encryption and Insert into the db
	msgId, err := app.saveMessage(chatID, userID, []byte(input.Text))
	if err != nil {
		return err
	}

	var resp = &Response{
		Id:       msgId,
		MsgText:  input.Text,
		ChatId:   chatID,
		SenderId: userID,
		MsgType:  messageType,
	}

	return app.broadcastMessage(resp)
}

func (app *Application) saveMessage(chatId, senderId string, payload []byte) (string, error) {
	chatObjectId, err := convertHelper.StringToObjectId(chatId)
	if err != nil {
		return "", err
	}

	senderObjectId, err := convertHelper.StringToObjectId(senderId)
	if err != nil {
		return "", err
	}

	cipherText, err := app.Cipher.Encrypt(payload)
	if err != nil {
		return "", err
	}

	msgId, err := app.Models.Message.InsertMessageInstance(chatObjectId, senderObjectId, cipherText)
	if err != nil {
		return "", err
	}

	return msgId, nil
}

func (app *Application) broadcastMessage(resp *Response) error {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	connections, ok := app.ChatConnections[resp.ChatId]
	if !ok {
		return errors.New("connection does not exist")
	}

	for userId, conn := range connections {
		//if UserId == resp.SenderId {
		//	continue // Sender does not receive his own msg
		//}

		data, err := json.Marshal(resp)
		if err != nil {
			return err
		}

		if err := conn.WriteMessage(resp.MsgType, data); err != nil {
			app.removeWebSocketConn(resp.ChatId, userId)
			return err
		}
	}
	return nil
}

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
