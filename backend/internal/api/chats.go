package api

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	. "whatsapp_clone/internal/errors"
	convertHelper "whatsapp_clone/internal/helpers/convert"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func (app *Application) CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Participants []string `json:"participants"`
	}

	if err := jsonHelper.DeSerializeJSON(r.Body, 10000, &input); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	participants, err := getParticipants(input.Participants)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = app.Models.Chat.CreateChatInstance(participants, nil); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getParticipants(participants []string) ([]primitive.ObjectID, error) {
	if len(participants) != 2 {
		return nil, errors.New("participants number must be 2")
	}

	firstUser, err := convertHelper.StringToObjectId(participants[0])
	if err != nil {
		return nil, err
	}

	secondUser, err := convertHelper.StringToObjectId(participants[1])
	if err != nil {
		return nil, err
	}

	idSlice := []primitive.ObjectID{firstUser, secondUser}
	return idSlice, nil
}

func (app *Application) ChatSocketHandler(w http.ResponseWriter, r *http.Request) {
	chatID := r.URL.Query().Get("chat_id")
	userID := r.URL.Query().Get("user_id")

	if chatID == "" || userID == "" {
		ServerErrorResponse(w, http.StatusBadRequest, "chat_id and user_id are both required parameters")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, "websocket upgrade failed")
		return
	}

	app.addWebSocketConn(chatID, userID, conn)

	go func() {
		if err = app.handleWebSocketConnection(conn, chatID, userID); err != nil {
			ServerErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}()
}

func (app *Application) GetChatHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ChatId string `json:"chat_id"`
	}

	if err := jsonHelper.DeSerializeJSON(r.Body, 100000, &input); err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatId, err := convertHelper.StringToObjectId(input.ChatId)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chat, err := app.Models.Chat.GetChatInstance(chatId)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Println(chat.Participants[0].Hex())
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
