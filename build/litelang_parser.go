// Code generated from litelang/src/LiteLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package LiteLang // LiteLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LiteLangParser struct {
	*antlr.BaseParser
}

var LiteLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func litelangParserInit() {
	staticData := &LiteLangParserStaticData
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
		"program", "statements", "statement", "variableStatement", "assignStatement",
		"functionStatement", "ifStatement", "forStatement", "classicForParam",
		"iteratorForParam", "switchStatement", "deleteStatement", "unitStatement",
		"valueStatement", "unit", "accessLhs", "accessRhs", "accessorRhs", "accessorLhs",
		"accessProperty", "functionCall", "paramMaker", "params", "arrowFunction",
		"arrayLiteral", "objectLiteral", "key", "value", "keyValue", "spread",
		"objectItem", "block", "constant",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 333, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 3, 0, 68, 8, 0, 1, 0, 1, 0, 1, 1, 4, 1, 73, 8,
		1, 11, 1, 12, 1, 74, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 94, 8, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 100, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 125, 8, 6, 10, 6, 12, 6, 128, 9, 6,
		1, 6, 1, 6, 3, 6, 132, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 138, 8, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 3, 8, 145, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 168, 8, 11, 5, 11, 170, 8,
		11, 10, 11, 12, 11, 173, 9, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 189, 8,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 206, 8, 14, 10, 14, 12, 14, 209,
		9, 14, 1, 15, 1, 15, 5, 15, 213, 8, 15, 10, 15, 12, 15, 216, 9, 15, 1,
		16, 1, 16, 3, 16, 220, 8, 16, 1, 16, 5, 16, 223, 8, 16, 10, 16, 12, 16,
		226, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 233, 8, 17, 1, 18,
		1, 18, 1, 18, 3, 18, 238, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 252, 8, 21, 10, 21,
		12, 21, 255, 9, 21, 3, 21, 257, 8, 21, 1, 22, 1, 22, 1, 22, 5, 22, 262,
		8, 22, 10, 22, 12, 22, 265, 9, 22, 3, 22, 267, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 279, 8, 24,
		10, 24, 12, 24, 282, 9, 24, 3, 24, 284, 8, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 5, 25, 292, 8, 25, 10, 25, 12, 25, 295, 9, 25, 3, 25,
		297, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 306,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 312, 8, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 325,
		8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 0, 1, 28, 33, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 3, 1, 0, 24, 31,
		2, 0, 34, 34, 52, 53, 2, 0, 32, 34, 53, 53, 343, 0, 67, 1, 0, 0, 0, 2,
		72, 1, 0, 0, 0, 4, 93, 1, 0, 0, 0, 6, 95, 1, 0, 0, 0, 8, 101, 1, 0, 0,
		0, 10, 105, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0, 14, 133, 1, 0, 0, 0, 16, 144,
		1, 0, 0, 0, 18, 151, 1, 0, 0, 0, 20, 156, 1, 0, 0, 0, 22, 162, 1, 0, 0,
		0, 24, 174, 1, 0, 0, 0, 26, 176, 1, 0, 0, 0, 28, 188, 1, 0, 0, 0, 30, 210,
		1, 0, 0, 0, 32, 219, 1, 0, 0, 0, 34, 232, 1, 0, 0, 0, 36, 237, 1, 0, 0,
		0, 38, 239, 1, 0, 0, 0, 40, 243, 1, 0, 0, 0, 42, 256, 1, 0, 0, 0, 44, 266,
		1, 0, 0, 0, 46, 268, 1, 0, 0, 0, 48, 274, 1, 0, 0, 0, 50, 287, 1, 0, 0,
		0, 52, 305, 1, 0, 0, 0, 54, 311, 1, 0, 0, 0, 56, 313, 1, 0, 0, 0, 58, 317,
		1, 0, 0, 0, 60, 324, 1, 0, 0, 0, 62, 326, 1, 0, 0, 0, 64, 330, 1, 0, 0,
		0, 66, 68, 3, 2, 1, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69,
		1, 0, 0, 0, 69, 70, 5, 0, 0, 1, 70, 1, 1, 0, 0, 0, 71, 73, 3, 4, 2, 0,
		72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1,
		0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 77, 3, 6, 3, 0, 77, 78, 5, 8, 0, 0, 78,
		94, 1, 0, 0, 0, 79, 80, 3, 8, 4, 0, 80, 81, 5, 8, 0, 0, 81, 94, 1, 0, 0,
		0, 82, 94, 3, 10, 5, 0, 83, 94, 3, 12, 6, 0, 84, 94, 3, 14, 7, 0, 85, 94,
		3, 20, 10, 0, 86, 94, 3, 26, 13, 0, 87, 88, 3, 22, 11, 0, 88, 89, 5, 8,
		0, 0, 89, 94, 1, 0, 0, 0, 90, 91, 3, 24, 12, 0, 91, 92, 5, 8, 0, 0, 92,
		94, 1, 0, 0, 0, 93, 76, 1, 0, 0, 0, 93, 79, 1, 0, 0, 0, 93, 82, 1, 0, 0,
		0, 93, 83, 1, 0, 0, 0, 93, 84, 1, 0, 0, 0, 93, 85, 1, 0, 0, 0, 93, 86,
		1, 0, 0, 0, 93, 87, 1, 0, 0, 0, 93, 90, 1, 0, 0, 0, 94, 5, 1, 0, 0, 0,
		95, 96, 5, 41, 0, 0, 96, 99, 5, 52, 0, 0, 97, 98, 5, 10, 0, 0, 98, 100,
		3, 54, 27, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 7, 1, 0, 0,
		0, 101, 102, 3, 30, 15, 0, 102, 103, 5, 10, 0, 0, 103, 104, 3, 54, 27,
		0, 104, 9, 1, 0, 0, 0, 105, 106, 5, 47, 0, 0, 106, 107, 5, 52, 0, 0, 107,
		108, 5, 4, 0, 0, 108, 109, 3, 42, 21, 0, 109, 110, 5, 5, 0, 0, 110, 111,
		3, 62, 31, 0, 111, 11, 1, 0, 0, 0, 112, 113, 5, 49, 0, 0, 113, 114, 5,
		4, 0, 0, 114, 115, 3, 28, 14, 0, 115, 116, 5, 5, 0, 0, 116, 126, 3, 62,
		31, 0, 117, 118, 5, 39, 0, 0, 118, 119, 5, 49, 0, 0, 119, 120, 5, 4, 0,
		0, 120, 121, 3, 28, 14, 0, 121, 122, 5, 5, 0, 0, 122, 123, 3, 62, 31, 0,
		123, 125, 1, 0, 0, 0, 124, 117, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 131, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 129, 130, 5, 39, 0, 0, 130, 132, 3, 62, 31, 0, 131, 129, 1,
		0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 13, 1, 0, 0, 0, 133, 134, 5, 44, 0,
		0, 134, 137, 5, 4, 0, 0, 135, 138, 3, 16, 8, 0, 136, 138, 3, 18, 9, 0,
		137, 135, 1, 0, 0, 0, 137, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 5, 5, 0, 0, 140, 141, 3, 62, 31, 0, 141, 15, 1, 0, 0, 0, 142, 145,
		3, 6, 3, 0, 143, 145, 3, 8, 4, 0, 144, 142, 1, 0, 0, 0, 144, 143, 1, 0,
		0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 5, 8, 0, 0, 147, 148, 3, 24, 12,
		0, 148, 149, 5, 8, 0, 0, 149, 150, 3, 8, 4, 0, 150, 17, 1, 0, 0, 0, 151,
		152, 5, 41, 0, 0, 152, 153, 5, 52, 0, 0, 153, 154, 5, 37, 0, 0, 154, 155,
		3, 32, 16, 0, 155, 19, 1, 0, 0, 0, 156, 157, 5, 45, 0, 0, 157, 158, 5,
		4, 0, 0, 158, 159, 3, 28, 14, 0, 159, 160, 5, 5, 0, 0, 160, 161, 3, 62,
		31, 0, 161, 21, 1, 0, 0, 0, 162, 163, 5, 51, 0, 0, 163, 171, 5, 52, 0,
		0, 164, 167, 5, 14, 0, 0, 165, 168, 3, 38, 19, 0, 166, 168, 5, 52, 0, 0,
		167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169,
		164, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 172,
		1, 0, 0, 0, 172, 23, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 175, 3, 28,
		14, 0, 175, 25, 1, 0, 0, 0, 176, 177, 3, 54, 27, 0, 177, 178, 5, 8, 0,
		0, 178, 27, 1, 0, 0, 0, 179, 180, 6, 14, -1, 0, 180, 181, 5, 4, 0, 0, 181,
		182, 3, 28, 14, 0, 182, 183, 5, 5, 0, 0, 183, 189, 1, 0, 0, 0, 184, 185,
		5, 19, 0, 0, 185, 189, 3, 28, 14, 3, 186, 189, 3, 32, 16, 0, 187, 189,
		3, 64, 32, 0, 188, 179, 1, 0, 0, 0, 188, 184, 1, 0, 0, 0, 188, 186, 1,
		0, 0, 0, 188, 187, 1, 0, 0, 0, 189, 207, 1, 0, 0, 0, 190, 191, 10, 8, 0,
		0, 191, 192, 5, 21, 0, 0, 192, 206, 3, 28, 14, 9, 193, 194, 10, 7, 0, 0,
		194, 195, 5, 20, 0, 0, 195, 206, 3, 28, 14, 8, 196, 197, 10, 6, 0, 0, 197,
		198, 5, 17, 0, 0, 198, 206, 3, 28, 14, 7, 199, 200, 10, 5, 0, 0, 200, 201,
		5, 18, 0, 0, 201, 206, 3, 28, 14, 6, 202, 203, 10, 4, 0, 0, 203, 204, 7,
		0, 0, 0, 204, 206, 3, 28, 14, 5, 205, 190, 1, 0, 0, 0, 205, 193, 1, 0,
		0, 0, 205, 196, 1, 0, 0, 0, 205, 199, 1, 0, 0, 0, 205, 202, 1, 0, 0, 0,
		206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208,
		29, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 214, 5, 52, 0, 0, 211, 213,
		3, 36, 18, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1,
		0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 31, 1, 0, 0, 0, 216, 214, 1, 0, 0,
		0, 217, 220, 3, 40, 20, 0, 218, 220, 5, 52, 0, 0, 219, 217, 1, 0, 0, 0,
		219, 218, 1, 0, 0, 0, 220, 224, 1, 0, 0, 0, 221, 223, 3, 34, 17, 0, 222,
		221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 33, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 233, 3, 38,
		19, 0, 228, 229, 5, 14, 0, 0, 229, 233, 5, 52, 0, 0, 230, 231, 5, 14, 0,
		0, 231, 233, 3, 40, 20, 0, 232, 227, 1, 0, 0, 0, 232, 228, 1, 0, 0, 0,
		232, 230, 1, 0, 0, 0, 233, 35, 1, 0, 0, 0, 234, 238, 3, 38, 19, 0, 235,
		236, 5, 14, 0, 0, 236, 238, 5, 52, 0, 0, 237, 234, 1, 0, 0, 0, 237, 235,
		1, 0, 0, 0, 238, 37, 1, 0, 0, 0, 239, 240, 5, 2, 0, 0, 240, 241, 7, 1,
		0, 0, 241, 242, 5, 3, 0, 0, 242, 39, 1, 0, 0, 0, 243, 244, 5, 52, 0, 0,
		244, 245, 5, 4, 0, 0, 245, 246, 3, 44, 22, 0, 246, 247, 5, 5, 0, 0, 247,
		41, 1, 0, 0, 0, 248, 253, 5, 52, 0, 0, 249, 250, 5, 9, 0, 0, 250, 252,
		5, 52, 0, 0, 251, 249, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0,
		0, 0, 253, 254, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0,
		256, 248, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 43, 1, 0, 0, 0, 258, 263,
		3, 28, 14, 0, 259, 260, 5, 9, 0, 0, 260, 262, 3, 28, 14, 0, 261, 259, 1,
		0, 0, 0, 262, 265, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0,
		0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 266, 258, 1, 0, 0, 0, 266,
		267, 1, 0, 0, 0, 267, 45, 1, 0, 0, 0, 268, 269, 5, 4, 0, 0, 269, 270, 3,
		44, 22, 0, 270, 271, 5, 5, 0, 0, 271, 272, 5, 1, 0, 0, 272, 273, 3, 62,
		31, 0, 273, 47, 1, 0, 0, 0, 274, 283, 5, 2, 0, 0, 275, 280, 3, 54, 27,
		0, 276, 277, 5, 9, 0, 0, 277, 279, 3, 54, 27, 0, 278, 276, 1, 0, 0, 0,
		279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281,
		284, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 275, 1, 0, 0, 0, 283, 284,
		1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 5, 3, 0, 0, 286, 49, 1, 0,
		0, 0, 287, 296, 5, 6, 0, 0, 288, 293, 3, 60, 30, 0, 289, 290, 5, 9, 0,
		0, 290, 292, 3, 60, 30, 0, 291, 289, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0,
		293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295,
		293, 1, 0, 0, 0, 296, 288, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298,
		1, 0, 0, 0, 298, 299, 5, 7, 0, 0, 299, 51, 1, 0, 0, 0, 300, 306, 5, 53,
		0, 0, 301, 302, 5, 2, 0, 0, 302, 303, 3, 32, 16, 0, 303, 304, 5, 3, 0,
		0, 304, 306, 1, 0, 0, 0, 305, 300, 1, 0, 0, 0, 305, 301, 1, 0, 0, 0, 306,
		53, 1, 0, 0, 0, 307, 312, 3, 40, 20, 0, 308, 312, 3, 28, 14, 0, 309, 312,
		3, 50, 25, 0, 310, 312, 3, 48, 24, 0, 311, 307, 1, 0, 0, 0, 311, 308, 1,
		0, 0, 0, 311, 309, 1, 0, 0, 0, 311, 310, 1, 0, 0, 0, 312, 55, 1, 0, 0,
		0, 313, 314, 3, 52, 26, 0, 314, 315, 5, 13, 0, 0, 315, 316, 3, 54, 27,
		0, 316, 57, 1, 0, 0, 0, 317, 318, 5, 14, 0, 0, 318, 319, 5, 14, 0, 0, 319,
		320, 5, 14, 0, 0, 320, 321, 3, 32, 16, 0, 321, 59, 1, 0, 0, 0, 322, 325,
		3, 56, 28, 0, 323, 325, 3, 58, 29, 0, 324, 322, 1, 0, 0, 0, 324, 323, 1,
		0, 0, 0, 325, 61, 1, 0, 0, 0, 326, 327, 5, 6, 0, 0, 327, 328, 3, 2, 1,
		0, 328, 329, 5, 7, 0, 0, 329, 63, 1, 0, 0, 0, 330, 331, 7, 2, 0, 0, 331,
		65, 1, 0, 0, 0, 29, 67, 74, 93, 99, 126, 131, 137, 144, 167, 171, 188,
		205, 207, 214, 219, 224, 232, 237, 253, 256, 263, 266, 280, 283, 293, 296,
		305, 311, 324,
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

// LiteLangParserInit initializes any static state used to implement LiteLangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLiteLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LiteLangParserInit() {
	staticData := &LiteLangParserStaticData
	staticData.once.Do(litelangParserInit)
}

// NewLiteLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLiteLangParser(input antlr.TokenStream) *LiteLangParser {
	LiteLangParserInit()
	this := new(LiteLangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LiteLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "LiteLang.g4"

	return this
}

// LiteLangParser tokens.
const (
	LiteLangParserEOF               = antlr.TokenEOF
	LiteLangParserT__0              = 1
	LiteLangParserOpenBracket       = 2
	LiteLangParserCloseBracket      = 3
	LiteLangParserOpenParen         = 4
	LiteLangParserCloseParen        = 5
	LiteLangParserOpenBrace         = 6
	LiteLangParserCloseBrace        = 7
	LiteLangParserSemiColon         = 8
	LiteLangParserComma             = 9
	LiteLangParserAssign            = 10
	LiteLangParserQuestionMark      = 11
	LiteLangParserQuestionMarkDot   = 12
	LiteLangParserColon             = 13
	LiteLangParserDot               = 14
	LiteLangParserPlusPlus          = 15
	LiteLangParserMinusMinus        = 16
	LiteLangParserPlus              = 17
	LiteLangParserMinus             = 18
	LiteLangParserNot               = 19
	LiteLangParserMultiply          = 20
	LiteLangParserDivide            = 21
	LiteLangParserModulus           = 22
	LiteLangParserPower             = 23
	LiteLangParserLessThan          = 24
	LiteLangParserMoreThan          = 25
	LiteLangParserLessThanEquals    = 26
	LiteLangParserGreaterThanEquals = 27
	LiteLangParserEquals            = 28
	LiteLangParserNotEquals         = 29
	LiteLangParserAnd               = 30
	LiteLangParserOr                = 31
	LiteLangParserNullLiteral       = 32
	LiteLangParserBooleanLiteral    = 33
	LiteLangParserDecimalLiteral    = 34
	LiteLangParserBreak             = 35
	LiteLangParserDo                = 36
	LiteLangParserOf                = 37
	LiteLangParserCase              = 38
	LiteLangParserElse              = 39
	LiteLangParserNew               = 40
	LiteLangParserVar               = 41
	LiteLangParserReturn            = 42
	LiteLangParserContinue          = 43
	LiteLangParserFor               = 44
	LiteLangParserSwitch            = 45
	LiteLangParserWhile             = 46
	LiteLangParserFunction          = 47
	LiteLangParserDefault           = 48
	LiteLangParserIf                = 49
	LiteLangParserElif              = 50
	LiteLangParserDelete            = 51
	LiteLangParserIdentifier        = 52
	LiteLangParserStringLiteral     = 53
	LiteLangParserWhiteSpaces       = 54
	LiteLangParserLineTerminator    = 55
)

