package main

func HGetAllCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk

	HSETsMu.RLock()
	hget, ok := HSETs[key]
	HSETsMu.RUnlock()

	if !ok {
		return Value{typ: "null"}
	}

	var fields []Value
	for field, value := range hget {
		fields = append(fields, Value{typ: "string", str: field}, Value{typ: "string", str: value})
	}

	return Value{typ: "array", array: fields}
}
