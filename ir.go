package pylot

import (
	"encoding/json"
	"fmt"
)

func Ir(j string) error {
	var maybeModuleDef map[string]any
	if err := json.Unmarshal([]byte(j), &maybeModuleDef); err != nil {
		return err
	}

	// type check
	typ, ok := maybeModuleDef["type"].(string)
	if !ok {
		return fmt.Errorf("moduleDef need type field")
	}
	if typ != "Module" {
		return fmt.Errorf("expect module but found: %v", typ)
	}

	fmt.Println("Module")

	// body
	bod, ok := maybeModuleDef["body"].([]any)
	if !ok {
		return fmt.Errorf("moduleDef need body field")
	}
	for _, b := range bod {
		bTyp, ok := b.(map[string]any)["type"].(string)
		if !ok {
			return fmt.Errorf("in modeuleDef.body, anyone need type field")
		}
		if bTyp == "ClassDef" {
			if err := cClassDef(b.(map[string]any)); err != nil {
				return err
			}
		} else {
			panic("unsupported")
		}
		// グローバル関数のみ
		// 直接実行スクリプト printとか
	}

	return nil
}

func cClassDef(m map[string]any) error {
	name, ok := m["name"].(string)
	if !ok {
		return fmt.Errorf("classDef need name field")
	}
	fmt.Printf("Class: %v\n", name)

	body, ok := m["body"].([]any)
	if !ok {
		return fmt.Errorf("classDef need body field")
	}

	for _, content := range body {
		contentTyp, ok := content.(map[string]any)["type"].(string)
		if !ok {
			return fmt.Errorf("in classDef.body, anyone need type field")
		}
		if contentTyp == "FunctionDef" {
			if err := cFunctionDef(name, content.(map[string]any)); err != nil {
				return err
			}
		} else {
			panic("unsupported")
		}
	}

	return nil
}

func cFunctionDef(className string, m map[string]any) error {
	name, ok := m["name"].(string)
	if !ok {
		return fmt.Errorf("functionDef need name field")
	}

	// return check
	// 定数ならcConstant、それ以外はそもそも存在しない場合(void)、Name()が入る場合がある
	var returnType string
	rtInfo, ok := m["returns"].(map[string]any)
	if !ok {
		// それ以外はそもそも存在しない場合(void)
		returnType = "void"
	} else {
		rt_, err := analyzeRtInfo(rtInfo)
		if err != nil {
			return err
		}
		returnType = rt_
	}

	fmt.Printf("define %v @%v.%v(", returnType, className, name)

	// param
	args, ok := m["args"].(map[string]any)
	if !ok {
		panic("unsupported")
	}
	arguments, err := cArguments(args)
	if err != nil {
		return err
	}
	fmt.Printf("%v", arguments)
	fmt.Printf(") ")

	fmt.Printf("{\n")

	// body
	bodyM, ok := m["body"]
	if !ok {
		panic("unsupported")
	}
	if err := analyzeBodyExpressions(bodyM.([]any)); err != nil {
		return err
	}

	fmt.Printf("}\n")
	return nil
}

func analyzeBodyExpressions(m []any) error {
	//fmt.Printf("\n")
	for _, exprLine := range m {
		_, err := analyzeBodyExpression(exprLine.(map[string]any))
		if err != nil {
			return err
		}
	}

	return nil
}

func analyzeBodyExpression(m map[string]any) (string, error) {
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("function body expr line need type field")
	}
	switch typ {
	case "Expr":
		if err := cExpr(m); err != nil {
			return "", err
		}
	case "AnnAssign", "Assign":
		if err := cAnnAssign(m); err != nil {
			return "", err
		}
	case "BinOp":
		s, err := cBinOp(m)
		if err != nil {
			return s, err
		}
		fmt.Printf("%v", s)
	default:
		panic("unsupported " + typ)
	}
	return "", nil
}

func analyzeRtInfo(m map[string]any) (string, error) {
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("returns need type field")
	}
	if typ == "Name" {
		return cName(m)
	} else if typ == "Constant" {
		return cConstant(m)
	} else {
		panic("unsupported")
	}
}

// 定数
func cConstant(m map[string]any) (string, error) {
	val, ok := m["value"]
	if !ok {
		return "", fmt.Errorf("constatn need value field")
	}

	if val == nil {
		// null
		return "void", nil
	}

	// ?
	panic("unsupported")
	return "", nil
}

// ネームスペース？
func cName(m map[string]any) (string, error) {
	id, ok := m["id"].(string)
	if !ok {
		return "", fmt.Errorf("name need id field")
	}

	return id, nil
}

