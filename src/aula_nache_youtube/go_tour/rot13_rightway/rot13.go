package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(bytes []byte) (int, error) {
	count := 0
	n, err := rot.r.Read(bytes)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		count++
		if bytes[i] >= 97 && bytes[i] <= 122 {
			if bytes[i]+13 > 122 {
				bytes[i] = bytes[i] + 13 - 26
				continue
			}
			bytes[i] = bytes[i] + 13
		}
		if bytes[i] >= 65 && bytes[i] <= 90 {
			if bytes[i]+13 > 90 {
				bytes[i] = bytes[i] + 13 - 26
				continue
			}
			bytes[i] = bytes[i] + 13
		}
	}
	return count, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
