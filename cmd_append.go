package main

func AppendCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk
	appendValue := args[1].bulk

	SETsMu.Lock()
	defer SETsMu.Unlock()

	value, ok := SETs[key]
	if !ok {
		SETs[key] = appendValue
		return Value{typ: "integer", integer: len(appendValue)}
	}

	value += appendValue
	SETs[key] = value
	return Value{typ: "integer", integer: len(value)}
}
