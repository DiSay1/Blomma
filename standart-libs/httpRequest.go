package standart

import (
	"fmt"
	"net/http"

	"github.com/DiSay1/Blomma/console"
	lua "github.com/yuin/gopher-lua"
)

type BlommaHTTPRequest struct {
	rw  http.ResponseWriter
	req *http.Request
}

var log = console.NewLogger("library")

func (b *BlommaHTTPRequest) getQuery(l *lua.LState) int {
	args := l.ToString(1)
	query := b.req.URL.Query()

	if _, ok := query[args]; ok {
		l.Push(lua.LString(query.Get(args)))
	} else {
		l.Push(lua.LString(""))
	}

	return 1
}

func (b *BlommaHTTPRequest) getFormData(l *lua.LState) int {
	if err := b.req.ParseForm(); err != nil {
		log.Panic("An error occurred while trying to parse the form. Error:", err)
		return 0
	}

	obj := l.ToTable(1)

	t := l.NewTable()

	obj.ForEach(func(l1, l2 lua.LValue) {
		l.SetField(t, l2.String(), lua.LString(b.req.FormValue(l2.String())))
	})

	l.Push(t)

	return 1
}

func (b *BlommaHTTPRequest) getHeaders(l *lua.LState) int {
	args := l.ToString(1)

	l.Push(lua.LString(b.req.Header.Get(args)))

	return 1
}

func (b *BlommaHTTPRequest) write(l *lua.LState) int {
	if _, err := fmt.Fprint(b.rw, l.ToString(1)); err != nil {
		log.Panic("No response was sent. Error:", err)
	}
	return 0
}

func NewHTTPRequest(l *lua.LState, rw http.ResponseWriter, req *http.Request) *lua.LTable {
	request := BlommaHTTPRequest{rw: rw, req: req}

	var exports = map[string]lua.LGFunction{
		"write":       request.write,
		"getQuery":    request.getQuery,
		"getFormData": request.getFormData,
		"getHeader":   request.getHeaders,
	}

	t := l.SetFuncs(l.NewTable(), exports)

	l.SetField(t, "method", lua.LString(req.Method))

	return t
}