// LiteLangParser rules.
const (
	LiteLangParserRULE_program           = 0
	LiteLangParserRULE_statements        = 1
	LiteLangParserRULE_statement         = 2
	LiteLangParserRULE_variableStatement = 3
	LiteLangParserRULE_assignStatement   = 4
	LiteLangParserRULE_functionStatement = 5
	LiteLangParserRULE_ifStatement       = 6
	LiteLangParserRULE_forStatement      = 7
	LiteLangParserRULE_classicForParam   = 8
	LiteLangParserRULE_iteratorForParam  = 9
	LiteLangParserRULE_switchStatement   = 10
	LiteLangParserRULE_deleteStatement   = 11
	LiteLangParserRULE_unitStatement     = 12
	LiteLangParserRULE_valueStatement    = 13
	LiteLangParserRULE_unit              = 14
	LiteLangParserRULE_accessLhs         = 15
	LiteLangParserRULE_accessRhs         = 16
	LiteLangParserRULE_accessorRhs       = 17
	LiteLangParserRULE_accessorLhs       = 18
	LiteLangParserRULE_accessProperty    = 19
	LiteLangParserRULE_functionCall      = 20
	LiteLangParserRULE_paramMaker        = 21
	LiteLangParserRULE_params            = 22
	LiteLangParserRULE_arrowFunction     = 23
	LiteLangParserRULE_arrayLiteral      = 24
	LiteLangParserRULE_objectLiteral     = 25
	LiteLangParserRULE_key               = 26
	LiteLangParserRULE_value             = 27
	LiteLangParserRULE_keyValue          = 28
	LiteLangParserRULE_spread            = 29
	LiteLangParserRULE_objectItem        = 30
	LiteLangParserRULE_block             = 31
	LiteLangParserRULE_constant          = 32
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Statements() IStatementsContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(LiteLangParserEOF, 0)
}

func (s *ProgramContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LiteLangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16521291784257620) != 0 {
		{
			p.SetState(66)
			p.Statements()
		}

	}
	{
		p.SetState(69)
		p.Match(LiteLangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LiteLangParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16521291784257620) != 0) {
		{
			p.SetState(71)
			p.Statement()
		}

		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableStatement() IVariableStatementContext
	SemiColon() antlr.TerminalNode
	AssignStatement() IAssignStatementContext
	FunctionStatement() IFunctionStatementContext
	IfStatement() IIfStatementContext
	ForStatement() IForStatementContext
	SwitchStatement() ISwitchStatementContext
	ValueStatement() IValueStatementContext
	DeleteStatement() IDeleteStatementContext
	UnitStatement() IUnitStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableStatement() IVariableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableStatementContext)
}

func (s *StatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(LiteLangParserSemiColon, 0)
}

func (s *StatementContext) AssignStatement() IAssignStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *StatementContext) FunctionStatement() IFunctionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ValueStatement() IValueStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueStatementContext)
}

