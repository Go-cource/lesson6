package main

import "fmt"

func main() {
	for i := 0; i < 255; i++ {
		fmt.Printf("%d - %v \n", i, string(byte(i)))
	}
}
