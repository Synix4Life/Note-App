package main

import (
	"NoteApp/GUIHandler"
	"NoteApp/Note"
	_ "bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func serveUserPage(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("templates/user.html")
	if err != nil {
		http.Error(w, "Couldn't load page", 500)
	}
	w.Write(page)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("templates/index.html")
	if err != nil {
		http.Error(w, "Couldn't load page", 500)
		return
	}
	w.Write(page)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	const filename string = "data.json"

	data, err := Note.LoadNotes(filename)
	if err != nil {
		data = make(Note.UserNotes)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexPage)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/usr/", func(w http.ResponseWriter, r *http.Request) {
		_, err := GUIHandler.GetUsername(r)
		if err != nil {
			http.Error(w, "No username provided", http.StatusForbidden)
			return
		}
		serveUserPage(w, r)
	})
	mux.HandleFunc("/write", GUIHandler.MakeWriteHandler(data))
	mux.HandleFunc("/read", GUIHandler.MakeReadHandler(data))
	mux.HandleFunc("/delete", GUIHandler.MakeDeleteHandler(data))
	mux.HandleFunc("/delete_all", GUIHandler.MakeDeleteAllHandler(data))
	mux.HandleFunc("/shutdown", GUIHandler.ShutdownHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	GUIHandler.Srv = &http.Server{Addr: ":8080", Handler: mux}
	fmt.Println("Server started at http://localhost:8080")

	if err := GUIHandler.Srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %s", err)
	}

	err = Note.SaveNotes(filename, data)
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return
	}
}
