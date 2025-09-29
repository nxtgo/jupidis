package main

import "log"

func DelCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()
	HSETsMu.Lock()
	defer HSETsMu.Unlock()

	var deletedCount int
	for _, arg := range args {
		key := arg.bulk
		if _, ok := SETs[key]; ok {
			deletedCount++
			delete(SETs, key)
			continue
		} else if _, ok := HSETs[key]; ok {
			deletedCount++
			delete(HSETs, key)
			continue
		} else {
			log.Println("DEL: key not found:", key)
			continue
		}
	}

	return Value{typ: "integer", integer: deletedCount}
}
