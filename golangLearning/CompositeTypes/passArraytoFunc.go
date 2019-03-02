package main

import "fmt"

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func main() {
	zero()
}