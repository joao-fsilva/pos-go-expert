package math

type Match struct {
	A int
	B int
}

func (m Match) Add() int {
	return m.A + m.B
}
