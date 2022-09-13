package server

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

type Address struct {
	Address string
	Path    string

	State *lua.LState
}

var Paths []*Address

func loadPaths() error {
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

				address := l.GetGlobal("address")

				if address.Type() != lua.LTString && address.Type() != lua.LTNil {
					return fmt.Errorf("in file %v address is not a string", path)
				}

				webPath := strings.ReplaceAll(path, "web/", "/")

				if address.Type() == lua.LTString {
					Paths = append(Paths, &Address{
						Address: address.String(),
						Path:    path,

						State: l,
					})
				} else {
					Paths = append(Paths, &Address{
						Address: webPath,
						Path:    path,

						State: l,
					})
				}
			case "html":
				webPath := strings.ReplaceAll(path, "web/", "/")
				Paths = append(Paths, &Address{
					Address: webPath,
					Path:    path,

					State: nil,
				})
			}
		}

		sort.SliceStable(Paths, func(i, j int) bool {
			return Paths[i].Address < Paths[j].Address
		})

		return err
	})

	return err
}
