package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DiSay1/Blomma/server/states"
	"github.com/DiSay1/Blomma/standart-libs"
	lua "github.com/yuin/gopher-lua"
)

// Router for processing the "main" (/) page
func (h *Handler) indexRouter(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" { // Checking the path
		_, err := os.Stat("./static" + req.URL.Path) // Trying to find a file in this path
		if err != nil {                              // If the file is not found
			return
		}

		data, err := os.ReadFile("./static" + req.URL.Path) // If we find the file, we try to read it
		if err != nil {
			log.Panic("An error occurred while trying to get a file to upload. Error:", err)
			return
		}

		_, err = rw.Write(data) // and send the file
		if err != nil {
			log.Panic("An error occurred while trying to send the file:", err)
			return
		}
		return // We leave
	}

	h.addressHandler(rw, req) // If nothing happened, call the request processing function
}

// We process requests
func (h *Handler) addressHandler(rw http.ResponseWriter, req *http.Request) {
	switch h.Type { // Handler type
	case "lua": // If the function is responsible for the lua handler
		if states.DEV_MODE { // If DEV_MODE is enabled
			if err := h.State.DoFile(h.Path); err != nil { // Interpreting the handler file again
				log.Panic("File compilation error. Error:", err)

				if _, err := fmt.Fprintf(rw, "File compilation error. Error:\n%v", err); err != nil {
					log.Fatal("An error occurred while trying to submit error information. Error:")
				}
				return // If an error occurs, exit
			}
		}

		lState := *h.State

		defer lState.Close()

		if err := lState.CallByParam( // We call the function responsible for processing requests
			lua.P{
				Fn:      h.State.GetGlobal("Handler"),
				NRet:    1,
				Protect: true,
			}, standart.NewHTTPRequest(h.State, rw, req), // Passing information about the request as an argument
		); err != nil {
			log.Panic("The function cannot be executed. Error:", err)

			if states.DEV_MODE {
				if _, err := fmt.Fprintf(rw, "The function cannot be executed. Error:\n%v", err); err != nil {
					log.Fatal("An error occurred while trying to submit error information. Error:")
				}
			}
			return
		}
		return
	case "html": // If the function is responsible for the html handler
		data, err := os.ReadFile(h.Path) // Reading an HTML file
		if err != nil {
			log.Panic("An error occurred while trying to read the HTML file. Error:", err)

			if states.DEV_MODE {
				if _, err := fmt.Fprintf(rw, "An error occurred while trying to read the HTML file. Error:\n%v", err); err != nil {
					log.Fatal("An error occurred while trying to submit error information. Error:")
				}
			}
			return
		}

		_, err = fmt.Fprint(rw, string(data)) // Sending an HTML file as a string
		if err != nil {
			log.Panic("An error occurred while trying to send the HTML file. Error:", err)
		}
		return
	}
}
