package api

import (
	"github.com/gorilla/websocket"
	"log/slog"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *Application) handleWebSocketConnection(conn *websocket.Conn, chatID, userID string) error {
	defer func() {
		app.removeWebSocketConn(chatID, userID)
		if err := conn.Close(); err != nil {
			slog.Error("closing web socket conn", "error", err)
		}
	}()

	app.addWebSocketConn(chatID, userID, conn)

	for {
		err := app.HandleIncomingMessages(chatID, userID, conn)
		if err != nil {
			return err
		}
	}
}

func (app *Application) addWebSocketConn(chatID, userID string, conn *websocket.Conn) {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	if app.ChatConnections[chatID] == nil {
		app.ChatConnections[chatID] = make(map[string]*websocket.Conn)
	}
	app.ChatConnections[chatID][userID] = conn
}

func (app *Application) removeWebSocketConn(chatID, userID string) {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	if app.ChatConnections[chatID] == nil {
		return
	}

	delete(app.ChatConnections[chatID], userID)
	if len(app.ChatConnections[chatID]) == 0 {
		delete(app.ChatConnections, chatID)
	}
}
