package parser

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// DataMorfParserBase implementation.
type DataMorfParserBase struct {
	*antlr.BaseParser
}

// Short for p.prev(str string)
func (p *DataMorfParserBase) p(str string) bool {
	return p.prev(str)
}

// Whether the previous token value equals to str.
func (p *DataMorfParserBase) prev(str string) bool {
	return p.GetTokenStream().LT(-1).GetText() == str
}

// Short for p.next(str string)
func (p *DataMorfParserBase) n(str string) bool {
	return p.next(str)
}

// Whether the next token value equals to str.
func (p *DataMorfParserBase) next(str string) bool {
	return p.GetTokenStream().LT(1).GetText() == str
}

func (p *DataMorfParserBase) notLineTerminator() bool {
	return !p.lineTerminatorAhead()
}

func (p *DataMorfParserBase) notOpenBraceAndNotFunction() bool {
	nextTokenType := p.GetTokenStream().LT(1).GetTokenType()
	return nextTokenType != DataMorfParserOpenBrace && nextTokenType != DataMorfParserFunction_
}

func (p *DataMorfParserBase) closeBrace() bool {
	return p.GetTokenStream().LT(1).GetTokenType() == DataMorfParserCloseBrace
}

// Returns true if on the current index of the parser's
// token stream a token exists on the Hidden channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *DataMorfParserBase) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	if possibleIndexEosToken < 0 {
		return false
	}
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return false
	}

	if ahead.GetTokenType() == DataMorfParserLineTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == DataMorfParserWhiteSpaces {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	_type := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	return (_type == DataMorfParserMultiLineComment && (strings.Contains(text, "\r") || strings.Contains(text, "\n"))) ||
		(_type == DataMorfParserLineTerminator)
}
