package main

func SMembersCommandCheck(args []Value) bool {
	return len(args) == 1
}

func SMembersCommand(args []Value) Value {
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
