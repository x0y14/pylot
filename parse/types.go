package parse

type Module struct {
	Body        []any `json:"body"`
	TypeIgnores []any `json:"type_ignores"`
}

type ClassDef struct {
	Name          string `json:"name"`
	Bases         []any  `json:"bases"`
	Keywords      []any  `json:"keywords"`
	Body          []any  `json:"body"`
	DecoratorList []any  `json:"decorator_list"`
}

type FunctionDef struct {
	Name          string    `json:"name"`
	Args          Arguments `json:"args"`
	Body          []any     `json:"body"`
	DecoratorList []any     `json:"decorator_list"`
}

type Arguments struct {
	Args        []Arg `json:"args"`
	Posonlyargs []any `json:"posonlyargs"`
	Kwonlyargs  []any `json:"kwonlyargs"`
	KwDefaults  []any `json:"kw_defaults"`
	Defaults    []any `json:"defaults"`
}

type Arg struct {
	Arg        string `json:"arg"`
	Annotation Name   `json:"annotation"`
}

type Name struct {
	Id  string `json:"id"`
	Ctx string `json:"ctx"`
}

type Attribute struct {
	Value Name   `json:"value"`
	Attr  string `json:"attr"`
	Ctx   string `json:"ctx"`
}

type Expr struct {
	Value any `json:"value"`
}

type Constant struct {
	Value string `json:"value"`
}

type Call struct {
	Func     Name  `json:"func"`
	Args     []any `json:"args"`
	Keywords []any `json:"keywords"`
}

type AnnAssign struct {
	Target     Attribute `json:"target"`
	Annotation Name      `json:"annotation"`
	Value      Name      `json:"value"`
	Simple     int       `json:"simple"`
}

type BinOp struct {
	Left  any `json:"left"`
	Op    any `json:"op"`
	Right any `json:"right"`
}

type Add struct{}
