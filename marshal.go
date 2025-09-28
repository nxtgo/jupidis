package main

import "fmt"

func (v Value) Marshal() []byte {
	switch v.typ {
	case "string":
		return v.marshalString()
	case "bulk":
		return v.marshalBulk()
	case "array":
		return v.marshalArray()
	case "error":
		return v.marshallError()
	case "null":
		return v.marshallNull()
	}
	return nil
}

func (v Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, fmt.Appendf(nil, "%d", len(v.bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.bulk...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalArray() []byte {
	var bytes []byte
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, fmt.Appendf(nil, "%d", len(v.array))...)
	bytes = append(bytes, '\r', '\n')

	for _, elem := range v.array {
		bytes = append(bytes, elem.Marshal()...)
	}
	return bytes
}

func (v Value) marshallError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}
