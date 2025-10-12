package main

import "slices"

func SDiffCommandCheck(args []Value) bool {
	return len(args) >= 2
}

func SDiffCommand(args []Value) Value {
	SETsMu.Lock()
	defer SETsMu.Unlock()

	var biggestSet string
	for _, arg := range args {
		if _, ok := SETs[arg.str]; !ok {
			continue
		}

		if biggestSet == "" || len(SETs[arg.str]) > len(SETs[biggestSet]) {
			biggestSet = arg.str
		}
	}

	if biggestSet == "" {
		return Value{typ: "array", array: []Value{}}
	}

	var values []Value
	for _, member := range SETs[biggestSet] {
		var found bool
		for _, arg := range args {
			if arg.str == biggestSet {
				continue
			}
			if slices.Contains(SETs[arg.str], member) {
				found = true
				break
			}
		}
		if !found && !slices.ContainsFunc(values, func(v Value) bool {
			return v.str == member
		}) {
			values = append(values, Value{typ: "string", str: member})
		}
	}
	return Value{typ: "array", array: values}
}
