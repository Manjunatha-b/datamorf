// Code generated from src/DataMorf.g4 by ANTLR 4.13.1. DO NOT EDIT.

package build // DataMorf
import "github.com/antlr4-go/antlr/v4"

// DataMorfListener is a complete listener for a parse tree produced by DataMorfParser.
type DataMorfListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterVariableStatement is called when entering the variableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterFunctionStatement is called when entering the functionStatement production.
	EnterFunctionStatement(c *FunctionStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterUnitStatement is called when entering the unitStatement production.
	EnterUnitStatement(c *UnitStatementContext)

	// EnterValueStatement is called when entering the valueStatement production.
	EnterValueStatement(c *ValueStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterClassicForCondition is called when entering the classicForCondition production.
	EnterClassicForCondition(c *ClassicForConditionContext)

	// EnterIteratorForCondition is called when entering the iteratorForCondition production.
	EnterIteratorForCondition(c *IteratorForConditionContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterAccessor is called when entering the accessor production.
	EnterAccessor(c *AccessorContext)

	// EnterAccessProperty is called when entering the accessProperty production.
	EnterAccessProperty(c *AccessPropertyContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterSignatureParams is called when entering the signatureParams production.
	EnterSignatureParams(c *SignatureParamsContext)

	// EnterSendingParams is called when entering the sendingParams production.
	EnterSendingParams(c *SendingParamsContext)

	// EnterArrowFunction is called when entering the arrowFunction production.
	EnterArrowFunction(c *ArrowFunctionContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterKeyValue is called when entering the keyValue production.
	EnterKeyValue(c *KeyValueContext)

	// EnterObjectItem is called when entering the objectItem production.
	EnterObjectItem(c *ObjectItemContext)

	// EnterSpread is called when entering the spread production.
	EnterSpread(c *SpreadContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitVariableStatement is called when exiting the variableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitFunctionStatement is called when exiting the functionStatement production.
	ExitFunctionStatement(c *FunctionStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitUnitStatement is called when exiting the unitStatement production.
	ExitUnitStatement(c *UnitStatementContext)

	// ExitValueStatement is called when exiting the valueStatement production.
	ExitValueStatement(c *ValueStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitClassicForCondition is called when exiting the classicForCondition production.
	ExitClassicForCondition(c *ClassicForConditionContext)

	// ExitIteratorForCondition is called when exiting the iteratorForCondition production.
	ExitIteratorForCondition(c *IteratorForConditionContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitAccessor is called when exiting the accessor production.
	ExitAccessor(c *AccessorContext)

	// ExitAccessProperty is called when exiting the accessProperty production.
	ExitAccessProperty(c *AccessPropertyContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitSignatureParams is called when exiting the signatureParams production.
	ExitSignatureParams(c *SignatureParamsContext)

	// ExitSendingParams is called when exiting the sendingParams production.
	ExitSendingParams(c *SendingParamsContext)

	// ExitArrowFunction is called when exiting the arrowFunction production.
	ExitArrowFunction(c *ArrowFunctionContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitKeyValue is called when exiting the keyValue production.
	ExitKeyValue(c *KeyValueContext)

	// ExitObjectItem is called when exiting the objectItem production.
	ExitObjectItem(c *ObjectItemContext)

	// ExitSpread is called when exiting the spread production.
	ExitSpread(c *SpreadContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)
}
