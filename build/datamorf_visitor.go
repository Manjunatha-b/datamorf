// Code generated from src/DataMorf.g4 by ANTLR 4.13.1. DO NOT EDIT.

package build // DataMorf
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DataMorfParser.
type DataMorfVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DataMorfParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by DataMorfParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by DataMorfParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#variableStatement.
	VisitVariableStatement(ctx *VariableStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#functionStatement.
	VisitFunctionStatement(ctx *FunctionStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#unitStatement.
	VisitUnitStatement(ctx *UnitStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#valueStatement.
	VisitValueStatement(ctx *ValueStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by DataMorfParser#classicForCondition.
	VisitClassicForCondition(ctx *ClassicForConditionContext) interface{}

	// Visit a parse tree produced by DataMorfParser#iteratorForCondition.
	VisitIteratorForCondition(ctx *IteratorForConditionContext) interface{}

	// Visit a parse tree produced by DataMorfParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by DataMorfParser#reference.
	VisitReference(ctx *ReferenceContext) interface{}

	// Visit a parse tree produced by DataMorfParser#accessor.
	VisitAccessor(ctx *AccessorContext) interface{}

	// Visit a parse tree produced by DataMorfParser#accessProperty.
	VisitAccessProperty(ctx *AccessPropertyContext) interface{}

	// Visit a parse tree produced by DataMorfParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by DataMorfParser#signatureParams.
	VisitSignatureParams(ctx *SignatureParamsContext) interface{}

	// Visit a parse tree produced by DataMorfParser#sendingParams.
	VisitSendingParams(ctx *SendingParamsContext) interface{}

	// Visit a parse tree produced by DataMorfParser#arrowFunction.
	VisitArrowFunction(ctx *ArrowFunctionContext) interface{}

	// Visit a parse tree produced by DataMorfParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by DataMorfParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by DataMorfParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by DataMorfParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by DataMorfParser#keyValue.
	VisitKeyValue(ctx *KeyValueContext) interface{}

	// Visit a parse tree produced by DataMorfParser#objectItem.
	VisitObjectItem(ctx *ObjectItemContext) interface{}

	// Visit a parse tree produced by DataMorfParser#spread.
	VisitSpread(ctx *SpreadContext) interface{}

	// Visit a parse tree produced by DataMorfParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by DataMorfParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}
}
