package internal

import (
	"reflect"
	"testing"
)

var Data = []byte(`name: 'alt:V Server',
players: 4096,
description: 'Default alt:V Server',
announce: true,
modules: [
  js-module,
  csharp-module
],
resources: [
  resource1,
  resource2
],
voice: {
  bitrate: 64000,
  externalSecret: secret123,
  externalHost: xx.xx.xx.xx,
  externalPort: 7798,
  externalPublicHost: xx.xx.xx.xx,
  externalPublicPort: 7799
},
Test: true
`)
var DataSize = 357
var TokenSize = 37

func TestNewParser(t *testing.T) {
	p := NewParser(Data)

	if p == nil {
		t.Error("NewParser returned nil")
	}
}

func TestParser_Get(t *testing.T) {
	r := []struct {
		offset int
		char   rune
	}{
		{0, 'n'},
		{1, 'a'},
		{2, 'm'},
		{3, 'e'},
		{4, ':'},
		{5, ' '},
		{6, '\''},
	}

	for _, s := range r {
		p := NewParser(Data)

		for i := 0; i < s.offset; i++ {
			p.Get()
		}

		c := p.Get()
		if c != s.char {
			t.Errorf("Parser_Get was incorrect, got: '%s', want: '%s'", string(c), string(s.char))
		}
	}
}

func TestParser_Peek(t *testing.T) {
	r := []struct {
		offset int
		char   rune
	}{
		{0, 'n'},
		{1, 'a'},
		{2, 'm'},
		{3, 'e'},
		{4, ':'},
		{5, ' '},
		{6, '\''},
	}

	for _, s := range r {
		p := NewParser(Data)

		c := p.Peek(s.offset)
		if c != s.char {
			t.Errorf("Parser_Peek was incorrect, got: '%s', want: '%s'", string(c), string(s.char))
		}
	}
}

func TestParser_Skip(t *testing.T) {
	s := []int{0, 12, 37, 42}

	for _, n := range s {
		p := NewParser(Data)

		p.Skip(n)

		if p.ReadPos != n {
			t.Errorf("Parser_Skip was incorrect, skipped: %v, expected: %v", p.ReadPos, n)
		}
	}
}

// FIXME:
func TestParser_SkipNextToken(t *testing.T) {
	p := NewParser(Data)

	unread := p.ReadPos

	p.SkipNextToken()

	if p.ReadPos == unread {
		t.Errorf("Parser_SkipNextToken failed, no bytes read")
	}
}

// TODO: multiple p.ReadPos values
func TestParser_Unread(t *testing.T) {
	p := NewParser(Data)

	u := p.Unread()
	if u != DataSize {
		t.Errorf("Parser_Unread was incorrect, got: %v, want: %v", u, DataSize)
	}
}

// TODO: check token contents
func TestParser_Tokenize(t *testing.T) {
	p := NewParser(Data)

	err := p.Tokenize()
	if err != nil {
		t.Error(err)
	}

	tSize := len(p.Tokens)
	if tSize != TokenSize {
		t.Errorf("Parser_Tokenize was incorrect, got: %v tokens, want: %v tokens", tSize, TokenSize)
	}
}

// TODO: more types
func TestParser_ParseToken(t *testing.T) {
	p := NewParser(Data)

	p.Tokens = append(p.Tokens, Token{
		Type:  Scalar,
		Value: "yes",
	})

	var b bool
	v := reflect.ValueOf(&b)

	err := p.ParseToken(v)
	if err != nil {
		t.Error(err)
	}

	if !b {
		t.Errorf("Parser_ParseToken was incorrect, got: %v, want: %v", b, true)
	}
}

// TODO:
func TestParser_Parse(t *testing.T) {

}
