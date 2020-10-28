package main

import (
	"books/config"
	"fmt"
)

func main()  {
	config.DatabaseInit()
	fmt.Println("Hello Books")
}
