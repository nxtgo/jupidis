package main

func AppendCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	key := args[0].bulk
	appendValue := args[1].bulk

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := VALUEs[key]
	if !ok {
		VALUEs[key] = appendValue
		return Value{typ: "integer", integer: len(appendValue)}
	}

	value += appendValue
	VALUEs[key] = value
	return Value{typ: "integer", integer: len(value)}
}
