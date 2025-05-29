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

	// Message
	router.HandlerFunc(http.MethodPost, "/messages/create/", app.CreateMessageHandler)

	// Chat
	router.HandlerFunc(http.MethodPost, "/chats/create/", app.CreateChatHandler)
	router.HandlerFunc(http.MethodDelete, "/chats/delete/", app.DeleteChatHandler)

	return router
}
