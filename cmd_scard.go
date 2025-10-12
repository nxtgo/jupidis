package main

func SCardCommandCheck(args []Value) error {
	if len(args) != 1 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func SCardCommand(args []Value) Value {
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "set"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	return Value{typ: "integer", integer: len(SETs[key])}
}
