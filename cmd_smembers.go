package main

func SMembersCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].bulk
	var value Value
	value.typ = "array"
	members, exists := SETs[key]
	if !exists {
		return value
	}

	for member := range members {
		value.array = append(value.array, Value{typ: "string", str: member})
	}
	return value
}
