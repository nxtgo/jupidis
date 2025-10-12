package main

func CommandCommandCheck(args []Value) error {
	if len(args) != 1 {
		return ErrWrongNumberOfArguments
	}
	return nil
}

func CommandCommand(args []Value) Value {
	return Value{typ: "string", str: args[0].str}
}
