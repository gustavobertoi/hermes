package config

import (
	"encoding/json"
	"os"
	"path"
)

var hermesJsonFileName = "hermes.json"

type Config struct {
	Signature  *PersonalSignature `json:"signature"`
	folderPath string
}

func NewConfig(folderPath string) *Config {
	return &Config{
		Signature:  nil,
		folderPath: folderPath,
	}
}

func InitConfigFromJSON(folderPath string) (*Config, error) {
	filePath := path.Join(folderPath, hermesJsonFileName)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	var config *Config
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	config.folderPath = folderPath
	return config, nil
}

func (c *Config) SaveConfig() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	filePath := path.Join(c.folderPath, hermesJsonFileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