func (s *StatementContext) DeleteStatement() IDeleteStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *StatementContext) UnitStatement() IUnitStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LiteLangParserRULE_statement)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.VariableStatement()
		}
		{
			p.SetState(77)
			p.Match(LiteLangParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.AssignStatement()
		}
		{
			p.SetState(80)
			p.Match(LiteLangParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.FunctionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.ForStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(85)
			p.SwitchStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(86)
			p.ValueStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(87)
			p.DeleteStatement()
		}
		{
			p.SetState(88)
			p.Match(LiteLangParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(90)
			p.UnitStatement()
		}
		{
			p.SetState(91)
			p.Match(LiteLangParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableStatementContext is an interface to support dynamic dispatch.
type IVariableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Assign() antlr.TerminalNode
	Value() IValueContext

	// IsVariableStatementContext differentiates from other interfaces.
	IsVariableStatementContext()
}

type VariableStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableStatementContext() *VariableStatementContext {
	var p = new(VariableStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_variableStatement
	return p
}

func InitEmptyVariableStatementContext(p *VariableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_variableStatement
}

func (*VariableStatementContext) IsVariableStatementContext() {}

func NewVariableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableStatementContext {
	var p = new(VariableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_variableStatement

	return p
}

func (s *VariableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(LiteLangParserVar, 0)
}

func (s *VariableStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *VariableStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(LiteLangParserAssign, 0)
}

func (s *VariableStatementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *VariableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

func (s *VariableStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitVariableStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) VariableStatement() (localctx IVariableStatementContext) {
	localctx = NewVariableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LiteLangParserRULE_variableStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(LiteLangParserVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(LiteLangParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LiteLangParserAssign {
		{
			p.SetState(97)
			p.Match(LiteLangParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Value()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignStatementContext is an interface to support dynamic dispatch.
type IAssignStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AccessLhs() IAccessLhsContext
	Assign() antlr.TerminalNode
	Value() IValueContext

	// IsAssignStatementContext differentiates from other interfaces.
	IsAssignStatementContext()
}

type AssignStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStatementContext() *AssignStatementContext {
	var p = new(AssignStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_assignStatement
	return p
}

func InitEmptyAssignStatementContext(p *AssignStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_assignStatement
}

func (*AssignStatementContext) IsAssignStatementContext() {}

func NewAssignStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStatementContext {
	var p = new(AssignStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_assignStatement

	return p
}

func (s *AssignStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStatementContext) AccessLhs() IAccessLhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessLhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessLhsContext)
}

func (s *AssignStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(LiteLangParserAssign, 0)
}

func (s *AssignStatementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) AssignStatement() (localctx IAssignStatementContext) {
	localctx = NewAssignStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LiteLangParserRULE_assignStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.AccessLhs()
	}
	{
		p.SetState(102)
		p.Match(LiteLangParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionStatementContext is an interface to support dynamic dispatch.
type IFunctionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	OpenParen() antlr.TerminalNode
	ParamMaker() IParamMakerContext
	CloseParen() antlr.TerminalNode
	Block() IBlockContext

	// IsFunctionStatementContext differentiates from other interfaces.
	IsFunctionStatementContext()
}

type FunctionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStatementContext() *FunctionStatementContext {
	var p = new(FunctionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_functionStatement
	return p
}

func InitEmptyFunctionStatementContext(p *FunctionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_functionStatement
}

func (*FunctionStatementContext) IsFunctionStatementContext() {}

func NewFunctionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStatementContext {
	var p = new(FunctionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_functionStatement

	return p
}

func (s *FunctionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStatementContext) Function() antlr.TerminalNode {
	return s.GetToken(LiteLangParserFunction, 0)
}

func (s *FunctionStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *FunctionStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, 0)
}

func (s *FunctionStatementContext) ParamMaker() IParamMakerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamMakerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamMakerContext)
}

func (s *FunctionStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, 0)
}

func (s *FunctionStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitFunctionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) FunctionStatement() (localctx IFunctionStatementContext) {
	localctx = NewFunctionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LiteLangParserRULE_functionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(LiteLangParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(LiteLangParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(LiteLangParserOpenParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.ParamMaker()
	}
	{
		p.SetState(109)
		p.Match(LiteLangParserCloseParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIf() []antlr.TerminalNode
	If(i int) antlr.TerminalNode
	AllOpenParen() []antlr.TerminalNode
	OpenParen(i int) antlr.TerminalNode
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	AllCloseParen() []antlr.TerminalNode
	CloseParen(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllElse() []antlr.TerminalNode
	Else(i int) antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) AllIf() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserIf)
}

func (s *IfStatementContext) If(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserIf, i)
}

func (s *IfStatementContext) AllOpenParen() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserOpenParen)
}

func (s *IfStatementContext) OpenParen(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, i)
}

func (s *IfStatementContext) AllUnit() []IUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitContext); ok {
			len++
		}
	}

	tst := make([]IUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitContext); ok {
			tst[i] = t.(IUnitContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Unit(i int) IUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *IfStatementContext) AllCloseParen() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserCloseParen)
}

func (s *IfStatementContext) CloseParen(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, i)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) AllElse() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserElse)
}

func (s *IfStatementContext) Else(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserElse, i)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LiteLangParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(LiteLangParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(LiteLangParserOpenParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.unit(0)
	}
	{
		p.SetState(115)
		p.Match(LiteLangParserCloseParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Block()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.Match(LiteLangParserElse)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(118)
				p.Match(LiteLangParserIf)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(119)
				p.Match(LiteLangParserOpenParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(120)
				p.unit(0)
			}
			{
				p.SetState(121)
				p.Match(LiteLangParserCloseParen)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(122)
				p.Block()
			}

		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LiteLangParserElse {
		{
			p.SetState(129)
			p.Match(LiteLangParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	OpenParen() antlr.TerminalNode
	CloseParen() antlr.TerminalNode
	Block() IBlockContext
	ClassicForParam() IClassicForParamContext
	IteratorForParam() IIteratorForParamContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(LiteLangParserFor, 0)
}

func (s *ForStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, 0)
}

func (s *ForStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, 0)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) ClassicForParam() IClassicForParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassicForParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassicForParamContext)
}

func (s *ForStatementContext) IteratorForParam() IIteratorForParamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorForParamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorForParamContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LiteLangParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(LiteLangParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(LiteLangParserOpenParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(135)
			p.ClassicForParam()
		}

	case 2:
		{
			p.SetState(136)
			p.IteratorForParam()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(139)
		p.Match(LiteLangParserCloseParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassicForParamContext is an interface to support dynamic dispatch.
type IClassicForParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondition returns the condition rule contexts.
	GetCondition() IUnitStatementContext

	// GetPostOp returns the postOp rule contexts.
	GetPostOp() IAssignStatementContext

	// SetCondition sets the condition rule contexts.
	SetCondition(IUnitStatementContext)

	// SetPostOp sets the postOp rule contexts.
	SetPostOp(IAssignStatementContext)

	// Getter signatures
	AllSemiColon() []antlr.TerminalNode
	SemiColon(i int) antlr.TerminalNode
	UnitStatement() IUnitStatementContext
	AllAssignStatement() []IAssignStatementContext
	AssignStatement(i int) IAssignStatementContext
	VariableStatement() IVariableStatementContext

	// IsClassicForParamContext differentiates from other interfaces.
	IsClassicForParamContext()
}

type ClassicForParamContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	condition IUnitStatementContext
	postOp    IAssignStatementContext
}

func NewEmptyClassicForParamContext() *ClassicForParamContext {
	var p = new(ClassicForParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_classicForParam
	return p
}

func InitEmptyClassicForParamContext(p *ClassicForParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_classicForParam
}

func (*ClassicForParamContext) IsClassicForParamContext() {}

func NewClassicForParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicForParamContext {
	var p = new(ClassicForParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_classicForParam

	return p
}

func (s *ClassicForParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicForParamContext) GetCondition() IUnitStatementContext { return s.condition }

func (s *ClassicForParamContext) GetPostOp() IAssignStatementContext { return s.postOp }

func (s *ClassicForParamContext) SetCondition(v IUnitStatementContext) { s.condition = v }

func (s *ClassicForParamContext) SetPostOp(v IAssignStatementContext) { s.postOp = v }

func (s *ClassicForParamContext) AllSemiColon() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserSemiColon)
}

func (s *ClassicForParamContext) SemiColon(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserSemiColon, i)
}

func (s *ClassicForParamContext) UnitStatement() IUnitStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitStatementContext)
}

func (s *ClassicForParamContext) AllAssignStatement() []IAssignStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignStatementContext); ok {
			len++
		}
	}

	tst := make([]IAssignStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignStatementContext); ok {
			tst[i] = t.(IAssignStatementContext)
			i++
		}
	}

	return tst
}

func (s *ClassicForParamContext) AssignStatement(i int) IAssignStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStatementContext)
}

func (s *ClassicForParamContext) VariableStatement() IVariableStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableStatementContext)
}

func (s *ClassicForParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicForParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicForParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterClassicForParam(s)
	}
}

func (s *ClassicForParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitClassicForParam(s)
	}
}

