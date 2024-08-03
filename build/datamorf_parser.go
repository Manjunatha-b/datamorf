// Code generated from src/DataMorf.g4 by ANTLR 4.13.1. DO NOT EDIT.

package build // DataMorf
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

type DataMorfParser struct {
	*antlr.BaseParser
}

var DataMorfParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func datamorfParserInit() {
	staticData := &DataMorfParserStaticData
	staticData.LiteralNames = []string{
		"", "'=>'", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','",
		"'='", "'?'", "'?.'", "':'", "'.'", "'+'", "'-'", "'!'", "'*'", "'/'",
		"'%'", "'**'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'",
		"'||'", "'null'", "", "", "'break'", "'do'", "'of'", "'case'", "'else'",
		"'var'", "'return'", "'continue'", "'for'", "'switch'", "'while'", "'function'",
		"'default'", "'if'", "'elif'", "'delete'",
	}
	staticData.SymbolicNames = []string{
		"", "", "OpenSquare", "CloseSquare", "OpenRound", "CloseRound", "OpenCurly",
		"CloseCurly", "SemiColon", "Comma", "Assign", "QuestionMark", "QuestionMarkDot",
		"Colon", "Dot", "Plus", "Minus", "Not", "Multiply", "Divide", "Modulus",
		"Power", "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals",
		"Equals", "NotEquals", "And", "Or", "NullLiteral", "BooleanLiteral",
		"DecimalLiteral", "Break", "Do", "Of", "Case", "Else", "Var", "Return",
		"Continue", "For", "Switch", "While", "Function", "Default", "If", "Elif",
		"Delete", "Identifier", "StringLiteral", "WhiteSpaces", "LineTerminator",
	}
	staticData.RuleNames = []string{
		"program", "statements", "statement", "variableStatement", "assignStatement",
		"functionStatement", "ifStatement", "switchStatement", "deleteStatement",
		"unitStatement", "valueStatement", "forStatement", "classicForCondition",
		"iteratorForCondition", "unit", "reference", "accessor", "accessProperty",
		"functionCall", "signatureParams", "sendingParams", "arrowFunction",
		"arrayLiteral", "objectLiteral", "key", "value", "keyValue", "objectItem",
		"spread", "block", "constant",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 52, 314, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 3,
		0, 64, 8, 0, 1, 0, 1, 0, 1, 1, 4, 1, 69, 8, 1, 11, 1, 12, 1, 70, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 96,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		5, 6, 121, 8, 6, 10, 6, 12, 6, 124, 9, 6, 1, 6, 1, 6, 3, 6, 128, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		141, 8, 8, 5, 8, 143, 8, 8, 10, 8, 12, 8, 146, 9, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 157, 8, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 182, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 199, 8,
		14, 10, 14, 12, 14, 202, 9, 14, 1, 15, 1, 15, 3, 15, 206, 8, 15, 1, 15,
		5, 15, 209, 8, 15, 10, 15, 12, 15, 212, 9, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 219, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 233, 8, 19, 10, 19, 12,
		19, 236, 9, 19, 3, 19, 238, 8, 19, 1, 20, 1, 20, 1, 20, 5, 20, 243, 8,
		20, 10, 20, 12, 20, 246, 9, 20, 3, 20, 248, 8, 20, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 260, 8, 22, 10,
		22, 12, 22, 263, 9, 22, 3, 22, 265, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 5, 23, 273, 8, 23, 10, 23, 12, 23, 276, 9, 23, 3, 23, 278,
		8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 287, 8,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 293, 8, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 3, 27, 301, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 0, 1, 28, 31, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 3, 1, 0, 22, 29, 2, 0, 32, 32,
		49, 50, 2, 0, 30, 32, 50, 50, 323, 0, 63, 1, 0, 0, 0, 2, 68, 1, 0, 0, 0,
		4, 89, 1, 0, 0, 0, 6, 91, 1, 0, 0, 0, 8, 97, 1, 0, 0, 0, 10, 101, 1, 0,
		0, 0, 12, 108, 1, 0, 0, 0, 14, 129, 1, 0, 0, 0, 16, 135, 1, 0, 0, 0, 18,
		147, 1, 0, 0, 0, 20, 149, 1, 0, 0, 0, 22, 152, 1, 0, 0, 0, 24, 161, 1,
		0, 0, 0, 26, 167, 1, 0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 205, 1, 0, 0, 0,
		32, 218, 1, 0, 0, 0, 34, 220, 1, 0, 0, 0, 36, 224, 1, 0, 0, 0, 38, 237,
		1, 0, 0, 0, 40, 247, 1, 0, 0, 0, 42, 249, 1, 0, 0, 0, 44, 255, 1, 0, 0,
		0, 46, 268, 1, 0, 0, 0, 48, 286, 1, 0, 0, 0, 50, 292, 1, 0, 0, 0, 52, 294,
		1, 0, 0, 0, 54, 300, 1, 0, 0, 0, 56, 302, 1, 0, 0, 0, 58, 307, 1, 0, 0,
		0, 60, 311, 1, 0, 0, 0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 63, 64,
		1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 5, 0, 0, 1, 66, 1, 1, 0, 0, 0,
		67, 69, 3, 4, 2, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68, 1,
		0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 3, 1, 0, 0, 0, 72, 73, 3, 6, 3, 0, 73,
		74, 5, 8, 0, 0, 74, 90, 1, 0, 0, 0, 75, 76, 3, 8, 4, 0, 76, 77, 5, 8, 0,
		0, 77, 90, 1, 0, 0, 0, 78, 90, 3, 10, 5, 0, 79, 90, 3, 12, 6, 0, 80, 90,
		3, 22, 11, 0, 81, 90, 3, 14, 7, 0, 82, 90, 3, 20, 10, 0, 83, 84, 3, 16,
		8, 0, 84, 85, 5, 8, 0, 0, 85, 90, 1, 0, 0, 0, 86, 87, 3, 18, 9, 0, 87,
		88, 5, 8, 0, 0, 88, 90, 1, 0, 0, 0, 89, 72, 1, 0, 0, 0, 89, 75, 1, 0, 0,
		0, 89, 78, 1, 0, 0, 0, 89, 79, 1, 0, 0, 0, 89, 80, 1, 0, 0, 0, 89, 81,
		1, 0, 0, 0, 89, 82, 1, 0, 0, 0, 89, 83, 1, 0, 0, 0, 89, 86, 1, 0, 0, 0,
		90, 5, 1, 0, 0, 0, 91, 92, 5, 38, 0, 0, 92, 95, 5, 49, 0, 0, 93, 94, 5,
		10, 0, 0, 94, 96, 3, 50, 25, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 7, 1, 0, 0, 0, 97, 98, 3, 30, 15, 0, 98, 99, 5, 10, 0, 0, 99, 100,
		3, 50, 25, 0, 100, 9, 1, 0, 0, 0, 101, 102, 5, 44, 0, 0, 102, 103, 5, 49,
		0, 0, 103, 104, 5, 4, 0, 0, 104, 105, 3, 38, 19, 0, 105, 106, 5, 5, 0,
		0, 106, 107, 3, 58, 29, 0, 107, 11, 1, 0, 0, 0, 108, 109, 5, 46, 0, 0,
		109, 110, 5, 4, 0, 0, 110, 111, 3, 28, 14, 0, 111, 112, 5, 5, 0, 0, 112,
		122, 3, 58, 29, 0, 113, 114, 5, 37, 0, 0, 114, 115, 5, 46, 0, 0, 115, 116,
		5, 4, 0, 0, 116, 117, 3, 28, 14, 0, 117, 118, 5, 5, 0, 0, 118, 119, 3,
		58, 29, 0, 119, 121, 1, 0, 0, 0, 120, 113, 1, 0, 0, 0, 121, 124, 1, 0,
		0, 0, 122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 127, 1, 0, 0, 0,
		124, 122, 1, 0, 0, 0, 125, 126, 5, 37, 0, 0, 126, 128, 3, 58, 29, 0, 127,
		125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 13, 1, 0, 0, 0, 129, 130, 5,
		42, 0, 0, 130, 131, 5, 4, 0, 0, 131, 132, 3, 28, 14, 0, 132, 133, 5, 5,
		0, 0, 133, 134, 3, 58, 29, 0, 134, 15, 1, 0, 0, 0, 135, 136, 5, 48, 0,
		0, 136, 144, 5, 49, 0, 0, 137, 140, 5, 14, 0, 0, 138, 141, 3, 34, 17, 0,
		139, 141, 5, 49, 0, 0, 140, 138, 1, 0, 0, 0, 140, 139, 1, 0, 0, 0, 141,
		143, 1, 0, 0, 0, 142, 137, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144, 142,
		1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 17, 1, 0, 0, 0, 146, 144, 1, 0,
		0, 0, 147, 148, 3, 28, 14, 0, 148, 19, 1, 0, 0, 0, 149, 150, 3, 50, 25,
		0, 150, 151, 5, 8, 0, 0, 151, 21, 1, 0, 0, 0, 152, 153, 5, 41, 0, 0, 153,
		156, 5, 4, 0, 0, 154, 157, 3, 24, 12, 0, 155, 157, 3, 26, 13, 0, 156, 154,
		1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 5, 5,
		0, 0, 159, 160, 3, 58, 29, 0, 160, 23, 1, 0, 0, 0, 161, 162, 3, 6, 3, 0,
		162, 163, 5, 8, 0, 0, 163, 164, 3, 18, 9, 0, 164, 165, 5, 8, 0, 0, 165,
		166, 3, 8, 4, 0, 166, 25, 1, 0, 0, 0, 167, 168, 5, 38, 0, 0, 168, 169,
		5, 49, 0, 0, 169, 170, 5, 35, 0, 0, 170, 171, 3, 30, 15, 0, 171, 27, 1,
		0, 0, 0, 172, 173, 6, 14, -1, 0, 173, 174, 5, 4, 0, 0, 174, 175, 3, 28,
		14, 0, 175, 176, 5, 5, 0, 0, 176, 182, 1, 0, 0, 0, 177, 178, 5, 17, 0,
		0, 178, 182, 3, 28, 14, 3, 179, 182, 3, 30, 15, 0, 180, 182, 3, 60, 30,
		0, 181, 172, 1, 0, 0, 0, 181, 177, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		180, 1, 0, 0, 0, 182, 200, 1, 0, 0, 0, 183, 184, 10, 8, 0, 0, 184, 185,
		5, 19, 0, 0, 185, 199, 3, 28, 14, 9, 186, 187, 10, 7, 0, 0, 187, 188, 5,
		18, 0, 0, 188, 199, 3, 28, 14, 8, 189, 190, 10, 6, 0, 0, 190, 191, 5, 15,
		0, 0, 191, 199, 3, 28, 14, 7, 192, 193, 10, 5, 0, 0, 193, 194, 5, 16, 0,
		0, 194, 199, 3, 28, 14, 6, 195, 196, 10, 4, 0, 0, 196, 197, 7, 0, 0, 0,
		197, 199, 3, 28, 14, 5, 198, 183, 1, 0, 0, 0, 198, 186, 1, 0, 0, 0, 198,
		189, 1, 0, 0, 0, 198, 192, 1, 0, 0, 0, 198, 195, 1, 0, 0, 0, 199, 202,
		1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 29, 1, 0,
		0, 0, 202, 200, 1, 0, 0, 0, 203, 206, 3, 36, 18, 0, 204, 206, 5, 49, 0,
		0, 205, 203, 1, 0, 0, 0, 205, 204, 1, 0, 0, 0, 206, 210, 1, 0, 0, 0, 207,
		209, 3, 32, 16, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208,
		1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 31, 1, 0, 0, 0, 212, 210, 1, 0,
		0, 0, 213, 219, 3, 34, 17, 0, 214, 215, 5, 14, 0, 0, 215, 219, 5, 49, 0,
		0, 216, 217, 5, 14, 0, 0, 217, 219, 3, 36, 18, 0, 218, 213, 1, 0, 0, 0,
		218, 214, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 33, 1, 0, 0, 0, 220, 221,
		5, 2, 0, 0, 221, 222, 7, 1, 0, 0, 222, 223, 5, 3, 0, 0, 223, 35, 1, 0,
		0, 0, 224, 225, 5, 49, 0, 0, 225, 226, 5, 4, 0, 0, 226, 227, 3, 40, 20,
		0, 227, 228, 5, 5, 0, 0, 228, 37, 1, 0, 0, 0, 229, 234, 5, 49, 0, 0, 230,
		231, 5, 9, 0, 0, 231, 233, 5, 49, 0, 0, 232, 230, 1, 0, 0, 0, 233, 236,
		1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 238, 1, 0,
		0, 0, 236, 234, 1, 0, 0, 0, 237, 229, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0,
		238, 39, 1, 0, 0, 0, 239, 244, 3, 28, 14, 0, 240, 241, 5, 9, 0, 0, 241,
		243, 3, 28, 14, 0, 242, 240, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242,
		1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244, 1, 0,
		0, 0, 247, 239, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 41, 1, 0, 0, 0,
		249, 250, 5, 4, 0, 0, 250, 251, 3, 38, 19, 0, 251, 252, 5, 5, 0, 0, 252,
		253, 5, 1, 0, 0, 253, 254, 3, 58, 29, 0, 254, 43, 1, 0, 0, 0, 255, 264,
		5, 2, 0, 0, 256, 261, 3, 50, 25, 0, 257, 258, 5, 9, 0, 0, 258, 260, 3,
		50, 25, 0, 259, 257, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0,
		0, 0, 261, 262, 1, 0, 0, 0, 262, 265, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0,
		264, 256, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266,
		267, 5, 3, 0, 0, 267, 45, 1, 0, 0, 0, 268, 277, 5, 6, 0, 0, 269, 274, 3,
		54, 27, 0, 270, 271, 5, 9, 0, 0, 271, 273, 3, 54, 27, 0, 272, 270, 1, 0,
		0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0,
		275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 269, 1, 0, 0, 0, 277,
		278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 5, 7, 0, 0, 280, 47, 1,
		0, 0, 0, 281, 287, 5, 50, 0, 0, 282, 283, 5, 2, 0, 0, 283, 284, 3, 30,
		15, 0, 284, 285, 5, 3, 0, 0, 285, 287, 1, 0, 0, 0, 286, 281, 1, 0, 0, 0,
		286, 282, 1, 0, 0, 0, 287, 49, 1, 0, 0, 0, 288, 293, 3, 36, 18, 0, 289,
		293, 3, 28, 14, 0, 290, 293, 3, 46, 23, 0, 291, 293, 3, 44, 22, 0, 292,
		288, 1, 0, 0, 0, 292, 289, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 292, 291,
		1, 0, 0, 0, 293, 51, 1, 0, 0, 0, 294, 295, 3, 48, 24, 0, 295, 296, 5, 13,
		0, 0, 296, 297, 3, 50, 25, 0, 297, 53, 1, 0, 0, 0, 298, 301, 3, 52, 26,
		0, 299, 301, 3, 56, 28, 0, 300, 298, 1, 0, 0, 0, 300, 299, 1, 0, 0, 0,
		301, 55, 1, 0, 0, 0, 302, 303, 5, 14, 0, 0, 303, 304, 5, 14, 0, 0, 304,
		305, 5, 14, 0, 0, 305, 306, 3, 30, 15, 0, 306, 57, 1, 0, 0, 0, 307, 308,
		5, 6, 0, 0, 308, 309, 3, 2, 1, 0, 309, 310, 5, 7, 0, 0, 310, 59, 1, 0,
		0, 0, 311, 312, 7, 2, 0, 0, 312, 61, 1, 0, 0, 0, 26, 63, 70, 89, 95, 122,
		127, 140, 144, 156, 181, 198, 200, 205, 210, 218, 234, 237, 244, 247, 261,
		264, 274, 277, 286, 292, 300,
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

// DataMorfParserInit initializes any static state used to implement DataMorfParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDataMorfParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DataMorfParserInit() {
	staticData := &DataMorfParserStaticData
	staticData.once.Do(datamorfParserInit)
}

// NewDataMorfParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDataMorfParser(input antlr.TokenStream) *DataMorfParser {
	DataMorfParserInit()
	this := new(DataMorfParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DataMorfParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DataMorf.g4"

	return this
}

// DataMorfParser tokens.
const (
	DataMorfParserEOF               = antlr.TokenEOF
	DataMorfParserT__0              = 1
	DataMorfParserOpenSquare        = 2
	DataMorfParserCloseSquare       = 3
	DataMorfParserOpenRound         = 4
	DataMorfParserCloseRound        = 5
	DataMorfParserOpenCurly         = 6
	DataMorfParserCloseCurly        = 7
	DataMorfParserSemiColon         = 8
	DataMorfParserComma             = 9
	DataMorfParserAssign            = 10
	DataMorfParserQuestionMark      = 11
	DataMorfParserQuestionMarkDot   = 12
	DataMorfParserColon             = 13
	DataMorfParserDot               = 14
	DataMorfParserPlus              = 15
	DataMorfParserMinus             = 16
	DataMorfParserNot               = 17
	DataMorfParserMultiply          = 18
	DataMorfParserDivide            = 19
	DataMorfParserModulus           = 20
	DataMorfParserPower             = 21
	DataMorfParserLessThan          = 22
	DataMorfParserMoreThan          = 23
	DataMorfParserLessThanEquals    = 24
	DataMorfParserGreaterThanEquals = 25
	DataMorfParserEquals            = 26
	DataMorfParserNotEquals         = 27
	DataMorfParserAnd               = 28
	DataMorfParserOr                = 29
	DataMorfParserNullLiteral       = 30
	DataMorfParserBooleanLiteral    = 31
	DataMorfParserDecimalLiteral    = 32
	DataMorfParserBreak             = 33
	DataMorfParserDo                = 34
	DataMorfParserOf                = 35
	DataMorfParserCase              = 36
	DataMorfParserElse              = 37
	DataMorfParserVar               = 38
	DataMorfParserReturn            = 39
	DataMorfParserContinue          = 40
	DataMorfParserFor               = 41
	DataMorfParserSwitch            = 42
	DataMorfParserWhile             = 43
	DataMorfParserFunction          = 44
	DataMorfParserDefault           = 45
	DataMorfParserIf                = 46
	DataMorfParserElif              = 47
	DataMorfParserDelete            = 48
	DataMorfParserIdentifier        = 49
	DataMorfParserStringLiteral     = 50
	DataMorfParserWhiteSpaces       = 51
	DataMorfParserLineTerminator    = 52
)

// DataMorfParser rules.
const (
	DataMorfParserRULE_program              = 0
	DataMorfParserRULE_statements           = 1
	DataMorfParserRULE_statement            = 2
	DataMorfParserRULE_variableStatement    = 3
	DataMorfParserRULE_assignStatement      = 4
	DataMorfParserRULE_functionStatement    = 5
	DataMorfParserRULE_ifStatement          = 6
	DataMorfParserRULE_switchStatement      = 7
	DataMorfParserRULE_deleteStatement      = 8
	DataMorfParserRULE_unitStatement        = 9
	DataMorfParserRULE_valueStatement       = 10
	DataMorfParserRULE_forStatement         = 11
	DataMorfParserRULE_classicForCondition  = 12
	DataMorfParserRULE_iteratorForCondition = 13
	DataMorfParserRULE_unit                 = 14
	DataMorfParserRULE_reference            = 15
	DataMorfParserRULE_accessor             = 16
	DataMorfParserRULE_accessProperty       = 17
	DataMorfParserRULE_functionCall         = 18
	DataMorfParserRULE_signatureParams      = 19
	DataMorfParserRULE_sendingParams        = 20
	DataMorfParserRULE_arrowFunction        = 21
	DataMorfParserRULE_arrayLiteral         = 22
	DataMorfParserRULE_objectLiteral        = 23
	DataMorfParserRULE_key                  = 24
	DataMorfParserRULE_value                = 25
	DataMorfParserRULE_keyValue             = 26
	DataMorfParserRULE_objectItem           = 27
	DataMorfParserRULE_spread               = 28
	DataMorfParserRULE_block                = 29
	DataMorfParserRULE_constant             = 30
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
	p.RuleIndex = DataMorfParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DataMorfParserEOF, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DataMorfParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2065165231194196) != 0 {
		{
			p.SetState(62)
			p.Statements()
		}

	}
	{
		p.SetState(65)
		p.Match(DataMorfParserEOF)
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
	p.RuleIndex = DataMorfParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_statements

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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DataMorfParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2065165231194196) != 0) {
		{
			p.SetState(67)
			p.Statement()
		}

		p.SetState(70)
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
	p.RuleIndex = DataMorfParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_statement

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
	return s.GetToken(DataMorfParserSemiColon, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DataMorfParserRULE_statement)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.VariableStatement()
		}
		{
			p.SetState(73)
			p.Match(DataMorfParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.AssignStatement()
		}
		{
			p.SetState(76)
			p.Match(DataMorfParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.FunctionStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(80)
			p.ForStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(81)
			p.SwitchStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(82)
			p.ValueStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(83)
			p.DeleteStatement()
		}
		{
			p.SetState(84)
			p.Match(DataMorfParserSemiColon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(86)
			p.UnitStatement()
		}
		{
			p.SetState(87)
			p.Match(DataMorfParserSemiColon)
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
	p.RuleIndex = DataMorfParserRULE_variableStatement
	return p
}

func InitEmptyVariableStatementContext(p *VariableStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_variableStatement
}

func (*VariableStatementContext) IsVariableStatementContext() {}

func NewVariableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableStatementContext {
	var p = new(VariableStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_variableStatement

	return p
}

func (s *VariableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(DataMorfParserVar, 0)
}

func (s *VariableStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *VariableStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(DataMorfParserAssign, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

func (s *VariableStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitVariableStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) VariableStatement() (localctx IVariableStatementContext) {
	localctx = NewVariableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DataMorfParserRULE_variableStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(DataMorfParserVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(DataMorfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DataMorfParserAssign {
		{
			p.SetState(93)
			p.Match(DataMorfParserAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
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
	Reference() IReferenceContext
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
	p.RuleIndex = DataMorfParserRULE_assignStatement
	return p
}

func InitEmptyAssignStatementContext(p *AssignStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_assignStatement
}

func (*AssignStatementContext) IsAssignStatementContext() {}

func NewAssignStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStatementContext {
	var p = new(AssignStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_assignStatement

	return p
}

func (s *AssignStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStatementContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *AssignStatementContext) Assign() antlr.TerminalNode {
	return s.GetToken(DataMorfParserAssign, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) AssignStatement() (localctx IAssignStatementContext) {
	localctx = NewAssignStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DataMorfParserRULE_assignStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Reference()
	}
	{
		p.SetState(98)
		p.Match(DataMorfParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
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
	OpenRound() antlr.TerminalNode
	SignatureParams() ISignatureParamsContext
	CloseRound() antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_functionStatement
	return p
}

func InitEmptyFunctionStatementContext(p *FunctionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_functionStatement
}

func (*FunctionStatementContext) IsFunctionStatementContext() {}

func NewFunctionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStatementContext {
	var p = new(FunctionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_functionStatement

	return p
}

func (s *FunctionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStatementContext) Function() antlr.TerminalNode {
	return s.GetToken(DataMorfParserFunction, 0)
}

func (s *FunctionStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *FunctionStatementContext) OpenRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, 0)
}

func (s *FunctionStatementContext) SignatureParams() ISignatureParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignatureParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignatureParamsContext)
}

func (s *FunctionStatementContext) CloseRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitFunctionStatement(s)
	}
}

func (s *FunctionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitFunctionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) FunctionStatement() (localctx IFunctionStatementContext) {
	localctx = NewFunctionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DataMorfParserRULE_functionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(DataMorfParserFunction)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(DataMorfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(DataMorfParserOpenRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.SignatureParams()
	}
	{
		p.SetState(105)
		p.Match(DataMorfParserCloseRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
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
	AllOpenRound() []antlr.TerminalNode
	OpenRound(i int) antlr.TerminalNode
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	AllCloseRound() []antlr.TerminalNode
	CloseRound(i int) antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) AllIf() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserIf)
}

func (s *IfStatementContext) If(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserIf, i)
}

func (s *IfStatementContext) AllOpenRound() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserOpenRound)
}

func (s *IfStatementContext) OpenRound(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, i)
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

func (s *IfStatementContext) AllCloseRound() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserCloseRound)
}

func (s *IfStatementContext) CloseRound(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, i)
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
	return s.GetTokens(DataMorfParserElse)
}

func (s *IfStatementContext) Else(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserElse, i)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DataMorfParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(DataMorfParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(DataMorfParserOpenRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.unit(0)
	}
	{
		p.SetState(111)
		p.Match(DataMorfParserCloseRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Block()
	}
	p.SetState(122)
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
				p.SetState(113)
				p.Match(DataMorfParserElse)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.Match(DataMorfParserIf)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(115)
				p.Match(DataMorfParserOpenRound)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(116)
				p.unit(0)
			}
			{
				p.SetState(117)
				p.Match(DataMorfParserCloseRound)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(118)
				p.Block()
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DataMorfParserElse {
		{
			p.SetState(125)
			p.Match(DataMorfParserElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
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

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Switch() antlr.TerminalNode
	OpenRound() antlr.TerminalNode
	Unit() IUnitContext
	CloseRound() antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(DataMorfParserSwitch, 0)
}

func (s *SwitchStatementContext) OpenRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, 0)
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

func (s *SwitchStatementContext) CloseRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DataMorfParserRULE_switchStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(DataMorfParserSwitch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(DataMorfParserOpenRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.unit(0)
	}
	{
		p.SetState(132)
		p.Match(DataMorfParserCloseRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
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
	p.RuleIndex = DataMorfParserRULE_deleteStatement
	return p
}

func InitEmptyDeleteStatementContext(p *DeleteStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_deleteStatement
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) Delete() antlr.TerminalNode {
	return s.GetToken(DataMorfParserDelete, 0)
}

func (s *DeleteStatementContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserIdentifier)
}

func (s *DeleteStatementContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, i)
}

func (s *DeleteStatementContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserDot)
}

func (s *DeleteStatementContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserDot, i)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitDeleteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DataMorfParserRULE_deleteStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(DataMorfParserDelete)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(136)
		p.Match(DataMorfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DataMorfParserDot {
		{
			p.SetState(137)
			p.Match(DataMorfParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case DataMorfParserOpenSquare:
			{
				p.SetState(138)
				p.AccessProperty()
			}

		case DataMorfParserIdentifier:
			{
				p.SetState(139)
				p.Match(DataMorfParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(146)
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
	p.RuleIndex = DataMorfParserRULE_unitStatement
	return p
}

func InitEmptyUnitStatementContext(p *UnitStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_unitStatement
}

func (*UnitStatementContext) IsUnitStatementContext() {}

func NewUnitStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitStatementContext {
	var p = new(UnitStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_unitStatement

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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterUnitStatement(s)
	}
}

func (s *UnitStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitUnitStatement(s)
	}
}

func (s *UnitStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitUnitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) UnitStatement() (localctx IUnitStatementContext) {
	localctx = NewUnitStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DataMorfParserRULE_unitStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
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
	p.RuleIndex = DataMorfParserRULE_valueStatement
	return p
}

func InitEmptyValueStatementContext(p *ValueStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_valueStatement
}

func (*ValueStatementContext) IsValueStatementContext() {}

func NewValueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueStatementContext {
	var p = new(ValueStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_valueStatement

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
	return s.GetToken(DataMorfParserSemiColon, 0)
}

func (s *ValueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterValueStatement(s)
	}
}

func (s *ValueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitValueStatement(s)
	}
}

func (s *ValueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitValueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ValueStatement() (localctx IValueStatementContext) {
	localctx = NewValueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DataMorfParserRULE_valueStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Value()
	}
	{
		p.SetState(150)
		p.Match(DataMorfParserSemiColon)
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

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	For() antlr.TerminalNode
	OpenRound() antlr.TerminalNode
	CloseRound() antlr.TerminalNode
	Block() IBlockContext
	ClassicForCondition() IClassicForConditionContext
	IteratorForCondition() IIteratorForConditionContext

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
	p.RuleIndex = DataMorfParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(DataMorfParserFor, 0)
}

func (s *ForStatementContext) OpenRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, 0)
}

func (s *ForStatementContext) CloseRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, 0)
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

func (s *ForStatementContext) ClassicForCondition() IClassicForConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassicForConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassicForConditionContext)
}

func (s *ForStatementContext) IteratorForCondition() IIteratorForConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorForConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorForConditionContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DataMorfParserRULE_forStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(DataMorfParserFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(DataMorfParserOpenRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(154)
			p.ClassicForCondition()
		}

	case 2:
		{
			p.SetState(155)
			p.IteratorForCondition()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(158)
		p.Match(DataMorfParserCloseRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
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

// IClassicForConditionContext is an interface to support dynamic dispatch.
type IClassicForConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IVariableStatementContext

	// GetCondition returns the condition rule contexts.
	GetCondition() IUnitStatementContext

	// GetPostOp returns the postOp rule contexts.
	GetPostOp() IAssignStatementContext

	// SetInit sets the init rule contexts.
	SetInit(IVariableStatementContext)

	// SetCondition sets the condition rule contexts.
	SetCondition(IUnitStatementContext)

	// SetPostOp sets the postOp rule contexts.
	SetPostOp(IAssignStatementContext)

	// Getter signatures
	AllSemiColon() []antlr.TerminalNode
	SemiColon(i int) antlr.TerminalNode
	VariableStatement() IVariableStatementContext
	UnitStatement() IUnitStatementContext
	AssignStatement() IAssignStatementContext

	// IsClassicForConditionContext differentiates from other interfaces.
	IsClassicForConditionContext()
}

type ClassicForConditionContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	init      IVariableStatementContext
	condition IUnitStatementContext
	postOp    IAssignStatementContext
}

func NewEmptyClassicForConditionContext() *ClassicForConditionContext {
	var p = new(ClassicForConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_classicForCondition
	return p
}

func InitEmptyClassicForConditionContext(p *ClassicForConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_classicForCondition
}

func (*ClassicForConditionContext) IsClassicForConditionContext() {}

func NewClassicForConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassicForConditionContext {
	var p = new(ClassicForConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_classicForCondition

	return p
}

func (s *ClassicForConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassicForConditionContext) GetInit() IVariableStatementContext { return s.init }

func (s *ClassicForConditionContext) GetCondition() IUnitStatementContext { return s.condition }

func (s *ClassicForConditionContext) GetPostOp() IAssignStatementContext { return s.postOp }

func (s *ClassicForConditionContext) SetInit(v IVariableStatementContext) { s.init = v }

func (s *ClassicForConditionContext) SetCondition(v IUnitStatementContext) { s.condition = v }

func (s *ClassicForConditionContext) SetPostOp(v IAssignStatementContext) { s.postOp = v }

func (s *ClassicForConditionContext) AllSemiColon() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserSemiColon)
}

func (s *ClassicForConditionContext) SemiColon(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserSemiColon, i)
}

func (s *ClassicForConditionContext) VariableStatement() IVariableStatementContext {
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

func (s *ClassicForConditionContext) UnitStatement() IUnitStatementContext {
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

func (s *ClassicForConditionContext) AssignStatement() IAssignStatementContext {
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

func (s *ClassicForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassicForConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassicForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterClassicForCondition(s)
	}
}

func (s *ClassicForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitClassicForCondition(s)
	}
}

func (s *ClassicForConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitClassicForCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ClassicForCondition() (localctx IClassicForConditionContext) {
	localctx = NewClassicForConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DataMorfParserRULE_classicForCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)

		var _x = p.VariableStatement()

		localctx.(*ClassicForConditionContext).init = _x
	}
	{
		p.SetState(162)
		p.Match(DataMorfParserSemiColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)

		var _x = p.UnitStatement()

		localctx.(*ClassicForConditionContext).condition = _x
	}
	{
		p.SetState(164)
		p.Match(DataMorfParserSemiColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)

		var _x = p.AssignStatement()

		localctx.(*ClassicForConditionContext).postOp = _x
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

// IIteratorForConditionContext is an interface to support dynamic dispatch.
type IIteratorForConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	Of() antlr.TerminalNode
	Reference() IReferenceContext

	// IsIteratorForConditionContext differentiates from other interfaces.
	IsIteratorForConditionContext()
}

type IteratorForConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteratorForConditionContext() *IteratorForConditionContext {
	var p = new(IteratorForConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_iteratorForCondition
	return p
}

func InitEmptyIteratorForConditionContext(p *IteratorForConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_iteratorForCondition
}

func (*IteratorForConditionContext) IsIteratorForConditionContext() {}

func NewIteratorForConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IteratorForConditionContext {
	var p = new(IteratorForConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_iteratorForCondition

	return p
}

func (s *IteratorForConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *IteratorForConditionContext) Var() antlr.TerminalNode {
	return s.GetToken(DataMorfParserVar, 0)
}

func (s *IteratorForConditionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *IteratorForConditionContext) Of() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOf, 0)
}

func (s *IteratorForConditionContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *IteratorForConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteratorForConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IteratorForConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterIteratorForCondition(s)
	}
}

func (s *IteratorForConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitIteratorForCondition(s)
	}
}

func (s *IteratorForConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitIteratorForCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) IteratorForCondition() (localctx IIteratorForConditionContext) {
	localctx = NewIteratorForConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DataMorfParserRULE_iteratorForCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(DataMorfParserVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(DataMorfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(DataMorfParserOf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Reference()
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
	OpenRound() antlr.TerminalNode
	CloseRound() antlr.TerminalNode
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	Not() antlr.TerminalNode
	Reference() IReferenceContext
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
	p.RuleIndex = DataMorfParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) GetLeft() IUnitContext { return s.left }

func (s *UnitContext) GetSingle() IUnitContext { return s.single }

func (s *UnitContext) GetRight() IUnitContext { return s.right }

func (s *UnitContext) SetLeft(v IUnitContext) { s.left = v }

func (s *UnitContext) SetSingle(v IUnitContext) { s.single = v }

func (s *UnitContext) SetRight(v IUnitContext) { s.right = v }

func (s *UnitContext) OpenRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, 0)
}

func (s *UnitContext) CloseRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, 0)
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
	return s.GetToken(DataMorfParserNot, 0)
}

func (s *UnitContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
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
	return s.GetToken(DataMorfParserDivide, 0)
}

func (s *UnitContext) Multiply() antlr.TerminalNode {
	return s.GetToken(DataMorfParserMultiply, 0)
}

func (s *UnitContext) Plus() antlr.TerminalNode {
	return s.GetToken(DataMorfParserPlus, 0)
}

func (s *UnitContext) Minus() antlr.TerminalNode {
	return s.GetToken(DataMorfParserMinus, 0)
}

func (s *UnitContext) LessThan() antlr.TerminalNode {
	return s.GetToken(DataMorfParserLessThan, 0)
}

func (s *UnitContext) MoreThan() antlr.TerminalNode {
	return s.GetToken(DataMorfParserMoreThan, 0)
}

func (s *UnitContext) LessThanEquals() antlr.TerminalNode {
	return s.GetToken(DataMorfParserLessThanEquals, 0)
}

func (s *UnitContext) GreaterThanEquals() antlr.TerminalNode {
	return s.GetToken(DataMorfParserGreaterThanEquals, 0)
}

func (s *UnitContext) Equals() antlr.TerminalNode {
	return s.GetToken(DataMorfParserEquals, 0)
}

func (s *UnitContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(DataMorfParserNotEquals, 0)
}

func (s *UnitContext) And() antlr.TerminalNode {
	return s.GetToken(DataMorfParserAnd, 0)
}

func (s *UnitContext) Or() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOr, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (s *UnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Unit() (localctx IUnitContext) {
	return p.unit(0)
}

func (p *DataMorfParser) unit(_p int) (localctx IUnitContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewUnitContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IUnitContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, DataMorfParserRULE_unit, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DataMorfParserOpenRound:
		{
			p.SetState(173)
			p.Match(DataMorfParserOpenRound)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)

			var _x = p.unit(0)

			localctx.(*UnitContext).single = _x
		}
		{
			p.SetState(175)
			p.Match(DataMorfParserCloseRound)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DataMorfParserNot:
		{
			p.SetState(177)
			p.Match(DataMorfParserNot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)

			var _x = p.unit(3)

			localctx.(*UnitContext).single = _x
		}

	case DataMorfParserIdentifier:
		{
			p.SetState(179)
			p.Reference()
		}

	case DataMorfParserNullLiteral, DataMorfParserBooleanLiteral, DataMorfParserDecimalLiteral, DataMorfParserStringLiteral:
		{
			p.SetState(180)
			p.Constant()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, DataMorfParserRULE_unit)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}

				{
					p.SetState(184)
					p.Match(DataMorfParserDivide)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(185)

					var _x = p.unit(9)

					localctx.(*UnitContext).right = _x
				}

			case 2:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, DataMorfParserRULE_unit)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}

				{
					p.SetState(187)
					p.Match(DataMorfParserMultiply)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(188)

					var _x = p.unit(8)

					localctx.(*UnitContext).right = _x
				}

			case 3:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, DataMorfParserRULE_unit)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}

				{
					p.SetState(190)
					p.Match(DataMorfParserPlus)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(191)

					var _x = p.unit(7)

					localctx.(*UnitContext).right = _x
				}

			case 4:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, DataMorfParserRULE_unit)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}

				{
					p.SetState(193)
					p.Match(DataMorfParserMinus)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(194)

					var _x = p.unit(6)

					localctx.(*UnitContext).right = _x
				}

			case 5:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, DataMorfParserRULE_unit)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(196)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1069547520) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(197)

					var _x = p.unit(5)

					localctx.(*UnitContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	Identifier() antlr.TerminalNode
	AllAccessor() []IAccessorContext
	Accessor(i int) IAccessorContext

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_reference
	return p
}

