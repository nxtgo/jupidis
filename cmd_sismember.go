package main

import "slices"

func SIsMemberCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	member := args[1].str

	if slices.Contains(SETs[key], member) {
		return Value{typ: "integer", integer: 1}
	}
	return Value{typ: "integer", integer: 0}
}
