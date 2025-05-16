package routes

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"whatsapp_clone/internal/api"
)

func Route(app *api.Application) *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/users/register/", app.RegisterUserHandler)
	router.HandlerFunc(http.MethodPost, "/users/login/", app.LoginUserHandler)

	return router
}
