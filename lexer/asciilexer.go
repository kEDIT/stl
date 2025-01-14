package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/kEDIT/stl/lexer/symbol"
)

type asciiLexer struct {
	rd    *bufio.Reader
	tbuf  strings.Builder
	toks  chan Token
	done  chan struct{}
	start int
	pos   int
}

// NewLexer returns an instance of an stl lexer
func NewAsciiLexer(input io.Reader) *asciiLexer {
	return &asciiLexer{
		rd:   bufio.NewReader(input),
		tbuf: strings.Builder{},
		toks: make(chan Token),
		done: make(chan struct{}),
	}
}

func (l *asciiLexer) tokens() (chan Token, chan struct{}) {
	return l.toks, l.done
}

func (l *asciiLexer) read() error {
	r, sz, err := l.rd.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.emit()
			return err
		}
		str := l.tbuf.String()
		e := fmt.Errorf("lex error: accumulated: %s, next rune: %v, err: %v", str, r, err)
		return e
	}
	if symbol.Ignore(r) {
		if l.start != l.pos {
			l.emit()
		}
		l.ignore(sz)
		return nil
	}
	l.accept(r, sz)
	return nil
}

func (l *asciiLexer) accept(r rune, sz int) {
	l.tbuf.WriteRune(r)
	l.pos += sz
}

func (l *asciiLexer) ignore(sz int) {
	l.pos += sz
	l.start = l.pos
	return
}

func (l *asciiLexer) emit() {
	l.toks <- toValidToken(l.tbuf.String(), l.start)
	l.tbuf.Reset()
}

func (l *asciiLexer) scan() {
	for {
		if err := l.read(); err != nil {
			if err != io.EOF {
				l.toks <- toErrorToken(err, l.start)
			}
			l.done <- struct{}{}
			close(l.toks)
			close(l.done)
			break
		}
	}
	return
}

func (l *asciiLexer) Lex() (chan Token, chan struct{}) {
	go l.scan()
	return l.toks, l.done
}
