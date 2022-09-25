package config

import (
	"os"

	"github.com/DiSay1/Blomma/console"
	"github.com/DiSay1/Blomma/server/states"
	"github.com/pelletier/go-toml"
)

// I think everything is clear here,
// I hope comments are not needed

type Config struct {
	ServerConfig `toml:"server"`

	SSLConfig `toml:"SSL"`

	DevMode bool `toml:"dev_mode"`
}

type ServerConfig struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

type SSLConfig struct {
	SSL      bool   `toml:"ssl"`
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

type DBConfig struct {
	//
}

var log = console.NewLogger("config")

var BlommaConfig Config

func LoadConfig() {
	if _, err := os.Stat("./config.toml"); err != nil {
		log.Info("The config file was not found.")
		log.Info("Creating a new config...")

		file, err := os.Create("./config.toml")
		if err != nil {
			log.Fatal("An error occurred while trying to create the config. Error:", err)
			return
		}

		cfg := Config{
			ServerConfig: ServerConfig{
				Address: "",
				Port:    8080,
			},

			SSLConfig: SSLConfig{
				SSL:      false,
				CertFile: "",
				KeyFile:  "",
			},
		}

		data, err := toml.Marshal(cfg)
		if err != nil {
			log.Fatal("An error occurred while trying to create the config. Error:", err)
			return
		}

		_, err = file.Write(data)
		if err != nil {
			log.Fatal("An error occurred while trying to create the config. Error:", err)
			return
		}
		log.Info("New config successfully created!")
	}

	data, err := os.ReadFile("./config.toml")
	if err != nil {
		log.Fatal("An error occurred while trying to load the config. Error:", err)
		return
	}

	if err := toml.Unmarshal(data, &BlommaConfig); err != nil {
		log.Fatal("An error occurred while trying to load the config. Error:", err)
		return
	}

	states.DEV_MODE = BlommaConfig.DevMode
}
