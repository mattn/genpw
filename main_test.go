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
	if buf.Len() != 6 {
		t.Fatalf("length should be 5 but %v", buf.Len())
	}
}
