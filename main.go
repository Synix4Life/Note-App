package main

import (
	"NoteApp/Note"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loop(data Note.UserNotes, reader *bufio.Reader, username string) {
	var key string
	for {
		fmt.Println("Please input the operation (read, write, delete, delete_all, exit):")
		key, _ = reader.ReadString('\n')
		key = strings.TrimSpace(key)
		if key == "exit" {
			break
		}
		switch key {
		case "read":
			Note.Print(data, username)
		case "write":
			data[username] = append(data[username], Note.CreateNote(reader))
		case "delete":
			var delKey string
			fmt.Println("Please input which to delete")
			delKey, _ = reader.ReadString('\n')
			if Note.DeleteNote(data, username, strings.TrimSpace(delKey)) {
				fmt.Println("Note deleted.")
			} else {
				fmt.Println("Note not found.")
			}
		case "delete_all":
			Note.ClearNotes(data, username)
		default:
			fmt.Println("No such operation.")
		}
	}
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

	loop(data, reader, strings.TrimSpace(username))

	err = Note.SaveNotes(filename, data)
	if err != nil {
		fmt.Println("Error saving notes:", err)
		return
	}
}
