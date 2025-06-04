package Note

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
