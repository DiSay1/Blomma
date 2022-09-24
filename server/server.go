package server

import (
	"net/http"

	"github.com/DiSay1/Blomma/console"
)

var log *console.Logger // Server logger

// Server start function
func StartServer() {
	log = console.NewLogger("SERVER") // Creating a logger

	// Loading handler paths
	log.Info("Loading Paths...")
	if err := LoadPaths(); err != nil {
		log.Fatal("An error occurred while trying to load router paths. Error:", err)
	}

	for _, handler := range Handlers { // Specifying paths for processing
		if handler.isWebSocket {
			http.HandleFunc(handler.Address, handler.websocketHandler) // For websocket paths
		} else if handler.Type == "lua" && !handler.isWebSocket { // For lua paths
			if err := handler.State.DoFile(handler.Path); err != nil { // Loading the handler file into the virtual machine stack
				log.Fatal("File compilation error. Error:", err)
				return
			}

			if handler.Address == "/" { // If it is the main path ("/")
				http.HandleFunc("/", handler.indexRouter)
			} else { // If not
				http.HandleFunc(handler.Address, handler.addressHandler)
			}
		} else if handler.Type == "html" { // If it is an HTML handler
			http.HandleFunc(handler.Address, handler.addressHandler)
		}
	}

	log.Info("Paths loaded successfully!")

	go func() { // Starting the HTTP Server
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("The server is not running. Error:", err)
		}
	}()
	log.Info("Server started successfully!")
}
