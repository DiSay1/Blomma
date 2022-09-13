package server

import (
	"fmt"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

func addressHandler(rw http.ResponseWriter, req *http.Request) {
	for _, a := range Paths {
		if a.Address == req.URL.Path {
			if err := a.State.DoFile(a.Path); err != nil {
				log.Panic("Error compiling file. Error:", err)
			}

			if err := a.State.CallByParam(
				lua.P{
					Fn:      a.State.GetGlobal("Get"),
					NRet:    1,
					Protect: true,
				},
			); err != nil {
				log.Panic("Failed to get handler function. Error:", err)
				//
			}
			res := a.State.Get(-1)
			fmt.Fprint(rw, res.String())
		}
	}
}
