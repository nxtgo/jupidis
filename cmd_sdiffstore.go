package main

import "slices"

func SDiffStoreCommand(args []Value) Value {
	if len(args) < 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	destination := args[0].str

	var biggestSet string
	for _, arg := range args[1:] {
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

	SETs[destination] = []string{}
	for _, member := range SETs[biggestSet] {
		var found bool
		for _, arg := range args[1:] {
			if arg.str == biggestSet {
				continue
			}
			if slices.Contains(SETs[arg.str], member) {
				found = true
				break
			}
		}
		if !found && !slices.Contains(SETs[destination], member) {
			SETs[destination] = append(SETs[destination], member)
		}
	}

	return Value{typ: "integer", integer: len(SETs[destination])}
}
