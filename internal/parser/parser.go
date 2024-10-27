package parser

import (
	"fmt"

	"github.com/voidwyrm-2/goconf/internal/lexer"
)

type Parser struct {
	tokens []lexer.Token
}

func NewParser(tokens []lexer.Token) Parser {
	return Parser{tokens: tokens}
}

func (p Parser) Parse() ([]GoconfField, error) {
	fields := []GoconfField{}

	err := func(token lexer.Token, msg string, a ...any) ([]GoconfField, error) {
		return []GoconfField{}, fmt.Errorf("error on %s: %s", token.Pos(), fmt.Sprintf(msg, a...))
	}

	tl := len(p.tokens)

	i := 0
	for i < tl {
		if p.tokens[i].IsType(lexer.IDENT) {
			flname := p.tokens[i]
			if i+1 >= tl {
				return err(p.tokens[i], "expected field type, but found EOF")
			}
			i++
			if p.tokens[i].IsType(lexer.TYPE) {
				fltype := p.tokens[i]
				if i+1 >= tl {
					return err(p.tokens[i], "expected field value, but found EOF")
				}
				i++
				if p.tokens[i].IsType(lexer.VALUE) {
					fields = append(fields, NewGoconfField(flname, fltype, p.tokens[i]))
					i++
				} else {
					return err(p.tokens[i], "expected field value, but found '%s' instead", p.tokens[i].Literal())
				}
			} else {
				return err(p.tokens[i], "expected field type, but found '%s' instead", p.tokens[i].Literal())
			}
		} else {
			return err(p.tokens[i], "expected field identifier, but found '%s' instead", p.tokens[i].Literal())
		}
	}

	return fields, nil
}
