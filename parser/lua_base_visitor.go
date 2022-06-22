// Code generated from Lua.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLuaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLuaVisitor) VisitTranslation(ctx *TranslationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitTopLevelStatement(ctx *TopLevelStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpressionBinary(ctx *ExpressionBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitExpressionConstant(ctx *ExpressionConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLuaVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
