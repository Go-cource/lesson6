package main

import "fmt"

func main() {
	message := "Hello!!"
	fmt.Printf("%c", message[0])
	message[0] = 'T' //ERROR!! СТРОКИ НЕИЗМЕНЯЕМЫ

}
