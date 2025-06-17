package main

import (
	"fmt"
	_ "log"
)

func main() {
	question := "¿Cómo estás?"
	for _, c := range question {
		fmt.Printf("%c\n", c)
	}
}
