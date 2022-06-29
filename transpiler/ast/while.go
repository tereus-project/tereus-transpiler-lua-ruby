package ast

import (
	"fmt"

	"github.com/tereus-project/tereus-transpiler-lua-ruby/utils"
)

type ASTWhile struct {
	Condition IASTExpression
	Do        *ASTChunk
}

func NewASTWhile(condition IASTExpression, do *ASTChunk) *ASTWhile {
	return &ASTWhile{
		Condition: condition,
		Do:        do,
	}
}

func (i *ASTWhile) String() string {
	var output string

	output += fmt.Sprintf("while %s do\n", i.Condition.String())
	output += utils.Indent(i.Do.String())
	output += "\nend"

	return output
}
