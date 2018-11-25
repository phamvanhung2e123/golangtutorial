package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("pro.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
	fmt.Println(err)
}