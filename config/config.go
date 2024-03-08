package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path"

	"github.com/gustavobertoi/hermes/pkg"
	"github.com/gustavobertoi/hermes/signatures"
)

var defaultFolderName = ".hermes"
var defaultFileName = "hermes.json"

var signaturesFolder = "signatures"

var config *Config

type Config struct {
	Signatures map[string]*PersonalSignature `json:"signatures"`
	folderPath string
}

func GetConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}

	// Get the user home directory and create the .hermes folder if it doesn't exist
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	folderPath := path.Join(usr.HomeDir, defaultFolderName)
	filePath := path.Join(folderPath, defaultFileName)

	if err := pkg.CreateFolderIfNotExists(folderPath); err != nil {
		return nil, err
	}

	if !pkg.ExistsFile(filePath) {
		defaultConfig := &Config{
			Signatures: make(map[string]*PersonalSignature),
		}
		if err := defaultConfig.SaveConfig(); err != nil {
			return nil, err
		}
		return defaultConfig, nil
	}

	// Open the hermes.json file and read the data
	config = &Config{}
	pkg.ReadFileAndBindToJson(filePath, config)

	config.folderPath = folderPath

	return config, nil
}

func (c *Config) SaveConfig() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	filePath := path.Join(c.folderPath, defaultFileName)
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

func (c *Config) AddSignature(algorithm string, sig signatures.Signature) error {
	if c.Signatures == nil {
		c.Signatures = make(map[string]*PersonalSignature)
	}

	folderPath := path.Join(c.folderPath, signaturesFolder, algorithm)

	c.Signatures[algorithm] = NewPersonalSignature(algorithm, folderPath, sig)

	if err := pkg.CreateFolderIfNotExists(folderPath); err != nil {
		return err
	}
	if err := sig.Save(folderPath); err != nil {
		return err
	}
	if err := c.SaveConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) GetSignature(algorithm string) (signatures.Signature, error) {
	if c.Signatures == nil {
		c.Signatures = make(map[string]*PersonalSignature)
	}
	if sig, ok := c.Signatures[algorithm]; ok {
		return sig.signature, nil
	}
	return nil, signatures.ErrSignatureNotFound
}
