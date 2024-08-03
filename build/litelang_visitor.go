// Code generated from litelang/src/LiteLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package LiteLang // LiteLang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LiteLangParser.
type LiteLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LiteLangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LiteLangParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by LiteLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#functionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#classicForParam.
	VisitClassicForParam(ctx *ClassicForParamContext) interface{}

	// Visit a parse tree produced by LiteLangParser#iteratorForParam.
	VisitIteratorForParam(ctx *IteratorForParamContext) interface{}

	// Visit a parse tree produced by LiteLangParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#unitStatement.
	VisitUnitStatement(ctx *UnitStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#valueStatement.
	VisitValueStatement(ctx *ValueStatementContext) interface{}

	// Visit a parse tree produced by LiteLangParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by LiteLangParser#accessLhs.
	VisitAccessLhs(ctx *AccessLhsContext) interface{}

	// Visit a parse tree produced by LiteLangParser#accessRhs.
	VisitAccessRhs(ctx *AccessRhsContext) interface{}

	// Visit a parse tree produced by LiteLangParser#accessorRhs.
	VisitAccessorRhs(ctx *AccessorRhsContext) interface{}

	// Visit a parse tree produced by LiteLangParser#accessorLhs.
	VisitAccessorLhs(ctx *AccessorLhsContext) interface{}

	// Visit a parse tree produced by LiteLangParser#accessProperty.
	VisitAccessProperty(ctx *AccessPropertyContext) interface{}

	// Visit a parse tree produced by LiteLangParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by LiteLangParser#paramMaker.
	VisitParamMaker(ctx *ParamMakerContext) interface{}

	// Visit a parse tree produced by LiteLangParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by LiteLangParser#arrowFunction.
	VisitArrowFunction(ctx *ArrowFunctionContext) interface{}

	// Visit a parse tree produced by LiteLangParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by LiteLangParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by LiteLangParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by LiteLangParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by LiteLangParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}

	// Visit a parse tree produced by LiteLangParser#spread.
	VisitSpread(ctx *SpreadContext) interface{}

	// Visit a parse tree produced by LiteLangParser#objectItem.
	VisitObjectItem(ctx *ObjectItemContext) interface{}

	// Visit a parse tree produced by LiteLangParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LiteLangParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}
}
