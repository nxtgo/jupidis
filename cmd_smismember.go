package main

import (
	"slices"
)

func SMIsMemberCommandCheck(args []Value) error {
	if len(args) < 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func SMIsMemberCommand(args []Value) Value {
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	members := args[1:]
	if _, available := IsKeyAvailable(key, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

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
