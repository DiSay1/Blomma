package main

import (
	"fmt"

	"github.com/DiSay1/Blomma/config"
	"github.com/DiSay1/Blomma/console"
	"github.com/DiSay1/Blomma/server"
)

func main() {
	log := console.NewLogger("APP")
	log.Info("Launching the console...")
	log.Info("Application launch...")

	log.Info("Loading config...")
	config.LoadConfig()
	log.Info("Config uploaded successfully")

	go server.StartServer()
	log.Info("Server start...")

	log.Info("Application launched successfully!")
	fmt.Scanln()
}
