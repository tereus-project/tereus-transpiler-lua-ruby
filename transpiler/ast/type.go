package ast

type ASTTypeKind int

const (
	ASTTypeKindUnknown ASTTypeKind = iota
	ASTTypeKindInt
	ASTTypeKindFloat
)

type ASTType struct {
	Kind ASTTypeKind
}

func NewASTType(kind ASTTypeKind) *ASTType {
	return &ASTType{
		Kind: kind,
	}
}
