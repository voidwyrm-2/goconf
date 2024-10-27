package lexer

import "fmt"

type Tokentype uint

const (
	NONE Tokentype = iota
	IDENT
	TYPE
	COLON
	VALUE
	COMMENT
)

type Token struct {
	_type                      Tokentype
	lit                        string
	start, end, startln, endln int
}

func NewToken(_type Tokentype, literal string, start, end, startln, endln int) Token {
	return Token{_type: _type, lit: literal, start: start, end: end, startln: startln, endln: endln}
}

func (t Token) IsType(types ...Tokentype) bool {
	for _, ty := range types {
		if t._type != ty {
			return false
		}
	}
	return true
}

func (t Token) IsLit(literals ...string) bool {
	for _, li := range literals {
		if t.lit != li {
			return false
		}
	}
	return true
}

func (t Token) Literal() string {
	return t.lit
}

func (t Token) Fmt() string {
	return fmt.Sprintf("Token{%s, '%s', %d..%d, %d..%d}", []string{"NONE", "IDENT", "TYPE", "COLON", "VALUE", "COMMENT"}[t._type], t.lit, t.start, t.end, t.startln, t.endln)
}

func (t Token) Pos() string {
	return fmt.Sprintf("line %d, col %d", t.startln+1, t.start+1)
}
