package ast

import "fmt"

type ASTExpressionBinary struct {
	Left       IASTExpression
	Operator   *ASTBinaryOperator
	Right      IASTExpression
	ReturnType *ASTType
}

func NewASTExpressionBinary(left IASTExpression, operator *ASTBinaryOperator, right IASTExpression, returnType *ASTType) *ASTExpressionBinary {
	return &ASTExpressionBinary{
		Left:       left,
		Operator:   operator,
		Right:      right,
		ReturnType: returnType,
	}
}

func (e *ASTExpressionBinary) String() string {
	return fmt.Sprintf("%s %s %s", e.Left.String(), e.Operator.String(), e.Right.String())
}

func (e *ASTExpressionBinary) GetType() *ASTType {
	return e.ReturnType
}
