package main

func AppendCommandCheck(args []Value) error {
	if len(args) != 2 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func AppendCommand(args []Value) Value {
	VALUEsMu.Lock()
	defer VALUEsMu.Unlock()

	key := args[0].str
	appendValue := args[1].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := VALUEs[key]
	if !ok {
		VALUEs[key] = appendValue
		return Value{typ: "integer", integer: len(appendValue)}
	}

	value += appendValue
	VALUEs[key] = value
	return Value{typ: "integer", integer: len(value)}
}
