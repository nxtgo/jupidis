package main

func HSetCommand(args []Value) Value {
	if (len(args)-1)%2 != 0 || len(args) < 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk

	HSETsMu.RLock()
	hset, ok := HSETs[key]
	HSETsMu.RUnlock()

	if !ok {
		hset = map[string]string{}
		HSETsMu.Lock()
		HSETs[key] = hset
		HSETsMu.Unlock()
	}

	KEYsMu.Lock()
	KEYs[key] = HashValueType
	KEYsMu.Unlock()

	HSETsMu.Lock()
	for i := 1; i < len(args); i += 2 {
		field := args[i].bulk
		value := args[i+1].bulk
		hset[field] = value
	}
	HSETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}
