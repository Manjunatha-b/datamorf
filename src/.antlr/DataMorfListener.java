// Generated from c://Users//chillTime//Documents//datamorf//src//DataMorf.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link DataMorfParser}.
 */
public interface DataMorfListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(DataMorfParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(DataMorfParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#statements}.
	 * @param ctx the parse tree
	 */
	void enterStatements(DataMorfParser.StatementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#statements}.
	 * @param ctx the parse tree
	 */
	void exitStatements(DataMorfParser.StatementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(DataMorfParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(DataMorfParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#variableStatement}.
	 * @param ctx the parse tree
	 */
	void enterVariableStatement(DataMorfParser.VariableStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#variableStatement}.
	 * @param ctx the parse tree
	 */
	void exitVariableStatement(DataMorfParser.VariableStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#assignStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignStatement(DataMorfParser.AssignStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#assignStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignStatement(DataMorfParser.AssignStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#functionStatement}.
	 * @param ctx the parse tree
	 */
	void enterFunctionStatement(DataMorfParser.FunctionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#functionStatement}.
	 * @param ctx the parse tree
	 */
	void exitFunctionStatement(DataMorfParser.FunctionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(DataMorfParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(DataMorfParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(DataMorfParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(DataMorfParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#deleteStatement}.
	 * @param ctx the parse tree
	 */
	void enterDeleteStatement(DataMorfParser.DeleteStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#deleteStatement}.
	 * @param ctx the parse tree
	 */
	void exitDeleteStatement(DataMorfParser.DeleteStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#unitStatement}.
	 * @param ctx the parse tree
	 */
	void enterUnitStatement(DataMorfParser.UnitStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#unitStatement}.
	 * @param ctx the parse tree
	 */
	void exitUnitStatement(DataMorfParser.UnitStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#valueStatement}.
	 * @param ctx the parse tree
	 */
	void enterValueStatement(DataMorfParser.ValueStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#valueStatement}.
	 * @param ctx the parse tree
	 */
	void exitValueStatement(DataMorfParser.ValueStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForStatement(DataMorfParser.ForStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForStatement(DataMorfParser.ForStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#classicForCondition}.
	 * @param ctx the parse tree
	 */
	void enterClassicForCondition(DataMorfParser.ClassicForConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#classicForCondition}.
	 * @param ctx the parse tree
	 */
	void exitClassicForCondition(DataMorfParser.ClassicForConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#iteratorForCondition}.
	 * @param ctx the parse tree
	 */
	void enterIteratorForCondition(DataMorfParser.IteratorForConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#iteratorForCondition}.
	 * @param ctx the parse tree
	 */
	void exitIteratorForCondition(DataMorfParser.IteratorForConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#unit}.
	 * @param ctx the parse tree
	 */
	void enterUnit(DataMorfParser.UnitContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#unit}.
	 * @param ctx the parse tree
	 */
	void exitUnit(DataMorfParser.UnitContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#reference}.
	 * @param ctx the parse tree
	 */
	void enterReference(DataMorfParser.ReferenceContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#reference}.
	 * @param ctx the parse tree
	 */
	void exitReference(DataMorfParser.ReferenceContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#accessor}.
	 * @param ctx the parse tree
	 */
	void enterAccessor(DataMorfParser.AccessorContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#accessor}.
	 * @param ctx the parse tree
	 */
	void exitAccessor(DataMorfParser.AccessorContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#accessProperty}.
	 * @param ctx the parse tree
	 */
	void enterAccessProperty(DataMorfParser.AccessPropertyContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#accessProperty}.
	 * @param ctx the parse tree
	 */
	void exitAccessProperty(DataMorfParser.AccessPropertyContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(DataMorfParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(DataMorfParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#signatureParams}.
	 * @param ctx the parse tree
	 */
	void enterSignatureParams(DataMorfParser.SignatureParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#signatureParams}.
	 * @param ctx the parse tree
	 */
	void exitSignatureParams(DataMorfParser.SignatureParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#sendingParams}.
	 * @param ctx the parse tree
	 */
	void enterSendingParams(DataMorfParser.SendingParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#sendingParams}.
	 * @param ctx the parse tree
	 */
	void exitSendingParams(DataMorfParser.SendingParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#arrowFunction}.
	 * @param ctx the parse tree
	 */
	void enterArrowFunction(DataMorfParser.ArrowFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#arrowFunction}.
	 * @param ctx the parse tree
	 */
	void exitArrowFunction(DataMorfParser.ArrowFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void enterArrayLiteral(DataMorfParser.ArrayLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void exitArrayLiteral(DataMorfParser.ArrayLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#objectLiteral}.
	 * @param ctx the parse tree
	 */
	void enterObjectLiteral(DataMorfParser.ObjectLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#objectLiteral}.
	 * @param ctx the parse tree
	 */
	void exitObjectLiteral(DataMorfParser.ObjectLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#key}.
	 * @param ctx the parse tree
	 */
	void enterKey(DataMorfParser.KeyContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#key}.
	 * @param ctx the parse tree
	 */
	void exitKey(DataMorfParser.KeyContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(DataMorfParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(DataMorfParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#keyValue}.
	 * @param ctx the parse tree
	 */
	void enterKeyValue(DataMorfParser.KeyValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#keyValue}.
	 * @param ctx the parse tree
	 */
	void exitKeyValue(DataMorfParser.KeyValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#objectItem}.
	 * @param ctx the parse tree
	 */
	void enterObjectItem(DataMorfParser.ObjectItemContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#objectItem}.
	 * @param ctx the parse tree
	 */
	void exitObjectItem(DataMorfParser.ObjectItemContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#spread}.
	 * @param ctx the parse tree
	 */
	void enterSpread(DataMorfParser.SpreadContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#spread}.
	 * @param ctx the parse tree
	 */
	void exitSpread(DataMorfParser.SpreadContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(DataMorfParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(DataMorfParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link DataMorfParser#constant}.
	 * @param ctx the parse tree
	 */
	void enterConstant(DataMorfParser.ConstantContext ctx);
	/**
	 * Exit a parse tree produced by {@link DataMorfParser#constant}.
	 * @param ctx the parse tree
	 */
	void exitConstant(DataMorfParser.ConstantContext ctx);
}