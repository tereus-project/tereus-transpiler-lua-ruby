package ast

type ASTBinaryOperatorKind int

const (
	ASTBinaryOperatorKindUnknown ASTBinaryOperatorKind = iota
	ASTBinaryOperatorKindAdd
	ASTBinaryOperatorKindSubtract
	ASTBinaryOperatorKindMultiply
	ASTBinaryOperatorKindDivide
)

type ASTBinaryOperator struct {
	Kind ASTBinaryOperatorKind
}

func NewASTBinaryOperator(kind ASTBinaryOperatorKind) *ASTBinaryOperator {
	return &ASTBinaryOperator{
		Kind: kind,
	}
}

func (o *ASTBinaryOperator) String() string {
	switch o.Kind {
	case ASTBinaryOperatorKindAdd:
		return "+"
	case ASTBinaryOperatorKindSubtract:
		return "-"
	case ASTBinaryOperatorKindMultiply:
		return "*"
	case ASTBinaryOperatorKindDivide:
		return "/"
	default:
		return "?"
	}
}
