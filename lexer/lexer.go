package lexer

// Lexer is an interface type defining the behavior of a lexer.
// It should return a channel of Token and an empty struct channel
// (to indicate when lexing has completed)
type Lexer interface {
	Lex() (chan Token, chan struct{})
}
