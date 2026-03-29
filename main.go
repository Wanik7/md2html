package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yuin/goldmark"
)

func main() {
	in := flag.String("in", "", "input markdown file (required)")
	out := flag.String("out", "", "output html file (required)")
	flag.Parse()

	if *in == "" || *out == "" {
		fmt.Fprintln(os.Stderr, "usage: md2html -in input.md -out output.html")
		os.Exit(2)
	}

	mdBytes, err := os.ReadFile(*in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input file: %v\n", err)
		os.Exit(1)
	}

	var htmlBytes []byte
	if err := goldmark.Convert(mdBytes, &sliceWriter{b: &htmlBytes}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert markdown: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*out, htmlBytes, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("converted %s -> %s\n", *in, *out)
}

// sliceWriter implements io.Writer into a byte slice pointer.
type sliceWriter struct {
	b *[]byte
}

func (w *sliceWriter) Write(p []byte) (n int, err error) {
	*w.b = append(*w.b, p...)
	return len(p), nil
}
