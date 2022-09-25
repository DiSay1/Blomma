package standart

import (
	"fmt"
	"net/http"

	"github.com/DiSay1/Blomma/console"
	lua "github.com/yuin/gopher-lua"
)

type BlommaHTTPRequest struct { // The structure responsible for storing information about the request
	rw  http.ResponseWriter
	req *http.Request
}

var log = console.NewLogger("library") // Library logger

// Function to get QUERY parameters
func (b *BlommaHTTPRequest) getQuery(l *lua.LState) int {
	args := l.ToString(1)      // Get the passed argument(Desired value key)
	query := b.req.URL.Query() // QUERY parameters passed in HTTP request

	if _, ok := query[args]; ok { // Does the desired value exist so the key?
		l.Push(lua.LString(query.Get(args))) // If yes, return it.
		return 1
	}

	return 0 // Number of return values
}

// The function responsible for receiving information from POST forms
func (b *BlommaHTTPRequest) getFormData(l *lua.LState) int {
	if err := b.req.ParseForm(); err != nil { // Parsing the POST form
		log.Panic("An error occurred while trying to parse the form. Error:", err)
		return 0
	}

	obj := l.ToTable(1) // Get table from desired values as function argument

	t := l.NewTable() // Return table

	obj.ForEach(func(l1, l2 lua.LValue) { //Filling the table with the necessary data
		l.SetField(t, l2.String(), lua.LString(b.req.FormValue(l2.String())))
	})

	l.Push(t) // Returning a table with data

	return 1
}

// Function for getting HTTP headers.
// It is as simple as possible, I think
// it is not necessary to describe it.
func (b *BlommaHTTPRequest) getHeaders(l *lua.LState) int {
	args := l.ToString(1)

	l.Push(lua.LString(b.req.Header.Get(args)))

	return 1
}

// Function for sending HTTP responses
func (b *BlommaHTTPRequest) write(l *lua.LState) int {
	if _, err := fmt.Fprint(b.rw, l.ToString(1)); err != nil {
		log.Panic("No response was sent. Error:", err)
	}

	return 0
}

func (b *BlommaHTTPRequest) httpRedirect(l *lua.LState) int {
	redirectURL := l.ToString(1)

	http.Redirect(b.rw, b.req, redirectURL, http.StatusSeeOther)

	return 0
}

func NewHTTPRequest(l *lua.LState, rw http.ResponseWriter, req *http.Request) *lua.LTable {
	request := BlommaHTTPRequest{rw: rw, req: req} // Saving information about an HTTP request

	var exports = map[string]lua.LGFunction{ // Functions to call
		"write":       request.write,
		"getQuery":    request.getQuery,
		"getFormData": request.getFormData,
		"getHeader":   request.getHeaders,

		"redirect": request.httpRedirect,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initializing them

	l.SetField(t, "method", lua.LString(req.Method)) // Request method

	return t
}
