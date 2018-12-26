package main

import "fmt"
import "time"

const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

const (
	a = 1
	b
	c = 2
	d
)
func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Printf("a = %d", a)
	fmt.Printf("d = %d", b)
	fmt.Printf("c = %d", c)
	fmt.Printf("d = %d", d)
	fmt.Println()
}

