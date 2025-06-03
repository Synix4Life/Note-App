package Note

import (
	"encoding/json"
	"os"
)

func LoadNotes(filename string) (UserNotes, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data UserNotes
	err = json.NewDecoder(file).Decode(&data)
	return data, err
}

func SaveNotes(filename string, data UserNotes) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
