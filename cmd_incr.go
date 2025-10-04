package main

import "strconv"

func IncrCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].str

	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := VALUEs[key]
	if !ok {
		VALUEs[key] = "1"
		return Value{typ: "integer", integer: 1}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue++
	VALUEs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: intValue}
}
