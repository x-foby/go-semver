package semver

import (
	"fmt"
)

type parser struct {
	scanner scanner
	pos     int
	tok     token
	lit     string
}

func newParser() *parser {
	return &parser{}
}

func (p *parser) next() {
	p.pos, p.tok, p.lit = p.scanner.scan()
}

// nolint:funlen,gocyclo
func (p *parser) parse(src string) (*Version, error) {
	p.scanner.Init([]rune(src))
	var v Version

	p.next()
	if p.tok != tokenIdent {
		return nil, unexpected(p.tok.String(), p.pos)
	}
	major, err := p.parseMajor(p.lit)
	if err != nil {
		return nil, err
	}
	v.Major = major

	p.next()
	if p.tok == tokenEOF {
		return &v, nil
	} else if p.tok != tokenDot {
		return nil, unexpected(p.tok.String(), p.pos)
	}

	p.next()
	minor, err := p.parseNumber(p.lit)
	if err != nil {
		return nil, err
	}
	v.Minor = minor

	p.next()
	if p.tok == tokenEOF {
		return &v, nil
	} else if p.tok != tokenDot {
		return nil, unexpected(p.tok.String(), p.pos)
	}

	p.next()
	patch, err := p.parseNumber(p.lit)
	if err != nil {
		return nil, err
	}
	v.Patch = patch

	p.next()
	// nolint:exhaustive
	switch p.tok {
	case tokenEOF:
		return &v, nil
	case tokenHyphen:
		p.next()
		preRelease, err := p.parsePreRelease()
		if err != nil {
			return nil, err
		}
		v.PreRelease = preRelease
	case tokenPlus:
		p.next()
		buildData, err := p.parseBuildData()
		if err != nil {
			return nil, err
		}
		v.BuildData = buildData
		return &v, nil
	default:
		return nil, unexpected(p.tok.String(), p.pos)
	}

	// nolint:exhaustive
	switch p.tok {
	case tokenEOF:
		return &v, nil
	case tokenPlus:
		p.next()
		buildData, err := p.parseBuildData()
		if err != nil {
			return nil, err
		}
		v.BuildData = buildData
		return &v, nil
	default:
		return nil, unexpected(p.tok.String(), p.pos)
	}
}

func (p *parser) parseMajor(src string) (uint, error) {
	var offset int
	if src != "" && src[0] == 'v' {
		offset = 1
		src = src[1:]
	}
	if src == "" {
		return 0, unexpected(tokenEOF.String(), p.pos+offset)
	}
	n, err := strToUint(src, p.pos, p.pos+offset)
	if err != nil {
		return 0, nil
	}
	return n, nil
}

func (p *parser) parseNumber(src string) (uint, error) {
	n, err := strToUint(src, p.pos, 0)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (p *parser) parsePreRelease() ([]string, error) {
	var idents []string
	for {
		if p.tok != tokenIdent {
			return nil, unexpected(p.tok.String(), p.pos)
		}
		ident, err := p.parseIdent(tokenPlus)
		if err != nil {
			return nil, err
		}
		idents = append(idents, ident)
		if p.tok == tokenEOF || p.tok == tokenPlus {
			return idents, nil
		}
		p.next()
	}
}

func (p *parser) parseBuildData() ([]string, error) {
	var idents []string
	for {
		if p.tok != tokenIdent {
			return nil, unexpected(p.tok.String(), p.pos)
		}
		ident, err := p.parseIdent(tokenEOF)
		if err != nil {
			return nil, err
		}
		idents = append(idents, ident)
		if p.tok == tokenEOF {
			return idents, nil
		}
		p.next()
	}
}

func (p *parser) parseIdent(stopToken token) (string, error) {
	var ident string
	for {
		// nolint:exhaustive
		switch p.tok {
		case tokenDot, tokenEOF, stopToken:
			return ident, nil
		case tokenIdent:
			ident += p.lit
		case tokenHyphen:
			ident += "-"

		default:
			return "", unexpected(p.tok.String(), p.pos)
		}
		p.next()
	}
}

func strToUint(src string, pos, offset int) (uint, error) {
	var n uint
	for i, ch := range src[offset:] {
		if ch < '0' || ch > '9' {
			return 0, unexpected(string(ch), i+pos+offset)
		}
		n = n*10 + uint(ch-'0')
	}
	return n, nil
}

func unexpected(got string, pos int) error {
	return fmt.Errorf("unexpected %s at %d", got, pos)
}
