package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func count(bs []byte, bp []byte) int {
	n := 0
	for _, b := range bs {
		if bytes.Index(bp, []byte{b}) != -1 {
			n++
		}
	}
	return n
}

func main() {
	var numbers = []byte("0123456789")
	var symbols = []byte("!\"#$%&'()=~|-^\\`{*}<>?_@[;:],./'")

	var n int
	var nc int
	var sc int
	flag.IntVar(&n, "n", 16, "number of characters")
	flag.IntVar(&nc, "nc", -1, "minimum count of numbers (default: any)")
	flag.IntVar(&sc, "sc", -1, "minimum count of symbols (default: any)")
	flag.Parse()

	var buf bytes.Buffer
	for r := 'a'; r < 'z'; r++ {
		buf.WriteRune(r)
	}
	for r := 'A'; r < 'Z'; r++ {
		buf.WriteRune(r)
	}
	if nc != 0 {
		buf.Write(numbers)
	}
	if sc != 0 {
		buf.Write(symbols)
	}
	chars := buf.Bytes()

	nn := 0
	for {
		var pw bytes.Buffer
		for {
			pw.Reset()
			for i := 0; i < n; i++ {
				r, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
				if err != nil {
					panic(err)
				}
				pw.WriteByte(chars[int(r.Int64())])
			}

			if nc > 0 && count(pw.Bytes(), numbers) < nc {
				continue
			}
			if sc > 0 && count(pw.Bytes(), symbols) < sc {
				continue
			}
			break
		}
		nn += n + 1
		fmt.Print(pw.String())

		if nn < 80-n {
			fmt.Print(" ")
		} else {
			fmt.Println()
			nn = 0
		}
	}
}
