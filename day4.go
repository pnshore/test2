package main

import "fmt"

type testInterface interface {
	testMethod()
}
type testStruct struct {
	j int
}

func (i testStruct) testMethod() {
	fmt.Println(i.j)
}

func main() {
	var v testInterface = testStruct{7}
	v.testMethod()
}
