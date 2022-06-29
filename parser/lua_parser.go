// Code generated from Lua.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lua

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LuaParser struct {
	*antlr.BaseParser
}

var luaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func luaParserInit() {
	staticData := &luaParserStaticData
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
		"translation", "chunk", "statement", "variableDeclaration", "expression",
		"binaryOperator", "ifStatement", "elseifStatement", "elseStatement",
		"whileStatement", "numericForStatement",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 97, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 48, 8, 4, 10, 4, 12, 4, 51, 9, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 60, 8, 6, 10, 6, 12, 6, 63, 9, 6,
		1, 6, 3, 6, 66, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 91, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		0, 1, 8, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 2, 1, 0, 29, 30,
		1, 0, 2, 5, 93, 0, 22, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 34, 1, 0, 0, 0,
		6, 36, 1, 0, 0, 0, 8, 40, 1, 0, 0, 0, 10, 52, 1, 0, 0, 0, 12, 54, 1, 0,
		0, 0, 14, 69, 1, 0, 0, 0, 16, 73, 1, 0, 0, 0, 18, 76, 1, 0, 0, 0, 20, 82,
		1, 0, 0, 0, 22, 23, 3, 2, 1, 0, 23, 1, 1, 0, 0, 0, 24, 26, 3, 4, 2, 0,
		25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1,
		0, 0, 0, 28, 3, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 35, 3, 6, 3, 0, 31,
		35, 3, 12, 6, 0, 32, 35, 3, 18, 9, 0, 33, 35, 3, 20, 10, 0, 34, 30, 1,
		0, 0, 0, 34, 31, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 33, 1, 0, 0, 0, 35,
		5, 1, 0, 0, 0, 36, 37, 5, 28, 0, 0, 37, 38, 5, 1, 0, 0, 38, 39, 3, 8, 4,
		0, 39, 7, 1, 0, 0, 0, 40, 41, 6, 4, -1, 0, 41, 42, 7, 0, 0, 0, 42, 49,
		1, 0, 0, 0, 43, 44, 10, 1, 0, 0, 44, 45, 3, 10, 5, 0, 45, 46, 3, 8, 4,
		2, 46, 48, 1, 0, 0, 0, 47, 43, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 9, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		52, 53, 7, 1, 0, 0, 53, 11, 1, 0, 0, 0, 54, 55, 5, 16, 0, 0, 55, 56, 3,
		8, 4, 0, 56, 57, 5, 24, 0, 0, 57, 61, 3, 2, 1, 0, 58, 60, 3, 14, 7, 0,
		59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 66, 3, 16, 8, 0, 65,
		64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 5, 12,
		0, 0, 68, 13, 1, 0, 0, 0, 69, 70, 5, 11, 0, 0, 70, 71, 3, 8, 4, 0, 71,
		72, 3, 2, 1, 0, 72, 15, 1, 0, 0, 0, 73, 74, 5, 10, 0, 0, 74, 75, 3, 2,
		1, 0, 75, 17, 1, 0, 0, 0, 76, 77, 5, 27, 0, 0, 77, 78, 3, 8, 4, 0, 78,
		79, 5, 9, 0, 0, 79, 80, 3, 2, 1, 0, 80, 81, 5, 12, 0, 0, 81, 19, 1, 0,
		0, 0, 82, 83, 5, 14, 0, 0, 83, 84, 5, 28, 0, 0, 84, 85, 5, 1, 0, 0, 85,
		86, 3, 8, 4, 0, 86, 87, 5, 6, 0, 0, 87, 90, 3, 8, 4, 0, 88, 89, 5, 6, 0,
		0, 89, 91, 3, 8, 4, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 93, 5, 9, 0, 0, 93, 94, 3, 2, 1, 0, 94, 95, 5, 12, 0, 0,
		95, 21, 1, 0, 0, 0, 6, 27, 34, 49, 61, 65, 90,
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

// LuaParserInit initializes any static state used to implement LuaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLuaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LuaParserInit() {
	staticData := &luaParserStaticData
	staticData.once.Do(luaParserInit)
}

// NewLuaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLuaParser(input antlr.TokenStream) *LuaParser {
	LuaParserInit()
	this := new(LuaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &luaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Lua.g4"

	return this
}

// LuaParser tokens.
const (
	LuaParserEOF             = antlr.TokenEOF
	LuaParserEqual           = 1
	LuaParserAdd             = 2
	LuaParserSubtract        = 3
	LuaParserMultiply        = 4
	LuaParserDivide          = 5
	LuaParserComma           = 6
	LuaParserAnd             = 7
	LuaParserBreak           = 8
	LuaParserDo              = 9
	LuaParserElse            = 10
	LuaParserElseIf          = 11
	LuaParserEnd             = 12
	LuaParserFalse           = 13
	LuaParserFor             = 14
	LuaParserFunction        = 15
	LuaParserIf              = 16
	LuaParserIn              = 17
	LuaParserLocal           = 18
	LuaParserNil             = 19
	LuaParserNot             = 20
	LuaParserOr              = 21
	LuaParserRepeat          = 22
	LuaParserReturn          = 23
	LuaParserThen            = 24
	LuaParserTrue            = 25
	LuaParserUntil           = 26
	LuaParserWhile           = 27
	LuaParserIdentifier      = 28
	LuaParserFloatConstant   = 29
	LuaParserIntegerConstant = 30
	LuaParserBlockComment    = 31
	LuaParserLineComment     = 32
	LuaParserWhitespace      = 33
	LuaParserNewline         = 34
)

// LuaParser rules.
const (
	LuaParserRULE_translation         = 0
	LuaParserRULE_chunk               = 1
	LuaParserRULE_statement           = 2
	LuaParserRULE_variableDeclaration = 3
	LuaParserRULE_expression          = 4
	LuaParserRULE_binaryOperator      = 5
	LuaParserRULE_ifStatement         = 6
	LuaParserRULE_elseifStatement     = 7
	LuaParserRULE_elseStatement       = 8
	LuaParserRULE_whileStatement      = 9
	LuaParserRULE_numericForStatement = 10
)

// ITranslationContext is an interface to support dynamic dispatch.
type ITranslationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslationContext differentiates from other interfaces.
	IsTranslationContext()
}

type TranslationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslationContext() *TranslationContext {
	var p = new(TranslationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_translation
	return p
}

func (*TranslationContext) IsTranslationContext() {}

func NewTranslationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslationContext {
	var p = new(TranslationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_translation

	return p
}

func (s *TranslationContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslationContext) Chunk() IChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChunkContext)
}

func (s *TranslationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitTranslation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Translation() (localctx ITranslationContext) {
	this := p
	_ = this

	localctx = NewTranslationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LuaParserRULE_translation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Chunk()
	}

	return localctx
}

// IChunkContext is an interface to support dynamic dispatch.
type IChunkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChunkContext differentiates from other interfaces.
	IsChunkContext()
}

type ChunkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChunkContext() *ChunkContext {
	var p = new(ChunkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_chunk
	return p
}

func (*ChunkContext) IsChunkContext() {}

func NewChunkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChunkContext {
	var p = new(ChunkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_chunk

	return p
}

func (s *ChunkContext) GetParser() antlr.Parser { return s.parser }

func (s *ChunkContext) AllStatement() []IStatementContext {
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

func (s *ChunkContext) Statement(i int) IStatementContext {
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

func (s *ChunkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChunkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChunkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitChunk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Chunk() (localctx IChunkContext) {
	this := p
	_ = this

	localctx = NewChunkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LuaParserRULE_chunk)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserFor)|(1<<LuaParserIf)|(1<<LuaParserWhile)|(1<<LuaParserIdentifier))) != 0 {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
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

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) NumericForStatement() INumericForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericForStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LuaParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LuaParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.VariableDeclaration()
		}

	case LuaParserIf:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.IfStatement()
		}

	case LuaParserWhile:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.WhileStatement()
		}

	case LuaParserFor:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(33)
			p.NumericForStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LuaParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Equal() antlr.TerminalNode {
	return s.GetToken(LuaParserEqual, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	this := p
	_ = this

	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LuaParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(LuaParserIdentifier)
	}
	{
		p.SetState(37)
		p.Match(LuaParserEqual)
	}
	{
		p.SetState(38)
		p.expression(0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionBinaryContext struct {
	*ExpressionContext
}

func NewExpressionBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionBinaryContext {
	var p = new(ExpressionBinaryContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionBinaryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionBinaryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionBinaryContext) BinaryOperator() IBinaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *ExpressionBinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpressionBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionConstantContext struct {
	*ExpressionContext
}

func NewExpressionConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionConstantContext {
	var p = new(ExpressionConstantContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressionConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionConstantContext) FloatConstant() antlr.TerminalNode {
	return s.GetToken(LuaParserFloatConstant, 0)
}

func (s *ExpressionConstantContext) IntegerConstant() antlr.TerminalNode {
	return s.GetToken(LuaParserIntegerConstant, 0)
}

func (s *ExpressionConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitExpressionConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LuaParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, LuaParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewExpressionConstantContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LuaParserFloatConstant || _la == LuaParserIntegerConstant) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LuaParserRULE_expression)
			p.SetState(43)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(44)
				p.BinaryOperator()
			}
			{
				p.SetState(45)
				p.expression(2)
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_binaryOperator
	return p
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) Add() antlr.TerminalNode {
	return s.GetToken(LuaParserAdd, 0)
}

func (s *BinaryOperatorContext) Subtract() antlr.TerminalNode {
	return s.GetToken(LuaParserSubtract, 0)
}

func (s *BinaryOperatorContext) Multiply() antlr.TerminalNode {
	return s.GetToken(LuaParserMultiply, 0)
}

func (s *BinaryOperatorContext) Divide() antlr.TerminalNode {
	return s.GetToken(LuaParserDivide, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	this := p
	_ = this

	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LuaParserRULE_binaryOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LuaParserAdd)|(1<<LuaParserSubtract)|(1<<LuaParserMultiply)|(1<<LuaParserDivide))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(LuaParserIf, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) Then() antlr.TerminalNode {
	return s.GetToken(LuaParserThen, 0)
}

func (s *IfStatementContext) Chunk() IChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChunkContext)
}

func (s *IfStatementContext) End() antlr.TerminalNode {
	return s.GetToken(LuaParserEnd, 0)
}

func (s *IfStatementContext) AllElseifStatement() []IElseifStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseifStatementContext); ok {
			len++
		}
	}

	tst := make([]IElseifStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseifStatementContext); ok {
			tst[i] = t.(IElseifStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseifStatement(i int) IElseifStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifStatementContext); ok {
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

	return t.(IElseifStatementContext)
}

func (s *IfStatementContext) ElseStatement() IElseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LuaParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(LuaParserIf)
	}
	{
		p.SetState(55)
		p.expression(0)
	}
	{
		p.SetState(56)
		p.Match(LuaParserThen)
	}
	{
		p.SetState(57)
		p.Chunk()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LuaParserElseIf {
		{
			p.SetState(58)
			p.ElseifStatement()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserElse {
		{
			p.SetState(64)
			p.ElseStatement()
		}

	}
	{
		p.SetState(67)
		p.Match(LuaParserEnd)
	}

	return localctx
}

// IElseifStatementContext is an interface to support dynamic dispatch.
type IElseifStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseifStatementContext differentiates from other interfaces.
	IsElseifStatementContext()
}

type ElseifStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseifStatementContext() *ElseifStatementContext {
	var p = new(ElseifStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_elseifStatement
	return p
}

func (*ElseifStatementContext) IsElseifStatementContext() {}

func NewElseifStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifStatementContext {
	var p = new(ElseifStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_elseifStatement

	return p
}

func (s *ElseifStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifStatementContext) ElseIf() antlr.TerminalNode {
	return s.GetToken(LuaParserElseIf, 0)
}

func (s *ElseifStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseifStatementContext) Chunk() IChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChunkContext)
}

func (s *ElseifStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitElseifStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) ElseifStatement() (localctx IElseifStatementContext) {
	this := p
	_ = this

	localctx = NewElseifStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LuaParserRULE_elseifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(LuaParserElseIf)
	}
	{
		p.SetState(70)
		p.expression(0)
	}
	{
		p.SetState(71)
		p.Chunk()
	}

	return localctx
}

