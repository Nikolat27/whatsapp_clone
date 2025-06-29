package api

import (
	"fmt"
	"net/http"
	"time"
	"whatsapp_clone/internal/errors"
	"whatsapp_clone/internal/helpers"
)

func (app *Application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if input.Username == "" || input.Password == "" {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "username or password is missing")
		return
	}

	user, err := app.Models.User.GetUserInstance(input.Username)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if user != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "this username exists")
		return
	}

	err = app.Models.User.CreateUserInstance(input.Username, input.Password)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, "user created successfully!")
}

func (app *Application) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserToken string `json:"userToken"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	isAuthenticated, err := helpers.IsUserAuthenticated(app.Rli, r.Header.Get("X-User-Token"))
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusInternalServerError, "Server Internal Error")
		return
	}
	
	if isAuthenticated == true {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "user is already logged in")
		return
	}

	if input.Username == "" || input.Password == "" {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "username or password is missing")
		return
	}

	user, err := app.Models.User.GetUserInstance(input.Username)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword := user.Password
	arePasswordsSame := helpers.VerifyPassword([]byte(input.Password), []byte(hashedPassword))
	if arePasswordsSame == false {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "invalid password")
		return
	}

	userToken := helpers.GenerateRandomString(10)
	userId := fmt.Sprintf("user-%s", user.Id)
	redisExpiration := 24 * time.Hour * 7 // 7 days

	app.Rli.Set(userToken, userId, redisExpiration)

	helpers.WriteJSON(w, http.StatusOK, userToken)
}

func (app *Application) GetUserChatsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserId string `json:"user_id"`
	}

	err := helpers.DeSerializeJSON(r.Body, 10000, &input)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := helpers.ConvertStringToObjectId(input.UserId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chats, err := app.Models.Chat.GetUserChats(userId)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, chats)
}
