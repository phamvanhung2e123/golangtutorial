package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var file_name = "pro.txt"
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	stat, err := os.Stat(file_name)
	start := stat.Size() - 1024
	fmt.Println(start)
	file.Seek(start,io.SeekEnd)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}