package main

func GetCommandCheck(args []Value) bool {
	return len(args) == 1
}

func GetCommand(args []Value) Value {
	VALUEsMu.RLock()
	defer VALUEsMu.RUnlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "string"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	value, ok := VALUEs[key]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "string", str: value}
}
