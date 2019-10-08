package lexer

import (
	"fmt"

	"github.com/kEDIT/stl/symbol"
)

type tokenType int

const (
	keywordTok tokenType = iota
	stringTok
	floatTok
	errTok
)

var tokenTypeStringer = []string{
	"keyword",
	"string",
	"float",
	"error",
}

// for debugging
func (t tokenType) String() string {
	return tokenTypeStringer[t]
}

type Token interface {
	Type() string
	Value() string
}

type token struct {
	typ tokenType
	val string
	pos int
}

func (t token) Type() string {
	return t.typ.String()
}

func (t token) Value() string {
	return t.val
}

func (t token) String() string {
	ts := t.typ.String()
	return fmt.Sprintf("token{ typ: %s, val: %s, pos: %d }", ts, t.val, t.pos)
}

func toValidToken(s string, start int) token {
	var t tokenType
	switch {
	case symbol.Keyword(s):
		t = keywordTok
		break
	case symbol.Float(s):
		t = floatTok
		break
	default:
		t = stringTok
	}

	return token{
		typ: t,
		val: s,
		pos: start,
	}
}

func toErrorToken(e error, start int) token {
	return token{
		typ: errTok,
		val: e.Error(),
		pos: start,
	}
}
