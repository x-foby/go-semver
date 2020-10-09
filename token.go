package semver

import "fmt"

type token int

const (
	tokenIllegal token = iota
	tokenEOF

	tokenIdent  // [0-9A-Za-z]
	tokenDot    // .
	tokenHyphen // -
	tokenPlus   // +
)

var tokens = [...]string{
	tokenIllegal: "ILLEGAL",
	tokenEOF:     "EOF",

	tokenIdent:  "IDENT",
	tokenDot:    ".",
	tokenHyphen: "-",
	tokenPlus:   "+",
}

func (t token) String() string {
	if t < 0 || t >= token(len(tokens)) {
		return fmt.Sprintf("undefined (%v)", int(t))
	}

	return tokens[t]
}
