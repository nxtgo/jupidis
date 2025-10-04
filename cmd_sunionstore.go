package main

func SUnionStoreCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	destination := args[0].str
	SETs[destination] = make(map[string]struct{})

	var values []string
	seen := make(map[string]struct{})
	for _, arg := range args[1:] {
		for member := range SETs[arg.str] {
			if _, ok := seen[member]; !ok {
				seen[member] = struct{}{}
				values = append(values, member)
				SETs[destination][member] = struct{}{}
			}
		}
	}

	return Value{typ: "integer", integer: len(values)}
}
