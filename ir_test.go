package pylot

//
//import (
//	"testing"
//)
//
//const (
//	testClass = `{
//    "type": "Module",
//    "body": [{
//        "type": "ClassDef",
//        "name": "Name",
//        "bases": [],
//        "keywords": [],
//        "body": [{
//            "type": "FunctionDef",
//            "name": "__init__",
//            "args": {
//                "type": "arguments",
//                "posonlyargs": [],
//                "args": [{
//                    "type": "arg",
//                    "arg": "self"
//                }, {
//                    "type": "arg",
//                    "arg": "first",
//                    "annotation": {
//                        "type": "Name",
//                        "id": "str",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    }
//                }, {
//                    "type": "arg",
//                    "arg": "middle",
//                    "annotation": {
//                        "type": "Name",
//                        "id": "str",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    }
//                }, {
//                    "type": "arg",
//                    "arg": "last",
//                    "annotation": {
//                        "type": "Name",
//                        "id": "str",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    }
//                }],
//                "kwonlyargs": [],
//                "kw_defaults": [],
//                "defaults": []
//            },
//            "body": [{
//                "type": "AnnAssign",
//                "target": {
//                    "type": "Attribute",
//                    "value": {
//                        "type": "Name",
//                        "id": "self",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    },
//                    "attr": "first",
//                    "ctx": {
//                        "type": "Store"
//                    }
//                },
//                "annotation": {
//                    "type": "Name",
//                    "id": "str",
//                    "ctx": {
//                        "type": "Load"
//                    }
//                },
//                "value": {
//                    "type": "Name",
//                    "id": "first",
//                    "ctx": {
//                        "type": "Load"
//                    }
//                },
//                "simple": "0"
//            }, {
//                "type": "AnnAssign",
//                "target": {
//                    "type": "Attribute",
//                    "value": {
//                        "type": "Name",
//                        "id": "self",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    },
//                    "attr": "middle",
//                    "ctx": {
//                        "type": "Store"
//                    }
//                },
//                "annotation": {
//                    "type": "Name",
//                    "id": "str",
//                    "ctx": {
//                        "type": "Load"
//                    }
//                },
//                "value": {
//                    "type": "Name",
//                    "id": "middle",
//                    "ctx": {
//                        "type": "Load"
//                    }
//                },
//                "simple": "0"
//            }, {
//                "type": "AnnAssign",
//                "target": {
//                    "type": "Attribute",
//                    "value": {
//                        "type": "Name",
//                        "id": "self",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    },
//                    "attr": "last",
//                    "ctx": {
//                        "type": "Store"
//                    }
//                },
//                "annotation": {
//                    "type": "Name",
//                    "id": "str",
//                    "ctx": {
//                        "type": "Load"
//                    }
//                },
//                "value": {
//                    "type": "Name",
//                    "id": "last",
//                    "ctx": {
//                        "type": "Load"
//                    }
//                },
//                "simple": "0"
//            }],
//            "decorator_list": []
//        }, {
//            "type": "FunctionDef",
//            "name": "to_s",
//            "args": {
//                "type": "arguments",
//                "posonlyargs": [],
//                "args": [{
//                    "type": "arg",
//                    "arg": "self"
//                }],
//                "kwonlyargs": [],
//                "kw_defaults": [],
//                "defaults": []
//            },
//            "body": [{
//                "type": "Expr",
//                "value": {
//                    "type": "Call",
//                    "func": {
//                        "type": "Name",
//                        "id": "print",
//                        "ctx": {
//                            "type": "Load"
//                        }
//                    },
//                    "args": [{
//                        "type": "BinOp",
//                        "left": {
//                            "type": "BinOp",
//                            "left": {
//                                "type": "BinOp",
//                                "left": {
//                                    "type": "BinOp",
//                                    "left": {
//                                        "type": "Attribute",
//                                        "value": {
//                                            "type": "Name",
//                                            "id": "self",
//                                            "ctx": {
//                                                "type": "Load"
//                                            }
//                                        },
//                                        "attr": "first",
//                                        "ctx": {
//                                            "type": "Load"
//                                        }
//                                    },
//                                    "op": {
//                                        "type": "Add"
//                                    },
//                                    "right": {
//                                        "type": "Constant",
//                                        "value": " "
//                                    }
//                                },
//                                "op": {
//                                    "type": "Add"
//                                },
//                                "right": {
//                                    "type": "Attribute",
//                                    "value": {
//                                        "type": "Name",
//                                        "id": "self",
//                                        "ctx": {
//                                            "type": "Load"
//                                        }
//                                    },
//                                    "attr": "middle",
//                                    "ctx": {
//                                        "type": "Load"
//                                    }
//                                }
//                            },
//                            "op": {
//                                "type": "Add"
//                            },
//                            "right": {
//                                "type": "Constant",
//                                "value": " "
//                            }
//                        },
//                        "op": {
//                            "type": "Add"
//                        },
//                        "right": {
//                            "type": "Attribute",
//                            "value": {
//                                "type": "Name",
//                                "id": "self",
//                                "ctx": {
//                                    "type": "Load"
//                                }
//                            },
//                            "attr": "last",
//                            "ctx": {
//                                "type": "Load"
//                            }
//                        }
//                    }],
//                    "keywords": []
//                }
//            }],
//            "decorator_list": [],
//            "returns": {
//                "type": "Constant",
//                "value": null
//            }
//        }],
//        "decorator_list": []
//    }],
//    "type_ignores": []
//}`
//)
//
//func TestIr(t *testing.T) {
//	//err := Ir(testClass)
//	//if err != nil {
//	//	t.Fatal(err)
//	//}
//	var tests = []struct {
//		name   string
//		in     string
//		expect string
//	}{
//		{
//			"a",
//			`{"type":"Module","body":[{"type":"Expr","value":{"type":"Call","func":{"type":"Name","id":"print","ctx":{"type":"Load"}},"args":[{"type":"Constant","value":"hello"}],"keywords":[]}}],"type_ignores":[]}`,
//			"",
//		},
//		{
//			"real",
//			testClass,
//			"",
//		},
//		{
//			"a2",
//			`{"type":"Module","body":[{"type":"ClassDef","name":"Human","bases":[],"keywords":[],"body":[{"type":"FunctionDef","name":"__init__","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"},{"type":"arg","arg":"name","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Assign","targets":[{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Store"}}],"value":{"type":"Name","id":"name","ctx":{"type":"Load"}}}],"decorator_list":[]},{"type":"FunctionDef","name":"say","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"},{"type":"arg","arg":"text","annotation":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Expr","value":{"type":"Call","func":{"type":"Name","id":"print","ctx":{"type":"Load"}},"args":[{"type":"BinOp","left":{"type":"BinOp","left":{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Load"}},"op":{"type":"Add"},"right":{"type":"Constant","value":"< "}},"op":{"type":"Add"},"right":{"type":"Name","id":"text","ctx":{"type":"Load"}}}],"keywords":[]}}],"decorator_list":[],"returns":{"type":"Constant","value":null}},{"type":"FunctionDef","name":"mr","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"BinOp","left":{"type":"Constant","value":"Mr."},"op":{"type":"Add"},"right":{"type":"Attribute","value":{"type":"Name","id":"self","ctx":{"type":"Load"}},"attr":"name","ctx":{"type":"Load"}}}}],"decorator_list":[],"returns":{"type":"Name","id":"str","ctx":{"type":"Load"}}},{"type":"FunctionDef","name":"my_name","args":{"type":"arguments","posonlyargs":[],"args":[{"type":"arg","arg":"self"}],"kwonlyargs":[],"kw_defaults":[],"defaults":[]},"body":[{"type":"Return","value":{"type":"Constant","value":"myname"}}],"decorator_list":[],"returns":{"type":"Name","id":"str","ctx":{"type":"Load"}}}],"decorator_list":[]}],"type_ignores":[]}`,
//			"",
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			err := Ir(tt.in)
//			if err != nil {
//				t.Fatal(err)
//			}
//		})
//	}
//}
