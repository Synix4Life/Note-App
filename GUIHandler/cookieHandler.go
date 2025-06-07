package GUIHandler

import "net/http"

func GetUsername(r *http.Request) (string, error) {
	cookie, err := r.Cookie("username")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
