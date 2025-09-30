package main

func HSetCommand(args []Value) Value {
	if (len(args)-1)%2 != 0 || len(args) < 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	HSETsMu.Lock()
	defer HSETsMu.Unlock()

	key := args[0].bulk

	if _, available := IsKeyAvailable(key, "hash"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	hset, ok := HSETs[key]
	if !ok {
		hset = map[string]string{}
		HSETs[key] = hset
	}

	for i := 1; i < len(args); i += 2 {
		field := args[i].bulk
		value := args[i+1].bulk
		hset[field] = value
	}

	return Value{typ: "string", str: "OK"}
}
