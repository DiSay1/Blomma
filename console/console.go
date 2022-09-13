package console

import (
	"fmt"
	"strings"

	"github.com/DiSay1/Blomma/server/states"
)

var logger = NewLogger("console")

func StartConsole() {
	logger.Info("Console loaded successfully")

	for {
		var cmd, args, values string

		fmt.Scan(&cmd, &args, &values)
		Handler(strings.ToLower(cmd), strings.ToUpper(args), strings.ToLower(values))
	}
}

func Handler(cmd, args, values string) {
	switch cmd {
	case "set":
		switch args {
		case "DEBUG_MOD":
			if values == "false" {
				states.DEBUG_MOD = false
				logger.Info("| DEBUG_MOD = " + values)
			} else {
				states.DEBUG_MOD = true
				logger.Info("| DEBUG_MOD = true")
			}
		}
	}
}
