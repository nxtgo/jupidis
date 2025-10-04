package main

func HGetAllCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

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
