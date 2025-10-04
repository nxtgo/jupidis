package main

func SMoveCommand(args []Value) Value {
	if len(args) != 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	srcKey := args[0].str
	destKey := args[1].str
	member := args[2].str

	SETsMu.Lock()
	defer SETsMu.Unlock()

	if _, ok := SETs[srcKey][member]; !ok {
		return Value{typ: "integer", integer: 0}
	}

	delete(SETs[srcKey], member)
	if SETs[srcKey] == nil {
		delete(SETs, srcKey)
	}

	if SETs[destKey] == nil {
		SETs[destKey] = make(map[string]struct{})
	}
	SETs[destKey][member] = struct{}{}

	return Value{typ: "integer", integer: 1}
}
