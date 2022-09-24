package server

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/DiSay1/Blomma/standart-libs"
	lua "github.com/yuin/gopher-lua"
)

// A structure for storing information about handlers.
type Handler struct {
	Address string // Handler address.
	Path    string // Path to handler

	isWebSocket bool   // Is it a websocket? To define a websocket handler
	Type        string // Handler type. HTML or Lua

	State *lua.LState // Lua machine states
}

var Handlers []*Handler // Handlers

// Function of loading and determining WEB paths
func LoadPaths() error {
	// We pass the folder ./web
	err := filepath.Walk("./web", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// To check file extension
		_, file := filepath.Split(path)
		pathElement := strings.Split(file, ".")

		path = strings.ReplaceAll(path, `\`, `/`) // We make the path clear and convenient for everyone who needs it

		if len(pathElement) > 1 { // Is this really a file?
			switch pathElement[1] {
			case "lua":
				l := lua.NewState() // Create a new Lua virtual machine state

				// Loading libraries
				l.PreloadModule("valueController", standart.InitGLLib)

				// We load into the virtual machine to get the handler settings
				if err := l.DoFile("./" + path); err != nil {
					return err
				}

				// We get the handler settings.
				options := l.GetGlobal("options")
				var address lua.LValue
				if options.Type() != lua.LTNil { // If the settings exist, find out the address of the handler
					address = l.GetField(options, "Address")
				}

				isWebSocket := false

				if options.Type() != lua.LTNil { // If the settings exist, find out if it's a WebSocket
					websocket := l.GetField(options, "WebSocket")
					if websocket.Type() == lua.LTBool {
						if websocket.String() == "true" {
							isWebSocket = true
						}
					}
				}

				// We remember the web path, in case it is not specified
				webPath := strings.ReplaceAll(path, "web/", "/")

				if address.Type() == lua.LTString { // Web path specified?
					Handlers = append(Handlers, &Handler{ // Store the handler information with the specified web developer path
						Address: address.String(),
						Path:    path,

						isWebSocket: isWebSocket,
						Type:        "lua",

						State: l,
					})
				} else { // If not specified
					Handlers = append(Handlers, &Handler{ // We save information about the handler with an automatically specified web path
						Address: webPath,
						Path:    path,

						isWebSocket: isWebSocket,
						Type:        "lua",

						State: l,
					})
				}
			case "html":
				webPath := strings.ReplaceAll(path, "web/", "/") // Specify the web path
				Handlers = append(Handlers, &Handler{            // Remembering information about the handler
					Address: webPath,
					Path:    path,

					isWebSocket: false,
					Type:        "html",

					State: nil,
				})
			}
		}

		return err
	})

	sort.SliceStable(Handlers, func(i, j int) bool { // Remembering information about the handler
		return Handlers[i].Address < Handlers[j].Address
	})

	return err
}
