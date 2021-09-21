package main

import "fmt"

func main() {
	fileExtensions := map[string]string{
		"Python" : ".py",
		"Golang" : ".go",
		"Java" : ".java",
		"Ansible" : ".yml",
		"C++" : ".cpp",
	}

	fmt.Println("Language Extensions are :" ,fileExtensions)

	delete(fileExtensions,"C++")

	fileExtensions["Julian"] = ".jl"

	fmt.Println("After deleting the C++ is :", fileExtensions)

}
