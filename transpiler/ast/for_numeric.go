package ast

import (
	"fmt"

	"github.com/tereus-project/tereus-transpiler-lua-ruby/utils"
)

type ASTNumericFor struct {
	VariableDeclaration *ASTVariableDeclaration
	Condition           IASTExpression
	Increment           IASTExpression
	Do                  *ASTChunk
}

func NewASTNumericFor(variableDeclaration *ASTVariableDeclaration, condition IASTExpression, increment IASTExpression, do *ASTChunk) *ASTNumericFor {
	return &ASTNumericFor{
		VariableDeclaration: variableDeclaration,
		Condition:           condition,
		Increment:           increment,
		Do:                  do,
	}
}

func (i *ASTNumericFor) String() string {
	var output string

	if i.Increment != nil {
		output += i.VariableDeclaration.String() + "\n"
		output += fmt.Sprintf("while %s != %s do\n", i.VariableDeclaration.Name, i.Condition.String())
		output += utils.Indent(i.Do.String())
		output += "\n\n"
		output += utils.Indent(fmt.Sprintf("%s = %s", i.VariableDeclaration.Name, i.Increment.String()))
	} else {
		output += fmt.Sprintf("for %s in %s..%s do\n", i.VariableDeclaration.Name, i.VariableDeclaration.Expression, i.Condition.String())
		output += utils.Indent(i.Do.String())
	}

	output += "\nend"

	return output
}
