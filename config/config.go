package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BaseURL string // Base URL (ex: https://site.com)
	Keys map[string]string // Map with license : user pair
	EmbedColor string // discord embed color
}

var config *Config

func init() { // Load config.json
	if data, err := ioutil.ReadFile("./config.json"); err == nil {
		if err = json.Unmarshal(data, &config); err != nil {
			log.Fatalf("Failed to deserialize config data: %s", err)
		}
	} else {
		log.Printf("Failed to read config.json: %s\nCreating config.json", err)

		config = &Config{BaseURL: "localhost:8080", Keys: map[string]string{"some_api_key": "api_key_owner"}, EmbedColor: "bf8dec"}
		config.Save()
	}
}

func Get() *Config {return config} // Config.Get() is more appealing than Config.Cfg

func (c *Config) Save() {
	if data, err := json.MarshalIndent(c, "", "\t"); err == nil {
		if err = ioutil.WriteFile("./config.json", data, os.ModePerm); err != nil {
			log.Printf("Failed to write config.json: %s\n", err)
		}
	} else {
		log.Printf("Failed to save config: %s\n", err)
	}
}