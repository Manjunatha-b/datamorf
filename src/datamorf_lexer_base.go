package parser

import "github.com/antlr4-go/antlr/v4"

// DataMorfLexerBase state
type DataMorfLexerBase struct {
	*antlr.BaseLexer

	scopeStrictModes []bool
	stackLength      int
	stackIx          int

	lastToken          antlr.Token
	useStrictDefault   bool
	useStrictCurrent   bool
	currentDepth       int
	templateDepthStack []int
}

func (l *DataMorfLexerBase) IsStartOfFile() bool {
	return l.lastToken == nil
}

func (l *DataMorfLexerBase) pushStrictModeScope(v bool) {
	if l.stackIx == l.stackLength {
		l.scopeStrictModes = append(l.scopeStrictModes, v)
		l.stackLength++
	} else {
		l.scopeStrictModes[l.stackIx] = v
	}
	l.stackIx++
}

func (l *DataMorfLexerBase) popStrictModeScope() bool {
	l.stackIx--
	v := l.scopeStrictModes[l.stackIx]
	l.scopeStrictModes[l.stackIx] = false
	return v
}

// IsStrictMode is self explanatory.
func (l *DataMorfLexerBase) IsStrictMode() bool {
	return l.useStrictCurrent
}

// NextToken from the character stream.
func (l *DataMorfLexerBase) NextToken() antlr.Token {
	next := l.BaseLexer.NextToken() // Get next token
	if next.GetChannel() == antlr.TokenDefaultChannel {
		// Keep track of the last token on default channel
		l.lastToken = next
	}
	return next
}

// ProcessOpenBrace is called when a { is encountered during
// lexing, we push a new scope everytime.
func (l *DataMorfLexerBase) ProcessOpenBrace() {
	l.currentDepth++
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 && l.scopeStrictModes[l.stackIx-1] {
		l.useStrictCurrent = true
	}
	l.pushStrictModeScope(l.useStrictCurrent)
}

// ProcessCloseBrace is called when a } is encountered during
// lexing, we pop a scope unless we're inside global scope.
func (l *DataMorfLexerBase) ProcessCloseBrace() {
	l.useStrictCurrent = l.useStrictDefault
	if l.stackIx > 0 {
		l.useStrictCurrent = l.popStrictModeScope()
	}
	l.currentDepth--
}

// ProcessStringLiteral is called when lexing a string literal.
func (l *DataMorfLexerBase) ProcessStringLiteral() {
	if l.lastToken == nil || l.lastToken.GetTokenType() == DataMorfLexerOpenBrace {
		if l.GetText() == `"use strict"` || l.GetText() == "'use strict'" {
			if l.stackIx > 0 {
				l.popStrictModeScope()
			}
			l.useStrictCurrent = true
			l.pushStrictModeScope(l.useStrictCurrent)
		}
	}
}

func (l *DataMorfLexerBase) ProcessTemplateOpenBrace() {
	l.currentDepth++
	l.templateDepthStack = append(l.templateDepthStack, l.currentDepth)
}

func (l *DataMorfLexerBase) ProcessTemplateCloseBrace() {
	l.templateDepthStack = l.templateDepthStack[:len(l.templateDepthStack)-1]
	l.currentDepth--
}

// IsRegexPossible returns true if the lexer can match a
// regex literal.
func (l *DataMorfLexerBase) IsRegexPossible() bool {
	if l.lastToken == nil {
		return true
	}
	switch l.lastToken.GetTokenType() {
	case DataMorfLexerIdentifier, DataMorfLexerNullLiteral,
		DataMorfLexerBooleanLiteral, DataMorfLexerThis,
		DataMorfLexerCloseBracket, DataMorfLexerCloseParen,
		DataMorfLexerOctalIntegerLiteral, DataMorfLexerDecimalLiteral,
		DataMorfLexerHexIntegerLiteral, DataMorfLexerStringLiteral,
		DataMorfLexerPlusPlus, DataMorfLexerMinusMinus:
		return false
	default:
		return true
	}
}

func (l *DataMorfLexerBase) IsInTemplateString() bool {
	return len(l.templateDepthStack) > 0 && l.templateDepthStack[len(l.templateDepthStack)-1] == l.currentDepth
}

func (l *DataMorfLexerBase) Reset() {
	l.scopeStrictModes = nil
	l.stackLength = 0
	l.stackIx = 0
	l.lastToken = nil
	l.useStrictDefault = false
	l.useStrictCurrent = false
	l.currentDepth = 0
	l.templateDepthStack = make([]int, 0)
	l.BaseLexer.Reset()
}
