package main

import "fmt"

type Currency int
const (
	USD Currency = iota
	EUR
	GBR
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q[2])

	var test int = 12
	fmt.Println(test)

	fromPair := [...]string{USD:"$", EUR:"9", GBR:"we", RMB:"yang"}
	fmt.Println(fromPair[0])

	unfill := [...]string{99:"liuzhuang"}
	fmt.Println(unfill[99])
	fmt.Println(len(unfill))

	//comparsion
	b := [2]int{1,2}
	c := [2]int{1,2}
	d := [...]int{1, 2}
	fmt.Println(b == c, c == d, b == d)

	e := [3]int {1, 2}
	fmt.Println(a == e)
}