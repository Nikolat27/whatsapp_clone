package auth

import (
	"errors"
	"fmt"
	"github.com/thanhpk/randstr"
	"net/http"
	"time"
)

func IsUserAuthenticated(r *http.Request) (bool, error) {
	cookie, err := r.Cookie("authToken")
	if errors.Is(err, http.ErrNoCookie) {
		return false, nil // user is not logged in
	}

	if err != nil {
		fmt.Println("Error: ", err)
		return false, err
	}

	return cookie != nil, nil
}

func GetUserAuthToken(r *http.Request) ([]byte, error) {
	cookie, err := r.Cookie("authToken")
	if errors.Is(err, http.ErrNoCookie) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return []byte(cookie.Value), nil
}

func GenerateRandomString(length int) string {
	return randstr.String(length)
}

func GenerateHTTPOnlyCookie(name, value string, duration time.Duration) *http.Cookie {
	expireTime := convertTimeToDuration(duration)

	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  expireTime,
	}
}

func DeleteHTTPOnlyCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}

func convertTimeToDuration(duration time.Duration) time.Time {
	timeNow := time.Now()
	return timeNow.Add(duration)
}
