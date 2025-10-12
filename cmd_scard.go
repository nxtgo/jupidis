package main

func SCardCommandCheck(args []Value) bool {
	return len(args) == 1
}

func SCardCommand(args []Value) Value {
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].str
	return Value{typ: "integer", integer: len(SETs[key])}
}