func (s *ClassicForParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitClassicForParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ClassicForParam() (localctx IClassicForParamContext) {
	localctx = NewClassicForParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LiteLangParserRULE_classicForParam)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LiteLangParserVar:
		{
			p.SetState(142)
			p.VariableStatement()
		}

	case LiteLangParserIdentifier:
		{
			p.SetState(143)
			p.AssignStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(146)
		p.Match(LiteLangParserSemiColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)

		var _x = p.UnitStatement()

		localctx.(*ClassicForParamContext).condition = _x
	}
	{
		p.SetState(148)
		p.Match(LiteLangParserSemiColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)

		var _x = p.AssignStatement()

		localctx.(*ClassicForParamContext).postOp = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIteratorForParamContext is an interface to support dynamic dispatch.
type IIteratorForParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Of() antlr.TerminalNode
	AccessRhs() IAccessRhsContext

	// IsIteratorForParamContext differentiates from other interfaces.
	IsIteratorForParamContext()
}

type IteratorForParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteratorForParamContext() *IteratorForParamContext {
	var p = new(IteratorForParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_iteratorForParam
	return p
}

func InitEmptyIteratorForParamContext(p *IteratorForParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_iteratorForParam
}

func (*IteratorForParamContext) IsIteratorForParamContext() {}

func NewIteratorForParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IteratorForParamContext {
	var p = new(IteratorForParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_iteratorForParam

	return p
}

func (s *IteratorForParamContext) GetParser() antlr.Parser { return s.parser }

func (s *IteratorForParamContext) Var() antlr.TerminalNode {
	return s.GetToken(LiteLangParserVar, 0)
}

func (s *IteratorForParamContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *IteratorForParamContext) Of() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOf, 0)
}

func (s *IteratorForParamContext) AccessRhs() IAccessRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessRhsContext)
}

func (s *IteratorForParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteratorForParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IteratorForParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterIteratorForParam(s)
	}
}

func (s *IteratorForParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitIteratorForParam(s)
	}
}

func (s *IteratorForParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitIteratorForParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) IteratorForParam() (localctx IIteratorForParamContext) {
	localctx = NewIteratorForParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LiteLangParserRULE_iteratorForParam)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(LiteLangParserVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(LiteLangParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(LiteLangParserOf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.AccessRhs()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Switch() antlr.TerminalNode
	OpenParen() antlr.TerminalNode
	Unit() IUnitContext
	CloseParen() antlr.TerminalNode
	Block() IBlockContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(LiteLangParserSwitch, 0)
}

func (s *SwitchStatementContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, 0)
}

func (s *SwitchStatementContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *SwitchStatementContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, 0)
}

func (s *SwitchStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LiteLangParserRULE_switchStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(LiteLangParserSwitch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.Match(LiteLangParserOpenParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.unit(0)
	}
	{
		p.SetState(159)
		p.Match(LiteLangParserCloseParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Delete() antlr.TerminalNode
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllDot() []antlr.TerminalNode
	Dot(i int) antlr.TerminalNode
	AllAccessProperty() []IAccessPropertyContext
	AccessProperty(i int) IAccessPropertyContext

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_deleteStatement
	return p
}

func InitEmptyDeleteStatementContext(p *DeleteStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_deleteStatement
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) Delete() antlr.TerminalNode {
	return s.GetToken(LiteLangParserDelete, 0)
}

func (s *DeleteStatementContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserIdentifier)
}

func (s *DeleteStatementContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, i)
}

func (s *DeleteStatementContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserDot)
}

func (s *DeleteStatementContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserDot, i)
}

func (s *DeleteStatementContext) AllAccessProperty() []IAccessPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessPropertyContext); ok {
			len++
		}
	}

	tst := make([]IAccessPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessPropertyContext); ok {
			tst[i] = t.(IAccessPropertyContext)
			i++
		}
	}

	return tst
}

func (s *DeleteStatementContext) AccessProperty(i int) IAccessPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessPropertyContext)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitDeleteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LiteLangParserRULE_deleteStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(LiteLangParserDelete)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(163)
		p.Match(LiteLangParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LiteLangParserDot {
		{
			p.SetState(164)
			p.Match(LiteLangParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LiteLangParserOpenBracket:
			{
				p.SetState(165)
				p.AccessProperty()
			}

		case LiteLangParserIdentifier:
			{
				p.SetState(166)
				p.Match(LiteLangParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnitStatementContext is an interface to support dynamic dispatch.
type IUnitStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unit() IUnitContext

	// IsUnitStatementContext differentiates from other interfaces.
	IsUnitStatementContext()
}

type UnitStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitStatementContext() *UnitStatementContext {
	var p = new(UnitStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_unitStatement
	return p
}

func InitEmptyUnitStatementContext(p *UnitStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_unitStatement
}

func (*UnitStatementContext) IsUnitStatementContext() {}

func NewUnitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitStatementContext {
	var p = new(UnitStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_unitStatement

	return p
}

func (s *UnitStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitStatementContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *UnitStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterUnitStatement(s)
	}
}

func (s *UnitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitUnitStatement(s)
	}
}

func (s *UnitStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitUnitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) UnitStatement() (localctx IUnitStatementContext) {
	localctx = NewUnitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LiteLangParserRULE_unitStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.unit(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueStatementContext is an interface to support dynamic dispatch.
type IValueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	SemiColon() antlr.TerminalNode

	// IsValueStatementContext differentiates from other interfaces.
	IsValueStatementContext()
}

type ValueStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueStatementContext() *ValueStatementContext {
	var p = new(ValueStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_valueStatement
	return p
}

func InitEmptyValueStatementContext(p *ValueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_valueStatement
}

func (*ValueStatementContext) IsValueStatementContext() {}

func NewValueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueStatementContext {
	var p = new(ValueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_valueStatement

	return p
}

func (s *ValueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueStatementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(LiteLangParserSemiColon, 0)
}

func (s *ValueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterValueStatement(s)
	}
}

func (s *ValueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitValueStatement(s)
	}
}

func (s *ValueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitValueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ValueStatement() (localctx IValueStatementContext) {
	localctx = NewValueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LiteLangParserRULE_valueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Value()
	}
	{
		p.SetState(177)
		p.Match(LiteLangParserSemiColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IUnitContext

	// GetSingle returns the single rule contexts.
	GetSingle() IUnitContext

	// GetRight returns the right rule contexts.
	GetRight() IUnitContext

	// SetLeft sets the left rule contexts.
	SetLeft(IUnitContext)

	// SetSingle sets the single rule contexts.
	SetSingle(IUnitContext)

	// SetRight sets the right rule contexts.
	SetRight(IUnitContext)

	// Getter signatures
	OpenParen() antlr.TerminalNode
	CloseParen() antlr.TerminalNode
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	Not() antlr.TerminalNode
	AccessRhs() IAccessRhsContext
	Constant() IConstantContext
	Divide() antlr.TerminalNode
	Multiply() antlr.TerminalNode
	Plus() antlr.TerminalNode
	Minus() antlr.TerminalNode
	LessThan() antlr.TerminalNode
	MoreThan() antlr.TerminalNode
	LessThanEquals() antlr.TerminalNode
	GreaterThanEquals() antlr.TerminalNode
	Equals() antlr.TerminalNode
	NotEquals() antlr.TerminalNode
	And() antlr.TerminalNode
	Or() antlr.TerminalNode

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IUnitContext
	single IUnitContext
	right  IUnitContext
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) GetLeft() IUnitContext { return s.left }

func (s *UnitContext) GetSingle() IUnitContext { return s.single }

func (s *UnitContext) GetRight() IUnitContext { return s.right }

func (s *UnitContext) SetLeft(v IUnitContext) { s.left = v }

func (s *UnitContext) SetSingle(v IUnitContext) { s.single = v }

func (s *UnitContext) SetRight(v IUnitContext) { s.right = v }

func (s *UnitContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, 0)
}

func (s *UnitContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, 0)
}

func (s *UnitContext) AllUnit() []IUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitContext); ok {
			len++
		}
	}

	tst := make([]IUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitContext); ok {
			tst[i] = t.(IUnitContext)
			i++
		}
	}

	return tst
}

func (s *UnitContext) Unit(i int) IUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *UnitContext) Not() antlr.TerminalNode {
	return s.GetToken(LiteLangParserNot, 0)
}

func (s *UnitContext) AccessRhs() IAccessRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessRhsContext)
}

