// Code generated from Lua.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LuaParser.
type LuaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LuaParser#translation.
	VisitTranslation(ctx *TranslationContext) interface{}

	// Visit a parse tree produced by LuaParser#topLevelStatement.
	VisitTopLevelStatement(ctx *TopLevelStatementContext) interface{}

	// Visit a parse tree produced by LuaParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpressionBinary.
	VisitExpressionBinary(ctx *ExpressionBinaryContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpressionConstant.
	VisitExpressionConstant(ctx *ExpressionConstantContext) interface{}

	// Visit a parse tree produced by LuaParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}
}
