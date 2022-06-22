package ast

type ASTExpressionConstant struct {
	Value string
	Type  *ASTType
}

func NewASTExpressionConstant(value string, typ *ASTType) *ASTExpressionConstant {
	return &ASTExpressionConstant{
		Value: value,
		Type:  typ,
	}
}

func (e *ASTExpressionConstant) String() string {
	return e.Value
}

func (e *ASTExpressionConstant) GetType() *ASTType {
	return e.Type
}
