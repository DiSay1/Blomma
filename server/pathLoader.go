package server

import (
	"fmt"
	"io/fs"
	"os"
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

// File example
var indexLua = `
-- Handler Options
options = {
    Address = "/", -- Web path to handler
}

-- Function called on request
function Handler (request)
    -- Outputting the values of the desired variable
    request.write("Hello world!")
end
`

// Function of loading and determining WEB paths
func LoadPaths() error {
	// Checking for the presence of the ./web folder
	if _, err := os.Stat("./web/"); err != nil {
		log.Panic("The ./web folder was not found. I create a new")

		// Folder creation
		if err := os.Mkdir("./web", os.ModeDir); err != nil {
			return fmt.Errorf("the ./web folder was not created. %v", err)
		}

		// Creating an index.lua file
		file, err := os.Create("./web/index.lua")
		if err != nil {
			return fmt.Errorf("the index.lua file was not created. %v", err)
		}

		// Writing an example to a file
		_, err = file.Write([]byte(indexLua))
		if err != nil {
			return fmt.Errorf("the index.lua file was not created. %v", err)
		}
	}

	// Checking for the presence of the ./satic folder
	if _, err := os.Stat("./static/"); err != nil {
		log.Panic("The ./static folder was not found. I create a new")

		// Folder creation
		if err := os.Mkdir("./static", os.ModeDir); err != nil {
			return fmt.Errorf("the ./static folder was not created. %v", err)
		}
	}

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
				l.PreloadModule("json", standart.InitJSONLib)

				// We load into the virtual machine to get the handler settings
				if err := l.DoFile("./" + path); err != nil {
					return err
				}

				// We get the handler settings.
				options := l.GetGlobal("options")
				var luaAddress lua.LValue
				if options.Type() != lua.LTNil { // If the settings exist, find out the address of the handler
					luaAddress = l.GetField(options, "Address")
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

				resAddress := ""

				if luaAddress == nil || luaAddress.Type() == lua.LTNil { // Check if the path was specified by the developer
					resAddress = strings.ReplaceAll(path, "web/", "/")
				} else {
					resAddress = luaAddress.String()
				}

				Handlers = append(Handlers, &Handler{ // Check if the path was specified by the developer
					Address: resAddress,
					Path:    path,

					isWebSocket: isWebSocket,
					Type:        "lua",

					State: l,
				})
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
