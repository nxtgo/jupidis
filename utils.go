package main

// string, list, set, zset, hash, stream, and vectorset.
func IsKeyAvailable(key string, t string) (string, bool) {
	var typeOfKey string = ""
	if _, ok := VALUEs[key]; ok {
		typeOfKey = "string"
	} else if _, ok := HSETs[key]; ok {
		typeOfKey = "hash"
	} else if _, ok := SETs[key]; ok {
		typeOfKey = "set"
	}

	return typeOfKey, typeOfKey == "" || typeOfKey == t
}

func LockAllMu() {
	VALUEsMu.Lock()
	HSETsMu.Lock()
	SETsMu.Lock()
}

func UnlockAllMu() {
	VALUEsMu.Unlock()
	HSETsMu.Unlock()
	SETsMu.Unlock()
}
