package main

import (
	"slices"
)

func SUnionStoreCommandCheck(args []Value) error {
	if len(args) < 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func SUnionStoreCommand(args []Value) Value {
	SETsMu.Lock()
	defer SETsMu.Unlock()

	destination := args[0].str
	if _, available := IsKeyAvailable(destination, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

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
