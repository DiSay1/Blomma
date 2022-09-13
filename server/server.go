package server

import (
	"net/http"

	"github.com/DiSay1/Blomma/console"
)

var log *console.Logger

func StartServer() {
	log = console.NewLogger("SERVER")

	log.Info("Loading Paths...")
	if err := loadPaths(); err != nil {
		log.Fatal("An error occurred while trying to load router paths. Error:", err)
	}

	http.HandleFunc("/", addressHandler)
	log.Info("Paths loaded successfully!")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("The server is not running. Error:", err)
	}
}
