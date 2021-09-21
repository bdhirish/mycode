package main

import "fmt"

type astro struct {
	name string
	age int
	mission string
	isneeded bool
}

type nasaMission struct {
	people []astro
	number int
	message string
}

func main() {

	ast1 := astro{"Kiran", 35, "ABB", false}
        ast2 := astro{"ketan", 41, "DBB", true}
        ast3 := astro{"Sachin", 50, "Space", true}

        fmt.Println(ast1)
        fmt.Println(ast2)
        fmt.Println(ast3)

	p := []astro{ast1,ast2,ast3}
	fmt.Println(p)

	fmt.Println(p[1].mission)

	s := nasaMission{p,3, "success"}

	fmt.Println(s)

	fmt.Printf("%+v", s)

}
