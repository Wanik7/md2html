package main

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark"
)

func TestMarkdownToHTML(t *testing.T) {
	input := "# Hello\n\nThis is **bold** text."

	var out []byte
	err := goldmark.Convert([]byte(input), &sliceWriter{b: &out})
	if err != nil {
		t.Fatalf("convert failed: %v", err)
	}

	html := string(out)

	if !strings.Contains(html, "<h1>Hello</h1>") {
		t.Fatalf("expected h1 in output, got: %s", html)
	}
	if !strings.Contains(html, "<strong>bold</strong>") {
		t.Fatalf("expected strong tag in output, got: %s", html)
	}
}
