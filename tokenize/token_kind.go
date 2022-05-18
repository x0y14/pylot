package tokenize

type TokenKind int

const (
	_ TokenKind = iota
	EOF
	IDENT  // keyword
	STRING // ""
	LBR    // (
	RBR    // )
	LSQB   // [
	RSQB   // ]
	COMMA  // ,
	EQU    // =
	WHITE  // " "
)

var kinds = [...]string{
	IDENT:  "IDENT",
	STRING: "STRING",
	LBR:    "LBR",
	RBR:    "RBR",
	LSQB:   "LSQB",
	RSQB:   "RSQB",
	COMMA:  "COMMA",
	EQU:    "EQU",
	WHITE:  "WHITE",
}

func (k TokenKind) String() string {
	return kinds[k]
}
