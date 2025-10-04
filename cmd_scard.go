package main

func SCardCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	return Value{typ: "integer", integer: len(SETs[key])}
}
