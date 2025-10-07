package main

import "slices"

func SRemCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	key := args[0].str
	members := args[1:]

	if _, ok := SETs[key]; !ok {
		return Value{typ: "integer", integer: 0}
	}

	var count int
	for _, member := range members {
		if slices.Contains(SETs[key], member.str) {
			index := slices.Index(SETs[key], member.str)
			SETs[key] = slices.Delete(SETs[key], index, index+1)
			count++
		}
	}

	if len(SETs[key]) == 0 {
		delete(SETs, key)
	}

	return Value{typ: "integer", integer: count}
}
