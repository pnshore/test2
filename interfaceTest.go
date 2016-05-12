package main

import "fmt"

type I interface {
	M()
}

type S struct {
	x int
}

func (t S) M() {
	fmt.Println(x)
}

func main() {
	var p I = S{9}
	p.M()
}
