package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read()( Config, error){
	config := Config{}
	homeDir, err := os.UserHomeDir() 
	if err != nil {
		return config, fmt.Errorf("error finding home directory %w", err)
	}

	data, err:= os.ReadFile(homeDir + "/.gatorconfig.json")
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("error unmarshalling config data: %w", err)
	}

	return config, nil
}

type Config struct {
	DbURL string `json:"db_url"`
}

