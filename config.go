package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

type state struct {
	config *Config
}

func getGatorConfigFilePath() string {
	const fileName string = ".gatorconfig.json"
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, fileName)
}

func ReadConfig() Config {
	fileToRead := getGatorConfigFilePath()
	var config Config

	configFile, err := os.Open(fileToRead)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return config
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return config
	}

	return config

}

func WriteConfig(config Config) {
	file, err := os.Create(getGatorConfigFilePath())
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print with indentation
	if err := encoder.Encode(config); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}
}

func SetUser(s *state, usr string) error {
	s.config.Current_user_name = usr
	WriteConfig(*s.config)
	return nil
}
