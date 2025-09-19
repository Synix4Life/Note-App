package GUIHandler

import (
	"NoteApp/SQLNote"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var Srv *http.Server

func MakeWriteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeHandler(w, r, db)
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	username, _ := GetUsername(r)
	var response = GUIRequest{}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(response.Title)
	content := strings.TrimSpace(response.Content)

	if title == "" || content == "" {
		w.Write([]byte("Needed data not provided"))
		return
	}
	SQLNote.Write(db, username, title, content)
	w.Write([]byte("Done!"))
}

func MakeReadHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readHandler(w, r, db)
	}
}

func readHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	username, _ := GetUsername(r)

	var response = GUIRequest{}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(response.Title)

	var messages []SQLNote.Message
	messages, err = SQLNote.Read(db, username, title)

	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(messages) == 0 {
		w.Write([]byte("No data found"))
		return
	}

	var t time.Time
	for _, m := range messages {
		t, err = time.Parse(time.RFC3339, m.Date)
		if err != nil {
			fmt.Println("Error Parsing Date: " + err.Error())
		}

		line := fmt.Sprintf("%s : \"%s\", on: %s\n", m.Title, m.Msg, t.Format("02 Jan 2006 15:04"))
		w.Write([]byte(line))
	}
}

func MakeDeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteHandler(w, r, db)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	username, _ := GetUsername(r)
	var response = GUIRequest{}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(response.Title)
	deleted, err := SQLNote.Delete(db, username, title)
	if err != nil {
		log.Println("Delete error:", err)
		return
	}
	if deleted {
		w.Write([]byte("Deleted Notes"))
	} else {
		w.Write([]byte("No such Note -> None Deleted"))
	}
}

func MakeHelpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Write: Title + Content -> New entry\n"))
		w.Write([]byte("Read: Title -> Read specific entry\n"))
		w.Write([]byte("Read:  -> Read all entries\n"))
		w.Write([]byte("Delete: Title -> Delete specific entry\n"))
		w.Write([]byte("Delete: -> Delete all entries\n"))
	}
}
