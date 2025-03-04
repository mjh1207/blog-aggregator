package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

func Read() (Config, error){
	filePath, err := getConfigFilePath()
	if err != nil { return Config{}, err}
	
	// read json from file path
	jsonFile, err := os.Open(filePath)
	if err != nil { return Config{}, err}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil { return Config{}, err}

	// unmarshal into config struct
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil { return Config{}, err}

	return cfg, nil
	
}

func (cfg *Config) SetUser(username string) error{
	cfg.CurrentUserName = username
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homeDir, configFileName)
	return fullPath, nil
}

func write(cfg Config) error {
	// get file path
	filePath, err := getConfigFilePath()
	if err != nil { return err }

	file, err := os.Create(filePath)
	if err != nil { return err }
	defer file.Close()

	// encode cfg back to json file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil { return err }

	return nil
}