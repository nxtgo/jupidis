package main

func HSetCommandCheck(args []Value) error {
	if (len(args)-1)%2 != 0 || len(args) < 3 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func HSetCommand(args []Value) Value {
	HSETsMu.Lock()
	defer HSETsMu.Unlock()

	key := args[0].str

	if _, available := IsKeyAvailable(key, "hash"); !available {
		return Value{typ: "error", str: "ERR key is not available"}
	}

	hset, ok := HSETs[key]
	if !ok {
		hset = map[string]string{}
		HSETs[key] = hset
	}

	for i := 1; i < len(args); i += 2 {
		field := args[i].str
		value := args[i+1].str
		hset[field] = value
	}

	return Value{typ: "string", str: "OK"}
}
