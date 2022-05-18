package parse

import (
	"fmt"
	"pylot/tokenize"
	"strconv"
)

type Parser struct {
	pos    int
	tokens []tokenize.Token
	out    string
}

func Parse(tokens []tokenize.Token) string {
	psr := NewParser(tokens)
	psr.parse()
	return psr.out
}

func NewParser(tokens []tokenize.Token) *Parser {
	return &Parser{
		pos:    0,
		tokens: tokens,
	}
}

func (p *Parser) print(s string) {
	fmt.Printf(s)
	p.out += s
}

func (p *Parser) parse() {
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == tokenize.IDENT && c.Raw == "Module":
			p.module()
		case c.TokenKind == tokenize.IDENT && c.Raw == "ClassDef":
			p.classDef()
		case c.TokenKind == tokenize.COMMA:
			p.print(",")
			p.consume(",")
		case c.TokenKind == tokenize.RBR:
			// ?
			return // )
		case c.TokenKind == tokenize.RSQB:
			// ?
			return // ]
		}
		//p.goNext()
	}
}

func (p *Parser) isEof() bool {
	return p.pos >= len(p.tokens)
}

func (p *Parser) curt() tokenize.Token {
	return p.tokens[p.pos]
}

func (p *Parser) goNext() {
	p.pos++
}

func (p *Parser) consume(raw string) {
	if p.curt().Raw != raw {
		panic(fmt.Errorf("consume expect %v, but found %v", raw, p.curt().Raw))
	}
	p.goNext() // raw
}

func (p *Parser) module() {
	p.consume("Module")
	p.consume("(")
	p.print("{")
	p.print("\"type\":\"module\",")

	// body or type_ignores
loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == tokenize.IDENT && c.Raw == "body":
			p.consume("body")
			p.consume("=")
			p.consume("[")
			p.print("\"body\":[")
			p.parse()
			p.consume("]")
			p.print("]")
		case c.TokenKind == tokenize.IDENT && c.Raw == "type_ignores":
			p.consume("type_ignores")
			p.consume("=")
			p.consume("[")
			p.print("\"type_ignores\":[")
			p.parse()
			p.consume("]")
			p.print("]")
		case c.TokenKind == tokenize.COMMA:
			p.print(",")
			p.consume(",")
		case c.TokenKind == tokenize.WHITE:
			p.goNext()
		case c.TokenKind == tokenize.RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.print("}")
	if p.curt().TokenKind != tokenize.EOF {
		panic("expect eof, but found" + p.curt().TokenKind.String())
	}
	p.consume("")
}

func (p *Parser) classDef() {
	p.consume("ClassDef")
	p.consume("(")
	p.print("{")
	p.print("\"type\":\"classDef\",")

	// name, bases, keywords, body, decorator_list
loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == tokenize.IDENT && c.Raw == "name":
			p.consume("name")
			p.consume("=")
			p.print(fmt.Sprintf("\"name\":%v", strconv.Quote(p.curt().Raw)))
			p.goNext() // string
		case c.TokenKind == tokenize.IDENT && c.Raw == "bases":
			p.consume("bases")
			p.consume("=")
			p.consume("[")
			p.print("\"bases\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.print("]")
		case c.TokenKind == tokenize.IDENT && c.Raw == "keywords":
			p.consume("keywords")
			p.consume("=")
			p.consume("[")
			p.print("\"keywords\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.print("]")
		case c.TokenKind == tokenize.IDENT && c.Raw == "body":
			p.consume("body")
			p.consume("=")
			p.consume("[")
			p.print("\"body\":[")
			p.parse()
			p.consume("]")
			p.print("]")
		case c.TokenKind == tokenize.IDENT && c.Raw == "decorator_list":
			p.consume("decorator_list")
			p.consume("=")
			p.consume("[")
			p.print("\"decorator_list\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.print("]")
		case c.TokenKind == tokenize.COMMA:
			p.print(",")
			p.consume(",")
		case c.TokenKind == tokenize.WHITE:
			p.goNext()
		case c.TokenKind == tokenize.RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.print("}")
}
