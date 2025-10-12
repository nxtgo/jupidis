package main

import (
	"strconv"
)

func DecrCommandCheck(args []Value) error {
	if len(args) != 1 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func DecrCommand(args []Value) Value {
	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := VALUEs[key]
	if !ok {
		VALUEs[key] = "-1"
		return Value{typ: "integer", integer: -1}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue--
	VALUEs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: intValue}
}
