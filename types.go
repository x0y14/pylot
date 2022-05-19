package pylot

type ModuleDef struct {
	Type        string `json:"type,omitempty"`
	Body        []any  `json:"body,omitempty"`
	TypeIgnores []any  `json:"type_ignores,omitempty"`
}

type ClassDef struct {
	Type          string `json:"type,omitempty"`
	Name          string `json:"name,omitempty"`
	Bases         []any  `json:"bases,omitempty"`
	Keywords      []any  `json:"keywords,omitempty"`
	Body          []any  `json:"body,omitempty"`
	DecoratorList []any  `json:"decorator_list,omitempty"`
}

type FunctionDef struct {
	Type          string       `json:"type,omitempty"`
	Name          string       `json:"name,omitempty"`
	Args          ArgumentsDef `json:"args,omitempty"`
	Body          []any        `json:"body,omitempty"`
	DecoratorList []any        `json:"decorator_list,omitempty"`
}

type ArgumentsDef struct {
	Type        string   `json:"type,omitempty"`
	Posonlyargs []any    `json:"posonlyargs,omitempty"`
	Args        []ArgDef `json:"args,omitempty"`
	Kwonlyargs  []any    `json:"kwonlyargs,omitempty"`
	KwDefaults  []any    `json:"kw_defaults,omitempty"`
	Defaults    []any    `json:"defaults,omitempty"`
}

type ArgDef struct {
	Type       string        `json:"type,omitempty"`
	Arg        string        `json:"arg,omitempty"`
	Annotation AnnotationDef `json:"annotation"`
}

type AnnotationDef NameDef

type NameDef struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
	Ctx  CtxDef `json:"ctx"`
}

type CtxDef struct {
	Type string `json:"type,omitempty"`
}

type TypeOnly struct {
	Type string `json:"type,omitempty"`
}
