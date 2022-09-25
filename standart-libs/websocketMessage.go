package standart

import (
	"github.com/gorilla/websocket"
	lua "github.com/yuin/gopher-lua"
)

type BlommaWSMessage struct { // Structure for storing information about a websocket message
	*websocket.Conn
}

// Function for sending websocket messages
func (ws *BlommaWSMessage) write(l *lua.LState) int {
	mt := l.ToInt(1)         // Getting message type from arguments
	message := l.ToString(2) // Getting message from arguments

	if err := ws.WriteMessage(mt, []byte(message)); err != nil { // send messages
		log.Panic("An error occurred while trying to send a WebSocket packet. Error:", err)
		return 0 // Number of return values
	}

	return 0
}

func (ws *BlommaWSMessage) read(l *lua.LState) int {
	mt, data, err := ws.ReadMessage() // Reading messages
	if err != nil {
		if !websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) { // Error due to closed message?
			log.Panic("An error occurred while trying to read the WebSocket packet. Error:", err) // If not, show an error
			return 0
		}
	}

	t := l.NewTable()

	// Saving message information
	l.SetField(t, "mt", lua.LNumber(mt))
	l.SetField(t, "data", lua.LString(data))

	l.Push(t)

	return 1
}

func (ws *BlommaWSMessage) close(l *lua.LState) int {
	if err := ws.Close(); err != nil {
		log.Panic("The connection was closed with an error. Error:", err)
	}
	return 0
}

func NewDataForWSHandler(l *lua.LState, c *websocket.Conn) *lua.LTable {
	request := BlommaWSMessage{Conn: c} // WebSocket message creation function

	var exports = map[string]lua.LGFunction{ //
		"write": request.write,
		"read":  request.read,
		"close": request.close,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initializing functions

	return t
}

// To keep information about closing a web socket connection
func NewWSOnCloseMessage(l *lua.LState, mt int, text string) *lua.LTable {
	t := l.NewTable()

	l.SetField(t, "mt", lua.LNumber(mt))
	l.SetField(t, "data", lua.LString(text))

	return t
}
