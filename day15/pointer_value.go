package main

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}
