package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	messageLength := len(message)
	for i := 0; i < messageLength; i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

}
