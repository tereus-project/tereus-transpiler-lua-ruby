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
		"", "'='", "'+'", "'-'", "'*'", "'/'", "'and'", "'break'", "'do'", "'else'",
		"'elseif'", "'end'", "'false'", "'for'", "'function'", "'if'", "'in'",
		"'local'", "'nil'", "'not'", "'or'", "'repeat'", "'return'", "'then'",
		"'true'", "'until'", "'while'",
	}
	staticData.symbolicNames = []string{
		"", "Equal", "Add", "Subtract", "Multiply", "Divide", "And", "Break",
		"Do", "Else", "ElseIf", "End", "False", "For", "Function", "If", "In",
		"Local", "Nil", "Not", "Or", "Repeat", "Return", "Then", "True", "Until",
		"While", "Identifier", "FloatConstant", "IntegerConstant", "BlockComment",
		"LineComment", "Whitespace", "Newline",
	}
	staticData.ruleNames = []string{
		"Equal", "Add", "Subtract", "Multiply", "Divide", "And", "Break", "Do",
		"Else", "ElseIf", "End", "False", "For", "Function", "If", "In", "Local",
		"Nil", "Not", "Or", "Repeat", "Return", "Then", "True", "Until", "While",
		"Identifier", "FloatConstant", "IntegerConstant", "BlockComment", "LineComment",
		"Whitespace", "Newline",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 254, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 5, 26, 187, 8, 26, 10, 26, 12, 26, 190, 9, 26, 1, 27,
		4, 27, 193, 8, 27, 11, 27, 12, 27, 194, 1, 27, 1, 27, 4, 27, 199, 8, 27,
		11, 27, 12, 27, 200, 1, 28, 1, 28, 5, 28, 205, 8, 28, 10, 28, 12, 28, 208,
		9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 216, 8, 29, 10,
		29, 12, 29, 219, 9, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 232, 8, 30, 10, 30, 12, 30, 235, 9,
		30, 1, 30, 1, 30, 1, 31, 4, 31, 240, 8, 31, 11, 31, 12, 31, 241, 1, 31,
		1, 31, 1, 32, 1, 32, 3, 32, 248, 8, 32, 1, 32, 3, 32, 251, 8, 32, 1, 32,
		1, 32, 1, 217, 0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8,
		17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 1, 0, 6, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48,
		57, 1, 0, 49, 57, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 262, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0,
		0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1,
		0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55,
		1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0,
		63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 1, 67, 1, 0, 0, 0, 3, 69, 1, 0, 0, 0,
		5, 71, 1, 0, 0, 0, 7, 73, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0, 11, 77, 1, 0,
		0, 0, 13, 81, 1, 0, 0, 0, 15, 87, 1, 0, 0, 0, 17, 90, 1, 0, 0, 0, 19, 95,
		1, 0, 0, 0, 21, 102, 1, 0, 0, 0, 23, 106, 1, 0, 0, 0, 25, 112, 1, 0, 0,
		0, 27, 116, 1, 0, 0, 0, 29, 125, 1, 0, 0, 0, 31, 128, 1, 0, 0, 0, 33, 131,
		1, 0, 0, 0, 35, 137, 1, 0, 0, 0, 37, 141, 1, 0, 0, 0, 39, 145, 1, 0, 0,
		0, 41, 148, 1, 0, 0, 0, 43, 155, 1, 0, 0, 0, 45, 162, 1, 0, 0, 0, 47, 167,
		1, 0, 0, 0, 49, 172, 1, 0, 0, 0, 51, 178, 1, 0, 0, 0, 53, 184, 1, 0, 0,
		0, 55, 192, 1, 0, 0, 0, 57, 202, 1, 0, 0, 0, 59, 209, 1, 0, 0, 0, 61, 227,
		1, 0, 0, 0, 63, 239, 1, 0, 0, 0, 65, 250, 1, 0, 0, 0, 67, 68, 5, 61, 0,
		0, 68, 2, 1, 0, 0, 0, 69, 70, 5, 43, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5,
		45, 0, 0, 72, 6, 1, 0, 0, 0, 73, 74, 5, 42, 0, 0, 74, 8, 1, 0, 0, 0, 75,
		76, 5, 47, 0, 0, 76, 10, 1, 0, 0, 0, 77, 78, 5, 97, 0, 0, 78, 79, 5, 110,
		0, 0, 79, 80, 5, 100, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5, 98, 0, 0, 82,
		83, 5, 114, 0, 0, 83, 84, 5, 101, 0, 0, 84, 85, 5, 97, 0, 0, 85, 86, 5,
		107, 0, 0, 86, 14, 1, 0, 0, 0, 87, 88, 5, 100, 0, 0, 88, 89, 5, 111, 0,
		0, 89, 16, 1, 0, 0, 0, 90, 91, 5, 101, 0, 0, 91, 92, 5, 108, 0, 0, 92,
		93, 5, 115, 0, 0, 93, 94, 5, 101, 0, 0, 94, 18, 1, 0, 0, 0, 95, 96, 5,
		101, 0, 0, 96, 97, 5, 108, 0, 0, 97, 98, 5, 115, 0, 0, 98, 99, 5, 101,
		0, 0, 99, 100, 5, 105, 0, 0, 100, 101, 5, 102, 0, 0, 101, 20, 1, 0, 0,
		0, 102, 103, 5, 101, 0, 0, 103, 104, 5, 110, 0, 0, 104, 105, 5, 100, 0,
		0, 105, 22, 1, 0, 0, 0, 106, 107, 5, 102, 0, 0, 107, 108, 5, 97, 0, 0,
		108, 109, 5, 108, 0, 0, 109, 110, 5, 115, 0, 0, 110, 111, 5, 101, 0, 0,
		111, 24, 1, 0, 0, 0, 112, 113, 5, 102, 0, 0, 113, 114, 5, 111, 0, 0, 114,
		115, 5, 114, 0, 0, 115, 26, 1, 0, 0, 0, 116, 117, 5, 102, 0, 0, 117, 118,
		5, 117, 0, 0, 118, 119, 5, 110, 0, 0, 119, 120, 5, 99, 0, 0, 120, 121,
		5, 116, 0, 0, 121, 122, 5, 105, 0, 0, 122, 123, 5, 111, 0, 0, 123, 124,
		5, 110, 0, 0, 124, 28, 1, 0, 0, 0, 125, 126, 5, 105, 0, 0, 126, 127, 5,
		102, 0, 0, 127, 30, 1, 0, 0, 0, 128, 129, 5, 105, 0, 0, 129, 130, 5, 110,
		0, 0, 130, 32, 1, 0, 0, 0, 131, 132, 5, 108, 0, 0, 132, 133, 5, 111, 0,
		0, 133, 134, 5, 99, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 108, 0, 0,
		136, 34, 1, 0, 0, 0, 137, 138, 5, 110, 0, 0, 138, 139, 5, 105, 0, 0, 139,
		140, 5, 108, 0, 0, 140, 36, 1, 0, 0, 0, 141, 142, 5, 110, 0, 0, 142, 143,
		5, 111, 0, 0, 143, 144, 5, 116, 0, 0, 144, 38, 1, 0, 0, 0, 145, 146, 5,
		111, 0, 0, 146, 147, 5, 114, 0, 0, 147, 40, 1, 0, 0, 0, 148, 149, 5, 114,
		0, 0, 149, 150, 5, 101, 0, 0, 150, 151, 5, 112, 0, 0, 151, 152, 5, 101,
		0, 0, 152, 153, 5, 97, 0, 0, 153, 154, 5, 116, 0, 0, 154, 42, 1, 0, 0,
		0, 155, 156, 5, 114, 0, 0, 156, 157, 5, 101, 0, 0, 157, 158, 5, 116, 0,
		0, 158, 159, 5, 117, 0, 0, 159, 160, 5, 114, 0, 0, 160, 161, 5, 110, 0,
		0, 161, 44, 1, 0, 0, 0, 162, 163, 5, 116, 0, 0, 163, 164, 5, 104, 0, 0,
		164, 165, 5, 101, 0, 0, 165, 166, 5, 110, 0, 0, 166, 46, 1, 0, 0, 0, 167,
		168, 5, 116, 0, 0, 168, 169, 5, 114, 0, 0, 169, 170, 5, 117, 0, 0, 170,
		171, 5, 101, 0, 0, 171, 48, 1, 0, 0, 0, 172, 173, 5, 117, 0, 0, 173, 174,
		5, 110, 0, 0, 174, 175, 5, 116, 0, 0, 175, 176, 5, 105, 0, 0, 176, 177,
		5, 108, 0, 0, 177, 50, 1, 0, 0, 0, 178, 179, 5, 119, 0, 0, 179, 180, 5,
		104, 0, 0, 180, 181, 5, 105, 0, 0, 181, 182, 5, 108, 0, 0, 182, 183, 5,
		101, 0, 0, 183, 52, 1, 0, 0, 0, 184, 188, 7, 0, 0, 0, 185, 187, 7, 1, 0,
		0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188,
		189, 1, 0, 0, 0, 189, 54, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191, 193, 7,
		2, 0, 0, 192, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 192, 1, 0, 0,
		0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 5, 46, 0, 0, 197,
		199, 7, 2, 0, 0, 198, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 198,
		1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 56, 1, 0, 0, 0, 202, 206, 7, 3,
		0, 0, 203, 205, 7, 2, 0, 0, 204, 203, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0,
		206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 58, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 209, 210, 5, 45, 0, 0, 210, 211, 5, 45, 0, 0, 211, 212, 5,
		91, 0, 0, 212, 213, 5, 91, 0, 0, 213, 217, 1, 0, 0, 0, 214, 216, 9, 0,
		0, 0, 215, 214, 1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0,
		217, 215, 1, 0, 0, 0, 218, 220, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220,
		221, 5, 45, 0, 0, 221, 222, 5, 45, 0, 0, 222, 223, 5, 93, 0, 0, 223, 224,
		5, 93, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 6, 29, 0, 0, 226, 60, 1, 0,
		0, 0, 227, 228, 5, 45, 0, 0, 228, 229, 5, 45, 0, 0, 229, 233, 1, 0, 0,
		0, 230, 232, 8, 4, 0, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233,
		231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 236, 1, 0, 0, 0, 235, 233,
		1, 0, 0, 0, 236, 237, 6, 30, 0, 0, 237, 62, 1, 0, 0, 0, 238, 240, 7, 5,
		0, 0, 239, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0,
		241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 6, 31, 0, 0, 244,
		64, 1, 0, 0, 0, 245, 247, 5, 13, 0, 0, 246, 248, 5, 10, 0, 0, 247, 246,
		1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 251, 5, 10,
		0, 0, 250, 245, 1, 0, 0, 0, 250, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0,
		252, 253, 6, 32, 0, 0, 253, 66, 1, 0, 0, 0, 10, 0, 188, 194, 200, 206,
		217, 233, 241, 247, 250, 1, 6, 0, 0,
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
	LuaLexerAnd             = 6
	LuaLexerBreak           = 7
	LuaLexerDo              = 8
	LuaLexerElse            = 9
	LuaLexerElseIf          = 10
	LuaLexerEnd             = 11
	LuaLexerFalse           = 12
	LuaLexerFor             = 13
	LuaLexerFunction        = 14
	LuaLexerIf              = 15
	LuaLexerIn              = 16
	LuaLexerLocal           = 17
	LuaLexerNil             = 18
	LuaLexerNot             = 19
	LuaLexerOr              = 20
	LuaLexerRepeat          = 21
	LuaLexerReturn          = 22
	LuaLexerThen            = 23
	LuaLexerTrue            = 24
	LuaLexerUntil           = 25
	LuaLexerWhile           = 26
	LuaLexerIdentifier      = 27
	LuaLexerFloatConstant   = 28
	LuaLexerIntegerConstant = 29
	LuaLexerBlockComment    = 30
	LuaLexerLineComment     = 31
	LuaLexerWhitespace      = 32
	LuaLexerNewline         = 33
)
