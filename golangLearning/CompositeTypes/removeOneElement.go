package main 
import "fmt"

func remmoveOneLlement(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s) - 1]
}

func main() {
	s := []string {"liuzhuang", "zhuan", "sfsf", "ddddddd"}
	s = remmoveOneLlement(s, 2)
	fmt.Println(s)
}