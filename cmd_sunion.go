package main

import (
	"slices"
)

func SUnionCommandCheck(args []Value) error {
	if len(args) < 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func SUnionCommand(args []Value) Value {
	SETsMu.Lock()
	defer SETsMu.Unlock()

	var values []Value
	for _, arg := range args {
		for _, member := range SETs[arg.str] {
			if !slices.ContainsFunc(values, func(v Value) bool {
				return v.str == member
			}) {
				values = append(values, Value{typ: "string", str: member})
			}
		}
	}
	return Value{typ: "array", array: values}
}
