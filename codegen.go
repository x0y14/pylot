package pylot

import (
	"encoding/json"
	"fmt"
	types2 "go/types"
	"strconv"
)

type CodeGen struct {
	code    string
	nest    int
	defines DataDefines
	//typeMap map[string]any
	group string
}

func NewCodeGen() *CodeGen {
	return &CodeGen{
		code:    "",
		nest:    0,
		defines: DataDefines{},
		//typeMap: map[string]any{},
		group: "",
	}
}
func (c *CodeGen) setGroupName(g string) {
	c.group = g
}
func (c *CodeGen) getGroupName() string {
	return c.group
}
func (c *CodeGen) concatGroupName(newLast string) {
	c.group += "." + newLast
}

func (c *CodeGen) identInsert(code string) {
	c.code += fmt.Sprintf("%v%v", Indent(c.nest), code)
}
func (c *CodeGen) rawInsert(code string) {
	c.code += code
}

func (c *CodeGen) Disclosure() {
	fmt.Printf("```\n%v\n```\n", c.code)
}

func (c *CodeGen) Gen(jsons string) error {
	var module map[string]any
	if err := json.Unmarshal([]byte(jsons), &module); err != nil {
		return err
	}

	c.identInsert("; module: filename.ext\n")
	c.setGroupName("undefined")

	body, ok := module["body"]
	if !ok {
		return fmt.Errorf("module need body field")
	}

	for _, statement := range body.([]any) {
		_, err := c.statement(statement.(map[string]any))
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CodeGen) statement(m map[string]any) (string, error) {
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("statement need type field")
	}

	switch typ {
	case "ClassDef":
		if err := c.classDef(m); err != nil {
			return "", err
		}
	case "FunctionDef":
		if err := c.functionDef(m); err != nil {
			return "", err
		}
	case "Constant":
		return c.constant(m)
	case "Name":
		return c.name(m)
	case "Attribute":
		return c.attribute(m)
	case "AnnAssign":
		return c.annAssign(m)
	case "Assign":
		return c.assign(m)
	case "Expr":
		return c.expr(m)
	case "BinOp":
		return c.binOp(m)
	case "Return":
		return c.return_(m)
	default:
		return "", fmt.Errorf("unsupported statement: %v", typ)
	}

	return "", nil
}

func (c *CodeGen) classDef(m map[string]any) error {
	name, ok := m["name"].(string)
	if !ok {
		return fmt.Errorf("classDef need name field")
	}
	c.identInsert(fmt.Sprintf("; class: %v\n", name))

	oldGroupName := c.getGroupName()
	c.concatGroupName(name)
	if err := c.defines.set(c.getGroupName(), DataDefine{
		//Scope: PUBLIC,
		Self: ObjectDataType{
			CLASS,
			"",
		},
		Ret: ObjectDataType{
			Types: UNDEFINED,
			Raw:   "",
		},
	}); err != nil {
		return err
	}

	body, ok := m["body"].([]any)
	if !ok {
		return fmt.Errorf("classDef need body field")
	}

	for _, statement := range body {
		_, err := c.statement(statement.(map[string]any))
		if err != nil {
			return err
		}
	}

	c.setGroupName(oldGroupName)

	return nil
}

