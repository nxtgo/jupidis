package main

func SetCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	key := args[0].bulk
	value := args[1].bulk

	SETs[key] = value

	return Value{typ: "string", str: "OK"}
}
