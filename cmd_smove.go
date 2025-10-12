package main

import "slices"

func SMoveCommandCheck(args []Value) bool {
	return len(args) == 3
}

func SMoveCommand(args []Value) Value {
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
