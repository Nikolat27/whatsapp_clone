package json

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// Use a buffer to encode first
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func DeSerializeJSON(body io.Reader, maxBytes int64, obj any) error {
	body = io.LimitReader(body, maxBytes)
	err := json.NewDecoder(body).Decode(&obj)
	return err
}
