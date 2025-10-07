package main

import "slices"

func SUnionStoreCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	destination := args[0].str
	SETs[destination] = []string{}

	for _, arg := range args[1:] {
		for _, member := range SETs[arg.str] {
			if !slices.Contains(SETs[destination], member) {
				SETs[destination] = append(SETs[destination], member)
			}
		}
	}

	return Value{typ: "integer", integer: len(SETs[destination])}
}
