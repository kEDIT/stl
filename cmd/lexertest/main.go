package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/kEDIT/stl/lexer"
)

const (
	fileName = "../data/example_ascii.stl"
)

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open file %s: %v", fileName, err)
	}
	defer f.Close()
	l := lexer.NewAsciiLexer(f)
	toks, done := l.Lex()
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case tok := <-toks:
				fmt.Println(tok.String())
			case _ = <-done:
				return
			}
		}
	}(wg)

	wg.Wait()
}
