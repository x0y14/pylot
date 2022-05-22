package pylot

import "fmt"

type DataDefine struct {
	//Scope // いらない気がする
	Types
	RetTypes Types
}

type DataDefines map[string]DataDefine

func (d DataDefines) set(name string, define DataDefine) error {
	_, ok := d[name]
	if ok {
		return fmt.Errorf("%v is already defined", name)
	}
	d[name] = define
	return nil
}

func (d DataDefines) get(name string) (DataDefine, error) {
	v, ok := d[name]
	if !ok {
		return DataDefine{}, fmt.Errorf("%v is undefined", name)
	}
	return v, nil
}
