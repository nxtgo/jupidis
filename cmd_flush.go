package main

func FlushCommand(args []Value) Value {
	LockAllMu()
	defer UnlockAllMu()

	clear(VALUEs)
	clear(HSETs)
	clear(SETs)

	err := AOF.Reset()
	if err != nil {
		return Value{typ: "error", str: "Error resetting AOF: " + err.Error()}
	}

	return Value{typ: "string", str: "OK"}
}
