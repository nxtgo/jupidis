package main

func SetCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	key := args[0].str
	value := args[1].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	VALUEs[key] = value

	return Value{typ: "string", str: "OK"}
}
