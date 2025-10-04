package main

func CommandCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	return Value{typ: "string", str: args[0].str}
}
