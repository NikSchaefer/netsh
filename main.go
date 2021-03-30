package main

import "fmt"

func main() {
	fmt.Println("HELLO WORLD")

	const name = "sideswipe"
	const password = ""

	passed := loginToNetwork(name, password)
	fmt.Println(passed)
}


func loginToNetwork(name, password string) bool {
	
	
	return false
}