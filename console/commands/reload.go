package commands

import (
	"strings"

	"github.com/DiSay1/Blomma/server"
)

func reload(args ...string) {
	if len(args) < 1 {
		logger.Panic("The number of arguments is less than 2.")
		return
	}

	args[0] = strings.ToUpper(args[0])

	switch args[0] {
	case "PATHS":
		logger.Info("Trying to reload paths...")

		server.Paths = nil
		if err := server.LoadPaths(); err != nil {
			logger.Info("An error occurred while trying to reload the paths. Error:", err)
			return
		}

		logger.Info("Paths updated successfully!")
	}
}
