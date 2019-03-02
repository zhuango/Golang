package main

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println(`\sdfasdf/ferf\afe`)

	fmt.Println(len("世界"))
	fmt.Println("世界")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\xe4\xb8\x96")

	s := "Hello, 世界"
	fmt.Println("s[7] is ")
	fmt.Println(s[7])
	for i , r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Println(n)

	n = 0
	for range s {
		n++
	}
	fmt.Println(n)
	fmt.Println(utf8.RuneCountInString(s))

	rawS := []rune(s)
	fmt.Println(rawS)
	fmt.Printf("% x\n", s)
	fmt.Printf("%x\n", rawS)
	fmt.Println(string(rawS))
	fmt.Println(string(65))
}