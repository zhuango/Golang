package main

import "fmt"
import "bufio"
import "os"

func main() {
	wordFreq := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordFreq[word] ++
	}
	for k, v := range wordFreq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}