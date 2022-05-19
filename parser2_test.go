package pylot

import (
	"github.com/stretchr/testify/assert"
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
		{
			"real",
			`Module(
    body=[
        ClassDef(
            name="Name",
            bases=[],
            keywords=[],
            body=[
                FunctionDef(
                    name="__init__",
                    args=arguments(
                        posonlyargs=[],
                        args=[
                            arg(arg="self"),
                            arg(arg="first", annotation=Name(id="str", ctx=Load())),
                            arg(arg="middle", annotation=Name(id="str", ctx=Load())),
                            arg(arg="last", annotation=Name(id="str", ctx=Load())),
                        ],
                        kwonlyargs=[],
                        kw_defaults=[],
                        defaults=[],
                    ),
                    body=[
                        AnnAssign(
                            target=Attribute(
                                value=Name(id="self", ctx=Load()),
                                attr="first",
                                ctx=Store(),
                            ),
                            annotation=Name(id="str", ctx=Load()),
                            value=Name(id="first", ctx=Load()),
                            simple=0,
                        ),
                        AnnAssign(
                            target=Attribute(
                                value=Name(id="self", ctx=Load()),
                                attr="middle",
                                ctx=Store(),
                            ),
                            annotation=Name(id="str", ctx=Load()),
                            value=Name(id="middle", ctx=Load()),
                            simple=0,
                        ),
                        AnnAssign(
                            target=Attribute(
                                value=Name(id="self", ctx=Load()),
                                attr="last",
                                ctx=Store(),
                            ),
                            annotation=Name(id="str", ctx=Load()),
                            value=Name(id="last", ctx=Load()),
                            simple=0,
                        ),
                    ],
                    decorator_list=[],
                ),
                FunctionDef(
                    name="to_s",
                    args=arguments(
                        posonlyargs=[],
                        args=[arg(arg="self")],
                        kwonlyargs=[],
                        kw_defaults=[],
                        defaults=[],
                    ),
                    body=[
                        Expr(
                            value=Call(
                                func=Name(id="print", ctx=Load()),
                                args=[
                                    BinOp(
                                        left=BinOp(
                                            left=BinOp(
                                                left=BinOp(
                                                    left=Attribute(
                                                        value=Name(
                                                            id="self", ctx=Load()
                                                        ),
                                                        attr="first",
                                                        ctx=Load(),
                                                    ),
                                                    op=Add(),
                                                    right=Constant(value=" "),
                                                ),
                                                op=Add(),
                                                right=Attribute(
                                                    value=Name(id="self", ctx=Load()),
                                                    attr="middle",
                                                    ctx=Load(),
                                                ),
                                            ),
                                            op=Add(),
                                            right=Constant(value=" "),
                                        ),
                                        op=Add(),
                                        right=Attribute(
                                            value=Name(id="self", ctx=Load()),
                                            attr="last",
                                            ctx=Load(),
                                        ),
                                    )
                                ],
                                keywords=[],
                            )
                        )
                    ],
                    decorator_list=[],
                    returns=Constant(value=None),
                ),
            ],
            decorator_list=[],
        )
    ],
    type_ignores=[],
)
`,
			"",
		},
	}

	tokenizer := NewTokenizer()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(tt.in, true)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, V2(tokens))
		})
	}
}
