package transpiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-transpiler-template/parser"
	"github.com/tereus-project/tereus-transpiler-template/transpiler/ast"
)

type Visitor struct {
	Path string
}

func NewVisitor(path string) *Visitor {
	return &Visitor{
		Path: path,
	}
}

func (v *Visitor) PositionedTranslationError(token antlr.Token, message string) error {
	return fmt.Errorf("%s:%d:%d: %s", v.Path, token.GetLine(), token.GetColumn(), message)
}

func (v *Visitor) NotImplementedError(token antlr.Token) error {
	return v.PositionedTranslationError(token, "not implemented")
}

func (v *Visitor) VisitTranslation(ctx *parser.TranslationContext) (string, error) {
	output := ""

	for _, topLevelStatement := range ctx.AllTopLevelStatement() {
		topLevelStatement, err := v.VisitTopLevelStatement(topLevelStatement.(*parser.TopLevelStatementContext))
		if err != nil {
			return "", err
		}

		output += topLevelStatement.String() + "\n"
	}

	return output, nil
}

func (v *Visitor) VisitTopLevelStatement(ctx *parser.TopLevelStatementContext) (ast.IASTItem, error) {
	if child := ctx.VariableDeclaration(); child != nil {
		return v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
	}

	return nil, v.NotImplementedError(ctx.GetStart())
}

func (v *Visitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) (*ast.ASTVariableDeclaration, error) {
	name := ctx.Identifier().GetText()

	expression, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	return ast.NewASTVariableDeclaration(name, expression.GetType(), expression), nil
}

func (v *Visitor) VisitExpression(ctx parser.IExpressionContext) (ast.IASTExpression, error) {
	switch child := ctx.(type) {
	case *parser.ExpressionConstantContext:
		return v.VisitExpressionConstant(child)
	case *parser.ExpressionBinaryContext:
		return v.VisitExpressionBinary(child)
	}

	return nil, v.NotImplementedError(ctx.GetStart())
}

func (v *Visitor) VisitExpressionConstant(ctx *parser.ExpressionConstantContext) (*ast.ASTExpressionConstant, error) {
	if floatConstant := ctx.FloatConstant(); floatConstant != nil {
		return ast.NewASTExpressionConstant(floatConstant.GetText(), ast.NewASTType(ast.ASTTypeKindFloat)), nil
	}

	if integerConstant := ctx.IntegerConstant(); integerConstant != nil {
		return ast.NewASTExpressionConstant(integerConstant.GetText(), ast.NewASTType(ast.ASTTypeKindInt)), nil
	}

	return nil, v.NotImplementedError(ctx.GetStart())
}

func (v *Visitor) VisitExpressionBinary(ctx *parser.ExpressionBinaryContext) (*ast.ASTExpressionBinary, error) {
	leftExpression, err := v.VisitExpression(ctx.Expression(0))
	if err != nil {
		return nil, err
	}

	rightExpression, err := v.VisitExpression(ctx.Expression(1))
	if err != nil {
		return nil, err
	}

	operator, err := v.VisitBinaryOperator(ctx.BinaryOperator().(*parser.BinaryOperatorContext))
	if err != nil {
		return nil, err
	}

	return ast.NewASTExpressionBinary(leftExpression, operator, rightExpression, leftExpression.GetType()), nil
}

func (v *Visitor) VisitBinaryOperator(ctx *parser.BinaryOperatorContext) (*ast.ASTBinaryOperator, error) {
	switch ctx.GetText() {
	case "+":
		return ast.NewASTBinaryOperator(ast.ASTBinaryOperatorKindAdd), nil
	case "-":
		return ast.NewASTBinaryOperator(ast.ASTBinaryOperatorKindSubtract), nil
	case "*":
		return ast.NewASTBinaryOperator(ast.ASTBinaryOperatorKindMultiply), nil
	case "/":
		return ast.NewASTBinaryOperator(ast.ASTBinaryOperatorKindDivide), nil
	}

	return nil, v.NotImplementedError(ctx.GetStart())
}
