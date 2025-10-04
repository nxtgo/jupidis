package main

func SIsMemberCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].bulk
	member := args[1].bulk

	_, ok := SETs[key][member]
	if ok {
		return Value{typ: "integer", integer: 1}
	}
	return Value{typ: "integer", integer: 0}
}
