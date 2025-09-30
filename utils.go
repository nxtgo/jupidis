package main

func IsKeyAvailable(key string, t string) bool {
	var typeOfKey string
	if _, ok := SETs[key]; ok {
		typeOfKey = "string"
	} else if _, ok := HSETs[key]; ok {
		typeOfKey = "hash"
	}

	return typeOfKey == "" || typeOfKey == t
}
