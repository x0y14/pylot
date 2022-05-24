package pylot

import (
	"fmt"
	"testing"
)

const (
	testClass = `{
	   "type": "Module",
	   "body": [{
	       "type": "ClassDef",
	       "name": "Name",
	       "bases": [],
	       "keywords": [],
	       "body": [{
	           "type": "FunctionDef",
	           "name": "__init__",
	           "args": {
	               "type": "arguments",
	               "posonlyargs": [],
	               "args": [{
	                   "type": "arg",
	                   "arg": "self"
	               }, {
	                   "type": "arg",
	                   "arg": "first",
	                   "annotation": {
	                       "type": "Name",
	                       "id": "str",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   }
	               }, {
	                   "type": "arg",
	                   "arg": "middle",
	                   "annotation": {
	                       "type": "Name",
	                       "id": "str",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   }
	               }, {
	                   "type": "arg",
	                   "arg": "last",
	                   "annotation": {
	                       "type": "Name",
	                       "id": "str",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   }
	               }],
	               "kwonlyargs": [],
	               "kw_defaults": [],
	               "defaults": []
	           },
	           "body": [{
	               "type": "AnnAssign",
	               "target": {
	                   "type": "Attribute",
	                   "value": {
	                       "type": "Name",
	                       "id": "self",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   },
	                   "attr": "first",
	                   "ctx": {
	                       "type": "Store"
	                   }
	               },
	               "annotation": {
	                   "type": "Name",
	                   "id": "str",
	                   "ctx": {
	                       "type": "Load"
	                   }
	               },
	               "value": {
	                   "type": "Name",
	                   "id": "first",
	                   "ctx": {
	                       "type": "Load"
	                   }
	               },
	               "simple": "0"
	           }, {
	               "type": "AnnAssign",
	               "target": {
	                   "type": "Attribute",
	                   "value": {
	                       "type": "Name",
	                       "id": "self",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   },
	                   "attr": "middle",
	                   "ctx": {
	                       "type": "Store"
	                   }
	               },
	               "annotation": {
	                   "type": "Name",
	                   "id": "str",
	                   "ctx": {
	                       "type": "Load"
	                   }
	               },
	               "value": {
	                   "type": "Name",
	                   "id": "middle",
	                   "ctx": {
	                       "type": "Load"
	                   }
	               },
	               "simple": "0"
	           }, {
	               "type": "AnnAssign",
	               "target": {
	                   "type": "Attribute",
	                   "value": {
	                       "type": "Name",
	                       "id": "self",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   },
	                   "attr": "last",
	                   "ctx": {
	                       "type": "Store"
	                   }
	               },
	               "annotation": {
	                   "type": "Name",
	                   "id": "str",
	                   "ctx": {
	                       "type": "Load"
	                   }
	               },
	               "value": {
	                   "type": "Name",
	                   "id": "last",
	                   "ctx": {
	                       "type": "Load"
	                   }
	               },
	               "simple": "0"
	           }],
	           "decorator_list": []
	       }, {
	           "type": "FunctionDef",
	           "name": "to_s",
	           "args": {
	               "type": "arguments",
	               "posonlyargs": [],
	               "args": [{
	                   "type": "arg",
	                   "arg": "self"
	               }],
	               "kwonlyargs": [],
	               "kw_defaults": [],
	               "defaults": []
	           },
	           "body": [{
	               "type": "Expr",
	               "value": {
	                   "type": "Call",
	                   "func": {
	                       "type": "Name",
	                       "id": "print",
	                       "ctx": {
	                           "type": "Load"
	                       }
	                   },
	                   "args": [{
	                       "type": "BinOp",
	                       "left": {
	                           "type": "BinOp",
	                           "left": {
	                               "type": "BinOp",
	                               "left": {
	                                   "type": "BinOp",
	                                   "left": {
	                                       "type": "Attribute",
	                                       "value": {
	                                           "type": "Name",
	                                           "id": "self",
	                                           "ctx": {
	                                               "type": "Load"
	                                           }
	                                       },
	                                       "attr": "first",
	                                       "ctx": {
	                                           "type": "Load"
	                                       }
	                                   },
	                                   "op": {
	                                       "type": "Add"
	                                   },
	                                   "right": {
	                                       "type": "Constant",
	                                       "value": " "
	                                   }
	                               },
	                               "op": {
	                                   "type": "Add"
	                               },
	                               "right": {
	                                   "type": "Attribute",
	                                   "value": {
	                                       "type": "Name",
	                                       "id": "self",
	                                       "ctx": {
	                                           "type": "Load"
	                                       }
	                                   },
	                                   "attr": "middle",
	                                   "ctx": {
	                                       "type": "Load"
	                                   }
	                               }
	                           },
	                           "op": {
	                               "type": "Add"
	                           },
	                           "right": {
	                               "type": "Constant",
	                               "value": " "
	                           }
	                       },
	                       "op": {
	                           "type": "Add"
	                       },
	                       "right": {
	                           "type": "Attribute",
	                           "value": {
	                               "type": "Name",
	                               "id": "self",
	                               "ctx": {
	                                   "type": "Load"
	                               }
	                           },
	                           "attr": "last",
	                           "ctx": {
	                               "type": "Load"
	                           }
	                       }
	                   }],
	                   "keywords": []
	               }
	           }],
	           "decorator_list": [],
	           "returns": {
	               "type": "Constant",
	               "value": null
	           }
	       }],
	       "decorator_list": []
	   }],
	   "type_ignores": []
	}`
)

