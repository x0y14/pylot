package tokenize

import (
	"fmt"
	"testing"
)

func TestTokenizer_Tokenize(t *testing.T) {
	tokenizer := NewTokenizer()
	tokens := tokenizer.Tokenize("print(\"hello\")\n")
	for _, tk := range tokens {
		fmt.Println(tk.String())
	}
}
