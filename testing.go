package main

import "fmt"

type structTest struct {
	X int
	Y int
}

func (s structTest) methodTest() int {
	return s.X * s.Y
}

func main() {
	v := structTest{3, 9}
	fmt.Println(v.methodTest())
}
