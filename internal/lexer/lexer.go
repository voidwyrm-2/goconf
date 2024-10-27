package lexer

import (
	"fmt"
)

type Lexer struct {
	text         string
	idx, ln, col int
	ch           rune
}

func NewLexer(text string) Lexer {
	l := Lexer{text: text, idx: -1, ln: 0, col: 0, ch: -1}
	l.advance()
	return l
}

func (l *Lexer) advance() {
	l.idx++
	l.col++
	if l.idx < len(l.text) {
		l.ch = rune(l.text[l.idx])
	} else {
		l.ch = -1
	}

	if l.ch == '\n' {
		l.ln++
		l.col = 0
	}
}

func (l Lexer) peek() rune {
	if l.idx+1 < len(l.text) {
		return rune(l.text[l.idx+1])
	}
	return -1
}

func (l Lexer) rpeek() rune {
	if l.idx-1 > -1 {
		return rune(l.text[l.idx-1])
	}
	return -1
}

func (l Lexer) identChar() bool {
	return (l.ch >= 'a' && l.ch <= 'z') || (l.ch >= 'A' && l.ch <= 'Z') || (l.ch >= '0' && l.ch <= '9') || l.ch == '_'
}

func (l *Lexer) Lex() ([]Token, error) {
	tokens := []Token{}

	/*
		var tpeek func() Token

		tpeek = func() Token {
			if len(tokens) > 0 {
				return tokens[len(tokens)-1]
			}
			return NewToken(NONE, "", -1, -1, -1, -1)
		}
	*/

	for l.ch != -1 {
		switch l.ch {
		case '\n', '\t', ' ':
			l.advance()
		case ':':
			{
				t1, t2, err := l.collectType()
				if err != nil {
					return []Token{}, err
				}
				tokens = append(tokens, t1, t2)
			}
		case '.':
			l.skipComment()
		default:
			if l.identChar() {
				tokens = append(tokens, l.collectIdent())
			} else {
				return []Token{}, fmt.Errorf("error on line %d, col %d: illegal character '%c'", l.ln+1, l.col+1, l.ch)
			}
		}
	}

	return tokens, nil
}

func (l *Lexer) collectIdent() Token {
	start := l.col
	startln := l.ln
	lit := ""

	for l.ch != -1 && l.identChar() {
		lit += string(l.ch)
		l.advance()
	}

	return NewToken(IDENT, lit, start, l.col, startln, l.ln)
}

func (l *Lexer) collectValue() (Token, error) {
	start := l.col
	startln := l.ln
	lit := ""
	escape := false

	l.advance()

	for l.ch != -1 {
		if escape {
			switch l.ch {
			case '\\', ';', ':':
				lit += string(l.ch)
			case 'n':
				lit += "\n"
			case 't':
				lit += "\t"
			default:
				return Token{}, fmt.Errorf("error on line %d, col %d: invalid escape character '%c'", l.ln+1, l.col+1, l.ch)
			}
			escape = false
		} else if l.ch == '\\' {
			escape = true
		} else {
			if l.ch == ';' {
				break
			}
			lit += string(l.ch)
		}
		l.advance()
	}

	if l.ch != ';' {
		return Token{}, fmt.Errorf("error on line %d, col %d: field values must be ended by ';'", l.ln+1, l.col+1)
	}

	token := NewToken(VALUE, lit, start, l.col, startln, l.ln)
	l.advance()
	return token, nil
}

func (l *Lexer) collectType() (Token, Token, error) {
	start := l.col
	startln := l.ln
	typeLit := ""

	l.advance()

	for l.ch != -1 && l.ch != ':' {
		typeLit += string(l.ch)
		l.advance()
	}

	if l.ch != ':' {
		return Token{}, Token{}, fmt.Errorf("error on line %d, col %d: field types must be followed by ':'", l.ln+1, l.col+1)
	}

	t1 := NewToken(TYPE, typeLit, start, l.col, startln, l.ln)

	t2, e := l.collectValue()

	return t1, t2, e
}

func (l *Lexer) skipComment() {
	for l.ch != -1 && l.ch != '\n' {
		l.advance()
	}
}
