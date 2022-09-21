package server

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type Handler struct {
	Address string
	Path    string

	isWebSocket bool
	Type        string

	State *lua.LState
}

var Handlers []*Handler

func LoadPaths() error {
	err := filepath.Walk("./web", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		_, file := filepath.Split(path)

		pathElement := strings.Split(file, ".")

		path = strings.ReplaceAll(path, `\`, `/`)

		if len(pathElement) > 1 {
			switch pathElement[1] {
			case "lua":
				l := lua.NewState()

				if err := l.DoFile("./" + strings.ReplaceAll(path, `\`, `/`)); err != nil {
					return err
				}

				options := l.GetGlobal("options")
				address := l.GetField(options, "Address")

				if address.Type() != lua.LTString && address.Type() != lua.LTNil {
					return fmt.Errorf("in file %v address is not a string", path)
				}

				isWebSocket := false
				websocket := l.GetField(options, "WebSocket")
				if websocket.Type() == lua.LTBool {
					if websocket.String() == "true" {
						isWebSocket = true
					}
				}

				webPath := strings.ReplaceAll(path, "web/", "/")

				if address.Type() == lua.LTString {
					Handlers = append(Handlers, &Handler{
						Address: address.String(),
						Path:    path,

						isWebSocket: isWebSocket,
						Type:        "lua",

						State: l,
					})
				} else {
					Handlers = append(Handlers, &Handler{
						Address: webPath,
						Path:    path,

						isWebSocket: isWebSocket,
						Type:        "lua",

						State: l,
					})
				}
			case "html":
				webPath := strings.ReplaceAll(path, "web/", "/")
				Handlers = append(Handlers, &Handler{
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

	sort.SliceStable(Handlers, func(i, j int) bool {
		return Handlers[i].Address < Handlers[j].Address
	})

	return err
}
