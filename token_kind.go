package pylot

type TokenKind int

const (
	_ TokenKind = iota
	EOF
	IDENT  // keyword
	STRING // ""
	NUMBER // 0123456789
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
	NUMBER: "NUMBER",
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
