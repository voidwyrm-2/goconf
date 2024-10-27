package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/voidwyrm-2/goconf/internal/lexer"
)

type GoconfField struct {
	name, _type, value lexer.Token
}

func NewGoconfField(name, _type, value lexer.Token) GoconfField {
	return GoconfField{name: name, _type: _type, value: value}
}

func (gf GoconfField) ToMap() (name string, a any, e error) {
	name = gf.name.Literal()

	ty := strings.ToLower(gf._type.Literal())
	va := gf.value.Literal()

	if ty == "str" || ty == "string" {
		a, e = va, nil
		return
	}

	va = strings.TrimSpace(va)

	switch ty {
	case "int", "integer", "int32", "int64":
		a, e = strconv.Atoi(va)
	case "uint", "uinteger", "uint32", "uint64":
		a, e = strconv.ParseUint(va, 10, 0)
		a = uint(a.(uint64))
	case "char", "byte":
		a, e = strconv.Atoi(va)
		a = byte(a.(int))
	case "bool", "boolean":
		a, e = strconv.ParseBool(va)
	case "float", "double", "float32", "float64":
		a, e = strconv.ParseFloat(va, 64)
	default:
		a, e = nil, fmt.Errorf("error on %s: invalid type '%s'", gf._type.Pos(), ty)
	}

	if e != nil {
		if !strings.HasPrefix(e.Error(), "error on line ") {
			e = fmt.Errorf("error on %s: %s", gf.value.Pos(), e.Error())
		}
	}

	return
}
