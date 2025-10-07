package main

import "slices"

func SMoveCommand(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	srcKey := args[0].str
	destKey := args[1].str
	member := args[2].str

	SETsMu.Lock()
	defer SETsMu.Unlock()

	if !slices.Contains(SETs[srcKey], member) {
		return Value{typ: "integer", integer: 0}
	}

	indexToRemove := slices.IndexFunc(SETs[srcKey], func(s string) bool {
		return s == member
	})
	SETs[srcKey] = append(SETs[srcKey][:indexToRemove], SETs[srcKey][indexToRemove+1:]...)

	if len(SETs[srcKey]) == 0 {
		delete(SETs, srcKey)
	}

	if SETs[destKey] == nil {
		SETs[destKey] = []string{}
	}
	SETs[destKey] = append(SETs[destKey], member)

	return Value{typ: "integer", integer: 1}
}