// IElseStatementContext is an interface to support dynamic dispatch.
type IElseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatementContext differentiates from other interfaces.
	IsElseStatementContext()
}

type ElseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatementContext() *ElseStatementContext {
	var p = new(ElseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_elseStatement
	return p
}

func (*ElseStatementContext) IsElseStatementContext() {}

func NewElseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatementContext {
	var p = new(ElseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_elseStatement

	return p
}

func (s *ElseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(LuaParserElse, 0)
}

func (s *ElseStatementContext) Chunk() IChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChunkContext)
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) ElseStatement() (localctx IElseStatementContext) {
	this := p
	_ = this

	localctx = NewElseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LuaParserRULE_elseStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(LuaParserElse)
	}
	{
		p.SetState(74)
		p.Chunk()
	}

	return localctx
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_whileStatement
	return p
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(LuaParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(LuaParserDo, 0)
}

func (s *WhileStatementContext) Chunk() IChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChunkContext)
}

func (s *WhileStatementContext) End() antlr.TerminalNode {
	return s.GetToken(LuaParserEnd, 0)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) WhileStatement() (localctx IWhileStatementContext) {
	this := p
	_ = this

	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LuaParserRULE_whileStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(LuaParserWhile)
	}
	{
		p.SetState(77)
		p.expression(0)
	}
	{
		p.SetState(78)
		p.Match(LuaParserDo)
	}
	{
		p.SetState(79)
		p.Chunk()
	}
	{
		p.SetState(80)
		p.Match(LuaParserEnd)
	}

	return localctx
}

// INumericForStatementContext is an interface to support dynamic dispatch.
type INumericForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericForStatementContext differentiates from other interfaces.
	IsNumericForStatementContext()
}

type NumericForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericForStatementContext() *NumericForStatementContext {
	var p = new(NumericForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LuaParserRULE_numericForStatement
	return p
}

func (*NumericForStatementContext) IsNumericForStatementContext() {}

func NewNumericForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericForStatementContext {
	var p = new(NumericForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LuaParserRULE_numericForStatement

	return p
}

func (s *NumericForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(LuaParserFor, 0)
}

func (s *NumericForStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LuaParserIdentifier, 0)
}

func (s *NumericForStatementContext) Equal() antlr.TerminalNode {
	return s.GetToken(LuaParserEqual, 0)
}

func (s *NumericForStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *NumericForStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *NumericForStatementContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(LuaParserComma)
}

func (s *NumericForStatementContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(LuaParserComma, i)
}

func (s *NumericForStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(LuaParserDo, 0)
}

func (s *NumericForStatementContext) Chunk() IChunkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChunkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChunkContext)
}

func (s *NumericForStatementContext) End() antlr.TerminalNode {
	return s.GetToken(LuaParserEnd, 0)
}

func (s *NumericForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LuaVisitor:
		return t.VisitNumericForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LuaParser) NumericForStatement() (localctx INumericForStatementContext) {
	this := p
	_ = this

	localctx = NewNumericForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LuaParserRULE_numericForStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(LuaParserFor)
	}
	{
		p.SetState(83)
		p.Match(LuaParserIdentifier)
	}
	{
		p.SetState(84)
		p.Match(LuaParserEqual)
	}
	{
		p.SetState(85)
		p.expression(0)
	}
	{
		p.SetState(86)
		p.Match(LuaParserComma)
	}
	{
		p.SetState(87)
		p.expression(0)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LuaParserComma {
		{
			p.SetState(88)
			p.Match(LuaParserComma)
		}
		{
			p.SetState(89)
			p.expression(0)
		}

	}
	{
		p.SetState(92)
		p.Match(LuaParserDo)
	}
	{
		p.SetState(93)
		p.Chunk()
	}
	{
		p.SetState(94)
		p.Match(LuaParserEnd)
	}

	return localctx
}

func (p *LuaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LuaParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
