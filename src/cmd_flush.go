package main

func FlushCommand(args []Value) Value {
	SETsMu.Lock()
	HSETsMu.Lock()
	defer SETsMu.Unlock()
	defer HSETsMu.Unlock()

	clear(SETs)
	clear(HSETs)

	err := AOF.Reset()
	if err != nil {
		return Value{typ: "error", str: "Error resetting AOF: " + err.Error()}
	}

	return Value{typ: "string", str: "OK"}
}
