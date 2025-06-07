package main

import (
	"NoteApp/GUIHandler"
	"NoteApp/Note"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const filename string = "data.json"

	data, err := Note.LoadNotes(filename)
	if err != nil {
		data = make(Note.UserNotes)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", GUIHandler.ServeIndexPage)
	mux.HandleFunc("/login", GUIHandler.LoginHandler)
	mux.HandleFunc("/usr/", func(w http.ResponseWriter, r *http.Request) {
		_, err := GUIHandler.GetUsername(r)
		if err != nil {
			http.Error(w, "No username provided", http.StatusForbidden)
			return
		}
		GUIHandler.ServeUserPage(w, r)
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
