package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"whatsapp_clone/internal/api"
)

func Route(app *api.Application) *httprouter.Router {
	router := httprouter.New()

	// Auth
	router.HandlerFunc(http.MethodPost, "/users/register/", app.RegisterUserHandler)
	router.HandlerFunc(http.MethodPost, "/users/login/", app.LoginUserHandler)
	router.HandlerFunc(http.MethodPost, "/users/chats/", app.GetUserChatsHandler)

	// Message
	router.HandlerFunc(http.MethodDelete, "/messages/delete/:id", app.DeleteMessageHandler)

	// Chat
	router.HandlerFunc(http.MethodPost, "/chats/create/", app.CreateChatHandler)
	router.HandlerFunc(http.MethodPost, "/chats/get/", app.GetChatHandler)
	router.HandlerFunc(http.MethodDelete, "/chats/delete/:id", app.DeleteChatHandler)

	// Web socket
	router.HandlerFunc(http.MethodGet, "/chats/open-socket", app.ChatSocketHandler)
	return router
}
