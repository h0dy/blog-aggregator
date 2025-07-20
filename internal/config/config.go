package config

import (
	"encoding/json"
	"os"
)

// Config struct holds a json file
type Config struct {
	DbURL  			string `json:"db_url"` // database connection
	CurrentUsername string `json:"current_user_name"` // current active user
}

// ReadConfigFile func reads the config from the json file and return it as Config struct
func ReadConfigFile() (Config, error) {
	// get the full path for the config
	filePath, err := getConfigFilePath() 
	if err != nil {
		return Config{}, err
	}

	// open the config file
	file, err := os.Open(filePath) 
	if err != nil {
		return Config{}, err
	}

	defer file.Close() // close the file after the function exits

	// decode the json data into a Config struct
	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	
	return config, nil
}

// SetUser func sets the current in the Config and update the json config
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username
	return write(cfg)
}

// write func saves the current config to the file in JSON format
func write(cfg *Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Create/overwrite the config file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close() // close the file after the function exits

	// encode the Config file as json and write to it
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	
	return nil
}

// getConfigFilePath returns the full path to the config file in the home directory.
func getConfigFilePath() (string, error) {
	configFileName := "/.gatorconfig.json"
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homePath + configFileName, nil
}