package main

func SInterStoreCommand(args []Value) Value {
	if len(args) < 3 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	destination := args[0].str

	var biggestSet string
	for _, arg := range args[1:] {
		if _, ok := SETs[arg.str]; !ok {
			continue
		}

		if biggestSet == "" || len(SETs[arg.str]) > len(SETs[biggestSet]) {
			biggestSet = arg.str
		}
	}

	if biggestSet == "" {
		return Value{typ: "array", array: []Value{}}
	}

	var values []string
	for member := range SETs[biggestSet] {
		var found = true
		for _, arg := range args[1:] {
			if arg.str == biggestSet {
				continue
			}
			if _, ok := SETs[arg.str][member]; !ok {
				found = false
				break
			}
		}
		if found {
			values = append(values, member)
		}
	}

	SETs[destination] = make(map[string]struct{})
	for _, member := range values {
		SETs[destination][member] = struct{}{}
	}

	return Value{typ: "integer", integer: len(values)}
}
