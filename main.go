package main

import (
	"fmt"
	"os"
)

func main() {
	helloString := "Hello world!"
	// 
	//
	fmt.Println(helloString)
	file, _ := os.Create("hello.txt")
	defer file.Close()
	_, _ = file.WriteString(helloString)

}
