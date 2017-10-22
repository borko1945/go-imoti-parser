package main

import (
	"encoding/json"
	"os"
)

var config *Config;

type Config struct {
	Email struct {
		Simulate bool
		From string `json:"from"`
		To []string `json:"to"`
		Pass string `json:"pass"`
	} `json:"email"`
	PagesToParse int `json:"pagesToParse"`
	ProcessAfterParse bool `json:"processAfterParse"`
	URL string `json:"url"`
}

type Configuration struct {
	Users    []string
	Groups   []string
}

func Cfg() *Config{
	return Load("");
}

func Load(file string) *Config {
	if (config != nil) {
		return config;
	}

	if (len(file) == 0) {
			file = "config.json"
	}

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		LogError("smtp error: " + err.Error())
		return nil;
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		LogError(err.Error())
		return nil;
	}
	
	return config
}