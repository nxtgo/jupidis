package main

func GetCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	value, ok := SETs[key]
	if !ok {
		return Value{typ: "null", str: ""}
	}

	return Value{typ: "string", str: value}
}
