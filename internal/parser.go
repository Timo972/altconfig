package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// TokenType enum
type TokenType = uint8

const (
	// ArrayStart array start
	ArrayStart TokenType = iota
	// ArrayEnd array end
	ArrayEnd
	// DictStart dict start
	DictStart
	// DictEnd dict end
	DictEnd
	// Key dict key
	Key
	// Scalar scalar value
	Scalar
)

// Token struct
type Token struct {
	Type  TokenType
	Value string
	Pos   int
	Line  int
	Col   int
}

type Parser struct {
	Buffer   []rune
	ReadPos  int
	Line     int
	Column   int
	Tokens   []Token
	iterator int
}

// NewParser create new parser
func NewParser(data []byte) *Parser {
	p := &Parser{Tokens: make([]Token, 0), Buffer: []rune(string(data)), ReadPos: 0, Line: 0, Column: 0, iterator: 0}
	return p
}

func (p *Parser) nextCol() {
	p.Column++
	if p.Peek(0) == '\n' {
		p.Line++
		p.Column = 0
	}
}

func (p *Parser) pushToken(t Token) {
	p.Tokens = append(p.Tokens, t)
}

func (p *Parser) findToken(key string) *Token {
	for i, t := range p.Tokens {
		if t.Type != Key {
			continue
		}

		if t.Value == key {
			return &p.Tokens[i+1]
		}
	}

	return nil
}

func (p *Parser) Unread() int {
	return len(p.Buffer) - p.ReadPos
}

func (p *Parser) Peek(offset int) rune {
	return p.Buffer[p.ReadPos+offset]
}

func (p *Parser) Get() rune {
	p.nextCol()

	pos := p.ReadPos
	p.ReadPos++
	return p.Buffer[pos]
}

func (p *Parser) Skip(n int) {
	for i := 0; i < n; i++ {
		p.nextCol()
	}

	p.ReadPos += n
}

