package main

func ExistsCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'exists' command"}
	}

	var count int64
	KEYsMu.RLock()
	for _, arg := range args {
		key := arg.bulk
		if _, ok := KEYs[key]; ok {
			count++
		}
	}
	KEYsMu.RUnlock()

	return Value{typ: "integer", integer: count}
}
