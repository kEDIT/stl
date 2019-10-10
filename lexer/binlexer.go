package lexer

// TODO: add binary lexer implementation

const (
	szHeader      int = 80 // number of bytes in ASCII header
	szFacetCount  int = 4  // number of bytes in facet count (uint64)
	szFacetRecord int = 50 // number of bytes per facet record
	szCoord       int = 4  // number of bytes per coordinte in a vertex (float64)
)

const (
	termSymbol = 0x0000
)

// type binLexer struct {
// 	rd *bufio.Reader
// 	tbuf strings.Builder
// 	toks chan Token
// }
