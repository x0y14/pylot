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
	fmt.Printf("@%v.%v\n", className, name)
	return nil
}
