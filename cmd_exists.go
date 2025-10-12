package main

func ExistsCommandCheck(args []Value) bool {
	return len(args) > 0
}

func ExistsCommand(args []Value) Value {
	LockAllMu()
	defer UnlockAllMu()

	var count int
	for _, arg := range args {
		key := arg.str
		if _, ok := VALUEs[key]; ok {
			count++
		} else if _, ok := HSETs[key]; ok {
			count++
		}
	}

	return Value{typ: "integer", integer: count}
}
