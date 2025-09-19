package main

import (
	"NoteApp/GUIHandler"
	"NoteApp/SQLNote"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite", "notes.db")
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer db.Close()

	if !SQLNote.CreateTable(db) {
		log.Fatal("Could not create table")
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
	mux.HandleFunc("/write", GUIHandler.MakeWriteHandler(db))
	mux.HandleFunc("/read", GUIHandler.MakeReadHandler(db))
	mux.HandleFunc("/delete", GUIHandler.MakeDeleteHandler(db))
	mux.HandleFunc("/help", GUIHandler.MakeHelpHandler())
	mux.HandleFunc("/shutdown", GUIHandler.ShutdownHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	GUIHandler.Srv = &http.Server{Addr: ":8080", Handler: mux}
	fmt.Println("Server started at http://localhost:8080")

	if err := GUIHandler.Srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %s", err)
	}
}
