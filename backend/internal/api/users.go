package api

import (
	"fmt"
	"log"
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

	passwordSame := helpers.VerifyPassword(input.Password, user.Password)
	if passwordSame == false {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "invalid password")
		return
	}
	
	log.Println(user.Id)
	redisKey := fmt.Sprintf("user-%s", user.Id)
	redisValue := helpers.GenerateRandomString(10)
	redisExpiration := 7 * 24 * time.Hour // 7 days

	app.Rli.Set(redisKey, redisValue, redisExpiration)

	helpers.WriteJSON(w, http.StatusOK, redisValue)
}
