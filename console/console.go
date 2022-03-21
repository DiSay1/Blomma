package console

import (
	"fmt"
	"log"

	"github.com/DiSay1/Tentanto/states"
)

func StartConsole() {
	for {
		var cmd, args, values string

		fmt.Scan(&cmd, &args, &values)
		Handler(cmd, args, values)
	}
}

func Handler(cmd, args, values string) {
	switch cmd {
	case "set":
		switch args {
		case "DEBUG_MOD":
			if values == "false" {
				states.DEBUG_MOD = false
				log.Println("| DEBUG_MOD = " + values)
			} else {
				states.DEBUG_MOD = true
				log.Println("| DEBUG_MOD = true")
			}
		}
	}
}
