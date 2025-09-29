package main

func ExistsCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'exists' command"}
	}

	KEYsMu.RLock()
	defer KEYsMu.RUnlock()

	var count int
	for _, arg := range args {
		key := arg.bulk
		if _, ok := KEYs[key]; ok {
			count++
		}
	}

	return Value{typ: "integer", integer: count}
}
