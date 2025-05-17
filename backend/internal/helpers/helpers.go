package helpers

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(msg + "\n"))
	if err != nil {
		w.WriteHeader(500)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func DeSerializeJSON(body io.Reader, maxBytes int64, obj any) error {
	body = io.LimitReader(body, maxBytes)
	err := json.NewDecoder(body).Decode(&obj)
	return err
}

func GenerateRandomString(length int) string {
	return randstr.String(length)
}

func IsUserAuthenticated(rli *redis.Client, userToken string) (bool, error) {
	if userToken == "" {
		return false, nil
	}
	val, err := rli.Get(userToken).Result()
	if err != nil {
		return false, err
	}
	return val != "", nil
}
