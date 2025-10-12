package main

import (
	"strconv"
)

func IncrByCommandCheck(args []Value) error {
	if len(args) != 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func IncrByCommand(args []Value) Value {
	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	key := args[0].str
	strIncrement := args[1].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	increment, err := strconv.Atoi(strIncrement)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	value, ok := VALUEs[key]
	if !ok {
		VALUEs[key] = strconv.Itoa(increment)
		return Value{typ: "integer", integer: increment}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue += increment
	VALUEs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: intValue}
}
