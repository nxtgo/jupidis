package main

import (
	"fmt"
	"strconv"
)

func DecrByCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk
	strDecrement := args[1].bulk

	decrement, err := strconv.Atoi(strDecrement)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	KEYsMu.Lock()
	defer KEYsMu.Unlock()
	SETsMu.Lock()
	defer SETsMu.Unlock()

	value, ok := SETs[key]
	if !ok {
		SETs[key] = fmt.Sprintf("%d", -decrement)
		KEYs[key] = StringValueType
		return Value{typ: "integer", integer: int64(-decrement)}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue -= decrement
	SETs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: int64(intValue)}
}
