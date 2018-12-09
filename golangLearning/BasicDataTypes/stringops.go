package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	// c := s[len(s)] //panic
	// fmt.Println(c)
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[:])

	fmt.Println("goobye" + s[5:])

	ss := "left foot"
	t := ss
	ss += ", right foot"

	fmt.Println(t)
	fmt.Println(ss)

	// ss[0] = 'L' // compile error : connot assign to s[0]

	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
		go command [arguments]
	...`

	fmt.Println(GoUsage)
}