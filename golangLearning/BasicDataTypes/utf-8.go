package main

import "unicode/utf8"
import "fmt"


func main() {
	s := "Hello, 世界"
	fmt.Println(s[7:])  // 世界
	fmt.Println(s[10:]) // 界
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	
	// H e l l o ,   世 界
	for _, char := range s {
		fmt.Printf("%c ", char)
	}
	fmt.Println()
}
