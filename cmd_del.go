package main

func DelCommandCheck(args []Value) bool {
	return len(args) > 0
}

func DelCommand(args []Value) Value {
	LockAllMu()
	defer UnlockAllMu()

	var deletedCount int
	for _, arg := range args {
		key := arg.str
		if _, ok := VALUEs[key]; ok {
			deletedCount++
			delete(VALUEs, key)
			continue
		} else if _, ok := HSETs[key]; ok {
			deletedCount++
			delete(HSETs, key)
			continue
		} else if _, ok := SETs[key]; ok {
			deletedCount++
			delete(SETs, key)
			continue
		}
	}

	return Value{typ: "integer", integer: deletedCount}
}
