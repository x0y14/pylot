package pylot

import "strings"

type Types int

const (
	UNKNOWN Types = iota
	CLASS
	FUNCTION

	INTEGER
	FLOAT

	STRING

	//VOID
	NULL
)

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
