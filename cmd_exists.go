package main

func ExistsCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'exists' command"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()
	HSETsMu.RLock()
	defer HSETsMu.RUnlock()

	var count int
	for _, arg := range args {
		key := arg.bulk
		if _, ok := SETs[key]; ok {
			count++
		} else if _, ok := HSETs[key]; ok {
			count++
		}
	}

	return Value{typ: "integer", integer: count}
}
