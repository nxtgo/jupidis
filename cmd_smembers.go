package main

func SMembersCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str

	var value Value
	value.typ = "array"

	for _, member := range SETs[key] {
		value.array = append(value.array, Value{typ: "string", str: member})
	}

	return value
}
