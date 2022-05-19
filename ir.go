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

	fmt.Printf("{")
	fmt.Printf("}\n")

	return nil
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

func cAttribute(m map[string]any) {}

func cAnnAssign(m map[string]any) (string, error) {
	// target : attribute
	// annotation : name
	// value : name
	// simple : int

	return "", nil
}
