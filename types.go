package pylot

import (
	"fmt"
	"strings"
)

type Types int

const (
	UNKNOWN Types = iota
	UNDEFINED
	CLASS
	FUNCTION

	INTEGER
	FLOAT

	STRING

	NULL
)

var types = [...]string{
	UNKNOWN:   "UNKNOWN",
	UNDEFINED: "UNDEFINED",
	CLASS:     "CLASS",
	FUNCTION:  "FUNCTION",
	INTEGER:   "INTEGER",
	FLOAT:     "FLOAT",
	STRING:    "STRING",
	NULL:      "NULL",
}

func (t Types) String() string {
	return types[t]
}

func ToTypes(s string) Types {
	switch strings.ToLower(s) {
	case "int":
		return INTEGER
	case "float", "double": // ????
		return FLOAT
	case "str", "string":
		return STRING
	case "null", "none":
		return NULL
	//case "void":
	//	return VOID
	default:
		return UNKNOWN
	}
}

type ObjectDataType struct {
	Types
	Raw string
}

func (o ObjectDataType) String() string {
	if o.Types == UNKNOWN {
		return fmt.Sprintf("%v(\"%v\")", o.Types.String(), o.Raw)
	}
	return o.Types.String()
}
