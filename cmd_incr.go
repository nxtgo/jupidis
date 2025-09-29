package main

import "strconv"

func IncrCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk

	KEYsMu.Lock()
	defer KEYsMu.Unlock()
	SETsMu.Lock()
	defer SETsMu.Unlock()

	value, ok := SETs[key]
	if !ok {
		SETs[key] = "1"
		KEYs[key] = StringValueType
		return Value{typ: "integer", integer: 1}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue++
	SETs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: intValue}
}
