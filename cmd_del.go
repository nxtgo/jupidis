package main

import "log"

func DelCommand(args []Value) Value {
	if len(args) < 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	KEYsMu.Lock()
	defer KEYsMu.Unlock()
	SETsMu.Lock()
	defer SETsMu.Unlock()
	HSETsMu.Lock()
	defer HSETsMu.Unlock()

	var deletedCount int
	for _, arg := range args {
		key := arg.bulk
		valueType, ok := KEYs[key]
		if !ok {
			continue
		}
		switch valueType {
		case StringValueType:
			delete(SETs, key)
		case HashValueType:
			delete(HSETs, key)
		default:
			log.Println("Unknown value type in DEL command:", valueType)
			// Should not happen
			continue
		}
		delete(KEYs, key)
		deletedCount++
	}

	return Value{typ: "integer", integer: int64(deletedCount)}
}
