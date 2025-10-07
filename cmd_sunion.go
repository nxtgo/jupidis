package main

import "slices"

func SUnionCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

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
