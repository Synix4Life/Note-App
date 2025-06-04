package main

import (
	"NoteApp/GUIHandler"
	"NoteApp/Note"
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func servePage(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("templates/index.html")
	if err != nil {
		http.Error(w, "Couldn't load page", 500)
	}
	w.Write(page)
}

func main() {
	const filename string = "data.json"

	data, err := Note.LoadNotes(filename)
	if err != nil {
		data = make(Note.UserNotes)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your username")
	var username, _ = reader.ReadString('\n')
	username = strings.TrimSpace(username)

	mux := http.NewServeMux()
	mux.HandleFunc("/", servePage)
	mux.HandleFunc("/write", GUIHandler.MakeWriteHandler(data, username))
	mux.HandleFunc("/read", GUIHandler.MakeReadHandler(data, username))
	mux.HandleFunc("/delete", GUIHandler.MakeDeleteHandler(data, username))
	mux.HandleFunc("/delete_all", GUIHandler.MakeDeleteAllHandler(data, username))
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
