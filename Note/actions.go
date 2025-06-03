package Note

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func CreateNote(reader *bufio.Reader) Note {
	fmt.Println("Please input the note title: ")
	var title, _ = reader.ReadString('\n')

	fmt.Println("Please input the note content: ")
	var content, _ = reader.ReadString('\n')

	now := time.Now()
	date := now.Format("2006-01-02")

	return Note{Title: strings.TrimSpace(title), Content: strings.TrimSpace(content), Date: date}
}

func DeleteNote(data UserNotes, username, title string) bool {
	notes := data[username]
	for i, n := range notes {
		if n.Title == title {
			data[username] = append(notes[:i], notes[i+1:]...)
			return true
		}
	}
	return false
}

func ClearNotes(data UserNotes, username string) bool {
	if _, ok := data[username]; ok {
		delete(data, username)
		return true
	}
	return false
}
