package pylot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizer_Tokenize(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect []Token
	}{
		{
			"[code] print hello dq",
			"print(\"hello\")\n",
			[]Token{
				{IDENT, "print", 0, 5},
				{LBR, "(", 5, 6},
				{STR, "hello", 6, 13},
				{RBR, ")", 13, 14},
				{WHITE, "\n", 14, 15},
				{EOF, "", 15, 15},
			},
		},
		{
			"[code] print hello sq",
			"print('hello')\n",
			[]Token{
				{IDENT, "print", 0, 5},
				{LBR, "(", 5, 6},
				{STR, "hello", 6, 13},
				{RBR, ")", 13, 14},
				{WHITE, "\n", 14, 15},
				{EOF, "", 15, 15},
			},
		},
		{
			"[code] print hello dq sq",
			"print('\"hello\"')\n",
			[]Token{
				{IDENT, "print", 0, 5},
				{LBR, "(", 5, 6},
				{STR, "\"hello\"", 6, 15},
				{RBR, ")", 15, 16},
				{WHITE, "\n", 16, 17},
				{EOF, "", 17, 17},
			},
		},
		{
			"[ast] module empty",
			`Module(body=[], type_ignores=[])`,
			[]Token{
				{IDENT, "Module", 0, 6},
				{LBR, "(", 6, 7},
				{IDENT, "body", 7, 11},
				{EQU, "=", 11, 12},
				{LSQB, "[", 12, 13},
				{RSQB, "]", 13, 14},
				{COMMA, ",", 14, 15},
				{WHITE, " ", 15, 16},
				{IDENT, "type_ignores", 16, 28},
				{EQU, "=", 28, 29},
				{LSQB, "[", 29, 30},
				{RSQB, "]", 30, 31},
				{RBR, ")", 31, 32},
				{EOF, "", 32, 32},
			},
		},
	}

	tokenizer := NewTokenizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := tokenizer.Tokenize(tt.in, false)
			if err != nil {
				t.Fatal(err)
			}
			//for _, tk := range actual {
			//	fmt.Println(tk.String())
			//}
			assert.Equal(t, tt.expect, actual)
		})
	}
}
