package goutil

import (
	"encoding/json"
	"os"
)

func LoadJsonFile(filePath string, output interface{}) (*[]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if output != nil {
		if err = json.Unmarshal(fileContent, output); err != nil {
			return nil, err
		}
	}

	return &fileContent, nil
}
