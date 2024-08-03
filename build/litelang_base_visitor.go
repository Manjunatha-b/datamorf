// Code generated from litelang/src/LiteLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package LiteLang // LiteLang
import "github.com/antlr4-go/antlr/v4"

type BaseLiteLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLiteLangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitFunctionStatement(ctx *FunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitClassicForParam(ctx *ClassicForParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitIteratorForParam(ctx *IteratorForParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitUnitStatement(ctx *UnitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitValueStatement(ctx *ValueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitAccessLhs(ctx *AccessLhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitAccessRhs(ctx *AccessRhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitAccessorRhs(ctx *AccessorRhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitAccessorLhs(ctx *AccessorLhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitAccessProperty(ctx *AccessPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitParamMaker(ctx *ParamMakerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitArrowFunction(ctx *ArrowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitKeyValue(ctx *KeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitSpread(ctx *SpreadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitObjectItem(ctx *ObjectItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteLangVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}
