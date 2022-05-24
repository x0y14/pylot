package pylot

import "fmt"

type DataDefine struct {
	//Scope // いらない気がする
	Types
	RetTypes Types
}

func (d DataDefine) String() string {
	return fmt.Sprintf("{ Types:%15v, RetTypes:%15v }", d.Types.String(), d.RetTypes.String())
}

type DataDefines map[string]DataDefine

func (d DataDefines) set(name string, define DataDefine) error {
	_, ok := d[name]
	if ok {
		return fmt.Errorf("define %v is already defined", name)
	}
	d[name] = define
	return nil
}

func (d DataDefines) get(name string) (DataDefine, bool) {
	v, ok := d[name]
	//if !ok {
	//	return DataDefine{}, fmt.Errorf("define %v is undefined", name)
	//}
	//return v, nil
	return v, ok
}

func (d DataDefines) String() string {
	fmt.Println("<DataDefines>")
	defs := ""
	for key, def := range d {
		defs += fmt.Sprintf("%-30v%v\n", key, def.String())
	}

	return defs
}
