package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage : " + os.Args[0] + " file name")
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	// do something with the file
	fmt.Print(string(file))
}