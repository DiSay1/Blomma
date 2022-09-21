package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DiSay1/Blomma/server/states"
	"github.com/DiSay1/Blomma/standart-libs"
	lua "github.com/yuin/gopher-lua"
)

func addressHandler(rw http.ResponseWriter, req *http.Request) {
	for _, handler := range Handlers {
		if handler.Address == req.URL.Path {
			if handler.Type == "lua" {
				if states.DEV_MODE {
					if err := handler.State.DoFile(handler.Path); err != nil {
						log.Panic("File compilation error. Error:", err)
						return
					}
				}

				if err := handler.State.CallByParam(
					lua.P{
						Fn:      handler.State.GetGlobal("Handler"),
						NRet:    1,
						Protect: true,
					}, standart.NewHTTPRequest(handler.State, rw, req),
				); err != nil {
					log.Panic("The function cannot be executed. Error:", err)
					return
				}
				return
			} else if handler.Type == "html" {
				data, err := os.ReadFile(handler.Path)
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
