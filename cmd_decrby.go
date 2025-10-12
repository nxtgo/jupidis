package main

import (
	"fmt"
	"strconv"
)

func DecrByCommandCheck(args []Value) error {
	if len(args) != 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func DecrByCommand(args []Value) Value {
	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	key := args[0].str
	strDecrement := args[1].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	decrement, err := strconv.Atoi(strDecrement)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	value, ok := VALUEs[key]
	if !ok {
		VALUEs[key] = fmt.Sprintf("%d", -decrement)
		return Value{typ: "integer", integer: -decrement}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue -= decrement
	VALUEs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: intValue}
}
