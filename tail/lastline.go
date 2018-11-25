package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func ScanLinesWithCR(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full carriage return-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func main() {
	f, _ := os.Open("test.log")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(ScanLinesWithCR)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}