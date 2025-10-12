package main

func HGetCommandCheck(args []Value) error {
	if len(args) != 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func HGetCommand(args []Value) Value {
	HSETsMu.RLock()
	defer HSETsMu.RUnlock()

	key := args[0].str
	field := args[1].str

	if _, available := IsKeyAvailable(key, "hash"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := HSETs[key][field]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}
