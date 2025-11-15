package config

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

func getGatorConfigFilePath() string {
	const fileName string = ".gatorconfig.json"
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, fileName)
}

func Read() Config {
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

func Write(config Config) {
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
	fmt.Println("Saved config file")
}

func SetUser(c Config, usr string) error {
	c.Current_user_name = usr
	Write(c)
	return nil
}
