package main

import "fmt"


func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
func main() {
	a := make(map[string] int)
	a["D"] = 23

	b := map[string] int {
		"D": 23,
		"eee": 231,
	}

	result := equal(a, b)
	fmt.Println(result)

	a["eee"] = 232
	result = equal(a, b)
	fmt.Println(result)

	b["eee"] = 0
	delete(a, "eee")
	result = equal(a, b)
	fmt.Println(result)
}