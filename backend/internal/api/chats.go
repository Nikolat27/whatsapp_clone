package api

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	. "whatsapp_clone/internal/errors"
	"whatsapp_clone/internal/helpers"
)

func getParticipants(participants []string) ([]primitive.ObjectID, error) {
	if len(participants) != 2 {
		return nil, errors.New("participants number must be 2")
	}

	firstUser, err := helpers.ConvertStringToObjectId(participants[0])
	if err != nil {
		return nil, err
	}

	secondUser, err := helpers.ConvertStringToObjectId(participants[1])
	if err != nil {
		return nil, err
	}

	idSlice := []primitive.ObjectID{firstUser, secondUser}
	return idSlice, nil
}

func (app *Application) GetChatHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ChatId string `json:"chat_id"`
	}

	err := helpers.DeSerializeJSON(r.Body, 100000, &input)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatId, err := helpers.ConvertStringToObjectId(input.ChatId)
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

func (app *Application) CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Participants []string `json:"participants"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	participants, err := getParticipants(input.Participants)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = app.Models.Chat.CreateChatInstance(participants, nil)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.WriteJSON(w, http.StatusCreated, "chat created successfully")
}

func (app *Application) DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ChatId string `json:"chat_id"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chatId, err := helpers.ConvertStringToObjectId(input.ChatId)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = app.Models.Chat.DeleteChatInstance(chatId)
	if err != nil {
		ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}
