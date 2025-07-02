package api

import (
	"net/http"
	"whatsapp_clone/internal/errors"
	jsonHelper "whatsapp_clone/internal/helpers/json"
)

func (app *Application) GetAllSaveMessagesHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := app.GetUserId(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	msgs, err := app.Models.SaveMessage.GetAllSaveMessages(userId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.Cipher.DecryptSaveMessages(&msgs); err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonHelper.WriteJSON(w, http.StatusOK, msgs)
}

func (app *Application) InsertSaveMessageHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := app.GetUserId(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		MsgContent string `json:"msg_content"`
	}

	msgId, err := app.Models.SaveMessage.InsertSaveMessageInstance(userId, []byte(input.MsgContent))
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(msgId))
}
