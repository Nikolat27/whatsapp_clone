package auth

import (
	"errors"
	"github.com/thanhpk/randstr"
	"net/http"
	"time"
)

func IsUserAuthenticated(r *http.Request) error {
	exist, err := r.Cookie("authToken")
	if err != nil {
		return err
	}

	if exist != nil {
		return errors.New("user is already logged in")
	}

	return nil
}

func GenerateRandomString(length int) string {
	return randstr.String(length)
}

func GenerateHTTPOnlyCookie(authToken string, duration time.Duration) *http.Cookie {
	expireTime := convertTimeToDuration(duration)

	return &http.Cookie{
		Name:     "authToken",
		Value:    authToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  expireTime,
	}
}

func convertTimeToDuration(duration time.Duration) time.Time {
	timeNow := time.Now()
	return timeNow.Add(duration)
}
