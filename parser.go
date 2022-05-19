package pylot

import (
	"fmt"
	"strconv"
)

type Parser struct {
	pos    int
	tokens []Token
	out    string
}

func Parse(tokens []Token) string {
	psr := NewParser(tokens)
	psr.parse()
	return psr.out
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		pos:    0,
		tokens: tokens,
	}
}

func (p *Parser) write(s string) {
	fmt.Printf(s)
	p.out += s
}

func (p *Parser) parse() {
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == IDENT && c.Raw == "Module":
			p.module()
		case c.TokenKind == IDENT && c.Raw == "ClassDef":
			p.classDef()
		case c.TokenKind == IDENT && c.Raw == "arguments":
			p.arguments()
		case c.TokenKind == IDENT && c.Raw == "arg":
			p.arg()
		case c.TokenKind == IDENT && c.Raw == "Attribute":
			panic("")
		case c.TokenKind == IDENT && c.Raw == "Constant":
			panic("")
		case c.TokenKind == IDENT && c.Raw == "Name":
			p.name()
		case c.TokenKind == IDENT && c.Raw == "Expr":
			panic("")
		case c.TokenKind == IDENT && c.Raw == "Call":
			panic("")
		case c.TokenKind == IDENT && c.Raw == "BinOp":
			panic("")
		case c.TokenKind == IDENT && c.Raw == "Add":
			p.add()
		case c.TokenKind == IDENT && c.Raw == "Store":
			p.store()
		case c.TokenKind == IDENT && c.Raw == "Load":
			p.load()
		case c.TokenKind == COMMA:
			p.write(",")
			p.consume(",")
		case c.TokenKind == RBR:
			// ?
			return // )
		case c.TokenKind == RSQB:
			// ?
			return // ]
		}
		//p.goNext()
	}
}

func (p *Parser) isEof() bool {
	return p.pos >= len(p.tokens)
}

