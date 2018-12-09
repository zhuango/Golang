package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i + 1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.'  {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("af/asdf/test.go"))
	fmt.Println(basename("a.b.s.go"))
	fmt.Println(basename("test"))
}