package main

func SDiffStoreCommand(args []Value) Value {
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

	SETs[destination] = make(map[string]struct{})
	for member := range SETs[biggestSet] {
		var found bool
		for _, arg := range args[1:] {
			if arg.str == biggestSet {
				continue
			}
			if _, ok := SETs[arg.str][member]; ok {
				found = true
				break
			}
		}
		if !found {
			SETs[destination][member] = struct{}{}
		}
	}

	return Value{typ: "integer", integer: len(SETs[destination])}
}
