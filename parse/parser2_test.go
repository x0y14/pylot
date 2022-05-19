package parse

import (
	"github.com/stretchr/testify/assert"
	"pylot/tokenize"
	"testing"
)

func TestParser2_Parse(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect string
	}{
		{
			"empty module",
			"Module(body=[], type_ignores=[])",
			`{"type":"Module","body":[],"type_ignores":[]}`,
		},
		{
			"class-def in module",
			"Module(body=[ClassDef(name=\"john\")], type_ignores=[])",
			`{"type":"Module","body":[{"type":"ClassDef","name":"john"}],"type_ignores":[]}`,
		},
		{
			"class-def in module",
			"Module(body=[ClassDef(name=\"john\"),ClassDef(name=\"tom\")], type_ignores=[])",
			`{"type":"Module","body":[{"type":"ClassDef","name":"john"},{"type":"ClassDef","name":"tom"}],"type_ignores":[]}`,
		},
	}

	tokenizer := tokenize.NewTokenizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, V2(tokens))
		})
	}
}
