package semver

type scanner struct {
	src    []rune
	len    int
	ch     rune
	offset int
}

func (s *scanner) Init(src []rune) {
	s.src = src
	s.len = len(s.src)
	s.offset = -1
	s.next()
}

func (s *scanner) next() {
	if s.offset < s.len-1 {
		s.offset++
		s.ch = s.src[s.offset]
	} else {
		s.offset = s.len
		s.ch = -1 // eof
	}
}

func (s *scanner) scanIdentifier() string {
	offs := s.offset
	for isASCII(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func isASCII(ch rune) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}

func (s *scanner) scan() (pos int, tok token, lit string) {
	tok = tokenIllegal
	pos = s.offset

	switch ch := s.ch; {
	case isASCII(ch):
		lit = s.scanIdentifier()
		tok = tokenIdent
	default:
		switch ch {
		case -1:
			tok = tokenEOF
		case '.':
			tok = tokenDot
		case '+':
			tok = tokenPlus
		case '-':
			tok = tokenHyphen
		}
		s.next()
	}
	return
}