func (p *Parser) curt() Token {
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
	p.write("{")
	p.write("\"type\":\"module\",")

	// body or type_ignores
loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == IDENT && c.Raw == "body":
			p.consume("body")
			p.consume("=")
			p.consume("[")
			p.write("\"body\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case c.TokenKind == IDENT && c.Raw == "type_ignores":
			p.consume("type_ignores")
			p.consume("=")
			p.consume("[")
			p.write("\"type_ignores\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case c.TokenKind == COMMA:
			p.write(",")
			p.consume(",")
		case c.TokenKind == WHITE:
			p.goNext()
		case c.TokenKind == RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.write("}")
	if p.curt().TokenKind != EOF {
		panic("expect eof, but found" + p.curt().TokenKind.String())
	}
	p.consume("")
}

func (p *Parser) classDef() {
	p.consume("ClassDef")
	p.consume("(")
	p.write("{")
	p.write("\"type\":\"classDef\",")

	// name, bases, keywords, body, decorator_list
loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == IDENT && c.Raw == "name":
			p.consume("name")
			p.consume("=")
			p.write(fmt.Sprintf("\"name\":%v", strconv.Quote(p.curt().Raw)))
			p.goNext() // string
		case c.TokenKind == IDENT && c.Raw == "bases":
			p.consume("bases")
			p.consume("=")
			p.consume("[")
			p.write("\"bases\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.write("]")
		case c.TokenKind == IDENT && c.Raw == "keywords":
			p.consume("keywords")
			p.consume("=")
			p.consume("[")
			p.write("\"keywords\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.write("]")
		case c.TokenKind == IDENT && c.Raw == "body":
			p.consume("body")
			p.consume("=")
			p.consume("[")
			p.write("\"body\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case c.TokenKind == IDENT && c.Raw == "decorator_list":
			p.consume("decorator_list")
			p.consume("=")
			p.consume("[")
			p.write("\"decorator_list\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.write("]")
		case c.TokenKind == COMMA:
			p.write(",")
			p.consume(",")
		case c.TokenKind == WHITE:
			p.goNext()
		case c.TokenKind == RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.write("}")
}

func (p *Parser) functionDef() {
	p.consume("FunctionDef")
	p.consume("(")
	p.write("{")
	p.write("\"type\":\"functionDef\",")

	// name, args, body, decorator_list
loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == IDENT && c.Raw == "name":
			p.consume("name")
			p.consume("=")
			p.write(fmt.Sprintf("\"name\":%v", strconv.Quote(p.curt().Raw)))
			p.goNext() // string
		case c.TokenKind == IDENT && c.Raw == "args":
			p.consume("args")
			p.consume("=")
			p.consume("[")
			p.write("\"args\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case c.TokenKind == IDENT && c.Raw == "body":
			p.consume("body")
			p.consume("=")
			p.consume("[")
			p.write("\"body\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case c.TokenKind == IDENT && c.Raw == "decorator_list":
			p.consume("decorator_list")
			p.consume("=")
			p.consume("[")
			p.write("\"decorator_list\":[")
			p.parse() // 中身不明
			p.consume("]")
			p.write("]")
		case c.TokenKind == COMMA:
			p.write(",")
			p.consume(",")
		case c.TokenKind == WHITE:
			p.goNext()
		case c.TokenKind == RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.write("}")
}

func (p *Parser) arguments() {
	p.consume("arguments")
	p.consume("(")
	p.write("{")
	p.write("\"type\":\"arguments\",")

	// posonlyargs, args, kwonlyargs, kw_defaults, defaults
	params := []string{"posonlyargs", "args", "kwonlyargs", "kw_defaults", "defaults"}
loop:
	for p.isEof() {
		c := p.curt()
		switch c.TokenKind {
		case IDENT:
			if !StringsContain(params, c.Raw) {
				panic("unexpected parameter: " + c.String())
			}
			p.consume(c.Raw)
			p.consume("=")
			p.consume("[")
			p.write("\"" + c.Raw + "\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case COMMA:
			p.write(",")
			p.consume(",")
		case WHITE:
			p.goNext()
		case RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.write("}")
}

func (p *Parser) arg() {
	p.consume("arg")
	p.consume("(")
	p.write("{")
	p.write("\"type\":\"arg\",")

loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == IDENT && c.Raw == "arg":
			p.consume("arg")
			p.consume("=")
			// string or number??
			p.write(fmt.Sprintf("\"arg\":%v", strconv.Quote(p.curt().Raw)))
			p.goNext()
		case c.TokenKind == IDENT && c.Raw == "annotation":
			p.consume(c.Raw)
			p.consume("=")
			p.consume("[")
			p.write("\"" + c.Raw + "\":[")
			p.parse()
			p.consume("]")
			p.write("]")
		case c.TokenKind == COMMA:
			p.write(",")
			p.consume(",")
		case c.TokenKind == WHITE:
			p.goNext()
		case c.TokenKind == RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.write("}")
}

func (p *Parser) name() {
	p.consume("Name")
	p.consume("(")
	p.write("{" + "\"type\":\"name\",")
loop:
	for !p.isEof() {
		c := p.curt()
		switch {
		case c.TokenKind == IDENT && c.Raw == "id":
			p.consume("id")
			p.consume("=")
			p.write(fmt.Sprintf("\"id\":%v", strconv.Quote(p.curt().Raw)))
			p.goNext() // string
		case c.TokenKind == IDENT && c.Raw == "ctx":
			p.consume("ctx")
			p.consume("=")
			p.write("\"ctx\":")
			p.parse()
		case c.TokenKind == COMMA:
			p.write(",")
			p.consume(",")
		case c.TokenKind == WHITE:
			p.goNext()
		case c.TokenKind == RBR:
			p.consume(")")
			break loop
		default:
			panic("syntax error: " + c.String())
		}
	}
	p.write("}")
}

func (p *Parser) store() {
	p.consume("Store")
	p.consume("(")
	p.consume(")")
	p.write("{\"type\":\"store\"}")
}

func (p *Parser) load() {
	p.consume("Load")
	p.consume("(")
	p.consume(")")
	p.write("{\"type\":\"load\"}")
}

func (p *Parser) add() {
	p.consume("Add")
	p.consume("(")
	p.consume(")")
	p.write("{\"type\":\"add\"}")
}
