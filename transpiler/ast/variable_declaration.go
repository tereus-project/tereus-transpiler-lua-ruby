package ast

import "fmt"

type ASTVariableDeclaration struct {
	Name       string
	Type       *ASTType
	Expression IASTExpression
}

func NewASTVariableDeclaration(name string, typ *ASTType, expression IASTExpression) *ASTVariableDeclaration {
	return &ASTVariableDeclaration{
		Name:       name,
		Type:       typ,
		Expression: expression,
	}
}

func (v *ASTVariableDeclaration) String() string {
	return fmt.Sprintf("%s = %s", v.Name, v.Expression.String())
}
