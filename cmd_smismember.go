package main

import "slices"

func SMIsMemberCommandCheck(args []Value) bool {
	return len(args) >= 2
}

func SMIsMemberCommand(args []Value) Value {
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	members := args[1:]
	var exists []Value
	for _, member := range members {
		if slices.Contains(SETs[key], member.str) {
			exists = append(exists, Value{typ: "integer", integer: 1})
		} else {
			exists = append(exists, Value{typ: "integer", integer: 0})
		}
	}
	return Value{typ: "array", array: exists}
}
