package funcexprtype

type Type int

const (
	STRING Type = iota
	INTEGER
	REF
	ARRAYREF
	FLOAT
	EXPRESSION
	BOOLEAN
	NIL
)

func (t Type) String() string {
	switch t {
	case STRING:
		return "string"
	case INTEGER:
		return "integer"
	case REF:
		return "ref"
	case ARRAYREF:
		return "arrayRef"
	case FLOAT:
		return "float"
	case EXPRESSION:
		return "expression"
	case BOOLEAN:
		return "boolean"
	case NIL:
		return "nil"
	}
	return ""
}
