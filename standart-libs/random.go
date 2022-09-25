package standart

import (
	"math/rand"

	lua "github.com/yuin/gopher-lua"
)

func randomString(l *lua.LState) int {
	args := l.ToString(1)
	n := l.ToInt(2)

	runes := []rune(args)

	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	l.Push(lua.LString(string(b)))

	return 1
}

func randomInt(l *lua.LState) int {
	min := l.ToInt(1)
	max := l.ToInt(2)

	l.Push(lua.LNumber(rand.Intn(max-min) + min))

	return 1
}

func InitRandomLIB(l *lua.LState) int {
	var exports = map[string]lua.LGFunction{ // Library Functions
		"randomString": randomString,
		"randomInt":    randomInt,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initialize functions

	l.Push(t) // Returning functions

	return 1
}
