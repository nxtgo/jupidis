package main

func SMIsMemberCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].bulk
	members := args[1:]
	var exists []Value
	for _, member := range members {
		if _, ok := SETs[key][member.bulk]; ok {
			exists = append(exists, Value{typ: "integer", integer: 1})
		} else {
			exists = append(exists, Value{typ: "integer", integer: 0})
		}
	}
	return Value{typ: "array", array: exists}
}
