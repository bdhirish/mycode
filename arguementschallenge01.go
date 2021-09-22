 package main

 import (
	 "fmt"
	 "os"
 )

 func main() {

	 arglength := len(os.Args[1:])

	 fmt.Printf("Arg length is %d", arglength)
 }
