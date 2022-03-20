package router

import (
	"fmt"
	"log"
	"net/http"

	lua "github.com/yuin/gopher-lua"
)

func Router(rw http.ResponseWriter, req *http.Request) {
	L := lua.NewState()

	defer L.Close()
	if err := L.DoFile("./web/index.lua"); err != nil {
		panic(err)
	}

	values := L.GetGlobal("adress").String()

	if values == req.URL.Path {
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
	} else {
		fmt.Fprintln(rw, "Page not found!")
	}
}
