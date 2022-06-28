package ast

import (
	"github.com/tereus-project/tereus-transpiler-lua-ruby/utils"
)

type ASTIf struct {
	Condition IASTExpression
	Then      *ASTChunk
	ElseIf    []*ASTIf
	Else      *ASTChunk
}

func NewASTIf(condition IASTExpression, then *ASTChunk, elseIf []*ASTIf, else_ *ASTChunk) *ASTIf {
	return &ASTIf{
		Condition: condition,
		Then:      then,
		ElseIf:    elseIf,
		Else:      else_,
	}
}

func (i *ASTIf) AddElseIf(elseIf *ASTIf) {
	i.ElseIf = append(i.ElseIf, elseIf)
}

func (i *ASTIf) String() string {
	var output string

	output += "if " + i.Condition.String() + " then\n"
	output += utils.Indent(i.Then.String())

	for _, elseIf := range i.ElseIf {
		output += "\n"
		output += "elsif " + elseIf.Condition.String() + " then\n"
		output += utils.Indent(elseIf.Then.String())
	}

	output += "\n"

	if i.Else != nil {
		output += "else\n"
		output += utils.Indent(i.Else.String())
		output += "\n"
	}

	output += "end"

	return output
}
