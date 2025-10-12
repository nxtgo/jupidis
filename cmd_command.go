package main

func CommandCommandCheck(args []Value) bool {
	return len(args) == 1
}

func CommandCommand(args []Value) Value {
	return Value{typ: "string", str: args[0].str}
}