func (s *UnitContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *UnitContext) Divide() antlr.TerminalNode {
	return s.GetToken(LiteLangParserDivide, 0)
}

func (s *UnitContext) Multiply() antlr.TerminalNode {
	return s.GetToken(LiteLangParserMultiply, 0)
}

func (s *UnitContext) Plus() antlr.TerminalNode {
	return s.GetToken(LiteLangParserPlus, 0)
}

func (s *UnitContext) Minus() antlr.TerminalNode {
	return s.GetToken(LiteLangParserMinus, 0)
}

func (s *UnitContext) LessThan() antlr.TerminalNode {
	return s.GetToken(LiteLangParserLessThan, 0)
}

func (s *UnitContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(LiteLangParserMoreThan, 0)
}

func (s *UnitContext) LessThanEquals() antlr.TerminalNode {
	return s.GetToken(LiteLangParserLessThanEquals, 0)
}

func (s *UnitContext) GreaterThanEquals() antlr.TerminalNode {
	return s.GetToken(LiteLangParserGreaterThanEquals, 0)
}

func (s *UnitContext) Equals() antlr.TerminalNode {
	return s.GetToken(LiteLangParserEquals, 0)
}

func (s *UnitContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(LiteLangParserNotEquals, 0)
}

func (s *UnitContext) And() antlr.TerminalNode {
	return s.GetToken(LiteLangParserAnd, 0)
}

func (s *UnitContext) Or() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOr, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (s *UnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Unit() (localctx IUnitContext) {
	return p.unit(0)
}

func (p *LiteLangParser) unit(_p int) (localctx IUnitContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewUnitContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IUnitContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, LiteLangParserRULE_unit, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LiteLangParserOpenParen:
		{
			p.SetState(180)
			p.Match(LiteLangParserOpenParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)

			var _x = p.unit(0)

			localctx.(*UnitContext).single = _x
		}
		{
			p.SetState(182)
			p.Match(LiteLangParserCloseParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LiteLangParserNot:
		{
			p.SetState(184)
			p.Match(LiteLangParserNot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)

			var _x = p.unit(3)

			localctx.(*UnitContext).single = _x
		}

	case LiteLangParserIdentifier:
		{
			p.SetState(186)
			p.AccessRhs()
		}

	case LiteLangParserNullLiteral, LiteLangParserBooleanLiteral, LiteLangParserDecimalLiteral, LiteLangParserStringLiteral:
		{
			p.SetState(187)
			p.Constant()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LiteLangParserRULE_unit)
				p.SetState(190)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}

				{
					p.SetState(191)
					p.Match(LiteLangParserDivide)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(192)

					var _x = p.unit(9)

					localctx.(*UnitContext).right = _x
				}

			case 2:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LiteLangParserRULE_unit)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}

				{
					p.SetState(194)
					p.Match(LiteLangParserMultiply)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(195)

					var _x = p.unit(8)

					localctx.(*UnitContext).right = _x
				}

			case 3:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LiteLangParserRULE_unit)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}

				{
					p.SetState(197)
					p.Match(LiteLangParserPlus)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(198)

					var _x = p.unit(7)

					localctx.(*UnitContext).right = _x
				}

			case 4:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LiteLangParserRULE_unit)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}

				{
					p.SetState(200)
					p.Match(LiteLangParserMinus)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(201)

					var _x = p.unit(6)

					localctx.(*UnitContext).right = _x
				}

			case 5:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LiteLangParserRULE_unit)
				p.SetState(202)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(203)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4278190080) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(204)

					var _x = p.unit(5)

					localctx.(*UnitContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessLhsContext is an interface to support dynamic dispatch.
type IAccessLhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	AllAccessorLhs() []IAccessorLhsContext
	AccessorLhs(i int) IAccessorLhsContext

	// IsAccessLhsContext differentiates from other interfaces.
	IsAccessLhsContext()
}

type AccessLhsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessLhsContext() *AccessLhsContext {
	var p = new(AccessLhsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessLhs
	return p
}

func InitEmptyAccessLhsContext(p *AccessLhsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessLhs
}

func (*AccessLhsContext) IsAccessLhsContext() {}

func NewAccessLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessLhsContext {
	var p = new(AccessLhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_accessLhs

	return p
}

func (s *AccessLhsContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessLhsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *AccessLhsContext) AllAccessorLhs() []IAccessorLhsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessorLhsContext); ok {
			len++
		}
	}

	tst := make([]IAccessorLhsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessorLhsContext); ok {
			tst[i] = t.(IAccessorLhsContext)
			i++
		}
	}

	return tst
}

func (s *AccessLhsContext) AccessorLhs(i int) IAccessorLhsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorLhsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorLhsContext)
}

func (s *AccessLhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessLhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessLhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterAccessLhs(s)
	}
}

func (s *AccessLhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitAccessLhs(s)
	}
}

func (s *AccessLhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitAccessLhs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) AccessLhs() (localctx IAccessLhsContext) {
	localctx = NewAccessLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LiteLangParserRULE_accessLhs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(LiteLangParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LiteLangParserOpenBracket || _la == LiteLangParserDot {
		{
			p.SetState(211)
			p.AccessorLhs()
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessRhsContext is an interface to support dynamic dispatch.
type IAccessRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	Identifier() antlr.TerminalNode
	AllAccessorRhs() []IAccessorRhsContext
	AccessorRhs(i int) IAccessorRhsContext

	// IsAccessRhsContext differentiates from other interfaces.
	IsAccessRhsContext()
}

type AccessRhsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessRhsContext() *AccessRhsContext {
	var p = new(AccessRhsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessRhs
	return p
}

func InitEmptyAccessRhsContext(p *AccessRhsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessRhs
}

func (*AccessRhsContext) IsAccessRhsContext() {}

func NewAccessRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessRhsContext {
	var p = new(AccessRhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_accessRhs

	return p
}

func (s *AccessRhsContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessRhsContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *AccessRhsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *AccessRhsContext) AllAccessorRhs() []IAccessorRhsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessorRhsContext); ok {
			len++
		}
	}

	tst := make([]IAccessorRhsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessorRhsContext); ok {
			tst[i] = t.(IAccessorRhsContext)
			i++
		}
	}

	return tst
}

func (s *AccessRhsContext) AccessorRhs(i int) IAccessorRhsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorRhsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessorRhsContext)
}

func (s *AccessRhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessRhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessRhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterAccessRhs(s)
	}
}

func (s *AccessRhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitAccessRhs(s)
	}
}

