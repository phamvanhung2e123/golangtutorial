package main

import (
	"fmt"
	"os"
	"time"
)

const MYFILE = "test.log"

func main() {
	c := time.Tick(1 * time.Second)
	for _ = range c {
		readFile(MYFILE)
	}
}

func readFile(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	stat, err := os.Stat(fname)
	start := stat.Size() - 1024
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("%s\n", buf)
	}
}