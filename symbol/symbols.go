package symbol

import (
	"unicode"
)

// keywords
const (
	solidKey    = "solid"
	endsolidKey = "endsolid"
	outerKey    = "outer"
	loopKey     = "loop"
	endloopKey  = "endloop"
	vertexKey   = "vertex"
	facetKey    = "facet"
	endfacetKey = "endfacet"
	normalKey   = "normal"
)

var keyword = map[string]struct{}{
	solidKey:    struct{}{},
	endsolidKey: struct{}{},
	outerKey:    struct{}{},
	loopKey:     struct{}{},
	endloopKey:  struct{}{},
	vertexKey:   struct{}{},
	facetKey:    struct{}{},
	endfacetKey: struct{}{},
	normalKey:   struct{}{},
}

func isFloatPunct(r rune) bool {
	return r == '.' || r == '-' || r == '+' || r == 'e' || r == 'E'
}

func Float(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && !isFloatPunct(r) {
			return false
		}
	}
	return true
}

func Keyword(s string) bool {
	if _, ok := keyword[s]; ok {
		return true
	}
	return false
}

func Ignore(r rune) bool {
	return unicode.IsSpace(r)
}