func (c *CodeGen) functionDef(m map[string]any) error {
	name, ok := m["name"].(string)
	if !ok {
		return fmt.Errorf("functionDef need name field")
	}

	oldGroupName := c.getGroupName()
	c.concatGroupName(name)

	var retType string
	retTypeM, ok := m["returns"].(map[string]any)
	if !ok {
		retType = "null"
	} else {
		retType_, err := c.statement(retTypeM)
		if err != nil {
			return err
		}
		retType = retType_
	}

	// 宣言を記憶
	if err := c.defines.set(c.getGroupName(), DataDefine{
		Self: ObjectDataType{
			Types: FUNCTION,
			Raw:   "",
		},
		Ret: ObjectDataType{
			Types: ToTypes(retType),
			Raw:   retType,
		},
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
	c.identInsert(fmt.Sprintf("define %v @%v", ToTypes(retType).String(), c.group))
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
		statement, err := c.statement(statementM.(map[string]any))
		if err != nil {
			return err
		}
		c.identInsert(statement + "\n")
	}
	c.nest--
	c.identInsert("}\n\n")

	c.setGroupName(oldGroupName)

	return nil
}

// 定数
func (c *CodeGen) constant(m map[string]any) (string, error) {
	// これ、定数。これだけでは型がわからない。
	// 推論するしかない。?
	// 呼び出し元がわかれば良い。 -> group..?

	//var typ string
	var val string

	value, ok := m["value"]
	if !ok {
		return "", fmt.Errorf("constant need value field")
	}

	switch value.(type) {
	case nil, types2.Nil:
		val = "null"
	case float32, float64, int:
		val = fmt.Sprintf("%v", value)
	case string:
		val = strconv.Quote(value.(string))
	}

	//typDef, ok := c.defines.get(c.getGroupName())
	//if !ok {
	//	typ = UNKNOWN.String()
	//} else {
	//	if typDef.Ret.Types != UNDEFINED {
	//		typ = typDef.Ret.String()
	//	} else {
	//		typ = typDef.Self.String()
	//	}
	//}

	return fmt.Sprintf("%v", val), nil
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

	oldGroupName := c.getGroupName()
	c.concatGroupName(name)

	var typ string
	annotation, ok := m["annotation"].(map[string]any)
	if ok {
		annTyp, err := c.annotation(annotation)
		if err != nil {
			return "", err
		}
		typ = annTyp
	} else {
		typ = UNKNOWN.String()
	}

	//typ = "&&" + typ

	err := c.defines.set(c.getGroupName(), DataDefine{
		Self: ObjectDataType{
			Types: ToTypes(typ),
			Raw:   typ,
		},
		Ret: ObjectDataType{
			Types: UNDEFINED,
			Raw:   "",
		},
	})
	if err != nil {
		return "", err
	}
	c.setGroupName(oldGroupName)

	return fmt.Sprintf("%v %%%v", ToTypes(typ).String(), name), nil
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

	// groupNameに.が２つ入っていたら、クラス内の関数だとわかる + selfだった場合 -> groupName + attrで定義されてる可能性あり
	// (selfは未使用推奨なのでselfだけで良い気がするが安全のため)

	//typ := UNKNOWN.String()
	//if strings.Count(c.getGroupName(), ".") == 2 && name == "self" {
	//	t, ok := c.defines.get(c.getGroupName() + "." + attr)
	//	if ok {
	//		typ = t.Self.String()
	//	}
	//}

	return fmt.Sprintf("%v.%v", name, attr), nil
}

func (c *CodeGen) annAssign(m map[string]any) (string, error) {
	//var annotation string
	//annotationM, ok := m["annotation"].(map[string]any)
	//if !ok {
	//	//return "", fmt.Errorf("annAssign need annotation field")
	//	annotation = UNKNOWN.String()
	//} else {
	//	ann, err := c.annotation(annotationM) // ?
	//	if err != nil {
	//		return "", err
	//	}
	//	annotation = ann
	//}

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

	// value annotation
	valueTyp := UNKNOWN.String()
	vt, ok := c.defines.get(c.getGroupName() + "." + value)
	if ok {
		valueTyp = vt.Self.String()
	}

	return fmt.Sprintf("%v = %v %v", target, valueTyp, value), nil
}

func (c *CodeGen) assign(m map[string]any) (string, error) {
	// s1, s2 = "", "" => (s1, s2) = ("", "")
	// targetsになっているのに１つしか入っていない、なぜ複数か不明。

	// pythonは複数の値の代入に対応しているため、
	// "target"ではなく"targets"になってる思われる。
	// んで、annAssignを流用できない。

	//// 型推論が必要、とりあえずプレースホルダー
	//annotation := UNKNOWN.String()

	targets := ""
	targetArr, ok := m["targets"].([]any)
	if !ok {
		return "", fmt.Errorf("assign need targets field")
	}
	for i, targetM := range targetArr {
		target, err := c.attribute(targetM.(map[string]any))
		if err != nil {
			return "", err
		}
		targets += target
		if i != len(targetArr)-1 {
			targets += ", "
		}
	}

	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("assign need value field")
	}
	value, err := c.name(valueM)
	if err != nil {
		return "", err
	}

	// value annotation
	valueTyp := UNKNOWN.String()
	vt, ok := c.defines.get(c.getGroupName() + "." + value)
	if ok {
		valueTyp = vt.Self.String()
	}

	return fmt.Sprintf("%v = %v %v", targets, valueTyp, value), nil
}

func (c *CodeGen) expr(m map[string]any) (string, error) {
	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("expr need value field")
	}

	return c.call(valueM)
}

func (c *CodeGen) call(m map[string]any) (string, error) {
	nameM, ok := m["func"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("call need func field")
	}
	name, err := c.name(nameM)
	if err != nil {
		return "", err
	}

	//fmt.Printf("  call $type @%v(", funcName)
	fArgs := ""
	funcArgs, ok := m["args"].([]any)
	if ok {
		for i, fArgM := range funcArgs {
			arg, err := c.statement(fArgM.(map[string]any))
			if err != nil {
				return "", err
			}
			fArgs += arg
			if i != len(funcArgs)-1 {
				fArgs += ", "
			}
		}
	}

	callRetTyp := UNKNOWN.String()

	return fmt.Sprintf("call %v %v(%v)", callRetTyp, name, fArgs), nil
}

func (c *CodeGen) binOp(m map[string]any) (string, error) {
	leftM, ok := m["left"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("binOp need left field")
	}
	left, err := c.value(leftM)
	if err != nil {
		return "", err
	}

	opM, ok := m["op"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("binOp need op field")
	}
	alphabetOp, err := c.op(opM)
	if err != nil {
		return "", err
	}
	symbolOp, err := c.alphabetOp2symbolOp(alphabetOp)
	if err != nil {
		return "", err
	}

	rightM, ok := m["right"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("binOp need op field")
	}
	right, err := c.value(rightM)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%v, %v, %v)", symbolOp, left, right), nil
}

func (c *CodeGen) alphabetOp2symbolOp(op string) (string, error) {
	switch op {
	case "Add":
		return "+", nil
	default:
		return "", fmt.Errorf("unsupported op: %v", op)
	}
}

func (c *CodeGen) value(m map[string]any) (string, error) {
	return c.statement(m)
}

func (c *CodeGen) op(m map[string]any) (string, error) {
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("op need type field")
	}
	return typ, nil
}

func (c *CodeGen) return_(m map[string]any) (string, error) {
	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return "", fmt.Errorf("return need value field")
	}

	stmt, err := c.statement(valueM)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("return %v", stmt), nil
}
