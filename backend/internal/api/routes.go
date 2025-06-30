package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func routes(app *Application) *httprouter.Router {
	router := httprouter.New()

	// Auth
	router.HandlerFunc(http.MethodPost, "/api/auth/register/", app.RegisterUserHandler)
	router.HandlerFunc(http.MethodPost, "/api/auth/login/", app.LoginUserHandler)
	router.HandlerFunc(http.MethodDelete, "/api/auth/logout", app.LogoutUserHandler)
	router.HandlerFunc(http.MethodGet, "/api/auth/check", app.CheckAuthHandler)

	// User
	router.HandlerFunc(http.MethodPost, "/api/users/chats/", app.GetUserChatsHandler)

	// Message
	router.HandlerFunc(http.MethodDelete, "/api/messages/delete/:id", app.DeleteMessageHandler)

	// Chat
	router.HandlerFunc(http.MethodPost, "/api/chats/create/", app.CreateChatHandler)
	router.HandlerFunc(http.MethodPost, "/api/chats/get/", app.GetChatHandler)
	router.HandlerFunc(http.MethodDelete, "/api/chats/delete/:id", app.DeleteChatHandler)

	// Web socket
	router.HandlerFunc(http.MethodGet, "/api/chats/open-socket", app.ChatSocketHandler)

	return router
}
