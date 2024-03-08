package pkg

import (
	"encoding/json"
	"os"
)

func CreateFolderIfNotExists(folderPath string) error {
	_, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
				return err
			}
		}
		return err
	}
	return nil
}

func ExistsFile(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ReadFileAndBindToJson(filePath string, v interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
