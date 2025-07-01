package api

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strings"
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
	router.HandlerFunc(http.MethodPost, "/api/users/", app.GetUserInfoHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/search/", app.SearchUserByUsernameHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/search-by-id/", app. SearchUserByIdHandler)
	router.HandlerFunc(http.MethodPut, "/api/users/update/", app.UpdateUserHandler)
	router.HandlerFunc(http.MethodPut, "/api/users/update-profile", app.UploadProfilePicHandler)

	// Message
	router.HandlerFunc(http.MethodDelete, "/api/messages/delete/:id", app.DeleteMessageHandler)

	// Chat
	router.HandlerFunc(http.MethodPost, "/api/chats/create/", app.CreateChatHandler)
	router.HandlerFunc(http.MethodPost, "/api/chats/get/", app.GetChatHandler)
	router.HandlerFunc(http.MethodDelete, "/api/chats/delete/:id", app.DeleteChatHandler)

	// Web socket
	router.HandlerFunc(http.MethodGet, "/api/chats/open-socket", app.ChatSocketHandler)

	// Static files
	fileServer(router, "/uploads/", http.Dir("./uploads"))

	return router
}

func fileServer(router *httprouter.Router, path string, root http.FileSystem) {
	if path[len(path)-1] != '/' {
		panic("path must end with slash")
	}
	// StripPrefix requires path to end with "/"
	handler := http.StripPrefix(path, http.FileServer(root))

	router.Handler("GET", path+"*filepath", handler)
}

func CORSMiddleware(next http.Handler) http.Handler {
	allowedOrigins, err := getAllowedCORS()
	if err != nil {
		panic(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				w.Header().Set("Access-Control-Allow-Origin", allowed)
				break
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getAllowedCORS() ([]string, error) {
	origins := os.Getenv("ALLOWED_CORS_ORIGINS")
	if origins == "" {
		return nil, errors.New("ALLOWED_CORS_ORIGINS env variable is empty")
	}

	return strings.Split(origins, ","), nil
}
