// Code generated from Lua.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LuaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lualexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lualexerLexerInit() {
	staticData := &lualexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'='", "'+'", "'-'", "'*'", "'/'", "','", "'and'", "'break'", "'do'",
		"'else'", "'elseif'", "'end'", "'false'", "'for'", "'function'", "'if'",
		"'in'", "'local'", "'nil'", "'not'", "'or'", "'repeat'", "'return'",
		"'then'", "'true'", "'until'", "'while'",
	}
	staticData.symbolicNames = []string{
		"", "Equal", "Add", "Subtract", "Multiply", "Divide", "Comma", "And",
		"Break", "Do", "Else", "ElseIf", "End", "False", "For", "Function",
		"If", "In", "Local", "Nil", "Not", "Or", "Repeat", "Return", "Then",
		"True", "Until", "While", "Identifier", "FloatConstant", "IntegerConstant",
		"BlockComment", "LineComment", "Whitespace", "Newline",
	}
	staticData.ruleNames = []string{
		"Equal", "Add", "Subtract", "Multiply", "Divide", "Comma", "And", "Break",
		"Do", "Else", "ElseIf", "End", "False", "For", "Function", "If", "In",
		"Local", "Nil", "Not", "Or", "Repeat", "Return", "Then", "True", "Until",
		"While", "Identifier", "FloatConstant", "IntegerConstant", "BlockComment",
		"LineComment", "Whitespace", "Newline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 258, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 191, 8, 27,
		10, 27, 12, 27, 194, 9, 27, 1, 28, 4, 28, 197, 8, 28, 11, 28, 12, 28, 198,
		1, 28, 1, 28, 4, 28, 203, 8, 28, 11, 28, 12, 28, 204, 1, 29, 1, 29, 5,
		29, 209, 8, 29, 10, 29, 12, 29, 212, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 5, 30, 220, 8, 30, 10, 30, 12, 30, 223, 9, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31,
		236, 8, 31, 10, 31, 12, 31, 239, 9, 31, 1, 31, 1, 31, 1, 32, 4, 32, 244,
		8, 32, 11, 32, 12, 32, 245, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 252, 8,
		33, 1, 33, 3, 33, 255, 8, 33, 1, 33, 1, 33, 1, 221, 0, 34, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 1, 0, 6, 3, 0, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 10,
		10, 13, 13, 2, 0, 9, 9, 32, 32, 266, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0,
		0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0,
		0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 73, 1, 0,
		0, 0, 7, 75, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79, 1, 0, 0, 0, 13, 81,
		1, 0, 0, 0, 15, 85, 1, 0, 0, 0, 17, 91, 1, 0, 0, 0, 19, 94, 1, 0, 0, 0,
		21, 99, 1, 0, 0, 0, 23, 106, 1, 0, 0, 0, 25, 110, 1, 0, 0, 0, 27, 116,
		1, 0, 0, 0, 29, 120, 1, 0, 0, 0, 31, 129, 1, 0, 0, 0, 33, 132, 1, 0, 0,
		0, 35, 135, 1, 0, 0, 0, 37, 141, 1, 0, 0, 0, 39, 145, 1, 0, 0, 0, 41, 149,
		1, 0, 0, 0, 43, 152, 1, 0, 0, 0, 45, 159, 1, 0, 0, 0, 47, 166, 1, 0, 0,
		0, 49, 171, 1, 0, 0, 0, 51, 176, 1, 0, 0, 0, 53, 182, 1, 0, 0, 0, 55, 188,
		1, 0, 0, 0, 57, 196, 1, 0, 0, 0, 59, 206, 1, 0, 0, 0, 61, 213, 1, 0, 0,
		0, 63, 231, 1, 0, 0, 0, 65, 243, 1, 0, 0, 0, 67, 254, 1, 0, 0, 0, 69, 70,
		5, 61, 0, 0, 70, 2, 1, 0, 0, 0, 71, 72, 5, 43, 0, 0, 72, 4, 1, 0, 0, 0,
		73, 74, 5, 45, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 42, 0, 0, 76, 8, 1,
		0, 0, 0, 77, 78, 5, 47, 0, 0, 78, 10, 1, 0, 0, 0, 79, 80, 5, 44, 0, 0,
		80, 12, 1, 0, 0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5, 110, 0, 0, 83, 84,
		5, 100, 0, 0, 84, 14, 1, 0, 0, 0, 85, 86, 5, 98, 0, 0, 86, 87, 5, 114,
		0, 0, 87, 88, 5, 101, 0, 0, 88, 89, 5, 97, 0, 0, 89, 90, 5, 107, 0, 0,
		90, 16, 1, 0, 0, 0, 91, 92, 5, 100, 0, 0, 92, 93, 5, 111, 0, 0, 93, 18,
		1, 0, 0, 0, 94, 95, 5, 101, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 115,
		0, 0, 97, 98, 5, 101, 0, 0, 98, 20, 1, 0, 0, 0, 99, 100, 5, 101, 0, 0,
		100, 101, 5, 108, 0, 0, 101, 102, 5, 115, 0, 0, 102, 103, 5, 101, 0, 0,
		103, 104, 5, 105, 0, 0, 104, 105, 5, 102, 0, 0, 105, 22, 1, 0, 0, 0, 106,
		107, 5, 101, 0, 0, 107, 108, 5, 110, 0, 0, 108, 109, 5, 100, 0, 0, 109,
		24, 1, 0, 0, 0, 110, 111, 5, 102, 0, 0, 111, 112, 5, 97, 0, 0, 112, 113,
		5, 108, 0, 0, 113, 114, 5, 115, 0, 0, 114, 115, 5, 101, 0, 0, 115, 26,
		1, 0, 0, 0, 116, 117, 5, 102, 0, 0, 117, 118, 5, 111, 0, 0, 118, 119, 5,
		114, 0, 0, 119, 28, 1, 0, 0, 0, 120, 121, 5, 102, 0, 0, 121, 122, 5, 117,
		0, 0, 122, 123, 5, 110, 0, 0, 123, 124, 5, 99, 0, 0, 124, 125, 5, 116,
		0, 0, 125, 126, 5, 105, 0, 0, 126, 127, 5, 111, 0, 0, 127, 128, 5, 110,
		0, 0, 128, 30, 1, 0, 0, 0, 129, 130, 5, 105, 0, 0, 130, 131, 5, 102, 0,
		0, 131, 32, 1, 0, 0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 110, 0, 0,
		134, 34, 1, 0, 0, 0, 135, 136, 5, 108, 0, 0, 136, 137, 5, 111, 0, 0, 137,
		138, 5, 99, 0, 0, 138, 139, 5, 97, 0, 0, 139, 140, 5, 108, 0, 0, 140, 36,
		1, 0, 0, 0, 141, 142, 5, 110, 0, 0, 142, 143, 5, 105, 0, 0, 143, 144, 5,
		108, 0, 0, 144, 38, 1, 0, 0, 0, 145, 146, 5, 110, 0, 0, 146, 147, 5, 111,
		0, 0, 147, 148, 5, 116, 0, 0, 148, 40, 1, 0, 0, 0, 149, 150, 5, 111, 0,
		0, 150, 151, 5, 114, 0, 0, 151, 42, 1, 0, 0, 0, 152, 153, 5, 114, 0, 0,
		153, 154, 5, 101, 0, 0, 154, 155, 5, 112, 0, 0, 155, 156, 5, 101, 0, 0,
		156, 157, 5, 97, 0, 0, 157, 158, 5, 116, 0, 0, 158, 44, 1, 0, 0, 0, 159,
		160, 5, 114, 0, 0, 160, 161, 5, 101, 0, 0, 161, 162, 5, 116, 0, 0, 162,
		163, 5, 117, 0, 0, 163, 164, 5, 114, 0, 0, 164, 165, 5, 110, 0, 0, 165,
		46, 1, 0, 0, 0, 166, 167, 5, 116, 0, 0, 167, 168, 5, 104, 0, 0, 168, 169,
		5, 101, 0, 0, 169, 170, 5, 110, 0, 0, 170, 48, 1, 0, 0, 0, 171, 172, 5,
		116, 0, 0, 172, 173, 5, 114, 0, 0, 173, 174, 5, 117, 0, 0, 174, 175, 5,
		101, 0, 0, 175, 50, 1, 0, 0, 0, 176, 177, 5, 117, 0, 0, 177, 178, 5, 110,
		0, 0, 178, 179, 5, 116, 0, 0, 179, 180, 5, 105, 0, 0, 180, 181, 5, 108,
		0, 0, 181, 52, 1, 0, 0, 0, 182, 183, 5, 119, 0, 0, 183, 184, 5, 104, 0,
		0, 184, 185, 5, 105, 0, 0, 185, 186, 5, 108, 0, 0, 186, 187, 5, 101, 0,
		0, 187, 54, 1, 0, 0, 0, 188, 192, 7, 0, 0, 0, 189, 191, 7, 1, 0, 0, 190,
		189, 1, 0, 0, 0, 191, 194, 1, 0, 0, 0, 192, 190, 1, 0, 0, 0, 192, 193,
		1, 0, 0, 0, 193, 56, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 195, 197, 7, 2,
		0, 0, 196, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0,
		198, 199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 5, 46, 0, 0, 201,
		203, 7, 2, 0, 0, 202, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 202,
		1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 58, 1, 0, 0, 0, 206, 210, 7, 3,
		0, 0, 207, 209, 7, 2, 0, 0, 208, 207, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0,
		210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 60, 1, 0, 0, 0, 212, 210,
		1, 0, 0, 0, 213, 214, 5, 45, 0, 0, 214, 215, 5, 45, 0, 0, 215, 216, 5,
		91, 0, 0, 216, 217, 5, 91, 0, 0, 217, 221, 1, 0, 0, 0, 218, 220, 9, 0,
		0, 0, 219, 218, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0,
		221, 219, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224,
		225, 5, 45, 0, 0, 225, 226, 5, 45, 0, 0, 226, 227, 5, 93, 0, 0, 227, 228,
		5, 93, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230, 6, 30, 0, 0, 230, 62, 1, 0,
		0, 0, 231, 232, 5, 45, 0, 0, 232, 233, 5, 45, 0, 0, 233, 237, 1, 0, 0,
		0, 234, 236, 8, 4, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237,
		235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237,
		1, 0, 0, 0, 240, 241, 6, 31, 0, 0, 241, 64, 1, 0, 0, 0, 242, 244, 7, 5,
		0, 0, 243, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0,
		245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 6, 32, 0, 0, 248,
		66, 1, 0, 0, 0, 249, 251, 5, 13, 0, 0, 250, 252, 5, 10, 0, 0, 251, 250,
		1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 255, 5, 10,
		0, 0, 254, 249, 1, 0, 0, 0, 254, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0,
		256, 257, 6, 33, 0, 0, 257, 68, 1, 0, 0, 0, 10, 0, 192, 198, 204, 210,
		221, 237, 245, 251, 254, 1, 6, 0, 0,
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

// LuaLexerInit initializes any static state used to implement LuaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLuaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LuaLexerInit() {
	staticData := &lualexerLexerStaticData
	staticData.once.Do(lualexerLexerInit)
}

// NewLuaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLuaLexer(input antlr.CharStream) *LuaLexer {
	LuaLexerInit()
	l := new(LuaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &lualexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Lua.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LuaLexer tokens.
const (
	LuaLexerEqual           = 1
	LuaLexerAdd             = 2
	LuaLexerSubtract        = 3
	LuaLexerMultiply        = 4
	LuaLexerDivide          = 5
	LuaLexerComma           = 6
	LuaLexerAnd             = 7
	LuaLexerBreak           = 8
	LuaLexerDo              = 9
	LuaLexerElse            = 10
	LuaLexerElseIf          = 11
	LuaLexerEnd             = 12
	LuaLexerFalse           = 13
	LuaLexerFor             = 14
	LuaLexerFunction        = 15
	LuaLexerIf              = 16
	LuaLexerIn              = 17
	LuaLexerLocal           = 18
	LuaLexerNil             = 19
	LuaLexerNot             = 20
	LuaLexerOr              = 21
	LuaLexerRepeat          = 22
	LuaLexerReturn          = 23
	LuaLexerThen            = 24
	LuaLexerTrue            = 25
	LuaLexerUntil           = 26
	LuaLexerWhile           = 27
	LuaLexerIdentifier      = 28
	LuaLexerFloatConstant   = 29
	LuaLexerIntegerConstant = 30
	LuaLexerBlockComment    = 31
	LuaLexerLineComment     = 32
	LuaLexerWhitespace      = 33
	LuaLexerNewline         = 34
)
