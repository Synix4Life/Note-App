package Note

import "fmt"

func Print(data UserNotes, username string) {
	notes := data[username]
	if len(notes) == 0 {
		fmt.Println("No notes found!")
	}
	for _, n := range notes {
		fmt.Println(n.Title + " : \"" + n.Content + "\", on: " + n.Date)
	}
}
