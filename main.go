package main

import (
	"log"
	"net/http"

	"github.com/DiSay1/Tentanto/console"
	"github.com/DiSay1/Tentanto/router"
)

func main() {
	http.HandleFunc("/", router.Router)

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalln("| Starting is not possible.\n Launch error:", err)
		}
	}()

	console.StartConsole()
}
