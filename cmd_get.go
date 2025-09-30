package main

func GetCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].bulk

	if !IsKeyAvailable(key, "string") {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := SETs[key]
	if !ok {
		return Value{typ: "null", str: ""}
	}

	return Value{typ: "string", str: value}
}
