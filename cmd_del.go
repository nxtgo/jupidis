package main

func DelCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

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
