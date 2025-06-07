package GUIHandler

import (
	"NoteApp/Note"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var Srv *http.Server

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		time.Sleep(1 * time.Second)
		Srv.Shutdown(context.Background())
	}()
	w.Write([]byte("Shutting down server..."))
}

func MakeWriteHandler(data Note.UserNotes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeHandler(w, r, data)
	}
}

func writeHandler(w http.ResponseWriter, r *http.Request, data Note.UserNotes) {
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
		w.Write([]byte("No data provided"))
		return
	}
	date := time.Now().Format("2006-01-02 15:04:05")
	data[username] = append(data[username], Note.Note{Title: strings.TrimSpace(title), Content: strings.TrimSpace(content), Date: date})
	w.Write([]byte("Done!"))
}

func MakeReadHandler(data Note.UserNotes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readHandler(w, r, data)
	}
}

func readHandler(w http.ResponseWriter, r *http.Request, data Note.UserNotes) {
	username, _ := GetUsername(r)
	notes := data[username]
	if len(notes) == 0 {
		w.Write([]byte("No data found"))
	}
	for _, n := range notes {
		w.Write([]byte(n.Title + " : \"" + n.Content + "\", on: " + n.Date + "\n"))
	}
}

func MakeDeleteHandler(data Note.UserNotes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteHandler(w, r, data)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request, data Note.UserNotes) {
	username, _ := GetUsername(r)
	var response = GUIRequest{}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(response.Title)
	if title == "" {
		w.Write([]byte("No data provided"))
		return
	}
	Note.DeleteNote(data, username, title)
	w.Write([]byte("Deleted"))
}

func MakeDeleteAllHandler(data Note.UserNotes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteAllHandler(w, r, data)
	}
}

func deleteAllHandler(w http.ResponseWriter, r *http.Request, data Note.UserNotes) {
	username, _ := GetUsername(r)
	Note.ClearNotes(data, username)
	w.Write([]byte("Deleted all"))
}
