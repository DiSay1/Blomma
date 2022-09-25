package standart

import (
	"encoding/json"
	"errors"

	lua "github.com/yuin/gopher-lua"
)

// I used the source project
// https://github.com/layeh/gopher-json
// to write this module.

func apiJSONDecode(l *lua.LState) int {
	str := l.ToString(1)

	value, err := decodeJSON(l, []byte(str))
	if err != nil {
		log.Panic("An error occurred while decoding JSON. Error:", err)
		return 0
	}

	l.Push(value)
	return 1
}

func apiJSONEncode(L *lua.LState) int {
	value := L.CheckAny(1)

	data, err := encodeJSON(value)
	if err != nil {
		log.Info("An error occurred while encoding JSON. Error:", err)
		return 0
	}

	L.Push(lua.LString(string(data)))
	return 1
}

func InitJSONLib(l *lua.LState) int {
	var exports = map[string]lua.LGFunction{ // Library Functions
		"Decode": apiJSONDecode,
		"Encode": apiJSONEncode,
	}

	t := l.SetFuncs(l.NewTable(), exports) // Initialize functions

	l.Push(t) // Returning functions

	return 1
}

func encodeJSON(value lua.LValue) ([]byte, error) {
	return json.Marshal(jsonValue{
		LValue:  value,
		visited: make(map[*lua.LTable]bool),
	})
}

type jsonValue struct {
	lua.LValue
	visited map[*lua.LTable]bool
}

var (
	errNested      = errors.New("cannot encode recursively nested tables to JSON")
	errSparseArray = errors.New("cannot encode sparse array")
	errInvalidKeys = errors.New("cannot encode mixed or invalid key types")
)

type invalidTypeError lua.LValueType

func (i invalidTypeError) Error() string {
	return `cannot encode ` + lua.LValueType(i).String() + ` to JSON`
}

func (j jsonValue) MarshalJSON() (data []byte, err error) {
	switch converted := j.LValue.(type) {
	case lua.LBool:
		data, err = json.Marshal(bool(converted))
	case lua.LNumber:
		data, err = json.Marshal(float64(converted))
	case *lua.LNilType:
		data = []byte(`null`)
	case lua.LString:
		data, err = json.Marshal(string(converted))
	case *lua.LTable:
		if j.visited[converted] {
			return nil, errNested
		}
		j.visited[converted] = true

		key, value := converted.Next(lua.LNil)

		switch key.Type() {
		case lua.LTNil: // empty table
			data = []byte(`[]`)
		case lua.LTNumber:
			arr := make([]jsonValue, 0, converted.Len())
			expectedKey := lua.LNumber(1)
			for key != lua.LNil {
				if key.Type() != lua.LTNumber {
					err = errInvalidKeys
					return
				}
				if expectedKey != key {
					err = errSparseArray
					return
				}
				arr = append(arr, jsonValue{value, j.visited})
				expectedKey++
				key, value = converted.Next(key)
			}
			data, err = json.Marshal(arr)
		case lua.LTString:
			obj := make(map[string]jsonValue)
			for key != lua.LNil {
				if key.Type() != lua.LTString {
					err = errInvalidKeys
					return
				}
				obj[key.String()] = jsonValue{value, j.visited}
				key, value = converted.Next(key)
			}
			data, err = json.Marshal(obj)
		default:
			err = errInvalidKeys
		}
	default:
		err = invalidTypeError(j.LValue.Type())
	}
	return
}

func decodeJSON(L *lua.LState, data []byte) (lua.LValue, error) {
	var value interface{}

	err := json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	return decodeJSONValue(L, value), nil
}

func decodeJSONValue(L *lua.LState, value interface{}) lua.LValue {
	switch converted := value.(type) {
	case bool:
		return lua.LBool(converted)
	case float64:
		return lua.LNumber(converted)
	case string:
		return lua.LString(converted)
	case json.Number:
		return lua.LString(converted)
	case []interface{}:
		arr := L.CreateTable(len(converted), 0)
		for _, item := range converted {
			arr.Append(decodeJSONValue(L, item))
		}
		return arr
	case map[string]interface{}:
		tbl := L.CreateTable(0, len(converted))
		for key, item := range converted {
			tbl.RawSetH(lua.LString(key), decodeJSONValue(L, item))
		}
		return tbl
	case nil:
		return lua.LNil
	}

	return lua.LNil
}
