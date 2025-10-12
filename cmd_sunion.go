package main

import "slices"

func SUnionCommandCheck(args []Value) bool {
	return len(args) >= 2
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
