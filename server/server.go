package server

import (
	"net/http"

	"github.com/DiSay1/Blomma/console"
)

var log *console.Logger

func StartServer() {
	log = console.NewLogger("SERVER")

	log.Info("Loading Paths...")
	if err := LoadPaths(); err != nil {
		log.Fatal("An error occurred while trying to load router paths. Error:", err)
	}

	for _, handler := range Handlers {
		if handler.isWebSocket {
			http.HandleFunc(handler.Address, handler.websocketHandler)
		} else if handler.Type == "lua" && !handler.isWebSocket {
			if err := handler.State.DoFile(handler.Path); err != nil {
				log.Fatal("File compilation error. Error:", err)
				return
			}
			if handler.Address == "/" {
				http.HandleFunc("/", handler.indexRouter)
			} else {
				http.HandleFunc(handler.Address, handler.addressHandler)
			}
		}
	}

	log.Info("Paths loaded successfully!")

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("The server is not running. Error:", err)
		}
	}()
	log.Info("Server started successfully!")
}
