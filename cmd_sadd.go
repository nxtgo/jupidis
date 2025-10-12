package main

import "slices"

func SAddCommandCheck(args []Value) bool {
	return len(args) >= 2
}

func SAddCommand(args []Value) Value {
	SETsMu.Lock()
	defer SETsMu.Unlock()

	key := args[0].str
	if _, available := IsKeyAvailable(key, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	var members = make([]string, 0, len(args)-1)
	for _, arg := range args[1:] {
		members = append(members, arg.str)
	}

	var count int
	for _, member := range members {
		if !slices.Contains(SETs[key], member) {
			SETs[key] = append(SETs[key], member)
			count++
		}
	}

	return Value{typ: "integer", integer: count}
}