func (p *Parser) SkipNextToken() {
	for p.Unread() > 0 {
		peek := p.Peek(0)
		if peek == ' ' || peek == '\n' || peek == '\r' || peek == '\t' || peek == ',' {
			p.Skip(1)
		} else if peek == '#' {
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

func (p *Parser) Tokenize() error {
	p.pushToken(Token{
		Type: DictStart,
		Pos:  0,
		Line: 0,
		Col:  0,
	})

	for p.Unread() > 0 {
		p.SkipNextToken()

		if p.Unread() == 0 {
			break
		}

		peek := p.Peek(0)
		if peek == '[' {
			p.Skip(1)
			p.pushToken(Token{Type: ArrayStart, Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else if peek == ']' {
			p.Skip(1)
			p.pushToken(Token{Type: ArrayEnd, Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else if peek == '{' {
			p.Skip(1)
			p.pushToken(Token{Type: DictStart, Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else if peek == '}' {
			p.Skip(1)
			p.pushToken(Token{Type: DictEnd, Pos: p.ReadPos, Line: p.Line, Col: p.Column})
		} else {
			var val string

			if peek == '\'' || peek == '"' {
				start := p.Get()

				if p.Peek(0) != start {
					for p.Unread() > 1 && (p.Peek(0) == '\\' || p.Peek(1) != start) {
						peek = p.Peek(0)
						if peek == '\n' || peek == '\r' {
							if p.Get() == '\r' && p.Peek(0) == '\n' {
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
						return fmt.Errorf("unexpected end of file at line %v, column %v", p.Line, p.Column)
					}
				}

				p.Skip(1)
			} else {
				for p.Unread() > 0 &&
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
				p.pushToken(Token{
					Type:  Key,
					Value: val,
					Pos:   p.ReadPos,
					Line:  p.Line,
					Col:   p.Column,
				})
			} else {
				p.pushToken(Token{
					Type:  Scalar,
					Value: val,
					Pos:   p.ReadPos,
					Line:  p.Line,
					Col:   p.Column,
				})
			}

			peek = p.Peek(0)
			if p.Unread() > 0 && (peek == ':' || peek == ',') {
				p.Skip(1)
			}
		}
	}

	p.pushToken(Token{
		Type: DictEnd,
		Pos:  p.ReadPos,
		Line: p.Line,
		Col:  p.Column,
	})

	// fmt.Println(p.Tokens)

	return nil
}

func (p *Parser) ParseToken(v reflect.Value) error {
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	tok := p.Tokens[p.iterator]
	switch tok.Type {
	case Scalar:
		if !v.IsValid() {
			// fmt.Printf("skipping value %v\n", p.iterator)
			return nil
		}
		// fmt.Println("parsing string:", tok.Value)

		switch v.Kind() {
		case reflect.String:
			v.SetString(tok.Value)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			i, err := strconv.ParseInt(tok.Value, 10, 64)
			if err != nil {
				return errors.New("config value does not match struct type")
			}
			v.SetInt(i)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			u64, err := strconv.ParseUint(tok.Value, 10, 64)
			if err != nil {
				return errors.New("config value does not match struct type")
			}
			v.SetUint(u64)
		case reflect.Float32, reflect.Float64:
			f, err := strconv.ParseFloat(tok.Value, 64)
			if err != nil {
				return errors.New("config value does not match struct type")
			}
			v.SetFloat(f)
		case reflect.Bool:
			v.SetBool(tok.Value == "true" || tok.Value == "yes")
		}

		return nil
	case ArrayStart:
		if !v.IsValid() || (v.Kind() != reflect.Slice) {
			// config struct does not match actual config

			for i := p.iterator; i < len(p.Tokens); i++ {
				if p.Tokens[i].Type == ArrayEnd {
					// fmt.Printf("skipping array %v of length %v\n", p.iterator, i-p.iterator-1)
					p.iterator = i
					break
				}
			}
			return nil
		}
		// fmt.Println("parsing list")

		//list := make([]*Node, 0)
		et := v.Type().Elem()
		i := 0
		for p.iterator < len(p.Tokens) && p.Tokens[p.iterator+1].Type != ArrayEnd {
			// fmt.Println("parse list value")
			p.iterator++

			ev := reflect.New(et).Elem()

			err := p.ParseToken(ev)
			if err != nil {
				return err
			}

			v.Set(reflect.Append(v, ev))

			i++
		}
		p.iterator++

		return nil
	case DictStart:
		if !v.IsValid() || (v.Kind() != reflect.Struct /*&& v.Kind() != reflect.Map*/) {
			// config struct does not match actual config
			for i := p.iterator; i < len(p.Tokens); i++ {
				if p.Tokens[i].Type == DictEnd {
					// fmt.Printf("skipping dict %v of length %v\n", p.iterator, i-p.iterator-1)
					p.iterator = i
					break
				}
			}
			return nil
		}
		// fmt.Println("parsing dict")

		// todo: add map support
		dt := v.Type()
		//dict := make(Dict)
		for p.iterator < len(p.Tokens) && p.Tokens[p.iterator+1].Type != DictEnd {
			p.iterator++
			nextTok := p.Tokens[p.iterator]
			if nextTok.Type != Key {
				return fmt.Errorf("key expected, got %v %v", nextTok.Type, nextTok.Value)
			}
			key := nextTok.Value
			var fv reflect.Value
			for i := 0; i < dt.NumField(); i++ {
				f := dt.Field(i)
				if GetTagName(f) == key && f.IsExported() {
					fv = v.Field(i)
					break
				}
			}

			p.iterator++
			err := p.ParseToken(fv)
			if err != nil {
				return err
			}
			//	dict[key] = val
		}

		return nil
	}
	return errors.New("invalid token")
}

func (p *Parser) Parse(v interface{}) error {
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	if rt.Kind() != reflect.Pointer {
		return errors.New("invalid out")
	}

	rt = rt.Elem()
	rv = rv.Elem()

	return p.ParseToken(rv)
}
