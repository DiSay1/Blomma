package main

import (
	"github.com/DiSay1/Blomma/console"
	"github.com/DiSay1/Blomma/server"
)

func main() {
	log := console.NewLogger("APP")
	log.Info("Started app.")

	go server.StartServer()

	console.StartConsole()
}