func TestCodeGen_Gen(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect string
	}{
		{
			"a",
			`{"type":"Module","body":[{"type":"Expr","value":{"type":"Call","func":{"type":"Name","id":"print","ctx":{"type":"Load"}},"args":[{"type":"Constant","value":"hello"}],"keywords":[]}}],"type_ignores":[]}`,
			"",
		},
		{
			"real",
			testClass,
			"",
		},
		{
			"a2",
			`{"type":"Module","body":[{"type":"ClassDef","name":"Human","bases":[],"keywords":[],"body":[{"type":"FunctionDef","name":"__init__","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"},{"type":"arg","arg":"name","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Assign","targets":[{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Store"}}],"value":{"type":"Name","id":"name","ctx":{"type":"Load"}}}],"decorator_list":[]},{"type":"FunctionDef","name":"say","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"},{"type":"arg","arg":"text","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Expr","value":{"type":"Call","func":{"type":"Name","id":"print","ctx":{"type":"Load"}},"args":[{"type":"BinOp","left":{"type":"BinOp","left":{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Load"}},"op":{"type":"Add"},"right":{"type":"Constant","value":"< "}},"op":{"type":"Add"},"right":{"type":"Name","id":"text","ctx":{"type":"Load"}}}],"keywords":[]}}],"decorator_list":[],"returns":{"type":"Constant","value":null}},{"type":"FunctionDef","name":"mr","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"BinOp","left":{"type":"Constant","value":"Mr."},"op":{"type":"Add"},"right":{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Load"}}}}],"decorator_list":[],"returns":{"type":"Name","id":"str","ctx":{"type":"Load"}}},{"type":"FunctionDef","name":"my_name","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"Constant","value":"myname"}}],"decorator_list":[],"returns":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"decorator_list":[]}],"type_ignores":[]}`,
			"",
		},
		{
			"a3",
			`{"type":"Module","body":[{"type":"FunctionDef","name":"global_func","args":{"type":"arguments","posonlyargs":[],"args":[],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"Constant","value":"1"}}],"decorator_list":[],"returns":{"type":"Name","id":"int","ctx":{"type":"Load"}}},{"type":"ClassDef","name":"Human","bases":[],"keywords":[],"body":[{"type":"FunctionDef","name":"__init__","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"},{"type":"arg","arg":"name","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Assign","targets":[{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Store"}}],"value":{"type":"Name","id":"name","ctx":{"type":"Load"}}}],"decorator_list":[]},{"type":"FunctionDef","name":"say","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"},{"type":"arg","arg":"text","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Expr","value":{"type":"Call","func":{"type":"Name","id":"print","ctx":{"type":"Load"}},"args":[{"type":"BinOp","left":{"type":"BinOp","left":{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Load"}},"op":{"type":"Add"},"right":{"type":"Constant","value":"< "}},"op":{"type":"Add"},"right":{"type":"Name","id":"text","ctx":{"type":"Load"}}}],"keywords":[]}}],"decorator_list":[],"returns":{"type":"Constant","value":null}},{"type":"FunctionDef","name":"mr","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"BinOp","left":{"type":"Constant","value":"Mr."},"op":{"type":"Add"},"right":{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Load"}}}}],"decorator_list":[],"returns":{"type":"Name","id":"str","ctx":{"type":"Load"}}},{"type":"FunctionDef","name":"my_name","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"Constant","value":"myname"}}],"decorator_list":[],"returns":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"decorator_list":[]}],"type_ignores":[]}`,
			"",
		},
		{
			"a4",
			`{"type":"Module","body":[{"type":"FunctionDef","name":"say","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"text","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"Name","id":"text","ctx":{"type":"Load"}}}],"decorator_list":[]},{"type":"FunctionDef","name":"hello","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"name","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Assign","targets":[{"type":"Name","id":"you","ctx":{"type":"Store"}}],"value":{"type":"Name","id":"name","ctx":{"type":"Load"}}},{"type":"Expr","value":{"type":"Call","func":{"type":"Name","id":"print","ctx":{"type":"Load"}},"args":[{"type":"Name","id":"you","ctx":{"type":"Load"}}],"keywords":[]}}],"decorator_list":[]}],"type_ignores":[]}`,
			"",
		},
	}

	cg := NewCodeGen()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cg.Gen(tt.in)
			cg.Disclosure()
			fmt.Printf("%v", cg.defines.String())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestCodeGen_Gen_Origin(t *testing.T) {
	var tests = []struct {
		name   string
		in     string
		expect string
	}{
		{
			"int string test",
			`def global_func() -> int:
	return 1
class Human:
    def __init__(self, name: str):
        self.name = name

    def say(self, text: str) -> None:
        print(self.name + "< " + text)
    
    def mr(self) -> str:
        return "Mr."+self.name
    
    def my_name(self) -> str:
        return "myname"`,
			"",
		},
	}

	cg := NewCodeGen()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := ToJson(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			err = cg.Gen(code)
			cg.Disclosure()
			fmt.Printf("%v", cg.defines.String())
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