//
func cArguments(m map[string]any) (string, error) {
	arguments := ""

	args, ok := m["args"].([]any)
	if !ok {
		return "", fmt.Errorf("arguments need args field")
	}
	for i, arg := range args {
		argStr, err := cArg(arg.(map[string]any))
		if err != nil {
			return "", err
		}
		arguments += argStr
		// 表示用コンマ
		if i != len(args)-1 {
			arguments += ", "
		}
	}
	return arguments, nil
}

func cArg(m map[string]any) (string, error) {
	argName, ok := m["arg"].(string)
	if !ok {
		return "", fmt.Errorf("arg need arg field")
	}
	var argType string
	annotation, ok := m["annotation"].(map[string]any)
	if ok {
		aTyp, err := cAnnotation(annotation)
		if err != nil {
			return "", err
		}
		argType = aTyp
	} else {
		argType = "unknown"
	}

	return fmt.Sprintf("%v %%%v", argType, argName), nil
}

func cAnnotation(m map[string]any) (string, error) {
	typ, err := cName(m)
	if err != nil {
		return "", err
	}
	return typ, nil
}

func cAttribute(m map[string]any) (string, error) {
	// value

	nameM, ok := m["value"]
	if !ok {
		panic("attribute need value field")
	}

	name, err := cName(nameM.(map[string]any))
	if err != nil {
		return "", err
	}
	// attr
	attr, ok := m["attr"].(string)
	if !ok {
		return "", fmt.Errorf("attribute need attr field")
	}

	return fmt.Sprintf("%v.%v", name, attr), nil
}

func cAnnAssign(m map[string]any) error {
	// [annotation, target, value]
	//var result [3]string

	// annotation : name
	var annotation string
	annotationM, ok := m["annotation"].(map[string]any)
	if !ok {
		annotation = "unknown"
		//return nil, fmt.Errorf("annAssign need annotation field")
	}
	anno, err := cName(annotationM)
	annotation = anno

	// target : attribute
	attrM, ok := m["target"].(map[string]any)
	if !ok {
		return fmt.Errorf("annAssign need target field")
	}
	target, err := cAttribute(attrM)
	if err != nil {
		return err
	}

	// value : name
	valueM, ok := m["value"].(map[string]any)
	if !ok {
		return fmt.Errorf("annAssign need value field")
	}
	value, err := cName(valueM)
	if err != nil {
		return err
	}

	// simple : int ?

	//result[0] = annotation
	//result[1] = target
	//result[2] = value
	fmt.Printf("  %v: %v = %v\n", target, annotation, value)

	return nil
}

func cExpr(m map[string]any) error {
	valueM, ok := m["value"]
	if !ok {
		return fmt.Errorf("expr need value field")
	}

	return cCall(valueM.(map[string]any))
}

func cCall(m map[string]any) error {
	funcNameM, ok := m["func"]
	if !ok {
		return fmt.Errorf("")
	}
	funcName, err := cName(funcNameM.(map[string]any))
	if err != nil {
		return err
	}

	fmt.Printf("  call $type @%v(", funcName)
	funcArgs, ok := m["args"].([]any)
	if ok {
		err = analyzeBodyExpressions(funcArgs)
		if err != nil {
			return err
		}
	}
	fmt.Printf(")\n")
	return nil
}

func cOp(m map[string]any) (string, error) {
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("op need type field")
	}

	return typ, nil
}

func cValue(m map[string]any) (string, error) {
	// cAttribute
	// cConstant
	typ, ok := m["type"].(string)
	if !ok {
		return "", fmt.Errorf("$value need type field")
	}

	switch typ {
	case "Attribute":
		return cAttribute(m)
	case "Constant":
		return cConstant(m)
	default:
		fmt.Printf("cValueType: %v\n", typ)
		panic("unsupported")
	}

	return "", nil
}

func cBinOp(m map[string]any) (string, error) {
	leftM, ok := m["left"]
	if !ok {
		return "", fmt.Errorf("binOp need left field")
	}
	left, err := cValue(leftM.(map[string]any))
	if err != nil {
		return "", err
	}

	opM, ok := m["op"]
	if !ok {
		return "", fmt.Errorf("binOp need op field")
	}
	op, err := cOp(opM.(map[string]any))
	if err != nil {
		return "", err
	}

	rightM, ok := m["right"]
	if !ok {
		return "", fmt.Errorf("binOp need right field")
	}
	right, err := cValue(rightM.(map[string]any))
	if err != nil {
		return "", err
	}

	//fmt.Printf("(%v, %v, %v)\n", op, left, right)
	return fmt.Sprintf("(%v, %v, %v)\n", op, left, right), nil
}
