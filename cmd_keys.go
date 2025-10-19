package main

import "jupidis/pkgs/golb"

func KeysCommandCheck(args []Value) error {
	if len(args) != 1 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func KeysCommand(args []Value) Value {
	pattern := args[0].str
	g := golb.Compile(pattern)

	var result []Value

	VALUEsMu.RLock()
	for key := range VALUEs {
		if g.Match(key) {
			result = append(result, Value{typ: "bulk", str: key})
		}
	}
	VALUEsMu.RUnlock()

	HSETsMu.RLock()
	for key := range HSETs {
		if g.Match(key) {
			result = append(result, Value{typ: "bulk", str: key})
		}
	}
	HSETsMu.RUnlock()

	SETsMu.RLock()
	for key := range SETs {
		if g.Match(key) {
			result = append(result, Value{typ: "bulk", str: key})
		}
	}
	SETsMu.RUnlock()

	LISTsMu.RLock()
	for key := range LISTs {
		if g.Match(key) {
			result = append(result, Value{typ: "bulk", str: key})
		}
	}
	LISTsMu.RUnlock()

	return Value{typ: "array", array: result}
}
