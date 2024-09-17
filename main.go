package main

import (
	"os"
)

func main() {
	helloString := "Hello world!"
	file, _ := os.Create("hello.txt")
	defer file.Close()
	_, _ = file.WriteString(helloString)

}
