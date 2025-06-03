package Note

type Note struct {
	Title   string "json:\"title\""
	Content string "json:\"content\""
	Date    string "json:\"date\""
}

type UserNotes map[string][]Note
