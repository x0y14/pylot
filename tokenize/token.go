package tokenize

import (
	"fmt"
	"strconv"
)

type Token struct {
	TokenKind
	Raw string
	S   int
	E   int
}

func NewToken(kind TokenKind, raw string, s, e int) *Token {
	return &Token{
		TokenKind: kind,
		Raw:       raw,
		S:         s,
		E:         e,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("{ TYPE: %7v, RAW: %20v, POS: %10v }", t.TokenKind.String(), strconv.Quote(t.Raw), fmt.Sprintf("%d:%d", t.S, t.E))
}
