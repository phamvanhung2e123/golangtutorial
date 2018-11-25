package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	numLine := flag.Int64("n", 10, "an int64")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Please input file name")
	}
	for i := 0; i < flag.NArg(); i++ {
		fileName := flag.Arg(i)
		tailFile(fileName, *numLine)
	}
}

func tailFile(fileName string, numLine int64) {
	fOffset, err := fileOffset(fileName, numLine)
	fmt.Println("file name:", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return
		panic(err)
	}
	file.Seek(fOffset, io.SeekEnd)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func fileOffset(fileName string, numLine int64) (int64, error) {
	var cr_counter, offset int64 = 0, 0
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Cannot open file: ", fileName)
		return 0, err
	}
	defer file.Close()
	stat, err := os.Stat(fileName)

	for cr_counter <= numLine {
		buf := make([]byte, 1)
		offset -= 1
		start := stat.Size() + offset
		_, err = file.ReadAt(buf, start)
		if err != nil {
			break
		}
		if bytes.Equal(buf[:1],[]byte{'\n'}) {
			cr_counter += 1
		}
	}
	return offset + 1, nil
}
