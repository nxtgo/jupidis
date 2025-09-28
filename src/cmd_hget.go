package main

func HGetCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk
	field := args[1].bulk

	HSETsMu.Lock()
	defer HSETsMu.Unlock()

	hget, ok := HSETs[key]
	if !ok {
		return Value{typ: "null"}
	}

	value, ok := hget[field]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}
