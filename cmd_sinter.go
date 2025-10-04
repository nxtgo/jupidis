package main

func SInterCommand(args []Value) Value {
	if len(args) < 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	SETsMu.Lock()
	defer SETsMu.Unlock()

	var biggestSet string
	for _, arg := range args {
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

	var values []Value
	for member := range SETs[biggestSet] {
		var found = true
		for _, arg := range args {
			if arg.str == biggestSet {
				continue
			}
			if _, ok := SETs[arg.str][member]; !ok {
				found = false
				break
			}
		}
		if found {
			values = append(values, Value{typ: "string", str: member})
		}
	}
	return Value{typ: "array", array: values}
}
