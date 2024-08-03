// Code generated from src/DataMorf.g4 by ANTLR 4.13.1. DO NOT EDIT.

package build // DataMorf
import "github.com/antlr4-go/antlr/v4"

// BaseDataMorfListener is a complete listener for a parse tree produced by DataMorfParser.
type BaseDataMorfListener struct{}

var _ DataMorfListener = &BaseDataMorfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDataMorfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDataMorfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDataMorfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDataMorfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseDataMorfListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseDataMorfListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseDataMorfListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseDataMorfListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDataMorfListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDataMorfListener) ExitStatement(ctx *StatementContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseDataMorfListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseDataMorfListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseDataMorfListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseDataMorfListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterFunctionStatement is called when production functionStatement is entered.
func (s *BaseDataMorfListener) EnterFunctionStatement(ctx *FunctionStatementContext) {}

// ExitFunctionStatement is called when production functionStatement is exited.
func (s *BaseDataMorfListener) ExitFunctionStatement(ctx *FunctionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseDataMorfListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseDataMorfListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseDataMorfListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseDataMorfListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseDataMorfListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseDataMorfListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterUnitStatement is called when production unitStatement is entered.
func (s *BaseDataMorfListener) EnterUnitStatement(ctx *UnitStatementContext) {}

// ExitUnitStatement is called when production unitStatement is exited.
func (s *BaseDataMorfListener) ExitUnitStatement(ctx *UnitStatementContext) {}

// EnterValueStatement is called when production valueStatement is entered.
func (s *BaseDataMorfListener) EnterValueStatement(ctx *ValueStatementContext) {}

// ExitValueStatement is called when production valueStatement is exited.
func (s *BaseDataMorfListener) ExitValueStatement(ctx *ValueStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseDataMorfListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseDataMorfListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterClassicForCondition is called when production classicForCondition is entered.
func (s *BaseDataMorfListener) EnterClassicForCondition(ctx *ClassicForConditionContext) {}

// ExitClassicForCondition is called when production classicForCondition is exited.
func (s *BaseDataMorfListener) ExitClassicForCondition(ctx *ClassicForConditionContext) {}

// EnterIteratorForCondition is called when production iteratorForCondition is entered.
func (s *BaseDataMorfListener) EnterIteratorForCondition(ctx *IteratorForConditionContext) {}

// ExitIteratorForCondition is called when production iteratorForCondition is exited.
func (s *BaseDataMorfListener) ExitIteratorForCondition(ctx *IteratorForConditionContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseDataMorfListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseDataMorfListener) ExitUnit(ctx *UnitContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseDataMorfListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseDataMorfListener) ExitReference(ctx *ReferenceContext) {}

// EnterAccessor is called when production accessor is entered.
func (s *BaseDataMorfListener) EnterAccessor(ctx *AccessorContext) {}

// ExitAccessor is called when production accessor is exited.
func (s *BaseDataMorfListener) ExitAccessor(ctx *AccessorContext) {}

// EnterAccessProperty is called when production accessProperty is entered.
func (s *BaseDataMorfListener) EnterAccessProperty(ctx *AccessPropertyContext) {}

// ExitAccessProperty is called when production accessProperty is exited.
func (s *BaseDataMorfListener) ExitAccessProperty(ctx *AccessPropertyContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseDataMorfListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseDataMorfListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterSignatureParams is called when production signatureParams is entered.
func (s *BaseDataMorfListener) EnterSignatureParams(ctx *SignatureParamsContext) {}

// ExitSignatureParams is called when production signatureParams is exited.
func (s *BaseDataMorfListener) ExitSignatureParams(ctx *SignatureParamsContext) {}

// EnterSendingParams is called when production sendingParams is entered.
func (s *BaseDataMorfListener) EnterSendingParams(ctx *SendingParamsContext) {}

// ExitSendingParams is called when production sendingParams is exited.
func (s *BaseDataMorfListener) ExitSendingParams(ctx *SendingParamsContext) {}

// EnterArrowFunction is called when production arrowFunction is entered.
func (s *BaseDataMorfListener) EnterArrowFunction(ctx *ArrowFunctionContext) {}

// ExitArrowFunction is called when production arrowFunction is exited.
func (s *BaseDataMorfListener) ExitArrowFunction(ctx *ArrowFunctionContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseDataMorfListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseDataMorfListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseDataMorfListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseDataMorfListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterKey is called when production key is entered.
func (s *BaseDataMorfListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseDataMorfListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseDataMorfListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseDataMorfListener) ExitValue(ctx *ValueContext) {}

// EnterKeyValue is called when production keyValue is entered.
func (s *BaseDataMorfListener) EnterKeyValue(ctx *KeyValueContext) {}

// ExitKeyValue is called when production keyValue is exited.
func (s *BaseDataMorfListener) ExitKeyValue(ctx *KeyValueContext) {}

// EnterObjectItem is called when production objectItem is entered.
func (s *BaseDataMorfListener) EnterObjectItem(ctx *ObjectItemContext) {}

// ExitObjectItem is called when production objectItem is exited.
func (s *BaseDataMorfListener) ExitObjectItem(ctx *ObjectItemContext) {}

// EnterSpread is called when production spread is entered.
func (s *BaseDataMorfListener) EnterSpread(ctx *SpreadContext) {}

// ExitSpread is called when production spread is exited.
func (s *BaseDataMorfListener) ExitSpread(ctx *SpreadContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseDataMorfListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseDataMorfListener) ExitBlock(ctx *BlockContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseDataMorfListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseDataMorfListener) ExitConstant(ctx *ConstantContext) {}
