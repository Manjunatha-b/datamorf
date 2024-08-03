// Generated from /Users/manjunatha/Documents/Research/litefunc-backend/litelang/src/litelang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link litelangParser}.
 */
public interface litelangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link litelangParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(litelangParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(litelangParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#statements}.
	 * @param ctx the parse tree
	 */
	void enterStatements(litelangParser.StatementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#statements}.
	 * @param ctx the parse tree
	 */
	void exitStatements(litelangParser.StatementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(litelangParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(litelangParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#variableStatement}.
	 * @param ctx the parse tree
	 */
	void enterVariableStatement(litelangParser.VariableStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#variableStatement}.
	 * @param ctx the parse tree
	 */
	void exitVariableStatement(litelangParser.VariableStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#assignStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignStatement(litelangParser.AssignStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#assignStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignStatement(litelangParser.AssignStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#functionStatement}.
	 * @param ctx the parse tree
	 */
	void enterFunctionStatement(litelangParser.FunctionStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#functionStatement}.
	 * @param ctx the parse tree
	 */
	void exitFunctionStatement(litelangParser.FunctionStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(litelangParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(litelangParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForStatement(litelangParser.ForStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForStatement(litelangParser.ForStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(litelangParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(litelangParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#unitStatement}.
	 * @param ctx the parse tree
	 */
	void enterUnitStatement(litelangParser.UnitStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#unitStatement}.
	 * @param ctx the parse tree
	 */
	void exitUnitStatement(litelangParser.UnitStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#unit}.
	 * @param ctx the parse tree
	 */
	void enterUnit(litelangParser.UnitContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#unit}.
	 * @param ctx the parse tree
	 */
	void exitUnit(litelangParser.UnitContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#accessLhs}.
	 * @param ctx the parse tree
	 */
	void enterAccessLhs(litelangParser.AccessLhsContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#accessLhs}.
	 * @param ctx the parse tree
	 */
	void exitAccessLhs(litelangParser.AccessLhsContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#accessRhs}.
	 * @param ctx the parse tree
	 */
	void enterAccessRhs(litelangParser.AccessRhsContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#accessRhs}.
	 * @param ctx the parse tree
	 */
	void exitAccessRhs(litelangParser.AccessRhsContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#accessorRhs}.
	 * @param ctx the parse tree
	 */
	void enterAccessorRhs(litelangParser.AccessorRhsContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#accessorRhs}.
	 * @param ctx the parse tree
	 */
	void exitAccessorRhs(litelangParser.AccessorRhsContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#accessorLhs}.
	 * @param ctx the parse tree
	 */
	void enterAccessorLhs(litelangParser.AccessorLhsContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#accessorLhs}.
	 * @param ctx the parse tree
	 */
	void exitAccessorLhs(litelangParser.AccessorLhsContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#accessProperty}.
	 * @param ctx the parse tree
	 */
	void enterAccessProperty(litelangParser.AccessPropertyContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#accessProperty}.
	 * @param ctx the parse tree
	 */
	void exitAccessProperty(litelangParser.AccessPropertyContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(litelangParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(litelangParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#paramMaker}.
	 * @param ctx the parse tree
	 */
	void enterParamMaker(litelangParser.ParamMakerContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#paramMaker}.
	 * @param ctx the parse tree
	 */
	void exitParamMaker(litelangParser.ParamMakerContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#params}.
	 * @param ctx the parse tree
	 */
	void enterParams(litelangParser.ParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#params}.
	 * @param ctx the parse tree
	 */
	void exitParams(litelangParser.ParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#arrowFunction}.
	 * @param ctx the parse tree
	 */
	void enterArrowFunction(litelangParser.ArrowFunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#arrowFunction}.
	 * @param ctx the parse tree
	 */
	void exitArrowFunction(litelangParser.ArrowFunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void enterArrayLiteral(litelangParser.ArrayLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void exitArrayLiteral(litelangParser.ArrayLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#objectLiteral}.
	 * @param ctx the parse tree
	 */
	void enterObjectLiteral(litelangParser.ObjectLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#objectLiteral}.
	 * @param ctx the parse tree
	 */
	void exitObjectLiteral(litelangParser.ObjectLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#key}.
	 * @param ctx the parse tree
	 */
	void enterKey(litelangParser.KeyContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#key}.
	 * @param ctx the parse tree
	 */
	void exitKey(litelangParser.KeyContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(litelangParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(litelangParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(litelangParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(litelangParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link litelangParser#constant}.
	 * @param ctx the parse tree
	 */
	void enterConstant(litelangParser.ConstantContext ctx);
	/**
	 * Exit a parse tree produced by {@link litelangParser#constant}.
	 * @param ctx the parse tree
	 */
	void exitConstant(litelangParser.ConstantContext ctx);
}