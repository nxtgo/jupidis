package main

func SetCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk
	value := args[1].bulk

	KEYsMu.Lock()
	KEYs[key] = StringValueType
	KEYsMu.Unlock()

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}
