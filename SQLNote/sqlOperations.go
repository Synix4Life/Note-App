package SQLNote

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

func CreateTable(db *sql.DB) bool {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Note (
		user TEXT,
		title TEXT,
		msg TEXT,
		date DATETIME
	);`)
	if err != nil {
		log.Println("Create table:", err)
		return false
	}
	return true
}

func Write(db *sql.DB, user string, title string, msg string) bool {
	date := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec(`INSERT INTO Note (user, title, msg, date) VALUES (?, ?, ?, ?)`, user, title, msg, date)
	if err != nil {
		log.Println("Insert:", err)
		return false
	}
	return true
}

func Read(db *sql.DB, user string, title string) ([]Message, error) {
	var rows *sql.Rows
	var err error
	if title == "" {
		rows, err = db.Query(`SELECT title, msg, date FROM Note WHERE user = ?`, user)
	} else {
		rows, err = db.Query(`SELECT title, msg, date FROM Note WHERE title = ? AND user = ?`, title, user)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.Title, &m.Msg, &m.Date); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	return results, nil
}

func Delete(db *sql.DB, user string, title string) (bool, error) {
	var result sql.Result
	var err error
	if title == "" {
		result, err = db.Exec(`DELETE FROM Note WHERE user = ?`, user)
	} else {
		result, err = db.Exec(`DELETE FROM Note WHERE user = ? AND title = ?`, user, title)
	}

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
