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

	var matchingKeys []string

	VALUEsMu.RLock()
	for key := range VALUEs {
		if g.Match(key) {
			matchingKeys = append(matchingKeys, key)
		}
	}
	VALUEsMu.RUnlock()

	HSETsMu.RLock()
	for key := range HSETs {
		if g.Match(key) {
			matchingKeys = append(matchingKeys, key)
		}
	}
	HSETsMu.RUnlock()

	SETsMu.RLock()
	for key := range SETs {
		if g.Match(key) {
			matchingKeys = append(matchingKeys, key)
		}
	}
	SETsMu.RUnlock()

	LISTsMu.RLock()
	for key := range LISTs {
		if g.Match(key) {
			matchingKeys = append(matchingKeys, key)
		}
	}
	LISTsMu.RUnlock()

	result := make([]Value, len(matchingKeys))
	for i, key := range matchingKeys {
		result[i] = Value{typ: "bulk", str: key}
	}

	return Value{typ: "array", array: result}
}
