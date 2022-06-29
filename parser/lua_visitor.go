// Code generated from Lua.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LuaParser.
type LuaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LuaParser#translation.
	VisitTranslation(ctx *TranslationContext) interface{}

	// Visit a parse tree produced by LuaParser#chunk.
	VisitChunk(ctx *ChunkContext) interface{}

	// Visit a parse tree produced by LuaParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LuaParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpressionBinary.
	VisitExpressionBinary(ctx *ExpressionBinaryContext) interface{}

	// Visit a parse tree produced by LuaParser#ExpressionConstant.
	VisitExpressionConstant(ctx *ExpressionConstantContext) interface{}

	// Visit a parse tree produced by LuaParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}

	// Visit a parse tree produced by LuaParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by LuaParser#elseifStatement.
	VisitElseifStatement(ctx *ElseifStatementContext) interface{}

	// Visit a parse tree produced by LuaParser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by LuaParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by LuaParser#numericForStatement.
	VisitNumericForStatement(ctx *NumericForStatementContext) interface{}
}
