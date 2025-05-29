package api

import (
	"github.com/gorilla/websocket"
	"net/http"
	"whatsapp_clone/internal/errors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *Application) ChatSocketHandler(w http.ResponseWriter, r *http.Request) {
	chatID := r.URL.Query().Get("chat_id")
	userID := r.URL.Query().Get("user_id")

	if chatID == "" || userID == "" {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "chat_id and user_id are required query parameters")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errors.ServerErrorResponse(w, http.StatusBadRequest, "websocket upgrade failed")
		return
	}

	app.addConnection(chatID, userID, conn)

	go app.handleWebSocketConnection(conn, chatID, userID)
}

func (app *Application) addConnection(chatID, userID string, conn *websocket.Conn) {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	if app.ChatConnections[chatID] == nil {
		app.ChatConnections[chatID] = make(map[string]*websocket.Conn)
	}
	app.ChatConnections[chatID][userID] = conn
}

func (app *Application) removeConnection(chatID, userID string) {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	if app.ChatConnections[chatID] != nil {
		delete(app.ChatConnections[chatID], userID)
		if len(app.ChatConnections[chatID]) == 0 {
			delete(app.ChatConnections, chatID)
		}
	}
}

func (app *Application) handleWebSocketConnection(conn *websocket.Conn, chatID, userID string) {
	defer func() {
		app.removeConnection(chatID, userID)
		conn.Close()
	}()

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}

		err = app.CreateMessage(chatID, userID, payload)
		if err != nil {
			conn.WriteJSON(err.Error())
			break // terminate the socket
		}
		app.publishMessage(chatID, userID, messageType, payload)
	}
}

func (app *Application) publishMessage(chatID, senderUserID string, messageType int, message []byte) {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	connections, ok := app.ChatConnections[chatID]
	if !ok {
		return
	}

	for userId, conn := range connections {
		if userId == senderUserID {
			continue // Sender does not receive their own msg
		}
		err := conn.WriteMessage(messageType, message)
		if err != nil {
			app.removeConnection(chatID, userId)
		}
	}
}
