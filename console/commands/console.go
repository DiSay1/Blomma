package commands

import (
	"fmt"
	"strings"

	"github.com/DiSay1/Blomma/console"
)

var logger = console.NewLogger("console")

type Command struct {
	Command      string
	RequiredArgs bool

	ArgsCount byte

	CalledFunction func(args ...string)
}

var Commands = []Command{
	{
		Command:      "set",
		RequiredArgs: true,

		ArgsCount: 2,

		CalledFunction: set,
	}, {
		Command:      "reload",
		RequiredArgs: true,

		ArgsCount: 1,

		CalledFunction: reload,
	},
}

func StartConsole() {
	logger.Info("Console loaded successfully")

	for {
		var cmd string

		if _, err := fmt.Scan(&cmd); err != nil {
			logger.Panic("An error occurred while trying to read user input. Error:", err)
			return
		}

		cmd = strings.ToLower(cmd)

		argsResult := make([]string, 0)

		for _, c := range Commands {
			if cmd == c.Command {
				if c.RequiredArgs {
					for i := byte(0); i < c.ArgsCount; i++ {
						var args string

						logger.Info(fmt.Sprintf("Enter %v argument: ", i))
						if _, err := fmt.Scan(&args); err != nil {
							logger.Panic("An error occurred while trying to read user input. Error:", err)
							return
						}

						argsResult = append(argsResult, args)
					}

					c.CalledFunction(argsResult...)
				} else {
					c.CalledFunction()
				}
			}
		}
	}
}
