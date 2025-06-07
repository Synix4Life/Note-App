package GUIHandler

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	username := strings.TrimSpace(r.FormValue("username"))
	if username == "" {
		http.Error(w, "Username required", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "username",
		Value: username,
		Path:  "/",
	})

	http.Redirect(w, r, "/usr/?username="+username, http.StatusSeeOther)
}

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		time.Sleep(1 * time.Second)
		Srv.Shutdown(context.Background())
	}()
	w.Write([]byte("Shutting down server..."))
}
