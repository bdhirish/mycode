package main

import "fmt"

type  astro struct {
	name string
	age int
	lastmission string
	isneeded bool
}

func main() {

	ast1 := astro{"Kiran", 35, "ABB", false}
	ast2 := astro{"ketan", 41, "DBB", true}
	ast3 := astro{"Sachin", 50, "Space", true}

	fmt.Println(ast1)
	fmt.Println(ast2)
	fmt.Println(ast3)
	fmt.Printf("%+v",ast1)
}

