package main

// /usr/lib/go-1.8/src/tempconv0.go (from $GOROOT)
// /home/laboratory/go/src/tempconv0.go (from $GOPATH)
import (
	"fmt"
	//"tempconv"
	"./popcount"
)

func main() {
	//tempconv.CRoF(1.23)
	result := popcount.PopCountWithLoop(1000000)
	fmt.Println(result)
}