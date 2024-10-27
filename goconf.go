package goconf

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"

	"github.com/voidwyrm-2/goconf/internal/lexer"
	"github.com/voidwyrm-2/goconf/internal/mapgen"
	"github.com/voidwyrm-2/goconf/internal/parser"
)

var validGoTypes = []string{
	"int",
	"int32",
	"int64",
	"uint",
	"uint32",
	"uint64",
	"byte",
	"bool",
	"float",
	"float32",
	"float64",
}

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func writeFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

/*
Converts a map to the goconf format
*/
func FromMap(m map[string]any) (string, error) {
	fields := []string{}

	for k, v := range m {
		if !slices.Contains(validGoTypes, reflect.TypeOf(v).Kind().String()) {
			return "", fmt.Errorf("invalid type for goconf conversion '%s'(attached to key '%s')", reflect.TypeOf(v).Kind().String(), k)
		} else if sv, ok := v.(string); ok {
			v = strings.ReplaceAll(sv, "\n", "\\n")
		}
		fields = append(fields, fmt.Sprintf("%s:%s:%s;", k, reflect.TypeOf(v).Kind().String(), v))
	}

	return strings.Join(fields, "\n"), nil
}

/*
Loads a goconf file
*/
func Load(path string) (map[string]any, error) {
	fcontent, err := readFile(path)
	if err != nil {
		return map[string]any{}, err
	}

	l := lexer.NewLexer(fcontent)

	tokens, err := l.Lex()
	if err != nil {
		return map[string]any{}, err
	}

	p := parser.NewParser(tokens)

	fields, err := p.Parse()
	if err != nil {
		return map[string]any{}, err
	}

	m := mapgen.New(fields)

	result, err := m.Generate()
	if err != nil {
		return map[string]any{}, err
	}

	return result, nil
}

/*
Saves a map as a goconf file
*/
func Save(path string, m map[string]any) error {
	conf, err := FromMap(m)
	if err != nil {
		return err
	}

	return writeFile(path, conf)
}
