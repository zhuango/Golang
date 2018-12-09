package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func dup1(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] ++
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else{
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	
	for line, n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup3(){
	counts := make(map[string]int)
	dupFilename := make(map[string]string)
	hasDup := make(map[string]int)
	for _, filename := range os.Args[1:]{
		data, err := ioutil.ReadFile(filename)
		// fmt.Printf(string(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){
			counts[line] ++
			hasDup[line] = 1
		}
		for key, _ := range hasDup {
			dupFilename[key] += filename + " "
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\n", dupFilename[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()] ++
	}
}
func main(){
	//dup1()
	//dup2()
	dup3()
}