package standart

import (
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

type BlommaWSMessage struct { // Structure for storing information about a websocket message
	Conn *websocket.Conn
}

// Function for sending websocket messages
func (wsMessage *BlommaWSMessage) write(l *lua.LState) int {
	mt := l.ToInt(1)         // Getting message type from arguments
	message := l.ToString(2) // Getting message type from arguments

	if err := wsMessage.Conn.WriteMessage(mt, []byte(message)); err != nil { // send messages
		log.Panic("An error occurred while trying to send a WebSocket packet. Error:", err)
		return 0 // Number of return values
	}

	return 0
}

func NewWSMessage(l *lua.LState, mt int, data []byte, c *websocket.Conn) *lua.LTable {
	request := BlommaWSMessage{Conn: c} // WebSocket message creation function

	var exports = map[string]lua.LGFunction{ //
		"write": request.write,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initializing functions

	// Saving message information
	l.SetField(t, "mt", lua.LNumber(mt))
	l.SetField(t, "data", lua.LString(data))

	return t
}

// To keep information about closing a web socket connection
func NewWSOnCloseMessage(l *lua.LState, mt int, text string) *lua.LTable {
	t := l.NewTable()

	l.SetField(t, "mt", lua.LNumber(mt))
	l.SetField(t, "data", lua.LString(text))

	return t
}
