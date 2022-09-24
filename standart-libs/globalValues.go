package standart

import lua "github.com/yuin/gopher-lua"

/*
	This is a library for storing global values.
	Used to pass information between handlers.
	More can be found here: https://github.com/DiSay1/Blomma/wiki/Standart-Library#valuecontroller
*/

var values = make(map[string]lua.LValue) // For saved values

// The function responsible for loading the library
func InitGLLib(l *lua.LState) int {
	var exports = map[string]lua.LGFunction{ // Library Functions
		"newValue":    newValue,
		"getValue":    getValue,
		"updateValue": updateValue,
		"removeValue": removeValue,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initialize functions

	l.Push(t) // Returning functions

	return 1
}

// Save value function
func newValue(l *lua.LState) int {
	key := l.ToString(1) // Save value function

	values[key] = l.CheckAny(2) // Store values in a given key

	return 0 // How many values does the function return
}

// Value updates
func updateValue(l *lua.LState) int {
	key := l.ToString(1)

	values[key] = l.CheckAny(2)

	return 0
}

// Get value
func getValue(l *lua.LState) int {
	key := l.ToString(1)

	l.Push(values[key])

	return 1
}

// Removing a value
func removeValue(l *lua.LState) int {
	key := l.ToString(1)

	delete(values, key)

	return 0
}
