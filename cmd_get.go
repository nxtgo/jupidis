package main

func GetCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	VALUEsMu.RLock()
	defer VALUEsMu.RUnlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := VALUEs[key]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}
