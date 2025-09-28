package main

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ     string
	str     string
	bulk    string
	array   []Value
	integer int64
}
