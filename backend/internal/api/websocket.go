package api

import (
	"encoding/json"
	"errors"
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

func (app *Application) addConnection(chatID, userID string, conn *websocket.Conn) {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	if app.ChatConnections[chatID] == nil {
		app.ChatConnections[chatID] = make(map[string]*websocket.Conn)
	}
	app.ChatConnections[chatID][userID] = conn
}

func (app *Application) handleWebSocketConnection(conn *websocket.Conn, chatID, userID string) error {
	defer func() {
		app.removeConnection(chatID, userID)
		err := conn.Close()
		if err != nil {
			slog.Error("closing web socket conn", "error", err)
		}
	}()

	_, payload, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	var msg struct {
		Type      string `json:"type"`
		PublicKey string `json:"publicKey"`
	}
	if err := json.Unmarshal(payload, &msg); err != nil || msg.Type != "public_key" {
		return errors.New("expected public key message")
	}

	app.ConnMu.Lock()
	app.PublicKeys[userID] = msg.PublicKey
	app.ConnMu.Unlock()

	app.addConnection(chatID, userID, conn)

	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		err = app.CreateMessage(chatID, userID, payload) // Insert into the database
		if err != nil {
			err2 := conn.WriteJSON(err.Error())
			if err2 != nil {
				return err2
			}
			return err
		}
		if err = app.publishMessage(chatID, userID, messageType, payload); err != nil {
			return err
		}
	}
}

func (app *Application) publishMessage(chatID, senderUserID string, messageType int, message []byte) error {
	app.ConnMu.Lock()
	defer app.ConnMu.Unlock()

	connections, ok := app.ChatConnections[chatID]
	if !ok {
		return errors.New("connection does not exist")
	}

	for userId, conn := range connections {
		if userId == senderUserID {
			continue // Sender does not receive their own msg
		}
		
		if err := conn.WriteMessage(messageType, message); err != nil {
			app.removeConnection(chatID, userId)
			return err
		}
	}
	return nil
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