func (s *AccessRhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitAccessRhs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) AccessRhs() (localctx IAccessRhsContext) {
	localctx = NewAccessRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LiteLangParserRULE_accessRhs)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(217)
			p.FunctionCall()
		}

	case 2:
		{
			p.SetState(218)
			p.Match(LiteLangParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(221)
				p.AccessorRhs()
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessorRhsContext is an interface to support dynamic dispatch.
type IAccessorRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AccessProperty() IAccessPropertyContext
	Dot() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	FunctionCall() IFunctionCallContext

	// IsAccessorRhsContext differentiates from other interfaces.
	IsAccessorRhsContext()
}

type AccessorRhsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorRhsContext() *AccessorRhsContext {
	var p = new(AccessorRhsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessorRhs
	return p
}

func InitEmptyAccessorRhsContext(p *AccessorRhsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessorRhs
}

func (*AccessorRhsContext) IsAccessorRhsContext() {}

func NewAccessorRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorRhsContext {
	var p = new(AccessorRhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_accessorRhs

	return p
}

func (s *AccessorRhsContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorRhsContext) AccessProperty() IAccessPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessPropertyContext)
}

func (s *AccessorRhsContext) Dot() antlr.TerminalNode {
	return s.GetToken(LiteLangParserDot, 0)
}

func (s *AccessorRhsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *AccessorRhsContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *AccessorRhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorRhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorRhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterAccessorRhs(s)
	}
}

func (s *AccessorRhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitAccessorRhs(s)
	}
}

func (s *AccessorRhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitAccessorRhs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) AccessorRhs() (localctx IAccessorRhsContext) {
	localctx = NewAccessorRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LiteLangParserRULE_accessorRhs)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(227)
			p.AccessProperty()
		}

	case 2:
		{
			p.SetState(228)
			p.Match(LiteLangParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(LiteLangParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(230)
			p.Match(LiteLangParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.FunctionCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessorLhsContext is an interface to support dynamic dispatch.
type IAccessorLhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AccessProperty() IAccessPropertyContext
	Dot() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsAccessorLhsContext differentiates from other interfaces.
	IsAccessorLhsContext()
}

type AccessorLhsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorLhsContext() *AccessorLhsContext {
	var p = new(AccessorLhsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessorLhs
	return p
}

func InitEmptyAccessorLhsContext(p *AccessorLhsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessorLhs
}

func (*AccessorLhsContext) IsAccessorLhsContext() {}

func NewAccessorLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorLhsContext {
	var p = new(AccessorLhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_accessorLhs

	return p
}

func (s *AccessorLhsContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorLhsContext) AccessProperty() IAccessPropertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessPropertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessPropertyContext)
}

func (s *AccessorLhsContext) Dot() antlr.TerminalNode {
	return s.GetToken(LiteLangParserDot, 0)
}

func (s *AccessorLhsContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *AccessorLhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorLhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorLhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterAccessorLhs(s)
	}
}

func (s *AccessorLhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitAccessorLhs(s)
	}
}

func (s *AccessorLhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitAccessorLhs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) AccessorLhs() (localctx IAccessorLhsContext) {
	localctx = NewAccessorLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LiteLangParserRULE_accessorLhs)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LiteLangParserOpenBracket:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.AccessProperty()
		}

	case LiteLangParserDot:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(LiteLangParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(LiteLangParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAccessPropertyContext is an interface to support dynamic dispatch.
type IAccessPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenBracket() antlr.TerminalNode
	CloseBracket() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	DecimalLiteral() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsAccessPropertyContext differentiates from other interfaces.
	IsAccessPropertyContext()
}

type AccessPropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessPropertyContext() *AccessPropertyContext {
	var p = new(AccessPropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessProperty
	return p
}

func InitEmptyAccessPropertyContext(p *AccessPropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_accessProperty
}

func (*AccessPropertyContext) IsAccessPropertyContext() {}

func NewAccessPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessPropertyContext {
	var p = new(AccessPropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_accessProperty

	return p
}

func (s *AccessPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessPropertyContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenBracket, 0)
}

func (s *AccessPropertyContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseBracket, 0)
}

func (s *AccessPropertyContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserStringLiteral, 0)
}

func (s *AccessPropertyContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserDecimalLiteral, 0)
}

func (s *AccessPropertyContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *AccessPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterAccessProperty(s)
	}
}

func (s *AccessPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitAccessProperty(s)
	}
}

func (s *AccessPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitAccessProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) AccessProperty() (localctx IAccessPropertyContext) {
	localctx = NewAccessPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LiteLangParserRULE_accessProperty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(LiteLangParserOpenBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13510816061980672) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(241)
		p.Match(LiteLangParserCloseBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	OpenParen() antlr.TerminalNode
	CloseParen() antlr.TerminalNode
	Params() IParamsContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, 0)
}

func (s *FunctionCallContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, 0)
}

func (s *FunctionCallContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, 0)
}

func (s *FunctionCallContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LiteLangParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(LiteLangParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(LiteLangParserOpenParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(245)
		p.Params()
	}

	{
		p.SetState(246)
		p.Match(LiteLangParserCloseParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamMakerContext is an interface to support dynamic dispatch.
type IParamMakerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsParamMakerContext differentiates from other interfaces.
	IsParamMakerContext()
}

type ParamMakerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamMakerContext() *ParamMakerContext {
	var p = new(ParamMakerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_paramMaker
	return p
}

func InitEmptyParamMakerContext(p *ParamMakerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_paramMaker
}

func (*ParamMakerContext) IsParamMakerContext() {}

func NewParamMakerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamMakerContext {
	var p = new(ParamMakerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_paramMaker

	return p
}

func (s *ParamMakerContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamMakerContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserIdentifier)
}

func (s *ParamMakerContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserIdentifier, i)
}

func (s *ParamMakerContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserComma)
}

func (s *ParamMakerContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserComma, i)
}

func (s *ParamMakerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamMakerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamMakerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterParamMaker(s)
	}
}

func (s *ParamMakerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitParamMaker(s)
	}
}

