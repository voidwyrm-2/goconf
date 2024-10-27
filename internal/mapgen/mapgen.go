package mapgen

import (
	"github.com/voidwyrm-2/goconf/internal/parser"
)

type MapGenerator struct {
	fields []parser.GoconfField
}

func New(fields []parser.GoconfField) MapGenerator {
	return MapGenerator{fields: fields}
}

func (mp MapGenerator) Generate() (map[string]any, error) {
	out := make(map[string]any)

	for _, f := range mp.fields {
		k, v, err := f.ToMap()
		if err != nil {
			return map[string]any{}, err
		}
		out[k] = v
	}

	return out, nil
}