func InitEmptyReferenceContext(p *ReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_reference
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) FunctionCall() IFunctionCallContext {
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

func (s *ReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *ReferenceContext) AllAccessor() []IAccessorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessorContext); ok {
			len++
		}
	}

	tst := make([]IAccessorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessorContext); ok {
			tst[i] = t.(IAccessorContext)
			i++
		}
	}

	return tst
}

func (s *ReferenceContext) Accessor(i int) IAccessorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorContext); ok {
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

	return t.(IAccessorContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitReference(s)
	}
}

func (s *ReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DataMorfParserRULE_reference)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(203)
			p.FunctionCall()
		}

	case 2:
		{
			p.SetState(204)
			p.Match(DataMorfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(207)
				p.Accessor()
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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

// IAccessorContext is an interface to support dynamic dispatch.
type IAccessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AccessProperty() IAccessPropertyContext
	Dot() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	FunctionCall() IFunctionCallContext

	// IsAccessorContext differentiates from other interfaces.
	IsAccessorContext()
}

type AccessorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorContext() *AccessorContext {
	var p = new(AccessorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_accessor
	return p
}

func InitEmptyAccessorContext(p *AccessorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_accessor
}

func (*AccessorContext) IsAccessorContext() {}

func NewAccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorContext {
	var p = new(AccessorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_accessor

	return p
}

func (s *AccessorContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorContext) AccessProperty() IAccessPropertyContext {
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

func (s *AccessorContext) Dot() antlr.TerminalNode {
	return s.GetToken(DataMorfParserDot, 0)
}

func (s *AccessorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *AccessorContext) FunctionCall() IFunctionCallContext {
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

func (s *AccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterAccessor(s)
	}
}

func (s *AccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitAccessor(s)
	}
}

func (s *AccessorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitAccessor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Accessor() (localctx IAccessorContext) {
	localctx = NewAccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DataMorfParserRULE_accessor)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(213)
			p.AccessProperty()
		}

	case 2:
		{
			p.SetState(214)
			p.Match(DataMorfParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Match(DataMorfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(216)
			p.Match(DataMorfParserDot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
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

// IAccessPropertyContext is an interface to support dynamic dispatch.
type IAccessPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpenSquare() antlr.TerminalNode
	CloseSquare() antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_accessProperty
	return p
}

func InitEmptyAccessPropertyContext(p *AccessPropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_accessProperty
}

func (*AccessPropertyContext) IsAccessPropertyContext() {}

func NewAccessPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessPropertyContext {
	var p = new(AccessPropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_accessProperty

	return p
}

func (s *AccessPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessPropertyContext) OpenSquare() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenSquare, 0)
}

func (s *AccessPropertyContext) CloseSquare() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseSquare, 0)
}

func (s *AccessPropertyContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserStringLiteral, 0)
}

func (s *AccessPropertyContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserDecimalLiteral, 0)
}

func (s *AccessPropertyContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *AccessPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterAccessProperty(s)
	}
}

