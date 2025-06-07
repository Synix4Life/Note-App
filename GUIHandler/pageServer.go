package GUIHandler

import (
	"net/http"
	"os"
)

func ServeUserPage(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("templates/user.html")
	if err != nil {
		http.Error(w, "Couldn't load page", 500)
	}
	w.Write(page)
}

func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("templates/index.html")
	if err != nil {
		http.Error(w, "Couldn't load page", 500)
		return
	}
	w.Write(page)
}
