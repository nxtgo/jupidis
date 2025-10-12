package main

func SMembersCommandCheck(args []Value) bool {
	return len(args) == 1
}

func SMembersCommand(args []Value) Value {
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	if _, available := IsKeyAvailable(key, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	var value Value
	value.typ = "array"

	for _, member := range SETs[key] {
		value.array = append(value.array, Value{typ: "string", str: member})
	}

	return value
}
