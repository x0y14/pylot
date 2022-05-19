package parse

import (
	"fmt"
	"pylot"
	"pylot/tokenize"
	"strconv"
)

var (
	reserved = []string{"None", "True", "False"}
)

type Parser2 struct {
	pos    int
	tokens []tokenize.Token
	out    string
}

func V2(tokens []tokenize.Token) string {
	psr := NewParser2(tokens)
	psr.parse()
	return psr.out
}

func NewParser2(tokens []tokenize.Token) *Parser2 {
	return &Parser2{
		pos:    0,
		tokens: tokens,
	}
}
func (p *Parser2) write(s string) {
	fmt.Printf(s)
	p.out += s
}

func (p *Parser2) isEof() bool {
	return p.pos >= len(p.tokens)
}

func (p *Parser2) curt() tokenize.Token {
	return p.tokens[p.pos]
}
func (p *Parser2) next() tokenize.Token {
	return p.tokens[p.pos+1]
}

func (p *Parser2) goNext() {
	p.pos++
}

func (p *Parser2) consume(raw string) {
	if p.curt().Raw != raw {
		panic(fmt.Errorf("consume expect %v, but found %v", raw, p.curt().Raw))
	}
	p.goNext() // raw
}

func (p *Parser2) parse() {
loop:
	for !p.isEof() {
		curt := p.curt()
		switch {
		case curt.TokenKind == tokenize.IDENT:
			if pylot.StringsContain(reserved, curt.Raw) {
				p.keyword()
			} else if p.next().TokenKind == tokenize.LBR {
				p.def()
			} else {
				p.param()
			}
		case curt.TokenKind == tokenize.COMMA:
			//p.write(",")
			//fmt.Printf("$%v", p.next().TokenKind.String())
			if p.next().TokenKind != tokenize.RSQB && p.next().TokenKind != tokenize.RBR {
				p.write(",")
			}
			p.consume(",")
		case curt.TokenKind == tokenize.LSQB:
			// [
			p.list()
		case curt.TokenKind == tokenize.RSQB:
			// ]
			return
		case curt.TokenKind == tokenize.RBR:
			// )
			return
		case curt.TokenKind == tokenize.STRING || curt.TokenKind == tokenize.NUMBER:
			p.value()
		case curt.TokenKind == tokenize.WHITE:
			p.goNext()
		case curt.TokenKind == tokenize.EOF:
			break loop
		default:
			panic("syntax error: " + p.curt().Raw)
		}
	}
	p.write("")
}

func (p *Parser2) def() {
	ident := p.curt()
	p.goNext()
	p.consume("(")
	p.write(fmt.Sprintf(`{"type":"%v"`, ident.Raw))
	if p.curt().TokenKind == tokenize.IDENT {
		p.write(",")
	}
	p.parse()
	p.consume(")")
	p.write("}")
}

func (p *Parser2) param() {
	ident := p.curt()
	p.goNext()
	p.consume("=")
	p.write(fmt.Sprintf(`"%v":`, ident.Raw))
	p.parse()
}

func (p *Parser2) list() {
	p.consume("[")
	p.write("[")
	p.parse()
	p.consume("]")
	p.write("]")
}

func (p *Parser2) value() {
	v := p.curt()
	p.write(strconv.Quote(v.Raw))
	p.goNext()
}

func (p *Parser2) keyword() {
	switch p.curt().Raw {
	case "None":
		p.write("null")
	case "True":
		p.write("true")
	case "False":
		p.write("false")
	}
	p.goNext()
}
