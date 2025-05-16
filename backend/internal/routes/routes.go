package routes

import (
	"github.com/julienschmidt/httprouter"
	"whatsapp_clone/internal/api"
)

func Route(app *api.Application) *httprouter.Router {
	router := httprouter.New()

	return router
}