func (s *AccessPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitAccessProperty(s)
	}
}

func (s *AccessPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitAccessProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) AccessProperty() (localctx IAccessPropertyContext) {
	localctx = NewAccessPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DataMorfParserRULE_accessProperty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(DataMorfParserOpenSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1688854155231232) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(222)
		p.Match(DataMorfParserCloseSquare)
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
	OpenRound() antlr.TerminalNode
	CloseRound() antlr.TerminalNode
	SendingParams() ISendingParamsContext

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
	p.RuleIndex = DataMorfParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, 0)
}

func (s *FunctionCallContext) OpenRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, 0)
}

func (s *FunctionCallContext) CloseRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, 0)
}

func (s *FunctionCallContext) SendingParams() ISendingParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISendingParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISendingParamsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DataMorfParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(DataMorfParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(DataMorfParserOpenRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(226)
		p.SendingParams()
	}

	{
		p.SetState(227)
		p.Match(DataMorfParserCloseRound)
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

// ISignatureParamsContext is an interface to support dynamic dispatch.
type ISignatureParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsSignatureParamsContext differentiates from other interfaces.
	IsSignatureParamsContext()
}

type SignatureParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignatureParamsContext() *SignatureParamsContext {
	var p = new(SignatureParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_signatureParams
	return p
}

func InitEmptySignatureParamsContext(p *SignatureParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_signatureParams
}

func (*SignatureParamsContext) IsSignatureParamsContext() {}

func NewSignatureParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignatureParamsContext {
	var p = new(SignatureParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_signatureParams

	return p
}

func (s *SignatureParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *SignatureParamsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserIdentifier)
}

func (s *SignatureParamsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserIdentifier, i)
}

func (s *SignatureParamsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserComma)
}

