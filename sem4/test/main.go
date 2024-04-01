package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		if errors.Is(err, os.ErrPermission) {

		}
		log.Fatal(err)
	}

	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err := fmt.Errorf("%w", closeErr)
			fmt.Println(err)
		}
	}()

	b := make([]byte, 64*1024)

	var line []byte

	for {
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		if !bytes.Contains(b[:n], []byte{'\n'}) {
			line = append(line, b[:n]...)
		}

		fmt.Println(line)
		line = line[:0]
	}

	var line1 string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line1 = scanner.Text()
		fmt.Println(line1)
	}

	fmt.Fprintf(f, "printing %s", line1)
}
