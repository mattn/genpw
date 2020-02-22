package main

import (
	"bytes"
	"flag"
	"testing"
)

func TestFlags(t *testing.T) {
	var buf bytes.Buffer
	err := run(&buf, -1, 0, 0, 0)
	if err != flag.ErrHelp {
		t.Fatal(err)
	}

	err = run(&buf, -1, -1, 0, 0)
	if err != flag.ErrHelp {
		t.Fatal(err)
	}

	err = run(&buf, 1, 1, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRun(t *testing.T) {
	var buf bytes.Buffer
	err := run(&buf, 1, 5, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 7 {
		t.Fatalf("length should be 7 but %v", buf.Len())
	}

	buf.Reset()
	err = run(&buf, 2, 5, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 13 {
		t.Fatalf("length should be 13 but %v", buf.Len())
	}

	buf.Reset()
	err = run(&buf, 5, 5, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 31 {
		t.Fatalf("length should be 31 but %v", buf.Len())
	}

	tests := []struct {
		c  int
		n  int
		nc int
		sc int
	}{
		{1, 5, 1, 0},
		{1, 5, 0, 1},
		{1, 5, 1, 1},
		{1, 5, 0, 0},
		{1, 5, 3, 2},
	}
	for _, test := range tests {
		buf.Reset()
		_ = run(&buf, test.c, test.n, test.nc, test.sc)
		for _, field := range bytes.Fields(buf.Bytes()) {
			got := count(field, numbers)
			if got < test.nc {
				t.Fatalf("numbers: want %v for %q but got %v", test.nc, string(field), got)
			}
			got = count(field, symbols)
			if got < test.sc {
				t.Fatalf("symbols: want %v for %q but got %v", test.sc, string(field), got)
			}
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		pattern []byte
		input   []byte
		want    int
	}{
		{[]byte("abcda"), []byte("ab"), 2},
		{[]byte(""), []byte("ab"), 0},
		{[]byte(""), []byte(""), 0},
		{[]byte("abcde"), []byte("-"), 0},
		{[]byte("a"), []byte("ab"), 1},
	}

	for _, test := range tests {
		got := count(test.input, test.pattern)
		if got != test.want {
			t.Fatalf("want %v but got %v", test.want, got)
		}
	}
}
