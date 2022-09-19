package standart

import (
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

type BlommaWSMessage struct {
	Conn *websocket.Conn
}

func (wsMessage *BlommaWSMessage) write(l *lua.LState) int {
	mt := l.ToInt(1)
	message := l.ToString(2)

	if err := wsMessage.Conn.WriteMessage(mt, []byte(message)); err != nil {
		log.Panic("An error occurred while trying to send a WebSocket packet. Error:", err)
		return 0
	}

	return 0
}

func NewWebSocketMessage(l *lua.LState, mt int, data []byte, c *websocket.Conn) *lua.LTable {
	request := BlommaWSMessage{Conn: c}

	var exports = map[string]lua.LGFunction{
		"write": request.write,
	}

	t := l.SetFuncs(l.NewTable(), exports)

	l.SetField(t, "mt", lua.LNumber(mt))
	l.SetField(t, "data", lua.LString(data))

	return t
}
