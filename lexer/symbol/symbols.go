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

// check if input string contains valid characters
// for sigle precision float (as described in STL format spec)
func isFloatPunct(r rune) bool {
	return r == '.' || r == '-' || r == '+' || r == 'e' || r == 'E'
}

// Float checks if the input string is a valid float (as described in STL format spec)
func Float(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && !isFloatPunct(r) {
			return false
		}
	}
	return true
}

// Keyword checks if the input string is a valid keyword in STL format
func Keyword(s string) bool {
	if _, ok := keyword[s]; ok {
		return true
	}
	return false
}

// Ignore checks if the input rune can be ignored by a lexer
func Ignore(r rune) bool {
	return unicode.IsSpace(r)
}
