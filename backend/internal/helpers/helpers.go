package helpers

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log/slog"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		slog.Error("encoding json", "error", err.Error())
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

func ConvertStringToObjectId(str string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(str)
}

func ReadIDParam(r *http.Request) (string, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id := params.ByName("id")
	if len(id) < 1 {
		return "", errors.New("no id received")
	}
	return id, nil
}
