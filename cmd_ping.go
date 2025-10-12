package main

func PingCommandCheck(args []Value) error {
	if len(args) > 1 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func PingCommand(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].str}
}
