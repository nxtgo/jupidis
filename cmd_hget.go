package main

func HGetCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	HSETsMu.RLock()
	defer HSETsMu.RUnlock()

	key := args[0].str
	field := args[1].str

	if _, available := IsKeyAvailable(key, "hash"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := HSETs[key][field]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}
