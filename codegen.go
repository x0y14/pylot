package pylot

import (
	"encoding/json"
	"fmt"
)

type CodeGen struct {
	code    string
	nest    int
	defines DataDefines
}

func NewCodeGen() *CodeGen {
	return &CodeGen{
		code:    "",
		nest:    0,
		defines: DataDefines{},
	}
}

func (c *CodeGen) identInsert(code string) {
	c.code += fmt.Sprintf("%v%v", Indent(c.nest), code)
}
func (c *CodeGen) rawInsert(code string) {
	c.code += code
}

func (c *CodeGen) Disclosure() {
	fmt.Printf("```\n%v\n```", c.code)
}

func (c *CodeGen) Gen(jsons string) error {
	var module map[string]any
	if err := json.Unmarshal([]byte(jsons), &module); err != nil {
		return err
	}

	c.identInsert("; module: filename.ext\n")

	body, ok := module["body"]
	if !ok {
		return fmt.Errorf("module need body field")
	}

	for _, statement := range body.([]any) {
		_, err := c.statement("$MODULE", statement.(map[string]any))
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CodeGen) statement(group string, m map[string]any) (string, error) {
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("statement need type field")
	}

	switch typ {
	case "ClassDef":
		if err := c.classDef(group, m); err != nil {
			return "", err
		}
	case "FunctionDef":
		if err := c.functionDef(group, m); err != nil {
			return "", err
		}
	case "Constant":
		return c.constant(m)
	case "Name":
		return c.name(m)
	case "AnnAssign":
		return c.annAssign(m)
	default:
		return "", fmt.Errorf("unsupported statement: %v", typ)
	}

	return "", nil
}

func (c *CodeGen) classDef(group string, m map[string]any) error {
	name, ok := m["name"].(string)
	if !ok {
		return fmt.Errorf("classDef need name field")
	}
	c.identInsert(fmt.Sprintf("; class: %v\n", name))

	groupName := fmt.Sprintf("%v.%v", group, name)
	if err := c.defines.set(groupName, DataDefine{
		//Scope: PUBLIC,
		Types: CLASS,
	}); err != nil {
		return err
	}

	body, ok := m["body"].([]any)
	if !ok {
		return fmt.Errorf("classDef need body field")
	}

	//c.nest++

	for _, statement := range body {
		_, err := c.statement(groupName, statement.(map[string]any))
		if err != nil {
			return err
		}
	}

	//c.nest--

	return nil
}

func (c *CodeGen) functionDef(group string, m map[string]any) error {
	name, ok := m["name"].(string)
	if !ok {
		return fmt.Errorf("functionDef need name field")
	}

	groupName := fmt.Sprintf("%v.%v", group, name)

	var retType string
	retTypeM, ok := m["returns"].(map[string]any)
	if !ok {
		retType = "null"
	} else {
		retType_, err := c.statement(groupName, retTypeM)
		if err != nil {
			return err
		}
		retType = retType_
	}

	// 宣言を記憶
	if err := c.defines.set(groupName, DataDefine{
		//Scope: scope,
		RetTypes: ToTypes(retType),
		Types:    FUNCTION,
	}); err != nil {
		return err
	}

	// 引数解析
	var argumentStr string
	argsM, ok := m["args"].(map[string]any)
	if !ok {
		// 引数なし
		argumentStr = ""
	}
	arguments, err := c.arguments(argsM)
	if err != nil {
		return err
	}
	argumentStr = arguments

	// 出力1
	c.identInsert(fmt.Sprintf("define %v @%v.%v", retType, group, name))
	c.rawInsert("(")
	c.rawInsert(argumentStr)
	c.rawInsert(")")
	c.rawInsert(" {\n")

	// body解析
	c.nest++
	body, ok := m["body"].([]any)
	if !ok {
		return fmt.Errorf("functionDef need body field")
	}
	for _, statementM := range body {
		statement, err := c.statement(groupName, statementM.(map[string]any))
		if err != nil {
			return err
		}
		c.identInsert(statement + "\n")
	}
	c.nest--
	c.identInsert("}\n")

	return nil
}

func (c *CodeGen) constant(m map[string]any) (string, error) {
	value, ok := m["value"]
	if !ok {
		return "", fmt.Errorf("constant need value field")
	}

	if value == nil {
		return "null", nil
	}

	return "", fmt.Errorf("unsupported constant: %v", value)
}

func (c *CodeGen) name(m map[string]any) (string, error) {
	id, ok := m["id"].(string)
	if !ok {
		return "", fmt.Errorf("name need id field")
	}

	return id, nil
}

func (c *CodeGen) arguments(m map[string]any) (string, error) {
	argumentsStr := ""

	args, ok := m["args"].([]any)
	if !ok {
		return "", fmt.Errorf("arguments need args field")
	}

	for i, arg := range args {
		argStr, err := c.arg(arg.(map[string]any))
		if err != nil {
			return "", err
		}
		argumentsStr += argStr
		if i != len(args)-1 {
			argumentsStr += ", "
		}
	}

	return argumentsStr, nil
}

func (c *CodeGen) arg(m map[string]any) (string, error) {
	name, ok := m["arg"].(string)
	if !ok {
		return "", fmt.Errorf("arg need arg field")
	}

	var typ string
	annotation, ok := m["annotation"].(map[string]any)
	if ok {
		annTyp, err := c.annotation(annotation)
		if err != nil {
			return "", err
		}
		typ = annTyp
	} else {
		typ = "unknown"
	}

	return fmt.Sprintf("%v %%%v", typ, name), nil
}

func (c *CodeGen) annotation(m map[string]any) (string, error) {
	return c.name(m)
}

func (c *CodeGen) attribute(m map[string]any) (string, error) {
	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("attribute need value field")
	}

	name, err := c.name(valueM)
	if err != nil {
		return "", err
	}

	attr, ok := m["attr"].(string)
	if !ok {
		return "", fmt.Errorf("attribute need attr field")
	}

	return fmt.Sprintf("%v.%v", name, attr), nil
}

func (c *CodeGen) annAssign(m map[string]any) (string, error) {
	var annotation string
	annotationM, ok := m["annotation"].(map[string]any)
	if !ok {
		//return "", fmt.Errorf("annAssign need annotation field")
		annotation = "unknown"
	}
	ann, err := c.annotation(annotationM) // ?
	if err != nil {
		return "", err
	}
	annotation = ann

	targetM, ok := m["target"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("annAssign need target field")
	}
	target, err := c.attribute(targetM)
	if err != nil {
		return "", err
	}

	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("annAssign need value field")
	}
	value, err := c.name(valueM)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v %v = %v", annotation, target, value), nil
}

func (c *CodeGen) expr(m map[string]any) (string, error) {
	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("expr need value field")
	}

	return c.call(valueM)
}

func (c *CodeGen) call(m map[string]any) (string, error) {
	name, ok := m["func"]
	if !ok {
		return "", fmt.Errorf("call need func field")
	}

	return "", nil
}
