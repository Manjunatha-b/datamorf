// Code generated from litelang/src/LiteLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package LiteLang

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LiteLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LiteLangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func litelanglexerLexerInit() {
	staticData := &LiteLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'=>'", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','",
		"'='", "'?'", "'?.'", "':'", "'.'", "'++'", "'--'", "'+'", "'-'", "'!'",
		"'*'", "'/'", "'%'", "'**'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='",
		"'&&'", "'||'", "'null'", "", "", "'break'", "'do'", "'of'", "'case'",
		"'else'", "'new'", "'var'", "'return'", "'continue'", "'for'", "'switch'",
		"'while'", "'function'", "'default'", "'if'", "'elif'", "'delete'",
	}
	staticData.SymbolicNames = []string{
		"", "", "OpenBracket", "CloseBracket", "OpenParen", "CloseParen", "OpenBrace",
		"CloseBrace", "SemiColon", "Comma", "Assign", "QuestionMark", "QuestionMarkDot",
		"Colon", "Dot", "PlusPlus", "MinusMinus", "Plus", "Minus", "Not", "Multiply",
		"Divide", "Modulus", "Power", "LessThan", "MoreThan", "LessThanEquals",
		"GreaterThanEquals", "Equals", "NotEquals", "And", "Or", "NullLiteral",
		"BooleanLiteral", "DecimalLiteral", "Break", "Do", "Of", "Case", "Else",
		"New", "Var", "Return", "Continue", "For", "Switch", "While", "Function",
		"Default", "If", "Elif", "Delete", "Identifier", "StringLiteral", "WhiteSpaces",
		"LineTerminator",
	}
	staticData.RuleNames = []string{
		"T__0", "OpenBracket", "CloseBracket", "OpenParen", "CloseParen", "OpenBrace",
		"CloseBrace", "SemiColon", "Comma", "Assign", "QuestionMark", "QuestionMarkDot",
		"Colon", "Dot", "PlusPlus", "MinusMinus", "Plus", "Minus", "Not", "Multiply",
		"Divide", "Modulus", "Power", "LessThan", "MoreThan", "LessThanEquals",
		"GreaterThanEquals", "Equals", "NotEquals", "And", "Or", "NullLiteral",
		"BooleanLiteral", "DecimalLiteral", "Break", "Do", "Of", "Case", "Else",
		"New", "Var", "Return", "Continue", "For", "Switch", "While", "Function",
		"Default", "If", "Elif", "Delete", "Identifier", "StringLiteral", "WhiteSpaces",
		"LineTerminator", "DoubleStringCharacter", "LineContinuation", "DecimalIntegerLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 55, 356, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 3, 32, 205, 8, 32, 1, 33, 1, 33, 1, 33, 5, 33, 210, 8, 33, 10, 33,
		12, 33, 213, 9, 33, 1, 33, 3, 33, 216, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 5, 51, 315,
		8, 51, 10, 51, 12, 51, 318, 9, 51, 1, 52, 1, 52, 5, 52, 322, 8, 52, 10,
		52, 12, 52, 325, 9, 52, 1, 52, 1, 52, 1, 53, 4, 53, 330, 8, 53, 11, 53,
		12, 53, 331, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 3,
		55, 342, 8, 55, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 5, 57, 350, 8,
		57, 10, 57, 12, 57, 353, 9, 57, 3, 57, 355, 8, 57, 0, 0, 58, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97,
		49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 0, 113, 0,
		115, 0, 1, 0, 9, 1, 0, 48, 57, 1, 0, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 2, 0, 9, 9, 32, 32, 3, 0, 10, 10, 13, 13, 32, 32, 4, 0, 10,
		10, 13, 13, 34, 34, 92, 92, 3, 0, 10, 10, 13, 13, 8232, 8233, 1, 0, 49,
		57, 2, 0, 48, 57, 95, 95, 361, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
		0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0,
		0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0,
		0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1,
		0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0,
		105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 1, 117, 1, 0,
		0, 0, 3, 120, 1, 0, 0, 0, 5, 122, 1, 0, 0, 0, 7, 124, 1, 0, 0, 0, 9, 126,
		1, 0, 0, 0, 11, 128, 1, 0, 0, 0, 13, 130, 1, 0, 0, 0, 15, 132, 1, 0, 0,
		0, 17, 134, 1, 0, 0, 0, 19, 136, 1, 0, 0, 0, 21, 138, 1, 0, 0, 0, 23, 140,
		1, 0, 0, 0, 25, 143, 1, 0, 0, 0, 27, 145, 1, 0, 0, 0, 29, 147, 1, 0, 0,
		0, 31, 150, 1, 0, 0, 0, 33, 153, 1, 0, 0, 0, 35, 155, 1, 0, 0, 0, 37, 157,
		1, 0, 0, 0, 39, 159, 1, 0, 0, 0, 41, 161, 1, 0, 0, 0, 43, 163, 1, 0, 0,
		0, 45, 165, 1, 0, 0, 0, 47, 168, 1, 0, 0, 0, 49, 170, 1, 0, 0, 0, 51, 172,
		1, 0, 0, 0, 53, 175, 1, 0, 0, 0, 55, 178, 1, 0, 0, 0, 57, 181, 1, 0, 0,
		0, 59, 184, 1, 0, 0, 0, 61, 187, 1, 0, 0, 0, 63, 190, 1, 0, 0, 0, 65, 204,
		1, 0, 0, 0, 67, 215, 1, 0, 0, 0, 69, 217, 1, 0, 0, 0, 71, 223, 1, 0, 0,
		0, 73, 226, 1, 0, 0, 0, 75, 229, 1, 0, 0, 0, 77, 234, 1, 0, 0, 0, 79, 239,
		1, 0, 0, 0, 81, 243, 1, 0, 0, 0, 83, 247, 1, 0, 0, 0, 85, 254, 1, 0, 0,
		0, 87, 263, 1, 0, 0, 0, 89, 267, 1, 0, 0, 0, 91, 274, 1, 0, 0, 0, 93, 280,
		1, 0, 0, 0, 95, 289, 1, 0, 0, 0, 97, 297, 1, 0, 0, 0, 99, 300, 1, 0, 0,
		0, 101, 305, 1, 0, 0, 0, 103, 312, 1, 0, 0, 0, 105, 319, 1, 0, 0, 0, 107,
		329, 1, 0, 0, 0, 109, 335, 1, 0, 0, 0, 111, 341, 1, 0, 0, 0, 113, 343,
		1, 0, 0, 0, 115, 354, 1, 0, 0, 0, 117, 118, 5, 61, 0, 0, 118, 119, 5, 62,
		0, 0, 119, 2, 1, 0, 0, 0, 120, 121, 5, 91, 0, 0, 121, 4, 1, 0, 0, 0, 122,
		123, 5, 93, 0, 0, 123, 6, 1, 0, 0, 0, 124, 125, 5, 40, 0, 0, 125, 8, 1,
		0, 0, 0, 126, 127, 5, 41, 0, 0, 127, 10, 1, 0, 0, 0, 128, 129, 5, 123,
		0, 0, 129, 12, 1, 0, 0, 0, 130, 131, 5, 125, 0, 0, 131, 14, 1, 0, 0, 0,
		132, 133, 5, 59, 0, 0, 133, 16, 1, 0, 0, 0, 134, 135, 5, 44, 0, 0, 135,
		18, 1, 0, 0, 0, 136, 137, 5, 61, 0, 0, 137, 20, 1, 0, 0, 0, 138, 139, 5,
		63, 0, 0, 139, 22, 1, 0, 0, 0, 140, 141, 5, 63, 0, 0, 141, 142, 5, 46,
		0, 0, 142, 24, 1, 0, 0, 0, 143, 144, 5, 58, 0, 0, 144, 26, 1, 0, 0, 0,
		145, 146, 5, 46, 0, 0, 146, 28, 1, 0, 0, 0, 147, 148, 5, 43, 0, 0, 148,
		149, 5, 43, 0, 0, 149, 30, 1, 0, 0, 0, 150, 151, 5, 45, 0, 0, 151, 152,
		5, 45, 0, 0, 152, 32, 1, 0, 0, 0, 153, 154, 5, 43, 0, 0, 154, 34, 1, 0,
		0, 0, 155, 156, 5, 45, 0, 0, 156, 36, 1, 0, 0, 0, 157, 158, 5, 33, 0, 0,
		158, 38, 1, 0, 0, 0, 159, 160, 5, 42, 0, 0, 160, 40, 1, 0, 0, 0, 161, 162,
		5, 47, 0, 0, 162, 42, 1, 0, 0, 0, 163, 164, 5, 37, 0, 0, 164, 44, 1, 0,
		0, 0, 165, 166, 5, 42, 0, 0, 166, 167, 5, 42, 0, 0, 167, 46, 1, 0, 0, 0,
		168, 169, 5, 60, 0, 0, 169, 48, 1, 0, 0, 0, 170, 171, 5, 62, 0, 0, 171,
		50, 1, 0, 0, 0, 172, 173, 5, 60, 0, 0, 173, 174, 5, 61, 0, 0, 174, 52,
		1, 0, 0, 0, 175, 176, 5, 62, 0, 0, 176, 177, 5, 61, 0, 0, 177, 54, 1, 0,
		0, 0, 178, 179, 5, 61, 0, 0, 179, 180, 5, 61, 0, 0, 180, 56, 1, 0, 0, 0,
		181, 182, 5, 33, 0, 0, 182, 183, 5, 61, 0, 0, 183, 58, 1, 0, 0, 0, 184,
		185, 5, 38, 0, 0, 185, 186, 5, 38, 0, 0, 186, 60, 1, 0, 0, 0, 187, 188,
		5, 124, 0, 0, 188, 189, 5, 124, 0, 0, 189, 62, 1, 0, 0, 0, 190, 191, 5,
		110, 0, 0, 191, 192, 5, 117, 0, 0, 192, 193, 5, 108, 0, 0, 193, 194, 5,
		108, 0, 0, 194, 64, 1, 0, 0, 0, 195, 196, 5, 116, 0, 0, 196, 197, 5, 114,
		0, 0, 197, 198, 5, 117, 0, 0, 198, 205, 5, 101, 0, 0, 199, 200, 5, 102,
		0, 0, 200, 201, 5, 97, 0, 0, 201, 202, 5, 108, 0, 0, 202, 203, 5, 115,
		0, 0, 203, 205, 5, 101, 0, 0, 204, 195, 1, 0, 0, 0, 204, 199, 1, 0, 0,
		0, 205, 66, 1, 0, 0, 0, 206, 207, 3, 115, 57, 0, 207, 211, 5, 46, 0, 0,
		208, 210, 7, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211,
		209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 216, 1, 0, 0, 0, 213, 211,
		1, 0, 0, 0, 214, 216, 3, 115, 57, 0, 215, 206, 1, 0, 0, 0, 215, 214, 1,
		0, 0, 0, 216, 68, 1, 0, 0, 0, 217, 218, 5, 98, 0, 0, 218, 219, 5, 114,
		0, 0, 219, 220, 5, 101, 0, 0, 220, 221, 5, 97, 0, 0, 221, 222, 5, 107,
		0, 0, 222, 70, 1, 0, 0, 0, 223, 224, 5, 100, 0, 0, 224, 225, 5, 111, 0,
		0, 225, 72, 1, 0, 0, 0, 226, 227, 5, 111, 0, 0, 227, 228, 5, 102, 0, 0,
		228, 74, 1, 0, 0, 0, 229, 230, 5, 99, 0, 0, 230, 231, 5, 97, 0, 0, 231,
		232, 5, 115, 0, 0, 232, 233, 5, 101, 0, 0, 233, 76, 1, 0, 0, 0, 234, 235,
		5, 101, 0, 0, 235, 236, 5, 108, 0, 0, 236, 237, 5, 115, 0, 0, 237, 238,
		5, 101, 0, 0, 238, 78, 1, 0, 0, 0, 239, 240, 5, 110, 0, 0, 240, 241, 5,
		101, 0, 0, 241, 242, 5, 119, 0, 0, 242, 80, 1, 0, 0, 0, 243, 244, 5, 118,
		0, 0, 244, 245, 5, 97, 0, 0, 245, 246, 5, 114, 0, 0, 246, 82, 1, 0, 0,
		0, 247, 248, 5, 114, 0, 0, 248, 249, 5, 101, 0, 0, 249, 250, 5, 116, 0,
		0, 250, 251, 5, 117, 0, 0, 251, 252, 5, 114, 0, 0, 252, 253, 5, 110, 0,
		0, 253, 84, 1, 0, 0, 0, 254, 255, 5, 99, 0, 0, 255, 256, 5, 111, 0, 0,
		256, 257, 5, 110, 0, 0, 257, 258, 5, 116, 0, 0, 258, 259, 5, 105, 0, 0,
		259, 260, 5, 110, 0, 0, 260, 261, 5, 117, 0, 0, 261, 262, 5, 101, 0, 0,
		262, 86, 1, 0, 0, 0, 263, 264, 5, 102, 0, 0, 264, 265, 5, 111, 0, 0, 265,
		266, 5, 114, 0, 0, 266, 88, 1, 0, 0, 0, 267, 268, 5, 115, 0, 0, 268, 269,
		5, 119, 0, 0, 269, 270, 5, 105, 0, 0, 270, 271, 5, 116, 0, 0, 271, 272,
		5, 99, 0, 0, 272, 273, 5, 104, 0, 0, 273, 90, 1, 0, 0, 0, 274, 275, 5,
		119, 0, 0, 275, 276, 5, 104, 0, 0, 276, 277, 5, 105, 0, 0, 277, 278, 5,
		108, 0, 0, 278, 279, 5, 101, 0, 0, 279, 92, 1, 0, 0, 0, 280, 281, 5, 102,
		0, 0, 281, 282, 5, 117, 0, 0, 282, 283, 5, 110, 0, 0, 283, 284, 5, 99,
		0, 0, 284, 285, 5, 116, 0, 0, 285, 286, 5, 105, 0, 0, 286, 287, 5, 111,
		0, 0, 287, 288, 5, 110, 0, 0, 288, 94, 1, 0, 0, 0, 289, 290, 5, 100, 0,
		0, 290, 291, 5, 101, 0, 0, 291, 292, 5, 102, 0, 0, 292, 293, 5, 97, 0,
		0, 293, 294, 5, 117, 0, 0, 294, 295, 5, 108, 0, 0, 295, 296, 5, 116, 0,
		0, 296, 96, 1, 0, 0, 0, 297, 298, 5, 105, 0, 0, 298, 299, 5, 102, 0, 0,
		299, 98, 1, 0, 0, 0, 300, 301, 5, 101, 0, 0, 301, 302, 5, 108, 0, 0, 302,
		303, 5, 105, 0, 0, 303, 304, 5, 102, 0, 0, 304, 100, 1, 0, 0, 0, 305, 306,
		5, 100, 0, 0, 306, 307, 5, 101, 0, 0, 307, 308, 5, 108, 0, 0, 308, 309,
		5, 101, 0, 0, 309, 310, 5, 116, 0, 0, 310, 311, 5, 101, 0, 0, 311, 102,
		1, 0, 0, 0, 312, 316, 7, 1, 0, 0, 313, 315, 7, 2, 0, 0, 314, 313, 1, 0,
		0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0,
		317, 104, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 323, 5, 34, 0, 0, 320,
		322, 3, 111, 55, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321,
		1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 326, 1, 0, 0, 0, 325, 323, 1, 0,
		0, 0, 326, 327, 5, 34, 0, 0, 327, 106, 1, 0, 0, 0, 328, 330, 7, 3, 0, 0,
		329, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331,
		332, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 6, 53, 0, 0, 334, 108,
		1, 0, 0, 0, 335, 336, 7, 4, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338, 6, 54,
		0, 0, 338, 110, 1, 0, 0, 0, 339, 342, 8, 5, 0, 0, 340, 342, 3, 113, 56,
		0, 341, 339, 1, 0, 0, 0, 341, 340, 1, 0, 0, 0, 342, 112, 1, 0, 0, 0, 343,
		344, 5, 92, 0, 0, 344, 345, 7, 6, 0, 0, 345, 114, 1, 0, 0, 0, 346, 355,
		5, 48, 0, 0, 347, 351, 7, 7, 0, 0, 348, 350, 7, 8, 0, 0, 349, 348, 1, 0,
		0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0,
		352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 354, 346, 1, 0, 0, 0, 354,
		347, 1, 0, 0, 0, 355, 116, 1, 0, 0, 0, 10, 0, 204, 211, 215, 316, 323,
		331, 341, 351, 354, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LiteLangLexerInit initializes any static state used to implement LiteLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLiteLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LiteLangLexerInit() {
	staticData := &LiteLangLexerLexerStaticData
	staticData.once.Do(litelanglexerLexerInit)
}

// NewLiteLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLiteLangLexer(input antlr.CharStream) *LiteLangLexer {
	LiteLangLexerInit()
	l := new(LiteLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LiteLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "LiteLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LiteLangLexer tokens.
const (
	LiteLangLexerT__0              = 1
	LiteLangLexerOpenBracket       = 2
	LiteLangLexerCloseBracket      = 3
	LiteLangLexerOpenParen         = 4
	LiteLangLexerCloseParen        = 5
	LiteLangLexerOpenBrace         = 6
	LiteLangLexerCloseBrace        = 7
	LiteLangLexerSemiColon         = 8
	LiteLangLexerComma             = 9
	LiteLangLexerAssign            = 10
	LiteLangLexerQuestionMark      = 11
	LiteLangLexerQuestionMarkDot   = 12
	LiteLangLexerColon             = 13
	LiteLangLexerDot               = 14
	LiteLangLexerPlusPlus          = 15
	LiteLangLexerMinusMinus        = 16
	LiteLangLexerPlus              = 17
	LiteLangLexerMinus             = 18
	LiteLangLexerNot               = 19
	LiteLangLexerMultiply          = 20
	LiteLangLexerDivide            = 21
	LiteLangLexerModulus           = 22
	LiteLangLexerPower             = 23
	LiteLangLexerLessThan          = 24
	LiteLangLexerMoreThan          = 25
	LiteLangLexerLessThanEquals    = 26
	LiteLangLexerGreaterThanEquals = 27
	LiteLangLexerEquals            = 28
	LiteLangLexerNotEquals         = 29
	LiteLangLexerAnd               = 30
	LiteLangLexerOr                = 31
	LiteLangLexerNullLiteral       = 32
	LiteLangLexerBooleanLiteral    = 33
	LiteLangLexerDecimalLiteral    = 34
	LiteLangLexerBreak             = 35
	LiteLangLexerDo                = 36
	LiteLangLexerOf                = 37
	LiteLangLexerCase              = 38
	LiteLangLexerElse              = 39
	LiteLangLexerNew               = 40
	LiteLangLexerVar               = 41
	LiteLangLexerReturn            = 42
	LiteLangLexerContinue          = 43
	LiteLangLexerFor               = 44
	LiteLangLexerSwitch            = 45
	LiteLangLexerWhile             = 46
	LiteLangLexerFunction          = 47
	LiteLangLexerDefault           = 48
	LiteLangLexerIf                = 49
	LiteLangLexerElif              = 50
	LiteLangLexerDelete            = 51
	LiteLangLexerIdentifier        = 52
	LiteLangLexerStringLiteral     = 53
	LiteLangLexerWhiteSpaces       = 54
	LiteLangLexerLineTerminator    = 55
)
