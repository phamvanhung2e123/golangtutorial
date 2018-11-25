package main

import (
	"io/ioutil"
	"os"
)

func main() {
	content := []byte(
		"hello world\n" +
			"hello world\n" +
			"hello world\n" +
			"hello world\n",
	)
	ioutil.WriteFile("/tmp/go-file", content, os.ModePerm)
}