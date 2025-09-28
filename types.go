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

type ValueType int

const (
	StringValueType ValueType = iota + 1
	HashValueType
	SetValueType
)
