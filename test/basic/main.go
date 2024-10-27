package main

import (
	"fmt"
	"reflect"

	"github.com/voidwyrm-2/goconf/internal/lexer"
	"github.com/voidwyrm-2/goconf/internal/mapgen"
	"github.com/voidwyrm-2/goconf/internal/parser"
)

type Integer int

func main() {
	l := lexer.NewLexer(`. config.txt
name:string:Jacob Thaumiel;
age:uint: 21;
has_cats:bool: true;
iq:float:1.46324;`)

	tokens, err := l.Lex()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, t := range tokens {
		fmt.Println(t.Fmt())
	}

	fmt.Println()

	p := parser.NewParser(tokens)

	fields, err := p.Parse()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, f := range fields {
		fmt.Println(f)
	}

	fmt.Println()

	m := mapgen.New(fields)

	result, err := m.Generate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for k, v := range result {
		fmt.Println(k, v, reflect.TypeOf(v).Kind().String())
	}
}
