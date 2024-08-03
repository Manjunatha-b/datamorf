// Code generated from litelang/src/LiteLang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package LiteLang // LiteLang
import "github.com/antlr4-go/antlr/v4"

// BaseLiteLangListener is a complete listener for a parse tree produced by LiteLangParser.
type BaseLiteLangListener struct{}

var _ LiteLangListener = &BaseLiteLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLiteLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLiteLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLiteLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLiteLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLiteLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLiteLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseLiteLangListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseLiteLangListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLiteLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLiteLangListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseLiteLangListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseLiteLangListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseLiteLangListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseLiteLangListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterFunctionStatement is called when production functionStatement is entered.
func (s *BaseLiteLangListener) EnterFunctionStatement(ctx *FunctionStatementContext) {}

// ExitFunctionStatement is called when production functionStatement is exited.
func (s *BaseLiteLangListener) ExitFunctionStatement(ctx *FunctionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseLiteLangListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseLiteLangListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseLiteLangListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseLiteLangListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterClassicForParam is called when production classicForParam is entered.
func (s *BaseLiteLangListener) EnterClassicForParam(ctx *ClassicForParamContext) {}

// ExitClassicForParam is called when production classicForParam is exited.
func (s *BaseLiteLangListener) ExitClassicForParam(ctx *ClassicForParamContext) {}

// EnterIteratorForParam is called when production iteratorForParam is entered.
func (s *BaseLiteLangListener) EnterIteratorForParam(ctx *IteratorForParamContext) {}

// ExitIteratorForParam is called when production iteratorForParam is exited.
func (s *BaseLiteLangListener) ExitIteratorForParam(ctx *IteratorForParamContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseLiteLangListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseLiteLangListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseLiteLangListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseLiteLangListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterUnitStatement is called when production unitStatement is entered.
func (s *BaseLiteLangListener) EnterUnitStatement(ctx *UnitStatementContext) {}

// ExitUnitStatement is called when production unitStatement is exited.
func (s *BaseLiteLangListener) ExitUnitStatement(ctx *UnitStatementContext) {}

// EnterValueStatement is called when production valueStatement is entered.
func (s *BaseLiteLangListener) EnterValueStatement(ctx *ValueStatementContext) {}

// ExitValueStatement is called when production valueStatement is exited.
func (s *BaseLiteLangListener) ExitValueStatement(ctx *ValueStatementContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseLiteLangListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseLiteLangListener) ExitUnit(ctx *UnitContext) {}

// EnterAccessLhs is called when production accessLhs is entered.
func (s *BaseLiteLangListener) EnterAccessLhs(ctx *AccessLhsContext) {}

// ExitAccessLhs is called when production accessLhs is exited.
func (s *BaseLiteLangListener) ExitAccessLhs(ctx *AccessLhsContext) {}

// EnterAccessRhs is called when production accessRhs is entered.
func (s *BaseLiteLangListener) EnterAccessRhs(ctx *AccessRhsContext) {}

// ExitAccessRhs is called when production accessRhs is exited.
func (s *BaseLiteLangListener) ExitAccessRhs(ctx *AccessRhsContext) {}

// EnterAccessorRhs is called when production accessorRhs is entered.
func (s *BaseLiteLangListener) EnterAccessorRhs(ctx *AccessorRhsContext) {}

// ExitAccessorRhs is called when production accessorRhs is exited.
func (s *BaseLiteLangListener) ExitAccessorRhs(ctx *AccessorRhsContext) {}

// EnterAccessorLhs is called when production accessorLhs is entered.
func (s *BaseLiteLangListener) EnterAccessorLhs(ctx *AccessorLhsContext) {}

// ExitAccessorLhs is called when production accessorLhs is exited.
func (s *BaseLiteLangListener) ExitAccessorLhs(ctx *AccessorLhsContext) {}

// EnterAccessProperty is called when production accessProperty is entered.
func (s *BaseLiteLangListener) EnterAccessProperty(ctx *AccessPropertyContext) {}

// ExitAccessProperty is called when production accessProperty is exited.
func (s *BaseLiteLangListener) ExitAccessProperty(ctx *AccessPropertyContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseLiteLangListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseLiteLangListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterParamMaker is called when production paramMaker is entered.
func (s *BaseLiteLangListener) EnterParamMaker(ctx *ParamMakerContext) {}

// ExitParamMaker is called when production paramMaker is exited.
func (s *BaseLiteLangListener) ExitParamMaker(ctx *ParamMakerContext) {}

// EnterParams is called when production params is entered.
func (s *BaseLiteLangListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseLiteLangListener) ExitParams(ctx *ParamsContext) {}

// EnterArrowFunction is called when production arrowFunction is entered.
func (s *BaseLiteLangListener) EnterArrowFunction(ctx *ArrowFunctionContext) {}

// ExitArrowFunction is called when production arrowFunction is exited.
func (s *BaseLiteLangListener) ExitArrowFunction(ctx *ArrowFunctionContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseLiteLangListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseLiteLangListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseLiteLangListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseLiteLangListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterKey is called when production key is entered.
func (s *BaseLiteLangListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseLiteLangListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseLiteLangListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseLiteLangListener) ExitValue(ctx *ValueContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseLiteLangListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseLiteLangListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterSpread is called when production spread is entered.
func (s *BaseLiteLangListener) EnterSpread(ctx *SpreadContext) {}

// ExitSpread is called when production spread is exited.
func (s *BaseLiteLangListener) ExitSpread(ctx *SpreadContext) {}

// EnterObjectItem is called when production objectItem is entered.
func (s *BaseLiteLangListener) EnterObjectItem(ctx *ObjectItemContext) {}

// ExitObjectItem is called when production objectItem is exited.
func (s *BaseLiteLangListener) ExitObjectItem(ctx *ObjectItemContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLiteLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLiteLangListener) ExitBlock(ctx *BlockContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseLiteLangListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseLiteLangListener) ExitConstant(ctx *ConstantContext) {}
