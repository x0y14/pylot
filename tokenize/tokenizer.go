package tokenize

import (
	"fmt"
	"unicode"
)

type Tokenizer struct {
	pos   int
	runes []rune
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		pos:   0,
		runes: nil,
	}
}

func (t *Tokenizer) Tokenize(text string) ([]Token, error) {
	t.runes = []rune(text)

	var result []Token

	for !t.isEof() {
		c := t.curt()
		switch {
		case 65 <= c && c <= 90 || 97 <= c && c <= 122 || c == '_':
			// A-Z || a-z || _
			tok := t.ident()
			result = append(result, *tok)
			continue // 新しい領域に踏み込んでるので goNextは不要
		case c == '"' || c == '\'':
			// string
			tok := t.str()
			result = append(result, *tok)
			continue // 新しい領域に踏み込んでるので goNextは不要
		case c == '-' || 48 <= c && c <= 57:
			// - || 0-9
			tok := t.number()
			result = append(result, *tok)
			continue // 新しい領域に踏み込んでるので goNextは不要
		case unicode.IsSpace(c):
			// white
			tok := t.white()
			result = append(result, *tok)
			continue // 新しい領域に踏み込んでるので goNextは不要
		case c == '(':
			tok := NewToken(LBR, string(c), t.pos, t.pos+1)
			result = append(result, *tok)
		case c == ')':
			tok := NewToken(RBR, string(c), t.pos, t.pos+1)
			result = append(result, *tok)
		case c == '[':
			tok := NewToken(LSQB, string(c), t.pos, t.pos+1)
			result = append(result, *tok)
		case c == ']':
			tok := NewToken(RSQB, string(c), t.pos, t.pos+1)
			result = append(result, *tok)
		case c == ',':
			tok := NewToken(COMMA, string(c), t.pos, t.pos+1)
			result = append(result, *tok)
		case c == '=':
			tok := NewToken(EQU, string(c), t.pos, t.pos+1)
			result = append(result, *tok)

		default:
			return nil, fmt.Errorf("unexpected token[%d]: %v", t.pos, string(c))
		}
		t.goNext()
	}

	result = append(result, *NewToken(EOF, "", t.pos, t.pos))
	return result, nil
}

func (t *Tokenizer) prev() rune {
	return t.runes[t.pos-1]
}

func (t *Tokenizer) curt() rune {
	return t.runes[t.pos]
}

func (t *Tokenizer) goNext() {
	t.pos++
}

func (t *Tokenizer) isEof() bool {
	return t.pos >= len(t.runes)
}

func (t *Tokenizer) ident() *Token {
	raw := ""
	s := t.pos
	for !t.isEof() {
		c := t.curt()
		if 65 <= c && c <= 90 || 97 <= c && c <= 122 || c == '_' {
			raw += string(c)
		} else {
			break
		}
		t.goNext()
	}
	e := t.pos
	return NewToken(IDENT, raw, s, e)
}

func (t *Tokenizer) str() *Token {
	raw := ""
	s := t.pos

	var dq bool
	if t.curt() == '"' {
		dq = true
	} else {
		dq = false
	}

	t.goNext() // " or '
	for !t.isEof() {
		c := t.curt()
		if c == '"' && dq || c == '\'' && !dq {
			if t.pos != 0 && t.prev() == '\\' {
				raw += string(c)
				t.goNext()
				continue
			} else {
				break
			}
		}
		raw += string(c)
		t.goNext()
	}
	t.goNext()
	e := t.pos
	return NewToken(STRING, raw, s, e)
}

func (t *Tokenizer) number() *Token {
	raw := ""
	s := t.pos
	for !t.isEof() {
		c := t.curt()
		if c == '-' || c == '.' || 48 <= c && c <= 57 {
			raw += string(c)
		} else {
			break
		}
		t.goNext()
	}
	e := t.pos
	return NewToken(NUMBER, raw, s, e)
}

func (t *Tokenizer) white() *Token {
	raw := ""
	s := t.pos
	for !t.isEof() {
		c := t.curt()
		if !unicode.IsSpace(c) {
			break
		}
		raw += string(c)
		t.goNext()
	}
	e := t.pos
	// 新しい領域に踏み込んでる
	return NewToken(WHITE, raw, s, e)
}
