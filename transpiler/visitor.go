package transpiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tereus-project/tereus-transpiler-lua-ruby/parser"
	"github.com/tereus-project/tereus-transpiler-lua-ruby/transpiler/ast"
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
	chunk, err := v.VisitChunk(ctx.Chunk().(*parser.ChunkContext))
	if err != nil {
		return "", err
	}

	return chunk.String(), nil
}

func (v *Visitor) VisitChunk(ctx *parser.ChunkContext) (*ast.ASTChunk, error) {
	chunk := ast.NewASTChunk()

	for _, child := range ctx.AllStatement() {
		statement, err := v.VisitStatement(child.(*parser.StatementContext))
		if err != nil {
			return nil, err
		}

		chunk.Add(statement)
	}

	return chunk, nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) (ast.IASTItem, error) {
	if child := ctx.VariableDeclaration(); child != nil {
		return v.VisitVariableDeclaration(child.(*parser.VariableDeclarationContext))
	}

	if child := ctx.IfStatement(); child != nil {
		return v.VisitIfStatement(child.(*parser.IfStatementContext))
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

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) (*ast.ASTIf, error) {
	condition, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	then, err := v.VisitChunk(ctx.Chunk().(*parser.ChunkContext))
	if err != nil {
		return nil, err
	}

	ifItem := ast.NewASTIf(condition, then, nil, nil)

	for _, child := range ctx.AllElseifStatement() {
		elseIf, err := v.VisitElseifStatement(child.(*parser.ElseifStatementContext))
		if err != nil {
			return nil, err
		}

		ifItem.AddElseIf(elseIf)
	}

	if child := ctx.ElseStatement(); child != nil {
		elseStatement, err := v.VisitElseStatement(child.(*parser.ElseStatementContext))
		if err != nil {
			return nil, err
		}

		ifItem.Else = elseStatement
	}

	return ifItem, nil
}

func (v *Visitor) VisitElseifStatement(ctx *parser.ElseifStatementContext) (*ast.ASTIf, error) {
	condition, err := v.VisitExpression(ctx.Expression())
	if err != nil {
		return nil, err
	}

	then, err := v.VisitChunk(ctx.Chunk().(*parser.ChunkContext))
	if err != nil {
		return nil, err
	}

	return ast.NewASTIf(condition, then, nil, nil), nil
}

func (v *Visitor) VisitElseStatement(ctx *parser.ElseStatementContext) (*ast.ASTChunk, error) {
	return v.VisitChunk(ctx.Chunk().(*parser.ChunkContext))
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
