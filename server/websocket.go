package server

import (
	"net/http"

	"github.com/DiSay1/Blomma/server/states"
	"github.com/DiSay1/Blomma/standart-libs"
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

var upgrader = websocket.Upgrader{} // use default options

type blommaWS struct {
	luaState lua.LState
}

func (h *Handler) websocketHandler(rw http.ResponseWriter, req *http.Request) {
	c, err := upgrader.Upgrade(rw, req, nil) // Updating to a WebSocket Connection
	if err != nil {
		log.Fatal("An error occurred while trying to create a WebSocket address. Error:", err)
		return
	}

	if states.DEV_MODE { // If DEV_MODE is enabled
		if err := h.State.DoFile(h.Path); err != nil { // Interpreting the handler file again
			log.Panic("File compilation error. Error:", err)
			return // If an error occurs, exit
		}
	}

	ws := blommaWS{ // Save the handler
		luaState: *h.State,
	}

	c.SetCloseHandler(ws.closeHandler) // Connection close handler

	if err := ws.luaState.CallByParam( // Calling the message handler
		lua.P{
			Fn:      ws.luaState.GetGlobal("WSHandler"),
			NRet:    1,
			Protect: true,
		}, standart.NewDataForWSHandler(&ws.luaState, c), // Transferring connection information
	); err != nil {
		log.Panic("The function cannot be executed. Error:", err)
		return
	}
}

// processing function closed
func (ws *blommaWS) closeHandler(code int, text string) error {
	if err := ws.luaState.CallByParam(lua.P{ // We call the connection closing handler function
		Fn:      ws.luaState.GetGlobal("onClose"),
		NRet:    1,
		Protect: true,
	}, standart.NewWSOnCloseMessage(&ws.luaState, code, text), // We call the connection closing handler function
	); err != nil {
		return err
	}

	return nil
}
