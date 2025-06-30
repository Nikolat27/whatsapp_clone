package api

import (
	"fmt"
	"net/http"
	"time"
	"whatsapp_clone/internal/errors"
	authHelper "whatsapp_clone/internal/helpers/auth"
	jsonHelper "whatsapp_clone/internal/helpers/json"
	passwordHelper "whatsapp_clone/internal/helpers/password"
)

func (app *Application) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	err := authHelper.IsUserAuthenticated(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		FullName string `json:"fullName"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := jsonHelper.DeSerializeJSON(r.Body, 10000, &input); err != nil {
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

	if err = app.Models.User.CreateUserInstance(input.Username, input.Password); err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *Application) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	err := authHelper.IsUserAuthenticated(r)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := jsonHelper.DeSerializeJSON(r.Body, 10000, &input); err != nil {
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

	hashedPassword := user.Password
	arePasswordsSame := passwordHelper.VerifyPassword([]byte(input.Password), []byte(hashedPassword))
	if arePasswordsSame == false {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "invalid password")
		return
	}

	userToken := authHelper.GenerateRandomString(10)
	userId := fmt.Sprintf("user-%s", user.Id)

	tokenExpireTime := 24 * time.Hour * 7 // 7 days

	app.Rli.Set(userToken, userId, tokenExpireTime)

	cookie := authHelper.GenerateHTTPOnlyCookie(userToken, tokenExpireTime)
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
}

func (app *Application) CheckAuthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("authToken")
	if err != nil {
		http.Error(w, "Error reading cookie", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
