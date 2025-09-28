package main

func HGetCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk
	field := args[1].bulk

	HSETsMu.RLock()
	value, ok := HSETs[key][field]
	HSETsMu.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}