func (s *SignatureParamsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserComma, i)
}

func (s *SignatureParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignatureParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignatureParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterSignatureParams(s)
	}
}

func (s *SignatureParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitSignatureParams(s)
	}
}

func (s *SignatureParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitSignatureParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) SignatureParams() (localctx ISignatureParamsContext) {
	localctx = NewSignatureParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DataMorfParserRULE_signatureParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DataMorfParserIdentifier {
		{
			p.SetState(229)
			p.Match(DataMorfParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DataMorfParserComma {
			{
				p.SetState(230)
				p.Match(DataMorfParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(231)
				p.Match(DataMorfParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(236)
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

// ISendingParamsContext is an interface to support dynamic dispatch.
type ISendingParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsSendingParamsContext differentiates from other interfaces.
	IsSendingParamsContext()
}

type SendingParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendingParamsContext() *SendingParamsContext {
	var p = new(SendingParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_sendingParams
	return p
}

func InitEmptySendingParamsContext(p *SendingParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_sendingParams
}

func (*SendingParamsContext) IsSendingParamsContext() {}

func NewSendingParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendingParamsContext {
	var p = new(SendingParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_sendingParams

	return p
}

func (s *SendingParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *SendingParamsContext) AllUnit() []IUnitContext {
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

func (s *SendingParamsContext) Unit(i int) IUnitContext {
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

func (s *SendingParamsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserComma)
}

func (s *SendingParamsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserComma, i)
}

func (s *SendingParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendingParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SendingParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterSendingParams(s)
	}
}

func (s *SendingParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitSendingParams(s)
	}
}

func (s *SendingParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitSendingParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) SendingParams() (localctx ISendingParamsContext) {
	localctx = NewSendingParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DataMorfParserRULE_sendingParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1688857376587792) != 0 {
		{
			p.SetState(239)
			p.unit(0)
		}
		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DataMorfParserComma {
			{
				p.SetState(240)
				p.Match(DataMorfParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(241)
				p.unit(0)
			}

			p.SetState(246)
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
	OpenRound() antlr.TerminalNode
	SignatureParams() ISignatureParamsContext
	CloseRound() antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_arrowFunction
	return p
}

func InitEmptyArrowFunctionContext(p *ArrowFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_arrowFunction
}

func (*ArrowFunctionContext) IsArrowFunctionContext() {}

func NewArrowFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowFunctionContext {
	var p = new(ArrowFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_arrowFunction

	return p
}

func (s *ArrowFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowFunctionContext) OpenRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenRound, 0)
}

func (s *ArrowFunctionContext) SignatureParams() ISignatureParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignatureParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignatureParamsContext)
}

func (s *ArrowFunctionContext) CloseRound() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseRound, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterArrowFunction(s)
	}
}

func (s *ArrowFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitArrowFunction(s)
	}
}

func (s *ArrowFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitArrowFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ArrowFunction() (localctx IArrowFunctionContext) {
	localctx = NewArrowFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DataMorfParserRULE_arrowFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(DataMorfParserOpenRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.SignatureParams()
	}
	{
		p.SetState(251)
		p.Match(DataMorfParserCloseRound)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(DataMorfParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
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
	OpenSquare() antlr.TerminalNode
	CloseSquare() antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) OpenSquare() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenSquare, 0)
}

func (s *ArrayLiteralContext) CloseSquare() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseSquare, 0)
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
	return s.GetTokens(DataMorfParserComma)
}

func (s *ArrayLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserComma, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DataMorfParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(DataMorfParserOpenSquare)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1688857376587860) != 0 {
		{
			p.SetState(256)
			p.Value()
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DataMorfParserComma {
			{
				p.SetState(257)
				p.Match(DataMorfParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(258)
				p.Value()
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(266)
		p.Match(DataMorfParserCloseSquare)
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
	OpenCurly() antlr.TerminalNode
	CloseCurly() antlr.TerminalNode
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
	p.RuleIndex = DataMorfParserRULE_objectLiteral
	return p
}

func InitEmptyObjectLiteralContext(p *ObjectLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_objectLiteral
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenCurly, 0)
}

func (s *ObjectLiteralContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseCurly, 0)
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
	return s.GetTokens(DataMorfParserComma)
}

func (s *ObjectLiteralContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserComma, i)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitObjectLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DataMorfParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(DataMorfParserOpenCurly)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125899906859012) != 0 {
		{
			p.SetState(269)
			p.ObjectItem()
		}
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DataMorfParserComma {
			{
				p.SetState(270)
				p.Match(DataMorfParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(271)
				p.ObjectItem()
			}

			p.SetState(276)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(279)
		p.Match(DataMorfParserCloseCurly)
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
	OpenSquare() antlr.TerminalNode
	Reference() IReferenceContext
	CloseSquare() antlr.TerminalNode

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
	p.RuleIndex = DataMorfParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserStringLiteral, 0)
}

func (s *KeyContext) OpenSquare() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenSquare, 0)
}

func (s *KeyContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *KeyContext) CloseSquare() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseSquare, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitKey(s)
	}
}

