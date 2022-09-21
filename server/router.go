package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DiSay1/Blomma/server/states"
	"github.com/DiSay1/Blomma/standart-libs"
	lua "github.com/yuin/gopher-lua"
)

func (h *Handler) indexRouter(rw http.ResponseWriter, req *http.Request) {
	_, err := os.Stat("./static" + req.URL.Path)
	if err != nil {
		h.addressHandler(rw, req)
		return
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

func (h *Handler) addressHandler(rw http.ResponseWriter, req *http.Request) {
	if h.Type == "lua" {
		if states.DEV_MODE {
			if err := h.State.DoFile(h.Path); err != nil {
				log.Panic("File compilation error. Error:", err)
				return
			}
		}

		if err := h.State.CallByParam(
			lua.P{
				Fn:      h.State.GetGlobal("Handler"),
				NRet:    1,
				Protect: true,
			}, standart.NewHTTPRequest(h.State, rw, req),
		); err != nil {
			log.Panic("The function cannot be executed. Error:", err)
			return
		}
		return
	} else if h.Type == "html" {
		data, err := os.ReadFile(h.Path)
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
