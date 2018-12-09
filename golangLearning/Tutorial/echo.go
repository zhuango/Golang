package main

import "os"
import "fmt"
import "strings"

func sayHi(){
	fmt.Println("Hello, 世界")
}

func echo(){
	var s, seq string
	for i := 1; i < len(os.Args); i++{
		s += seq + os.Args[i]
		seq = " "
	}
	fmt.Println(s)
}

func echo2(){
	s, seq := "", ""
	for _, arg := range os.Args[1:]{
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}

func echo3(){
	fmt.Println(strings.Join(os.Args[1:], " "))

}

func echo4(){
	fmt.Println(os.Args[1:])
}

func echoWithIndex(){
	for i, arg := range os.Args[1:]{	
		fmt.Println(i)
		fmt.Println(arg)
	}
}

func main() {
	sayHi()
	echo()
	echo2()
	echo3()
	echo4()
	echoWithIndex()
}