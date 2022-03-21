package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/DiSay1/Tentanto/states"
	lua "github.com/yuin/gopher-lua"
)

func Router(rw http.ResponseWriter, req *http.Request) {
	L := lua.NewState()

	defer L.Close()

	path := PathChecker(req.URL.Path)
	if path == "" {
		fmt.Fprintln(rw, "Page not found.")
		return
	}

	re := regexp.MustCompile(`.static/`)
	if re.FindString(path) == ".*static/" {
		return
	}

	if err := L.DoFile(path); err != nil {
		if states.DEBUG_MOD == true {
			fmt.Fprintln(rw, "Page loaded with an error.\n Err:", err)
		}

		return
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("Get"),
		NRet:    1,
		Protect: true,
	}, lua.LString("Good work!")); err != nil {
		log.Panicln("| Err call lua func\n Err:", err)
	}

	ret := L.Get(-1)
	fmt.Fprintln(rw, ret.String())
	L.Pop(1)
}

func PathChecker(url string) string {
	var path string

	re := regexp.MustCompile(`\.lua|\.html`)
	if re.FindString(url) == ".lua" || re.FindString(url) == ".html" {
		path = "./web" + url
	} else {
		path = "./static" + url
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Println("| URL:", path+" not found")
		return ""
	} else {
		return path
	}
}
