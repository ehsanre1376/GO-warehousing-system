package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteString("Hello World")
	// defer will be executed at the end of the function
	defer file.Close()
}
