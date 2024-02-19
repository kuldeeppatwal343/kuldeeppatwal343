package main

import (
	routes "Gol/Routes"
	"Gol/check"
	"fmt"
)


func main() {
	fmt.Print("hello world")
	check.CheckImportFunction(5)
	routes.Get()
}

