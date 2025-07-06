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

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// security headers for all responses
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	})

	// Auth
	router.HandlerFunc(http.MethodPost, "/api/auth/register/", app.RegisterUserHandler)
	router.HandlerFunc(http.MethodPost, "/api/auth/login/", app.LoginUserHandler)
	router.HandlerFunc(http.MethodDelete, "/api/auth/logout", app.LogoutUserHandler)
	router.HandlerFunc(http.MethodGet, "/api/auth/check", app.CheckAuthHandler)

	// User
	router.HandlerFunc(http.MethodPost, "/api/users/chats/", app.GetUserChatsHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/", app.GetUserInfoHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/search/", app.SearchUserByUsernameHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/search-by-id/", app.SearchUserByIdHandler)
	router.HandlerFunc(http.MethodPut, "/api/users/update/", app.UpdateUserHandler)
	router.HandlerFunc(http.MethodPut, "/api/users/update-profile", app.UploadProfilePicHandler)

	router.HandlerFunc(http.MethodGet, "/api/users/get/save-messages", app.GetAllSaveMessagesHandler)
	router.HandlerFunc(http.MethodPost, "/api/users/insert/save-messages/", app.InsertSaveMessageHandler)

	// Message
	router.HandlerFunc(http.MethodDelete, "/api/messages/delete/:id", app.DeleteMessageHandler)

	// Chat
	router.HandlerFunc(http.MethodPost, "/api/chats/create/", app.CreateChatHandler)
	router.HandlerFunc(http.MethodPost, "/api/chats/get/", app.GetChatMessagesHandler)
	router.HandlerFunc(http.MethodDelete, "/api/chats/delete/:id", app.DeleteChatHandler)

	// Web socket
	router.HandlerFunc(http.MethodGet, "/api/chats/open-socket", app.ChatSocketHandler)

	// Static files
	secureFileServer(router, "/uploads/", http.Dir("./uploads"))

	return router
}


func secureFileServer(router *httprouter.Router, path string, root http.FileSystem) {
	if path[len(path)-1] != '/' {
		panic("path must end with slash")
	}

	handler := http.StripPrefix(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		http.FileServer(root).ServeHTTP(w, r)
	}))

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
				w.Header().Set("Vary", "Origin")
				break
			}
		}
		
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
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

	originList := strings.Split(origins, ",")
	var cleanedOrigins []string
	
	for _, origin := range originList {
		clean := strings.TrimSpace(origin)
		if clean != "" {
			cleanedOrigins = append(cleanedOrigins, clean)
		}
	}

	if len(cleanedOrigins) == 0 {
		return nil, errors.New("no valid origins in ALLOWED_CORS_ORIGINS")
	}

	return cleanedOrigins, nil
}
