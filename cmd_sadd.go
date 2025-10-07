package main

import "slices"

func SAddCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	key := args[0].str
	var members []string
	for _, arg := range args[1:] {
		members = append(members, arg.str)
	}

	if _, available := IsKeyAvailable(key, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
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
