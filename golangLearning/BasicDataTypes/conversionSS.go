package main

import "fmt"
import "strconv"


func main() {

	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(12), 2))
	fmt.Println(fmt.Sprintf("x=%d", x))
	fmt.Println(fmt.Sprintf("x=%b", x))

	xx, err := strconv.Atoi("123")
	yy, err := strconv.ParseInt("123", 10, 64)
}