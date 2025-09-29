package main

import "strconv"

func IncrByCommand(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments"}
	}

	key := args[0].bulk
	strIncrement := args[1].bulk

	increment, err := strconv.Atoi(strIncrement)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	KEYsMu.Lock()
	defer KEYsMu.Unlock()
	SETsMu.Lock()
	defer SETsMu.Unlock()

	value, ok := SETs[key]
	if !ok {
		SETs[key] = strconv.Itoa(increment)
		KEYs[key] = StringValueType
		return Value{typ: "integer", integer: increment}
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return Value{typ: "error", str: "ERR value is not an integer or out of range"}
	}

	intValue += increment
	SETs[key] = strconv.Itoa(intValue)
	return Value{typ: "integer", integer: intValue}
}
