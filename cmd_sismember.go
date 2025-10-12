package main

import "slices"

func SIsMemberCommandCheck(args []Value) bool {
	return len(args) == 2
}

func SIsMemberCommand(args []Value) Value {
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	member := args[1].str

	if _, available := IsKeyAvailable(key, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	if slices.Contains(SETs[key], member) {
		return Value{typ: "integer", integer: 1}
	}
	return Value{typ: "integer", integer: 0}
}
