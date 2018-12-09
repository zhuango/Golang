package main

import "fmt"
import "strings"

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("af/asdf/test.go"))
	fmt.Println(basename("a.b.s.go"))
	fmt.Println(basename("test"))
}