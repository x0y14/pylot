package pylot

type Datatype int

const (
	Unknown Datatype = iota

	Void

	String
	Char

	Float
	Double

	I64
	I32
	Ui64
	Ui32
)

// \s?([A-z0-9]+)\s?\n?

var dtypes = [...]string{
	Unknown: "Unknown",
	Void:    "Void",
	String:  "String",
	Char:    "Char",
	Float:   "Float",
	Double:  "Double",
	I64:     "I64",
	I32:     "I32",
	Ui64:    "Ui64",
	Ui32:    "Ui32",
}

func (d Datatype) String() string {
	return dtypes[d]
}
