package main

func ExistsCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'exists' command"}
	}

	LockAllMu()
	defer UnlockAllMu()

	var count int
	for _, arg := range args {
		key := arg.bulk
		if _, ok := VALUEs[key]; ok {
			count++
		} else if _, ok := HSETs[key]; ok {
			count++
		}
	}

	return Value{typ: "integer", integer: count}
}
