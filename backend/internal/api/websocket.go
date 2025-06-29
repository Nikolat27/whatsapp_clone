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
		app.removeConnection(chatID, userID)
		if err := conn.Close(); err != nil {
			slog.Error("closing web socket conn", "error", err)
		}
	}()

	app.addConnection(chatID, userID, conn)

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		 // Insert into the db
		if err = app.CreateMessage(chatID, userID, payload); err != nil {
			if err2 := conn.WriteJSON(err.Error()); err2 != nil {
				return err2
			}
			return err
		}
		if err = app.PublishMessageToChat(chatID, userID, messageType, payload); err != nil {
			return err
		}
	}
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

	if app.ChatConnections[chatID] == nil {
		return
	}

	delete(app.ChatConnections[chatID], userID)
	if len(app.ChatConnections[chatID]) == 0 {
		delete(app.ChatConnections, chatID)
	}
}
