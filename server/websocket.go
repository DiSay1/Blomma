package server

import (
	"net/http"

	"github.com/DiSay1/Blomma/standart-libs"
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

var upgrader = websocket.Upgrader{} // use default options

type blommaWS struct {
	luaState *lua.LState
}

func websocketHandler(rw http.ResponseWriter, req *http.Request) {
	c, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Fatal("An error occurred while trying to create a WebSocket address. Error:", err)
		return
	}

	var ws blommaWS

	for _, a := range Paths {
		if req.URL.Path == a.Address && a.isWebSocket {
			if err := a.State.DoFile(a.Path); err != nil {
				log.Panic("File compilation error. Error:", err)
				return
			}

			ws = blommaWS{
				luaState: a.State,
			}
		}
	}

	c.SetCloseHandler(ws.closeHandler)

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				break
			}
			log.Panic("An error occurred while trying to read the WebSocket packet. Error:", err)
			break
		}

		if err := ws.luaState.CallByParam(
			lua.P{
				Fn:      ws.luaState.GetGlobal("onMessage"),
				NRet:    1,
				Protect: true,
			}, standart.NewWSMessage(ws.luaState, mt, message, c),
		); err != nil {
			log.Panic("The function cannot be executed. Error:", err)
			return
		}
	}
}

func (ws *blommaWS) closeHandler(code int, text string) error {
	if err := ws.luaState.CallByParam(lua.P{
		Fn:      ws.luaState.GetGlobal("onClose"),
		NRet:    1,
		Protect: true,
	}, standart.NewWSOnCloseMessage(ws.luaState, code, text)); err != nil {
		return err
	}

	return nil
}
