package mips32

import (
	"bufio"
	"os"
)

const (
	R = 0b_
)

func ReadFile(filename string) []byte {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(fd)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

	}
}
