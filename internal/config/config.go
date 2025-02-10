package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read()(Config, error){
	cfg := Config{}
	configPath, err := getConfigPath()
	if err != nil {
		return cfg, err
	}

	data, err:= os.ReadFile(configPath)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}


func SetUser(cfg *Config, userName string) error{	
	cfg.CurrentUserName = userName
	return write(cfg)
}


func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir() 
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homeDir + configFileName)
	return configPath, nil
}


func write (cfg *Config) error {
	configPath, err := getConfigPath() 
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	
	return nil
 }
