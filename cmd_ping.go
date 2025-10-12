package main

func PingCommandCheck(args []Value) bool {
	return len(args) <= 1
}

func PingCommand(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].str}
}
