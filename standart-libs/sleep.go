package standart

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func timeSleep(l *lua.LState) int {
	second := l.ToInt(1)

	time.Sleep(time.Duration(second) * time.Second)

	return 0
}

func InitTIMELib(l *lua.LState) int {
	var exports = map[string]lua.LGFunction{ // Library Functions
		"sleep": newValue,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initialize functions

	l.Push(t) // Returning functions

	return 1
}
