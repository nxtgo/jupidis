package main

func SUnionCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	var values []Value
	seen := make(map[string]struct{})
	for _, arg := range args {
		for member := range SETs[arg.str] {
			if _, ok := seen[member]; !ok {
				seen[member] = struct{}{}
				values = append(values, Value{typ: "string", str: member})
			}
		}
	}
	return Value{typ: "array", array: values}
}
