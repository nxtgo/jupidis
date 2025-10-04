package main

func TypeCommand(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	LockAllMu()
	defer UnlockAllMu()

	typeOfKey, _ := IsKeyAvailable(args[0].bulk, "")
	if typeOfKey == "" {
		typeOfKey = "none"
	}
	return Value{typ: "string", str: typeOfKey}
}
