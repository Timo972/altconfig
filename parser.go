package cfgreader

import (
	"errors"
	"fmt"
)

type TokenType = uint8

const (
	TOKEN_ARRAY_START TokenType = iota
	TOKEN_ARRAY_END
	TOKEN_DICT_START
	TOKEN_DICT_END
	TOKEN_KEY
	TOKEN_SCALAR
)

type Token struct {
	Type TokenType
	Value string
	Pos uint
	Line uint
	Col uint
}

type Parser struct {
	Tokens []Token
	Buffer []rune
	ReadPos uint
	Line uint
	Column uint
	TokIdx uint
}

// NewParser create new parser
func NewParser(content string) *Parser {
	p := &Parser{Tokens: make([]Token, 0), Buffer: []rune(content), ReadPos: 0, Line: 0, Column: 0, TokIdx: 0}
	return p
}

// Parse string to alt-config Node
func (p *Parser) Parse() (*Node, error) {
	err := p.Tokenize()
	if err != nil {
		return nil, err
	}
	node, err := p.ParseTok()
	return node, err
}

// Unread get size of unread bytes
func (p *Parser) Unread() uint {
	return uint(len(p.Buffer)) - p.ReadPos
}

// Peek char at current position with offset
func (p *Parser) Peek(offset uint) rune {
	idx := p.ReadPos + offset
	return rune(p.Buffer[idx])
}

// Get char at current position and update read position
func (p *Parser) Get() rune {
	p.Column++
	if p.Peek(0) == '\n' {
		p.Line++
		p.Column = 0
	}
	currPos := p.ReadPos
	p.ReadPos++
	return p.Buffer[currPos]
}

// Skip skip n chars
func (p *Parser) Skip(n uint) {
	for i := uint(0) ; i < n; i++ {
		p.Column++
		if p.Peek(i) == '\n' {
			p.Line++
			p.Column = 0
		}
	}
	p.ReadPos = p.ReadPos + n
}

// SkipToNextToken skip until next token
func (p *Parser) SkipToNextToken() {
	for p.Unread() > 0 {
		if p.Peek(0) == ' ' || p.Peek(0) == '\n' || p.Peek(0) == '\r' || p.Peek(0) == '\t' || p.Peek(0) == ',' {
			p.Skip(1)
		} else if p.Peek(0) == '#' {
			p.Skip(1)

			for p.Unread() > 0 && p.Peek(0) != '\n' && p.Peek(0) != '#' {
				p.Skip(1)
			}

			if p.Unread() > 0 {
				p.Skip(1)
			}
		} else {
			break
		}
	}
}

// Tokenize content
func (p *Parser) Tokenize() error {
	p.Tokens = append(p.Tokens, Token{Type:TOKEN_DICT_START, Value: "", Pos: 0, Line: 0, Col: 0})

	for p.Unread() > 0 {
		p.SkipToNextToken()

		if p.Unread() == 0 {
			break
		}

		if p.Peek(0) == '[' {
			p.Skip(1)
			p.Tokens = append(p.Tokens, Token{Type:TOKEN_ARRAY_START, Value: "", Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else if p.Peek(0) == ']' {
			p.Skip(1)
			p.Tokens = append(p.Tokens, Token{Type: TOKEN_ARRAY_END, Value: "", Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else if p.Peek(0) == '{' {
			p.Skip(1)
			p.Tokens = append(p.Tokens, Token{Type: TOKEN_DICT_START, Value: "", Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else if p.Peek(0) == '}' {
			p.Skip(1)
			p.Tokens = append(p.Tokens, Token{Type: TOKEN_DICT_END, Value: "", Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else {
			//TODO this is not how it should work because if string starts with " it gets added here. Bad.
			val := "" // string(p.Peek(0))

			if p.Peek(0) == '\'' || p.Peek(0) == '"' {
				start := p.Get()

				if p.Peek(0) != start {
					for p.Unread() > 1 && (p.Peek(0) == '\\' || p.Peek(1) != start) {
						if p.Peek(0) == '\n' || p.Peek(0) == '\r' {
							if p.Get() == '\n' || p.Peek(0) == '\n' {
								p.Skip(1)
							}
							val += "\n"
							continue
						}
						val += string(p.Get())
					}

					if p.Unread() > 0 {
						val += string(p.Get())
					}

					if p.Unread() == 0 {
						return errors.New(fmt.Sprintf("Unexpected end of file"))
					}
				}

				p.Skip(1)
			} else {
				for p.Unread() > 0 &&
					// TODO this check is not working -> even if Peek returns , it gets added
					p.Peek(0) != '\n' &&
					p.Peek(0) != ':' &&
					p.Peek(0) != ',' &&
					p.Peek(0) != ']' &&
					p.Peek(0) != '}' &&
					p.Peek(0) != '#' {

					val += string(p.Get())
				}
			}

			val = Unescape(val)

			if p.Unread() > 0 && p.Peek(0) == ':' {
				p.Tokens = append(p.Tokens, Token{Type: TOKEN_KEY, Value: val, Pos: p.ReadPos, Line: p.Line, Col: p.Column})
			} else {
				p.Tokens = append(p.Tokens, Token{Type: TOKEN_SCALAR, Value: val, Pos: p.ReadPos, Line: p.Line, Col: p.Column})
			}

			if p.Unread() > 0 && (p.Peek(0) == ':' || p.Peek(0) == ',') {
				p.Skip(1)
			}
		}
	}
	p.Tokens = append(p.Tokens, Token{Type: TOKEN_DICT_END, Value: "", Pos: 0, Line: 0, Col: 0})
	return nil
}

// ParseTok parse token to node
func (p *Parser) ParseTok() (*Node, error) {
	tok := p.Tokens[p.TokIdx]
	switch tok.Type {
	case TOKEN_SCALAR:
		return &Node{Type: SCALAR, Value: tok.Value}, nil
	case TOKEN_ARRAY_START:
		list := make([]*Node, 0)
		for p.TokIdx < uint(len(p.Tokens)) && p.Tokens[p.TokIdx + 1].Type != TOKEN_ARRAY_END {
			p.TokIdx++
			node, err := p.ParseTok()
			if err != nil {
				return nil, err
			}
			list = append(list, node)
		}
		p.TokIdx++

		return &Node{Type: LIST, Value: list}, nil
	case TOKEN_DICT_START:
		dict := make(Dict)
		for p.TokIdx < uint(len(p.Tokens)) && p.Tokens[p.TokIdx + 1].Type != TOKEN_DICT_END {
			p.TokIdx++
			nextTok := p.Tokens[p.TokIdx]
			if nextTok.Type != TOKEN_KEY {
				return nil, errors.New("key expected")
			}
			key := nextTok.Value
			p.TokIdx++
			val, err := p.ParseTok()
			if err != nil {
				return nil, err
			}
			dict[key] = val
		}

		return &Node{Type: DICT, Value: dict}, nil
	}
	return nil, errors.New("invalid token")
}