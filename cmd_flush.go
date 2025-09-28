package main

func FlushCommand(args []Value) Value {
	SETsMu.Lock()
	HSETsMu.Lock()
	KEYsMu.Lock()
	defer SETsMu.Unlock()
	defer HSETsMu.Unlock()
	defer KEYsMu.Unlock()

	clear(SETs)
	clear(HSETs)
	clear(KEYs)

	err := AOF.Reset()
	if err != nil {
		return Value{typ: "error", str: "Error resetting AOF: " + err.Error()}
	}

	return Value{typ: "string", str: "OK"}
}
