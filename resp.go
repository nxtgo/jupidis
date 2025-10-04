package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{
		reader: bufio.NewReader(rd),
	}
}

func (r *Resp) readLine() ([]byte, int, error) {
	var line []byte
	var n int
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, fmt.Errorf("error reading byte: %v", err)
		}
		n++
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line, n, nil
}

func (r *Resp) Read() (Value, error) {
	prefix, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, errors.Join(fmt.Errorf("error reading prefix"), err)
	}
	switch prefix {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		return Value{}, fmt.Errorf("unknown type: %v", string(prefix))
	}
}

func (r *Resp) readInteger() (int, int, error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, fmt.Errorf("error reading line: %v", err)
	}

	i64, err := strconv.ParseInt(string(line)[0:len(line)-2], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid integer: %v", err)
	}
	return int(i64), n, nil
}

func (r *Resp) readArray() (Value, error) {
	v := Value{typ: "array"}

	length, _, err := r.readInteger()
	if err != nil {
		return Value{}, fmt.Errorf("error reading array length: %v", err)
	}

	v.array = make([]Value, length)
	for i := range length {
		val, err := r.Read()
		if err != nil {
			return v, fmt.Errorf("error reading array element: %v", err)
		}
		v.array[i] = val
	}
	return v, nil
}

func (r *Resp) readBulk() (Value, error) {
	v := Value{typ: "bulk"}

	length, _, err := r.readInteger()
	if err != nil {
		return v, fmt.Errorf("error reading bulk length: %v", err)
	}

	bulk := make([]byte, length)
	_, err = r.reader.Read(bulk)
	if err != nil {
		return v, fmt.Errorf("error reading bulk data: %v", err)
	}
	v.str = string(bulk)
	r.readLine()
	return v, nil
}
