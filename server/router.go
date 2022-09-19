package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DiSay1/Blomma/standart-libs"
	lua "github.com/yuin/gopher-lua"
)

func addressHandler(rw http.ResponseWriter, req *http.Request) {
	for _, a := range Paths {
		if a.Address == req.URL.Path {
			if a.Type == "lua" {
				if err := a.State.DoFile(a.Path); err != nil {
					log.Panic("File compilation error. Error:", err)
					return
				}

				if err := a.State.CallByParam(
					lua.P{
						Fn:      a.State.GetGlobal("Handler"),
						NRet:    1,
						Protect: true,
					}, standart.NewHTTPRequest(a.State, rw, req),
				); err != nil {
					log.Panic("The function cannot be executed. Error:", err)
					return
				}
				return
			} else if a.Type == "html" {
				data, err := os.ReadFile(a.Path)
				if err != nil {
					log.Panic("Err:", err)
					return
				}

				_, err = fmt.Fprint(rw, string(data))
				if err != nil {
					log.Panic("Err", err)
				}
				return
			}
		}
	}

	data, err := os.ReadFile("./static" + req.URL.Path)
	if err != nil {
		log.Panic("Err:", err)
		return
	}

	_, err = fmt.Fprint(rw, string(data))
	if err != nil {
		log.Panic("Err", err)
		return
	}
}
