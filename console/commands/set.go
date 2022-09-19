package commands

import (
	"strings"

	"github.com/DiSay1/Blomma/server/states"
)

func set(args ...string) {
	if len(args) < 2 {
		logger.Panic("The number of arguments is less than 2.")
		return
	}

	args[0] = strings.ToUpper(args[0])
	args[1] = strings.ToLower(args[1])

	switch args[0] {
	case "DEV_MODE":
		if args[1] == "true" {
			states.DEV_MODE = true
			logger.Info("DEV_MODE = true")
			break
		} else if args[1] == "false" {
			states.DEV_MODE = false
			logger.Info("DEV_MODE = false")
			break
		}
	}
}
