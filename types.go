package pylot

import (
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
	//switch t {
	//case UNKNOWN, UNDEFINED:
	//	return color.HiRedString(types[t])
	//case CLASS:
	//	return color.HiWhiteString(types[t])
	//case FUNCTION:
	//	return color.HiCyanString(types[t])
	//case INTEGER, FLOAT, STRING, NULL:
	//	return color.HiGreenString(types[t])
	//default:
	//	panic("unsupported types")
	//}
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