func (s *KeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DataMorfParserRULE_key)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DataMorfParserStringLiteral:
		{
			p.SetState(281)
			p.Match(DataMorfParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DataMorfParserOpenSquare:
		{
			p.SetState(282)
			p.Match(DataMorfParserOpenSquare)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Reference()
		}
		{
			p.SetState(284)
			p.Match(DataMorfParserCloseSquare)
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
	p.RuleIndex = DataMorfParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_value

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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DataMorfParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(288)
			p.FunctionCall()
		}

	case 2:
		{
			p.SetState(289)
			p.unit(0)
		}

	case 3:
		{
			p.SetState(290)
			p.ObjectLiteral()
		}

	case 4:
		{
			p.SetState(291)
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
	p.RuleIndex = DataMorfParserRULE_keyValue
	return p
}

func InitEmptyKeyValueContext(p *KeyValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_keyValue
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_keyValue

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
	return s.GetToken(DataMorfParserColon, 0)
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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (s *KeyValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitKeyValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DataMorfParserRULE_keyValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)
		p.Key()
	}
	{
		p.SetState(295)
		p.Match(DataMorfParserColon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
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
	p.RuleIndex = DataMorfParserRULE_objectItem
	return p
}

func InitEmptyObjectItemContext(p *ObjectItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_objectItem
}

func (*ObjectItemContext) IsObjectItemContext() {}

func NewObjectItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectItemContext {
	var p = new(ObjectItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_objectItem

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
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterObjectItem(s)
	}
}

func (s *ObjectItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitObjectItem(s)
	}
}

func (s *ObjectItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitObjectItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) ObjectItem() (localctx IObjectItemContext) {
	localctx = NewObjectItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DataMorfParserRULE_objectItem)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DataMorfParserOpenSquare, DataMorfParserStringLiteral:
		{
			p.SetState(298)
			p.KeyValue()
		}

	case DataMorfParserDot:
		{
			p.SetState(299)
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

// ISpreadContext is an interface to support dynamic dispatch.
type ISpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDot() []antlr.TerminalNode
	Dot(i int) antlr.TerminalNode
	Reference() IReferenceContext

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
	p.RuleIndex = DataMorfParserRULE_spread
	return p
}

func InitEmptySpreadContext(p *SpreadContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_spread
}

func (*SpreadContext) IsSpreadContext() {}

func NewSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpreadContext {
	var p = new(SpreadContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_spread

	return p
}

func (s *SpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *SpreadContext) AllDot() []antlr.TerminalNode {
	return s.GetTokens(DataMorfParserDot)
}

func (s *SpreadContext) Dot(i int) antlr.TerminalNode {
	return s.GetToken(DataMorfParserDot, i)
}

func (s *SpreadContext) Reference() IReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *SpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterSpread(s)
	}
}

func (s *SpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitSpread(s)
	}
}