func (s *ParamMakerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitParamMaker(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ParamMaker() (localctx IParamMakerContext) {
	localctx = NewParamMakerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LiteLangParserRULE_paramMaker)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LiteLangParserIdentifier {
		{
			p.SetState(248)
			p.Match(LiteLangParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LiteLangParserComma {
			{
				p.SetState(249)
				p.Match(LiteLangParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(250)
				p.Match(LiteLangParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllUnit() []IUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitContext); ok {
			len++
		}
	}

	tst := make([]IUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitContext); ok {
			tst[i] = t.(IUnitContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Unit(i int) IUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *ParamsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserComma)
}

func (s *ParamsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserComma, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LiteLangParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13510828947406864) != 0 {
		{
			p.SetState(258)
			p.unit(0)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LiteLangParserComma {
			{
				p.SetState(259)
				p.Match(LiteLangParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(260)
				p.unit(0)
			}

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrowFunctionContext is an interface to support dynamic dispatch.
type IArrowFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenParen() antlr.TerminalNode
	Params() IParamsContext
	CloseParen() antlr.TerminalNode
	Block() IBlockContext

	// IsArrowFunctionContext differentiates from other interfaces.
	IsArrowFunctionContext()
}

type ArrowFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowFunctionContext() *ArrowFunctionContext {
	var p = new(ArrowFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_arrowFunction
	return p
}

func InitEmptyArrowFunctionContext(p *ArrowFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_arrowFunction
}

func (*ArrowFunctionContext) IsArrowFunctionContext() {}

func NewArrowFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowFunctionContext {
	var p = new(ArrowFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_arrowFunction

	return p
}

func (s *ArrowFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowFunctionContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenParen, 0)
}

func (s *ArrowFunctionContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *ArrowFunctionContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseParen, 0)
}

func (s *ArrowFunctionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ArrowFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrowFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterArrowFunction(s)
	}
}

func (s *ArrowFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitArrowFunction(s)
	}
}

func (s *ArrowFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitArrowFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ArrowFunction() (localctx IArrowFunctionContext) {
	localctx = NewArrowFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LiteLangParserRULE_arrowFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(LiteLangParserOpenParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Params()
	}
	{
		p.SetState(270)
		p.Match(LiteLangParserCloseParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(LiteLangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenBracket() antlr.TerminalNode
	CloseBracket() antlr.TerminalNode
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenBracket, 0)
}

func (s *ArrayLiteralContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseBracket, 0)
}

func (s *ArrayLiteralContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayLiteralContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserComma)
}

func (s *ArrayLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserComma, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LiteLangParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(LiteLangParserOpenBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13510828947406932) != 0 {
		{
			p.SetState(275)
			p.Value()
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LiteLangParserComma {
			{
				p.SetState(276)
				p.Match(LiteLangParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(277)
				p.Value()
			}

			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(285)
		p.Match(LiteLangParserCloseBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenBrace() antlr.TerminalNode
	CloseBrace() antlr.TerminalNode
	AllObjectItem() []IObjectItemContext
	ObjectItem(i int) IObjectItemContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_objectLiteral
	return p
}

func InitEmptyObjectLiteralContext(p *ObjectLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_objectLiteral
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenBrace, 0)
}

func (s *ObjectLiteralContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseBrace, 0)
}

func (s *ObjectLiteralContext) AllObjectItem() []IObjectItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectItemContext); ok {
			len++
		}
	}

	tst := make([]IObjectItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectItemContext); ok {
			tst[i] = t.(IObjectItemContext)
			i++
		}
	}

	return tst
}

func (s *ObjectLiteralContext) ObjectItem(i int) IObjectItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectItemContext)
}

func (s *ObjectLiteralContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserComma)
}

func (s *ObjectLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserComma, i)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitObjectLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LiteLangParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(LiteLangParserOpenBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9007199254757380) != 0 {
		{
			p.SetState(288)
			p.ObjectItem()
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == LiteLangParserComma {
			{
				p.SetState(289)
				p.Match(LiteLangParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(290)
				p.ObjectItem()
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(298)
		p.Match(LiteLangParserCloseBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	OpenBracket() antlr.TerminalNode
	AccessRhs() IAccessRhsContext
	CloseBracket() antlr.TerminalNode

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserStringLiteral, 0)
}

func (s *KeyContext) OpenBracket() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenBracket, 0)
}

func (s *KeyContext) AccessRhs() IAccessRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessRhsContext)
}

func (s *KeyContext) CloseBracket() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseBracket, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LiteLangParserRULE_key)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LiteLangParserStringLiteral:
		{
			p.SetState(300)
			p.Match(LiteLangParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LiteLangParserOpenBracket:
		{
			p.SetState(301)
			p.Match(LiteLangParserOpenBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.AccessRhs()
		}
		{
			p.SetState(303)
			p.Match(LiteLangParserCloseBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	Unit() IUnitContext
	ObjectLiteral() IObjectLiteralContext
	ArrayLiteral() IArrayLiteralContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ValueContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *ValueContext) ObjectLiteral() IObjectLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ValueContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LiteLangParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(307)
			p.FunctionCall()
		}

	case 2:
		{
			p.SetState(308)
			p.unit(0)
		}

	case 3:
		{
			p.SetState(309)
			p.ObjectLiteral()
		}

	case 4:
		{
			p.SetState(310)
			p.ArrayLiteral()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	Colon() antlr.TerminalNode
	Value() IValueContext

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_keyValue
	return p
}

func InitEmptyKeyValueContext(p *KeyValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_keyValue
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *KeyValueContext) Colon() antlr.TerminalNode {
	return s.GetToken(LiteLangParserColon, 0)
}

func (s *KeyValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (s *KeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LiteLangParserRULE_keyValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Key()
	}
	{
		p.SetState(314)
		p.Match(LiteLangParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISpreadContext is an interface to support dynamic dispatch.
type ISpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDot() []antlr.TerminalNode
	Dot(i int) antlr.TerminalNode
	AccessRhs() IAccessRhsContext

	// IsSpreadContext differentiates from other interfaces.
	IsSpreadContext()
}

type SpreadContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpreadContext() *SpreadContext {
	var p = new(SpreadContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_spread
	return p
}

func InitEmptySpreadContext(p *SpreadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_spread
}

func (*SpreadContext) IsSpreadContext() {}

func NewSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpreadContext {
	var p = new(SpreadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_spread

	return p
}

func (s *SpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *SpreadContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(LiteLangParserDot)
}

func (s *SpreadContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(LiteLangParserDot, i)
}

func (s *SpreadContext) AccessRhs() IAccessRhsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessRhsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessRhsContext)
}

func (s *SpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterSpread(s)
	}
}

func (s *SpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitSpread(s)
	}
}

func (s *SpreadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitSpread(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Spread() (localctx ISpreadContext) {
	localctx = NewSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LiteLangParserRULE_spread)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(LiteLangParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(LiteLangParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(LiteLangParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.AccessRhs()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectItemContext is an interface to support dynamic dispatch.
type IObjectItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KeyValue() IKeyValueContext
	Spread() ISpreadContext

	// IsObjectItemContext differentiates from other interfaces.
	IsObjectItemContext()
}

type ObjectItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectItemContext() *ObjectItemContext {
	var p = new(ObjectItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_objectItem
	return p
}

func InitEmptyObjectItemContext(p *ObjectItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_objectItem
}

func (*ObjectItemContext) IsObjectItemContext() {}

func NewObjectItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectItemContext {
	var p = new(ObjectItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_objectItem

	return p
}

func (s *ObjectItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectItemContext) KeyValue() IKeyValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
}

func (s *ObjectItemContext) Spread() ISpreadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpreadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpreadContext)
}

func (s *ObjectItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterObjectItem(s)
	}
}

func (s *ObjectItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitObjectItem(s)
	}
}

func (s *ObjectItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitObjectItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) ObjectItem() (localctx IObjectItemContext) {
	localctx = NewObjectItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LiteLangParserRULE_objectItem)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LiteLangParserOpenBracket, LiteLangParserStringLiteral:
		{
			p.SetState(322)
			p.KeyValue()
		}

	case LiteLangParserDot:
		{
			p.SetState(323)
			p.Spread()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenBrace() antlr.TerminalNode
	Statements() IStatementsContext
	CloseBrace() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) OpenBrace() antlr.TerminalNode {
	return s.GetToken(LiteLangParserOpenBrace, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) CloseBrace() antlr.TerminalNode {
	return s.GetToken(LiteLangParserCloseBrace, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LiteLangParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(LiteLangParserOpenBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.Statements()
	}
	{
		p.SetState(328)
		p.Match(LiteLangParserCloseBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NullLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	DecimalLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LiteLangParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiteLangParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserNullLiteral, 0)
}

func (s *ConstantContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserBooleanLiteral, 0)
}

func (s *ConstantContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserDecimalLiteral, 0)
}

func (s *ConstantContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LiteLangParserStringLiteral, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiteLangListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LiteLangVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LiteLangParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LiteLangParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9007229319512064) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *LiteLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *UnitContext = nil
		if localctx != nil {
			t = localctx.(*UnitContext)
		}
		return p.Unit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LiteLangParser) Unit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
