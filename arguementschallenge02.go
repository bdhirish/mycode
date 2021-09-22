package main

import (
	"fmt"
	"os"
)

func main() {

	argument := len(os.Args[1:])
	fmt.Printf("Arg lenght is %d \n", argument)

	for i,a := range os.Args[1:] {
		fmt.Printf("Arg %d is %s\n", i+1, a)
	}
}
