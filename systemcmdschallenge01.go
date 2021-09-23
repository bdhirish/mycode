package main

import (
	"fmt"
	"log"
	"os/exec"
)


func main() {

	cmd,err := exec.Command("ping","-c","4").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(cmd))
}

