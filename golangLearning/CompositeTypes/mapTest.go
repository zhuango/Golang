package main

import "fmt"
import "sort"

func main() {
	a := make(map[string] int)
	a["S"] = 23
	fmt.Println(a)

	b := map[string] int {
		"D": 23,
		"eee": 231,
	}
	fmt.Println(b)
	fmt.Println(a["S"])

	delete(b, "bucunzai")
	delete(b, "D")
	fmt.Println(b)
	fmt.Println(b["sssss"])

	fmt.Println("retrive+++++++++++++")
	for key, value := range b {
		fmt.Printf("%s\t%d\n", key, value)
	}

	b["liuzhuang"] = 27
	b["Allen"] = 78

	var names []string
	for key, _ := range b {
		names = append(names, key)
	}

	sort.Strings(names)
	fmt.Println("print in order-------------------")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, b[name])
	}

	fmt.Println("=======================")
	fmt.Println(b["liuzhuang"])
	age, ok := b["liuzhuang"]
	fmt.Println(age, ok)
	age, ok = b["ss"]
	if !ok {
		fmt.Println("ss is not a key in b")
	}
	fmt.Println(age, ok)
}