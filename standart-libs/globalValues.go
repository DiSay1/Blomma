package standart

import lua "github.com/yuin/gopher-lua"

var values = make(map[string]lua.LValue)

func InitGLLib(l *lua.LState) int {
	var exports = map[string]lua.LGFunction{
		"newValue":    newValue,
		"getValue":    getValue,
		"updateValue": updateValue,
		"removeValue": removeValue,
	}

	t := l.SetFuncs(l.NewTable(), exports)

	l.Push(t)

	return 1
}

func newValue(l *lua.LState) int {
	key := l.ToString(1)

	values[key] = l.CheckAny(2)

	return 0
}

func updateValue(l *lua.LState) int {
	key := l.ToString(1)

	values[key] = l.CheckAny(2)

	return 0
}

func getValue(l *lua.LState) int {
	key := l.ToString(1)

	l.Push(values[key])

	return 1
}

func removeValue(l *lua.LState) int {
	key := l.ToString(1)

	delete(values, key)

	return 0
}
