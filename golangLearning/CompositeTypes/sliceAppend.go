package main

import "fmt"


// what the build-in function append do inside
func self_append(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
		fmt.Printf("cap: %d, len: %d\n", cap(runes), len(runes))
	}
	fmt.Printf("%q\n", runes)
}