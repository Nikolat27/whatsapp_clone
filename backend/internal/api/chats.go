package api

import (
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	. "whatsapp_clone/internal/errors"
	convertHelper "whatsapp_clone/internal/helpers/convert"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func (app *Application) CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := app.GetUserId(r)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		ReceiverUsername string `json:"receiver_username"`
	}

	if err := jsonHelper.DeSerialize(r.Body, 10000, &input); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	receiverUserId, err := app.Models.User.GetUserInstanceByUsername(input.ReceiverUsername)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var participants = []primitive.ObjectID{userId, receiverUserId.Id}

	chatId, err := app.Models.Chat.CheckChatExists(participants)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if chatId == "" {
		chatId, err = app.Models.Chat.CreateChatInstance(participants, nil)
		if err != nil {
			ServerErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	data := map[string]any{
		"chat_id": chatId,
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}

func (app *Application) ChatSocketHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := app.GetUserId(r)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		ServerErrorResponse(w, http.StatusBadRequest, "chat_id is a required parameter")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, "websocket upgrade failed")
		return
	}

	app.addWebSocketConn(chatID, userID.Hex(), conn)

	go func() {
		if err = app.handleWebSocketConnection(conn, chatID, userID.Hex()); err != nil {
			ServerErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}()
}

func (app *Application) GetChatMessagesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ChatId string `json:"chat_id"`
	}

	if err := jsonHelper.DeSerialize(r.Body, 100000, &input); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatId, err := convertHelper.StringToObjectId(input.ChatId)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	msgs, err := app.Models.Message.GetAllMessages(chatId)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.Cipher.DecryptChatMessages(&msgs); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonHelper.WriteJSON(w, http.StatusOK, msgs)
}

func (app *Application) DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatId, err := convertHelper.StringToObjectId(id)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = app.Models.Chat.DeleteChatInstance(chatId); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}

func readIDParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	if len(id) < 1 {
		return "", errors.New("no id received")
	}
	return id, nil
}
