package internal

import (
	"encoding/json"
	"os"
)

func ParseJson(fileName string) (DynamicContent, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return DynamicContent{}, err
	}
	defer file.Close()

	var content DynamicContent

	err = json.NewDecoder(file).Decode(&content)
	if err != nil {
		return DynamicContent{}, err
	}
	return content, nil
}
