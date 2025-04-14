package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	// Expands ENV vars that are in the ${var} format in the yaml file
	bytes = []byte(os.ExpandEnv(string(bytes)))

	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server           *Server   `yaml:"server,omitempty"`
	DB               *Database `yaml:"database,omitempty"`
	TelegramBotToken *Telegram `yaml:"bot_token,omitempty"`
	TelegramChatID   *Telegram `yaml:"chat_id,omitempty"`
}

type Database struct {
	PSN string `yaml:"psn,omitempty"`
	DB  string `yaml:"db,omitempty"`
}

type Telegram struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN,omitempty"`
	TelegramChatID   string `env:"TELEGRAM_CHAT_ID,omitempty"`
}

type Server struct {
	Port string `yaml:"port,omitempty"`
}
