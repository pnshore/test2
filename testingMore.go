package main

import "fmt"

var m map[string]int

func main() {

	var interT interface{} = 234
	s, je := interT.(string)
	fmt.Println(s, je)

	m = make(map[string]int)
	m["mapTest"] = 40
	fmt.Println(m["mapTest"])
	i, ok := m["mapTest"]
	fmt.Println(i, ok)
	y := []int{1, 2, 5, 90}
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}
	commits := map[string]int{
		"rsc": 234,
		"r":   23,
		"sf":  45634,
	}

	fmt.Println(commits["rsc"])
}
