package main

import (
	"strings"
	"fmt"
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNoRecur(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	rest := n % 3
	var buf bytes.Buffer
	if rest == 0 {
		buf.WriteString(s[:3])
	} else {
		buf.WriteString(s[:rest])
	}
	for i := rest; i < n; i += 3 {
		buf.WriteRune(',')
		buf.WriteString(s[i:i+3])
	}
	return buf.String()
}

func similiarString(s1, s2 string) bool {
	for _, r := range s1 {
		if !strings.Contains(s2, string(r)) {
			return false
		}
	}
	for _, r := range s2 {
		if !strings.Contains(s1, string(r)) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(comma("12343145464234564"))
	fmt.Println(commaNoRecur("12343145464234564"))
	fmt.Println(commaNoRecur("1"))
	fmt.Println(commaNoRecur("12"))

	fmt.Println(commaNoRecur("123"))
	fmt.Println(commaNoRecur("123456"))
	fmt.Println(commaNoRecur(""))

	fmt.Println("================================")
	fmt.Println(similiarString("asdf", "fdas"))
	fmt.Println(similiarString("asdf", "fdassess"))
}