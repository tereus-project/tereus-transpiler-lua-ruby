package transpiler

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-transpiler-template/parser"
)

type RemixerErrorListener struct {
	filename string

	Errors []string
}

func NewRemixerErrorListener(filename string) *RemixerErrorListener {
	return &RemixerErrorListener{
		filename: filename,
	}
}

func (l *RemixerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, fmt.Sprintf("%s:%d:%d: %s", l.filename, line, column, msg))
}

func (l *RemixerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *RemixerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

func (l *RemixerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
}

func Transpile(entrypoint string) (string, error) {
	input, err := antlr.NewFileStream(entrypoint)
	if err != nil {
		return "", err
	}

	lexer := parser.NewLuaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(stream)
	p.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)
	p.RemoveErrorListeners()

	errorListener := NewRemixerErrorListener(entrypoint)
	p.AddErrorListener(errorListener)

	tree := p.Translation()

	if len(errorListener.Errors) > 0 {
		return "", fmt.Errorf("%s", strings.Join(errorListener.Errors, "\n"))
	}

	visitor := NewVisitor(entrypoint)
	output, err := visitor.VisitTranslation(tree.(*parser.TranslationContext))
	if err != nil {
		return "", err
	}

	return output, nil
}
