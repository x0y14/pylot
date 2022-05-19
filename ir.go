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

	fmt.Printf("%v @%v.%v\n", returnType, className, name)
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