func (s *SpreadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitSpread(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Spread() (localctx ISpreadContext) {
	localctx = NewSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DataMorfParserRULE_spread)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(DataMorfParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(DataMorfParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(DataMorfParserDot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Reference()
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
	OpenCurly() antlr.TerminalNode
	Statements() IStatementsContext
	CloseCurly() antlr.TerminalNode

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
	p.RuleIndex = DataMorfParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) OpenCurly() antlr.TerminalNode {
	return s.GetToken(DataMorfParserOpenCurly, 0)
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

func (s *BlockContext) CloseCurly() antlr.TerminalNode {
	return s.GetToken(DataMorfParserCloseCurly, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DataMorfParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(DataMorfParserOpenCurly)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Statements()
	}
	{
		p.SetState(309)
		p.Match(DataMorfParserCloseCurly)
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
	p.RuleIndex = DataMorfParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DataMorfParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataMorfParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserNullLiteral, 0)
}

func (s *ConstantContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserBooleanLiteral, 0)
}

func (s *ConstantContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserDecimalLiteral, 0)
}

func (s *ConstantContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DataMorfParserStringLiteral, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataMorfListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataMorfVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataMorfParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DataMorfParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1125907423035392) != 0) {
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

func (p *DataMorfParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *DataMorfParser) Unit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
