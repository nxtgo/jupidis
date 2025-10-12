package main

func TypeCommandCheck(args []Value) bool {
	return len(args) == 1
}

func TypeCommand(args []Value) Value {
	LockAllMu()
	defer UnlockAllMu()

	typeOfKey, _ := IsKeyAvailable(args[0].str, "")
	if typeOfKey == "" {
		typeOfKey = "none"
	}
	return Value{typ: "string", str: typeOfKey}
}
