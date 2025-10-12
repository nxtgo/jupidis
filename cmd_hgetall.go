package main

func HGetAllCommandCheck(args []Value) bool {
	return len(args) == 1
}

func HGetAllCommand(args []Value) Value {
	HSETsMu.RLock()
	defer HSETsMu.RUnlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "hash"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	hget, ok := HSETs[key]
	if !ok {
		return Value{typ: "null"}
	}

	var fields []Value
	for field, value := range hget {
		fields = append(fields, Value{typ: "string", str: field}, Value{typ: "string", str: value})
	}

	return Value{typ: "array", array: fields}
}
