package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var alphabetUpper = map[byte]byte{
	'A': 'N',
	'L': 'Y',
}

var alphabetLower = map[byte]byte{
	'a': 'n',
	'b': 'o',
	'l': 'y',
}

func (rot rot13Reader) Read(bytes []byte) (int, error) {
	count := 0
	n, err := rot.r.Read(bytes)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		count++
		v, ok := alphabetUpper[bytes[i]]
		if ok {
			bytes[i] = v
			continue
		}
		v, ok = alphabetLower[bytes[i]]
		if ok {
			bytes[i] = v
			continue
		}

	}
	return count, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
