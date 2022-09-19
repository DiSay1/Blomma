package server

import (
	"net/http"

	"github.com/DiSay1/Blomma/standart-libs"
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

var upgrader = websocket.Upgrader{} // use default options

func websocketHandler(rw http.ResponseWriter, req *http.Request) {
	c, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Fatal("An error occurred while trying to create a WebSocket address. Error:", err)
		return
	}

	var luaState *lua.LState

	for _, a := range Paths {
		if req.URL.Path == a.Address && a.isWebSocket {
			if err := a.State.DoFile(a.Path); err != nil {
				log.Panic("File compilation error. Error:", err)
				return
			}

			luaState = a.State
		}
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if mt == -1 {
			break
		}
		if err != nil {
			log.Panic("An error occurred while trying to read the WebSocket packet. Error:", err)
			break
		}

		if err := luaState.CallByParam(
			lua.P{
				Fn:      luaState.GetGlobal("Handler"),
				NRet:    1,
				Protect: true,
			}, standart.NewWebSocketMessage(luaState, mt, message, c),
		); err != nil {
			log.Panic("The function cannot be executed. Error:", err)
			return
		}
	}
}
