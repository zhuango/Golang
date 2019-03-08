package main

import "fmt"

func removedup(strings []string) []string {
	if len(strings) < 2 {
		return strings
	}
	i := 1
	for _, str := range strings[1:] {
		if str != strings[i-1]{
			strings[i] = str
			i++
		}
	}
	return strings[:i]
}

func main() {
	//strings := []string {"liu", "liu", "liuzhuang", "zha", "zha", "zha", "dabao", "dabao", "sd", "dddd", "dddd"}
	//strings := []string {"liu", "liuzhuang", "zha", "dabao", "sd", "dddd"}
	// strings := []string {"liu", "liu", "liu"}
	strings := []string {"e","e"}
	strings = removedup(strings)
	fmt.Println(strings)
}