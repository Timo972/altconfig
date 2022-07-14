package internal

import (
	"testing"
)

func TestEscape(t *testing.T) {
	target := "hello\\n"
	escaped := Escape("hello\n")

	if escaped != target {
		t.Errorf("Escaping failed, got: %s, want: %s", escaped, target)
	}
}

// FIXME:
func TestUnescape(t *testing.T) {
	target := "hello"
	unescaped := Unescape("hello\n")

	if unescaped != target {
		t.Errorf("Unescaping failed, got: %s, want: %s", unescaped, target)
	}
}
