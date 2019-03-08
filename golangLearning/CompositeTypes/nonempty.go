package main
import "fmt"

func nonempty(strings []string) []string{
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	strings := []string {"liuzhuang", "dsf", "", "fs", "eeeeeeeE"}
	for _, s := range nonempty(strings) {
		fmt.Println(s)
	}
}