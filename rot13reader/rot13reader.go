package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		prevChar := b[i]
		newChar := prevChar
		isAlphabetical := false
		lastChar := 'z'
		firstChar := 'a'
		if prevChar >= 'A' && prevChar <= 'Z' {
			lastChar = 'Z'
			firstChar = 'A'
			isAlphabetical = true
		}
		if prevChar >= 'a' && prevChar <= 'z' {
			isAlphabetical = true
		}
		if isAlphabetical {
			newChar = prevChar + 13
			if newChar > byte(lastChar) {
				newChar = newChar % byte(lastChar+1)
				newChar += byte(firstChar)
			}
		}
		fmt.Printf("Conveted %v to %v\n", string(prevChar), string(newChar))
		b[i] = newChar
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
