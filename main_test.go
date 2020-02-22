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
}
