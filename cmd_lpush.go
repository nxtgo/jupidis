package main

func LPushCommandCheck(args []Value) bool {
	return len(args) >= 2
}

func LPushCommand(args []Value) Value {
	LISTsMu.Lock()
	defer LISTsMu.Unlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "list"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	values := make([]string, 0, len(args)-1)
	for _, arg := range args[1:] {
		values = append(values, arg.str)
	}

	LISTs[key] = append(values, LISTs[key]...)
	return Value{typ: "integer", integer: len(LISTs[key])}
}
