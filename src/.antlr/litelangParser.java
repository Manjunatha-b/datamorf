// Generated from /Users/manjunatha/Documents/Research/litefunc-backend/litelang/src/litelang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class litelangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, OpenBracket=2, CloseBracket=3, OpenParen=4, CloseParen=5, OpenBrace=6, 
		CloseBrace=7, SemiColon=8, Comma=9, Assign=10, QuestionMark=11, QuestionMarkDot=12, 
		Colon=13, Dot=14, PlusPlus=15, MinusMinus=16, Plus=17, Minus=18, Not=19, 
		Multiply=20, Divide=21, Modulus=22, Power=23, LessThan=24, MoreThan=25, 
		LessThanEquals=26, GreaterThanEquals=27, Equals=28, NotEquals=29, And=30, 
		Or=31, NullLiteral=32, BooleanLiteral=33, DecimalLiteral=34, Break=35, 
		Do=36, Case=37, Else=38, New=39, Var=40, Return=41, Continue=42, For=43, 
		Switch=44, While=45, Function=46, Default=47, If=48, Elif=49, Delete=50, 
		Identifier=51, StringLiteral=52, WhiteSpaces=53, LineTerminator=54;
	public static final int
		RULE_program = 0, RULE_statements = 1, RULE_statement = 2, RULE_variableStatement = 3, 
		RULE_assignStatement = 4, RULE_functionStatement = 5, RULE_ifStatement = 6, 
		RULE_forStatement = 7, RULE_switchStatement = 8, RULE_unitStatement = 9, 
		RULE_unit = 10, RULE_accessLhs = 11, RULE_accessRhs = 12, RULE_accessorRhs = 13, 
		RULE_accessorLhs = 14, RULE_accessProperty = 15, RULE_functionCall = 16, 
		RULE_paramMaker = 17, RULE_params = 18, RULE_arrowFunction = 19, RULE_arrayLiteral = 20, 
		RULE_objectLiteral = 21, RULE_key = 22, RULE_value = 23, RULE_block = 24, 
		RULE_constant = 25;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statements", "statement", "variableStatement", "assignStatement", 
			"functionStatement", "ifStatement", "forStatement", "switchStatement", 
			"unitStatement", "unit", "accessLhs", "accessRhs", "accessorRhs", "accessorLhs", 
			"accessProperty", "functionCall", "paramMaker", "params", "arrowFunction", 
			"arrayLiteral", "objectLiteral", "key", "value", "block", "constant"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'=>'", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','", 
			"'='", "'?'", "'?.'", "':'", "'.'", "'++'", "'--'", "'+'", "'-'", "'!'", 
			"'*'", "'/'", "'%'", "'**'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", 
			"'&&'", "'||'", "'null'", null, null, "'break'", "'do'", "'case'", "'else'", 
			"'new'", "'var'", "'return'", "'continue'", "'for'", "'switch'", "'while'", 
			"'function'", "'default'", "'if'", "'elif'", "'delete'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "OpenBracket", "CloseBracket", "OpenParen", "CloseParen", 
			"OpenBrace", "CloseBrace", "SemiColon", "Comma", "Assign", "QuestionMark", 
			"QuestionMarkDot", "Colon", "Dot", "PlusPlus", "MinusMinus", "Plus", 
			"Minus", "Not", "Multiply", "Divide", "Modulus", "Power", "LessThan", 
			"MoreThan", "LessThanEquals", "GreaterThanEquals", "Equals", "NotEquals", 
			"And", "Or", "NullLiteral", "BooleanLiteral", "DecimalLiteral", "Break", 
			"Do", "Case", "Else", "New", "Var", "Return", "Continue", "For", "Switch", 
			"While", "Function", "Default", "If", "Elif", "Delete", "Identifier", 
			"StringLiteral", "WhiteSpaces", "LineTerminator"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "litelang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public litelangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(litelangParser.EOF, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 7134761017933840L) != 0)) {
				{
				setState(52);
				statements();
				}
			}

			setState(55);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	}

	public final StatementsContext statements() throws RecognitionException {
		StatementsContext _localctx = new StatementsContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(57);
				statement();
				}
				}
				setState(60); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 7134761017933840L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public VariableStatementContext variableStatement() {
			return getRuleContext(VariableStatementContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(litelangParser.SemiColon, 0); }
		public AssignStatementContext assignStatement() {
			return getRuleContext(AssignStatementContext.class,0);
		}
		public FunctionStatementContext functionStatement() {
			return getRuleContext(FunctionStatementContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public ForStatementContext forStatement() {
			return getRuleContext(ForStatementContext.class,0);
		}
		public SwitchStatementContext switchStatement() {
			return getRuleContext(SwitchStatementContext.class,0);
		}
		public UnitStatementContext unitStatement() {
			return getRuleContext(UnitStatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_statement);
		try {
			setState(75);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				variableStatement();
				setState(63);
				match(SemiColon);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(65);
				assignStatement();
				setState(66);
				match(SemiColon);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(68);
				functionStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(69);
				ifStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(70);
				forStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(71);
				switchStatement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(72);
				unitStatement();
				setState(73);
				match(SemiColon);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableStatementContext extends ParserRuleContext {
		public TerminalNode Var() { return getToken(litelangParser.Var, 0); }
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public TerminalNode Assign() { return getToken(litelangParser.Assign, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public VariableStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableStatement; }
	}

	public final VariableStatementContext variableStatement() throws RecognitionException {
		VariableStatementContext _localctx = new VariableStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_variableStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			match(Var);
			setState(78);
			match(Identifier);
			setState(81);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Assign) {
				{
				setState(79);
				match(Assign);
				setState(80);
				value();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignStatementContext extends ParserRuleContext {
		public AccessLhsContext accessLhs() {
			return getRuleContext(AccessLhsContext.class,0);
		}
		public TerminalNode Assign() { return getToken(litelangParser.Assign, 0); }
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public AssignStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignStatement; }
	}

	public final AssignStatementContext assignStatement() throws RecognitionException {
		AssignStatementContext _localctx = new AssignStatementContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assignStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(83);
			accessLhs();
			setState(84);
			match(Assign);
			setState(85);
			unit(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionStatementContext extends ParserRuleContext {
		public TerminalNode Function() { return getToken(litelangParser.Function, 0); }
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public TerminalNode OpenParen() { return getToken(litelangParser.OpenParen, 0); }
		public ParamMakerContext paramMaker() {
			return getRuleContext(ParamMakerContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(litelangParser.CloseParen, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FunctionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionStatement; }
	}

	public final FunctionStatementContext functionStatement() throws RecognitionException {
		FunctionStatementContext _localctx = new FunctionStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_functionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(87);
			match(Function);
			setState(88);
			match(Identifier);
			setState(89);
			match(OpenParen);
			setState(90);
			paramMaker();
			setState(91);
			match(CloseParen);
			setState(92);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode If() { return getToken(litelangParser.If, 0); }
		public List<TerminalNode> OpenParen() { return getTokens(litelangParser.OpenParen); }
		public TerminalNode OpenParen(int i) {
			return getToken(litelangParser.OpenParen, i);
		}
		public List<UnitContext> unit() {
			return getRuleContexts(UnitContext.class);
		}
		public UnitContext unit(int i) {
			return getRuleContext(UnitContext.class,i);
		}
		public List<TerminalNode> CloseParen() { return getTokens(litelangParser.CloseParen); }
		public TerminalNode CloseParen(int i) {
			return getToken(litelangParser.CloseParen, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> Elif() { return getTokens(litelangParser.Elif); }
		public TerminalNode Elif(int i) {
			return getToken(litelangParser.Elif, i);
		}
		public TerminalNode Else() { return getToken(litelangParser.Else, 0); }
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			match(If);
			setState(95);
			match(OpenParen);
			setState(96);
			unit(0);
			setState(97);
			match(CloseParen);
			setState(98);
			block();
			setState(107);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Elif) {
				{
				{
				setState(99);
				match(Elif);
				setState(100);
				match(OpenParen);
				setState(101);
				unit(0);
				setState(102);
				match(CloseParen);
				setState(103);
				block();
				}
				}
				setState(109);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Else) {
				{
				setState(110);
				match(Else);
				setState(111);
				block();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForStatementContext extends ParserRuleContext {
		public UnitStatementContext condition;
		public AssignStatementContext postOp;
		public TerminalNode For() { return getToken(litelangParser.For, 0); }
		public TerminalNode OpenParen() { return getToken(litelangParser.OpenParen, 0); }
		public List<TerminalNode> SemiColon() { return getTokens(litelangParser.SemiColon); }
		public TerminalNode SemiColon(int i) {
			return getToken(litelangParser.SemiColon, i);
		}
		public TerminalNode CloseParen() { return getToken(litelangParser.CloseParen, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public UnitStatementContext unitStatement() {
			return getRuleContext(UnitStatementContext.class,0);
		}
		public List<AssignStatementContext> assignStatement() {
			return getRuleContexts(AssignStatementContext.class);
		}
		public AssignStatementContext assignStatement(int i) {
			return getRuleContext(AssignStatementContext.class,i);
		}
		public VariableStatementContext variableStatement() {
			return getRuleContext(VariableStatementContext.class,0);
		}
		public ForStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStatement; }
	}

	public final ForStatementContext forStatement() throws RecognitionException {
		ForStatementContext _localctx = new ForStatementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_forStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(For);
			setState(115);
			match(OpenParen);
			setState(118);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Var:
				{
				setState(116);
				variableStatement();
				}
				break;
			case Identifier:
				{
				setState(117);
				assignStatement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(120);
			match(SemiColon);
			setState(121);
			((ForStatementContext)_localctx).condition = unitStatement();
			setState(122);
			match(SemiColon);
			setState(123);
			((ForStatementContext)_localctx).postOp = assignStatement();
			setState(124);
			match(CloseParen);
			setState(125);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStatementContext extends ParserRuleContext {
		public TerminalNode Switch() { return getToken(litelangParser.Switch, 0); }
		public TerminalNode OpenParen() { return getToken(litelangParser.OpenParen, 0); }
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(litelangParser.CloseParen, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public SwitchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStatement; }
	}

	public final SwitchStatementContext switchStatement() throws RecognitionException {
		SwitchStatementContext _localctx = new SwitchStatementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_switchStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			match(Switch);
			setState(128);
			match(OpenParen);
			setState(129);
			unit(0);
			setState(130);
			match(CloseParen);
			setState(131);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnitStatementContext extends ParserRuleContext {
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public UnitStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unitStatement; }
	}

	public final UnitStatementContext unitStatement() throws RecognitionException {
		UnitStatementContext _localctx = new UnitStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_unitStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			unit(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnitContext extends ParserRuleContext {
		public UnitContext left;
		public UnitContext single;
		public UnitContext right;
		public TerminalNode OpenParen() { return getToken(litelangParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(litelangParser.CloseParen, 0); }
		public List<UnitContext> unit() {
			return getRuleContexts(UnitContext.class);
		}
		public UnitContext unit(int i) {
			return getRuleContext(UnitContext.class,i);
		}
		public TerminalNode Not() { return getToken(litelangParser.Not, 0); }
		public AccessRhsContext accessRhs() {
			return getRuleContext(AccessRhsContext.class,0);
		}
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public TerminalNode Divide() { return getToken(litelangParser.Divide, 0); }
		public TerminalNode Multiply() { return getToken(litelangParser.Multiply, 0); }
		public TerminalNode Plus() { return getToken(litelangParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(litelangParser.Minus, 0); }
		public TerminalNode LessThan() { return getToken(litelangParser.LessThan, 0); }
		public TerminalNode MoreThan() { return getToken(litelangParser.MoreThan, 0); }
		public TerminalNode LessThanEquals() { return getToken(litelangParser.LessThanEquals, 0); }
		public TerminalNode GreaterThanEquals() { return getToken(litelangParser.GreaterThanEquals, 0); }
		public TerminalNode Equals() { return getToken(litelangParser.Equals, 0); }
		public TerminalNode NotEquals() { return getToken(litelangParser.NotEquals, 0); }
		public TerminalNode And() { return getToken(litelangParser.And, 0); }
		public TerminalNode Or() { return getToken(litelangParser.Or, 0); }
		public UnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unit; }
	}

	public final UnitContext unit() throws RecognitionException {
		return unit(0);
	}

	private UnitContext unit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		UnitContext _localctx = new UnitContext(_ctx, _parentState);
		UnitContext _prevctx = _localctx;
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_unit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenParen:
				{
				setState(136);
				match(OpenParen);
				setState(137);
				((UnitContext)_localctx).single = unit(0);
				setState(138);
				match(CloseParen);
				}
				break;
			case Not:
				{
				setState(140);
				match(Not);
				setState(141);
				((UnitContext)_localctx).single = unit(3);
				}
				break;
			case Identifier:
				{
				setState(142);
				accessRhs();
				}
				break;
			case NullLiteral:
			case BooleanLiteral:
			case DecimalLiteral:
			case StringLiteral:
				{
				setState(143);
				constant();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(163);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(161);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(146);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						{
						setState(147);
						match(Divide);
						}
						setState(148);
						((UnitContext)_localctx).right = unit(9);
						}
						break;
					case 2:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(149);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						{
						setState(150);
						match(Multiply);
						}
						setState(151);
						((UnitContext)_localctx).right = unit(8);
						}
						break;
					case 3:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(152);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						{
						setState(153);
						match(Plus);
						}
						setState(154);
						((UnitContext)_localctx).right = unit(7);
						}
						break;
					case 4:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(155);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						{
						setState(156);
						match(Minus);
						}
						setState(157);
						((UnitContext)_localctx).right = unit(6);
						}
						break;
					case 5:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(158);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(159);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4278190080L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(160);
						((UnitContext)_localctx).right = unit(5);
						}
						break;
					}
					} 
				}
				setState(165);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AccessLhsContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public List<AccessorLhsContext> accessorLhs() {
			return getRuleContexts(AccessorLhsContext.class);
		}
		public AccessorLhsContext accessorLhs(int i) {
			return getRuleContext(AccessorLhsContext.class,i);
		}
		public AccessLhsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessLhs; }
	}

	public final AccessLhsContext accessLhs() throws RecognitionException {
		AccessLhsContext _localctx = new AccessLhsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_accessLhs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(Identifier);
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OpenBracket || _la==Dot) {
				{
				{
				setState(167);
				accessorLhs();
				}
				}
				setState(172);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AccessRhsContext extends ParserRuleContext {
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public List<AccessorRhsContext> accessorRhs() {
			return getRuleContexts(AccessorRhsContext.class);
		}
		public AccessorRhsContext accessorRhs(int i) {
			return getRuleContext(AccessorRhsContext.class,i);
		}
		public AccessRhsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessRhs; }
	}

	public final AccessRhsContext accessRhs() throws RecognitionException {
		AccessRhsContext _localctx = new AccessRhsContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_accessRhs);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(173);
				functionCall();
				}
				break;
			case 2:
				{
				setState(174);
				match(Identifier);
				}
				break;
			}
			setState(180);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(177);
					accessorRhs();
					}
					} 
				}
				setState(182);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AccessorRhsContext extends ParserRuleContext {
		public AccessPropertyContext accessProperty() {
			return getRuleContext(AccessPropertyContext.class,0);
		}
		public TerminalNode Dot() { return getToken(litelangParser.Dot, 0); }
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public AccessorRhsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessorRhs; }
	}

	public final AccessorRhsContext accessorRhs() throws RecognitionException {
		AccessorRhsContext _localctx = new AccessorRhsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_accessorRhs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(183);
				accessProperty();
				}
				break;
			case 2:
				{
				{
				setState(184);
				match(Dot);
				setState(185);
				match(Identifier);
				}
				}
				break;
			case 3:
				{
				{
				setState(186);
				match(Dot);
				setState(187);
				functionCall();
				}
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AccessorLhsContext extends ParserRuleContext {
		public AccessPropertyContext accessProperty() {
			return getRuleContext(AccessPropertyContext.class,0);
		}
		public TerminalNode Dot() { return getToken(litelangParser.Dot, 0); }
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public AccessorLhsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessorLhs; }
	}

	public final AccessorLhsContext accessorLhs() throws RecognitionException {
		AccessorLhsContext _localctx = new AccessorLhsContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_accessorLhs);
		try {
			setState(193);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenBracket:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				accessProperty();
				}
				break;
			case Dot:
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(191);
				match(Dot);
				setState(192);
				match(Identifier);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AccessPropertyContext extends ParserRuleContext {
		public TerminalNode OpenBracket() { return getToken(litelangParser.OpenBracket, 0); }
		public TerminalNode CloseBracket() { return getToken(litelangParser.CloseBracket, 0); }
		public TerminalNode StringLiteral() { return getToken(litelangParser.StringLiteral, 0); }
		public TerminalNode DecimalLiteral() { return getToken(litelangParser.DecimalLiteral, 0); }
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public AccessPropertyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessProperty; }
	}

	public final AccessPropertyContext accessProperty() throws RecognitionException {
		AccessPropertyContext _localctx = new AccessPropertyContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_accessProperty);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			match(OpenBracket);
			setState(196);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 6755416620924928L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(197);
			match(CloseBracket);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(litelangParser.Identifier, 0); }
		public TerminalNode OpenParen() { return getToken(litelangParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(litelangParser.CloseParen, 0); }
		public ParamsContext params() {
			return getRuleContext(ParamsContext.class,0);
		}
		public ArrowFunctionContext arrowFunction() {
			return getRuleContext(ArrowFunctionContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_functionCall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(199);
			match(Identifier);
			setState(200);
			match(OpenParen);
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(201);
				params();
				}
				break;
			case 2:
				{
				setState(202);
				arrowFunction();
				}
				break;
			}
			setState(205);
			match(CloseParen);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamMakerContext extends ParserRuleContext {
		public List<TerminalNode> Identifier() { return getTokens(litelangParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(litelangParser.Identifier, i);
		}
		public List<TerminalNode> Comma() { return getTokens(litelangParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(litelangParser.Comma, i);
		}
		public ParamMakerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramMaker; }
	}

	public final ParamMakerContext paramMaker() throws RecognitionException {
		ParamMakerContext _localctx = new ParamMakerContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_paramMaker);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(207);
				match(Identifier);
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(208);
					match(Comma);
					setState(209);
					match(Identifier);
					}
					}
					setState(214);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamsContext extends ParserRuleContext {
		public List<UnitContext> unit() {
			return getRuleContexts(UnitContext.class);
		}
		public UnitContext unit(int i) {
			return getRuleContext(UnitContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(litelangParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(litelangParser.Comma, i);
		}
		public ParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_params; }
	}

	public final ParamsContext params() throws RecognitionException {
		ParamsContext _localctx = new ParamsContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6755429506351120L) != 0)) {
				{
				setState(217);
				unit(0);
				setState(222);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(218);
					match(Comma);
					setState(219);
					unit(0);
					}
					}
					setState(224);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrowFunctionContext extends ParserRuleContext {
		public TerminalNode OpenParen() { return getToken(litelangParser.OpenParen, 0); }
		public ParamsContext params() {
			return getRuleContext(ParamsContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(litelangParser.CloseParen, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ArrowFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrowFunction; }
	}

	public final ArrowFunctionContext arrowFunction() throws RecognitionException {
		ArrowFunctionContext _localctx = new ArrowFunctionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_arrowFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(227);
			match(OpenParen);
			setState(228);
			params();
			setState(229);
			match(CloseParen);
			setState(230);
			match(T__0);
			setState(231);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayLiteralContext extends ParserRuleContext {
		public TerminalNode OpenBracket() { return getToken(litelangParser.OpenBracket, 0); }
		public TerminalNode CloseBracket() { return getToken(litelangParser.CloseBracket, 0); }
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(litelangParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(litelangParser.Comma, i);
		}
		public ArrayLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteral; }
	}

	public final ArrayLiteralContext arrayLiteral() throws RecognitionException {
		ArrayLiteralContext _localctx = new ArrayLiteralContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
			match(OpenBracket);
			setState(242);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6755429506351188L) != 0)) {
				{
				setState(234);
				value();
				setState(239);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(235);
					match(Comma);
					setState(236);
					value();
					}
					}
					setState(241);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(244);
			match(CloseBracket);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ObjectLiteralContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(litelangParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(litelangParser.CloseBrace, 0); }
		public List<KeyContext> key() {
			return getRuleContexts(KeyContext.class);
		}
		public KeyContext key(int i) {
			return getRuleContext(KeyContext.class,i);
		}
		public List<TerminalNode> Colon() { return getTokens(litelangParser.Colon); }
		public TerminalNode Colon(int i) {
			return getToken(litelangParser.Colon, i);
		}
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(litelangParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(litelangParser.Comma, i);
		}
		public ObjectLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectLiteral; }
	}

	public final ObjectLiteralContext objectLiteral() throws RecognitionException {
		ObjectLiteralContext _localctx = new ObjectLiteralContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_objectLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			match(OpenBrace);
			setState(260);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OpenBracket || _la==StringLiteral) {
				{
				setState(247);
				key();
				setState(248);
				match(Colon);
				setState(249);
				value();
				setState(257);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(250);
					match(Comma);
					setState(251);
					key();
					setState(252);
					match(Colon);
					setState(253);
					value();
					}
					}
					setState(259);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(262);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeyContext extends ParserRuleContext {
		public TerminalNode StringLiteral() { return getToken(litelangParser.StringLiteral, 0); }
		public TerminalNode OpenBracket() { return getToken(litelangParser.OpenBracket, 0); }
		public AccessRhsContext accessRhs() {
			return getRuleContext(AccessRhsContext.class,0);
		}
		public TerminalNode CloseBracket() { return getToken(litelangParser.CloseBracket, 0); }
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case StringLiteral:
				{
				setState(264);
				match(StringLiteral);
				}
				break;
			case OpenBracket:
				{
				setState(265);
				match(OpenBracket);
				setState(266);
				accessRhs();
				setState(267);
				match(CloseBracket);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends ParserRuleContext {
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_value);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenParen:
			case Not:
			case NullLiteral:
			case BooleanLiteral:
			case DecimalLiteral:
			case Identifier:
			case StringLiteral:
				{
				setState(271);
				unit(0);
				}
				break;
			case OpenBrace:
				{
				setState(272);
				objectLiteral();
				}
				break;
			case OpenBracket:
				{
				setState(273);
				arrayLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(litelangParser.OpenBrace, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(litelangParser.CloseBrace, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(OpenBrace);
			setState(277);
			statements();
			setState(278);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConstantContext extends ParserRuleContext {
		public TerminalNode NullLiteral() { return getToken(litelangParser.NullLiteral, 0); }
		public TerminalNode BooleanLiteral() { return getToken(litelangParser.BooleanLiteral, 0); }
		public TerminalNode DecimalLiteral() { return getToken(litelangParser.DecimalLiteral, 0); }
		public TerminalNode StringLiteral() { return getToken(litelangParser.StringLiteral, 0); }
		public ConstantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant; }
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_constant);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(280);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4503629692141568L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 10:
			return unit_sempred((UnitContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean unit_sempred(UnitContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 8);
		case 1:
			return precpred(_ctx, 7);
		case 2:
			return precpred(_ctx, 6);
		case 3:
			return precpred(_ctx, 5);
		case 4:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u00016\u011b\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0001\u0000\u0003\u00006\b\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0004\u0001;\b\u0001\u000b\u0001\f\u0001<\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002L\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0003\u0003R\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0005\u0006j\b\u0006\n\u0006\f\u0006m\t\u0006\u0001\u0006\u0001"+
		"\u0006\u0003\u0006q\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0003\u0007w\b\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u0091\b\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0005\n\u00a2\b\n\n\n\f\n\u00a5\t\n\u0001\u000b"+
		"\u0001\u000b\u0005\u000b\u00a9\b\u000b\n\u000b\f\u000b\u00ac\t\u000b\u0001"+
		"\f\u0001\f\u0003\f\u00b0\b\f\u0001\f\u0005\f\u00b3\b\f\n\f\f\f\u00b6\t"+
		"\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00bd\b\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0003\u000e\u00c2\b\u000e\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0003\u0010\u00cc\b\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0005\u0011\u00d3\b\u0011\n\u0011\f\u0011\u00d6\t\u0011\u0003"+
		"\u0011\u00d8\b\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u00dd"+
		"\b\u0012\n\u0012\f\u0012\u00e0\t\u0012\u0003\u0012\u00e2\b\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u00ee\b\u0014\n"+
		"\u0014\f\u0014\u00f1\t\u0014\u0003\u0014\u00f3\b\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u0100\b\u0015\n"+
		"\u0015\f\u0015\u0103\t\u0015\u0003\u0015\u0105\b\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003"+
		"\u0016\u010e\b\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u0113"+
		"\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0000\u0001\u0014\u001a\u0000\u0002\u0004\u0006\b\n"+
		"\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02\u0000"+
		"\u0003\u0001\u0000\u0018\u001f\u0002\u0000\"\"34\u0002\u0000 \"44\u0126"+
		"\u00005\u0001\u0000\u0000\u0000\u0002:\u0001\u0000\u0000\u0000\u0004K"+
		"\u0001\u0000\u0000\u0000\u0006M\u0001\u0000\u0000\u0000\bS\u0001\u0000"+
		"\u0000\u0000\nW\u0001\u0000\u0000\u0000\f^\u0001\u0000\u0000\u0000\u000e"+
		"r\u0001\u0000\u0000\u0000\u0010\u007f\u0001\u0000\u0000\u0000\u0012\u0085"+
		"\u0001\u0000\u0000\u0000\u0014\u0090\u0001\u0000\u0000\u0000\u0016\u00a6"+
		"\u0001\u0000\u0000\u0000\u0018\u00af\u0001\u0000\u0000\u0000\u001a\u00bc"+
		"\u0001\u0000\u0000\u0000\u001c\u00c1\u0001\u0000\u0000\u0000\u001e\u00c3"+
		"\u0001\u0000\u0000\u0000 \u00c7\u0001\u0000\u0000\u0000\"\u00d7\u0001"+
		"\u0000\u0000\u0000$\u00e1\u0001\u0000\u0000\u0000&\u00e3\u0001\u0000\u0000"+
		"\u0000(\u00e9\u0001\u0000\u0000\u0000*\u00f6\u0001\u0000\u0000\u0000,"+
		"\u010d\u0001\u0000\u0000\u0000.\u0112\u0001\u0000\u0000\u00000\u0114\u0001"+
		"\u0000\u0000\u00002\u0118\u0001\u0000\u0000\u000046\u0003\u0002\u0001"+
		"\u000054\u0001\u0000\u0000\u000056\u0001\u0000\u0000\u000067\u0001\u0000"+
		"\u0000\u000078\u0005\u0000\u0000\u00018\u0001\u0001\u0000\u0000\u0000"+
		"9;\u0003\u0004\u0002\u0000:9\u0001\u0000\u0000\u0000;<\u0001\u0000\u0000"+
		"\u0000<:\u0001\u0000\u0000\u0000<=\u0001\u0000\u0000\u0000=\u0003\u0001"+
		"\u0000\u0000\u0000>?\u0003\u0006\u0003\u0000?@\u0005\b\u0000\u0000@L\u0001"+
		"\u0000\u0000\u0000AB\u0003\b\u0004\u0000BC\u0005\b\u0000\u0000CL\u0001"+
		"\u0000\u0000\u0000DL\u0003\n\u0005\u0000EL\u0003\f\u0006\u0000FL\u0003"+
		"\u000e\u0007\u0000GL\u0003\u0010\b\u0000HI\u0003\u0012\t\u0000IJ\u0005"+
		"\b\u0000\u0000JL\u0001\u0000\u0000\u0000K>\u0001\u0000\u0000\u0000KA\u0001"+
		"\u0000\u0000\u0000KD\u0001\u0000\u0000\u0000KE\u0001\u0000\u0000\u0000"+
		"KF\u0001\u0000\u0000\u0000KG\u0001\u0000\u0000\u0000KH\u0001\u0000\u0000"+
		"\u0000L\u0005\u0001\u0000\u0000\u0000MN\u0005(\u0000\u0000NQ\u00053\u0000"+
		"\u0000OP\u0005\n\u0000\u0000PR\u0003.\u0017\u0000QO\u0001\u0000\u0000"+
		"\u0000QR\u0001\u0000\u0000\u0000R\u0007\u0001\u0000\u0000\u0000ST\u0003"+
		"\u0016\u000b\u0000TU\u0005\n\u0000\u0000UV\u0003\u0014\n\u0000V\t\u0001"+
		"\u0000\u0000\u0000WX\u0005.\u0000\u0000XY\u00053\u0000\u0000YZ\u0005\u0004"+
		"\u0000\u0000Z[\u0003\"\u0011\u0000[\\\u0005\u0005\u0000\u0000\\]\u0003"+
		"0\u0018\u0000]\u000b\u0001\u0000\u0000\u0000^_\u00050\u0000\u0000_`\u0005"+
		"\u0004\u0000\u0000`a\u0003\u0014\n\u0000ab\u0005\u0005\u0000\u0000bk\u0003"+
		"0\u0018\u0000cd\u00051\u0000\u0000de\u0005\u0004\u0000\u0000ef\u0003\u0014"+
		"\n\u0000fg\u0005\u0005\u0000\u0000gh\u00030\u0018\u0000hj\u0001\u0000"+
		"\u0000\u0000ic\u0001\u0000\u0000\u0000jm\u0001\u0000\u0000\u0000ki\u0001"+
		"\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000lp\u0001\u0000\u0000\u0000"+
		"mk\u0001\u0000\u0000\u0000no\u0005&\u0000\u0000oq\u00030\u0018\u0000p"+
		"n\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000q\r\u0001\u0000\u0000"+
		"\u0000rs\u0005+\u0000\u0000sv\u0005\u0004\u0000\u0000tw\u0003\u0006\u0003"+
		"\u0000uw\u0003\b\u0004\u0000vt\u0001\u0000\u0000\u0000vu\u0001\u0000\u0000"+
		"\u0000wx\u0001\u0000\u0000\u0000xy\u0005\b\u0000\u0000yz\u0003\u0012\t"+
		"\u0000z{\u0005\b\u0000\u0000{|\u0003\b\u0004\u0000|}\u0005\u0005\u0000"+
		"\u0000}~\u00030\u0018\u0000~\u000f\u0001\u0000\u0000\u0000\u007f\u0080"+
		"\u0005,\u0000\u0000\u0080\u0081\u0005\u0004\u0000\u0000\u0081\u0082\u0003"+
		"\u0014\n\u0000\u0082\u0083\u0005\u0005\u0000\u0000\u0083\u0084\u00030"+
		"\u0018\u0000\u0084\u0011\u0001\u0000\u0000\u0000\u0085\u0086\u0003\u0014"+
		"\n\u0000\u0086\u0013\u0001\u0000\u0000\u0000\u0087\u0088\u0006\n\uffff"+
		"\uffff\u0000\u0088\u0089\u0005\u0004\u0000\u0000\u0089\u008a\u0003\u0014"+
		"\n\u0000\u008a\u008b\u0005\u0005\u0000\u0000\u008b\u0091\u0001\u0000\u0000"+
		"\u0000\u008c\u008d\u0005\u0013\u0000\u0000\u008d\u0091\u0003\u0014\n\u0003"+
		"\u008e\u0091\u0003\u0018\f\u0000\u008f\u0091\u00032\u0019\u0000\u0090"+
		"\u0087\u0001\u0000\u0000\u0000\u0090\u008c\u0001\u0000\u0000\u0000\u0090"+
		"\u008e\u0001\u0000\u0000\u0000\u0090\u008f\u0001\u0000\u0000\u0000\u0091"+
		"\u00a3\u0001\u0000\u0000\u0000\u0092\u0093\n\b\u0000\u0000\u0093\u0094"+
		"\u0005\u0015\u0000\u0000\u0094\u00a2\u0003\u0014\n\t\u0095\u0096\n\u0007"+
		"\u0000\u0000\u0096\u0097\u0005\u0014\u0000\u0000\u0097\u00a2\u0003\u0014"+
		"\n\b\u0098\u0099\n\u0006\u0000\u0000\u0099\u009a\u0005\u0011\u0000\u0000"+
		"\u009a\u00a2\u0003\u0014\n\u0007\u009b\u009c\n\u0005\u0000\u0000\u009c"+
		"\u009d\u0005\u0012\u0000\u0000\u009d\u00a2\u0003\u0014\n\u0006\u009e\u009f"+
		"\n\u0004\u0000\u0000\u009f\u00a0\u0007\u0000\u0000\u0000\u00a0\u00a2\u0003"+
		"\u0014\n\u0005\u00a1\u0092\u0001\u0000\u0000\u0000\u00a1\u0095\u0001\u0000"+
		"\u0000\u0000\u00a1\u0098\u0001\u0000\u0000\u0000\u00a1\u009b\u0001\u0000"+
		"\u0000\u0000\u00a1\u009e\u0001\u0000\u0000\u0000\u00a2\u00a5\u0001\u0000"+
		"\u0000\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a3\u00a4\u0001\u0000"+
		"\u0000\u0000\u00a4\u0015\u0001\u0000\u0000\u0000\u00a5\u00a3\u0001\u0000"+
		"\u0000\u0000\u00a6\u00aa\u00053\u0000\u0000\u00a7\u00a9\u0003\u001c\u000e"+
		"\u0000\u00a8\u00a7\u0001\u0000\u0000\u0000\u00a9\u00ac\u0001\u0000\u0000"+
		"\u0000\u00aa\u00a8\u0001\u0000\u0000\u0000\u00aa\u00ab\u0001\u0000\u0000"+
		"\u0000\u00ab\u0017\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000\u0000"+
		"\u0000\u00ad\u00b0\u0003 \u0010\u0000\u00ae\u00b0\u00053\u0000\u0000\u00af"+
		"\u00ad\u0001\u0000\u0000\u0000\u00af\u00ae\u0001\u0000\u0000\u0000\u00b0"+
		"\u00b4\u0001\u0000\u0000\u0000\u00b1\u00b3\u0003\u001a\r\u0000\u00b2\u00b1"+
		"\u0001\u0000\u0000\u0000\u00b3\u00b6\u0001\u0000\u0000\u0000\u00b4\u00b2"+
		"\u0001\u0000\u0000\u0000\u00b4\u00b5\u0001\u0000\u0000\u0000\u00b5\u0019"+
		"\u0001\u0000\u0000\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b7\u00bd"+
		"\u0003\u001e\u000f\u0000\u00b8\u00b9\u0005\u000e\u0000\u0000\u00b9\u00bd"+
		"\u00053\u0000\u0000\u00ba\u00bb\u0005\u000e\u0000\u0000\u00bb\u00bd\u0003"+
		" \u0010\u0000\u00bc\u00b7\u0001\u0000\u0000\u0000\u00bc\u00b8\u0001\u0000"+
		"\u0000\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000\u00bd\u001b\u0001\u0000"+
		"\u0000\u0000\u00be\u00c2\u0003\u001e\u000f\u0000\u00bf\u00c0\u0005\u000e"+
		"\u0000\u0000\u00c0\u00c2\u00053\u0000\u0000\u00c1\u00be\u0001\u0000\u0000"+
		"\u0000\u00c1\u00bf\u0001\u0000\u0000\u0000\u00c2\u001d\u0001\u0000\u0000"+
		"\u0000\u00c3\u00c4\u0005\u0002\u0000\u0000\u00c4\u00c5\u0007\u0001\u0000"+
		"\u0000\u00c5\u00c6\u0005\u0003\u0000\u0000\u00c6\u001f\u0001\u0000\u0000"+
		"\u0000\u00c7\u00c8\u00053\u0000\u0000\u00c8\u00cb\u0005\u0004\u0000\u0000"+
		"\u00c9\u00cc\u0003$\u0012\u0000\u00ca\u00cc\u0003&\u0013\u0000\u00cb\u00c9"+
		"\u0001\u0000\u0000\u0000\u00cb\u00ca\u0001\u0000\u0000\u0000\u00cc\u00cd"+
		"\u0001\u0000\u0000\u0000\u00cd\u00ce\u0005\u0005\u0000\u0000\u00ce!\u0001"+
		"\u0000\u0000\u0000\u00cf\u00d4\u00053\u0000\u0000\u00d0\u00d1\u0005\t"+
		"\u0000\u0000\u00d1\u00d3\u00053\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000"+
		"\u0000\u00d3\u00d6\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000"+
		"\u0000\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u00d8\u0001\u0000\u0000"+
		"\u0000\u00d6\u00d4\u0001\u0000\u0000\u0000\u00d7\u00cf\u0001\u0000\u0000"+
		"\u0000\u00d7\u00d8\u0001\u0000\u0000\u0000\u00d8#\u0001\u0000\u0000\u0000"+
		"\u00d9\u00de\u0003\u0014\n\u0000\u00da\u00db\u0005\t\u0000\u0000\u00db"+
		"\u00dd\u0003\u0014\n\u0000\u00dc\u00da\u0001\u0000\u0000\u0000\u00dd\u00e0"+
		"\u0001\u0000\u0000\u0000\u00de\u00dc\u0001\u0000\u0000\u0000\u00de\u00df"+
		"\u0001\u0000\u0000\u0000\u00df\u00e2\u0001\u0000\u0000\u0000\u00e0\u00de"+
		"\u0001\u0000\u0000\u0000\u00e1\u00d9\u0001\u0000\u0000\u0000\u00e1\u00e2"+
		"\u0001\u0000\u0000\u0000\u00e2%\u0001\u0000\u0000\u0000\u00e3\u00e4\u0005"+
		"\u0004\u0000\u0000\u00e4\u00e5\u0003$\u0012\u0000\u00e5\u00e6\u0005\u0005"+
		"\u0000\u0000\u00e6\u00e7\u0005\u0001\u0000\u0000\u00e7\u00e8\u00030\u0018"+
		"\u0000\u00e8\'\u0001\u0000\u0000\u0000\u00e9\u00f2\u0005\u0002\u0000\u0000"+
		"\u00ea\u00ef\u0003.\u0017\u0000\u00eb\u00ec\u0005\t\u0000\u0000\u00ec"+
		"\u00ee\u0003.\u0017\u0000\u00ed\u00eb\u0001\u0000\u0000\u0000\u00ee\u00f1"+
		"\u0001\u0000\u0000\u0000\u00ef\u00ed\u0001\u0000\u0000\u0000\u00ef\u00f0"+
		"\u0001\u0000\u0000\u0000\u00f0\u00f3\u0001\u0000\u0000\u0000\u00f1\u00ef"+
		"\u0001\u0000\u0000\u0000\u00f2\u00ea\u0001\u0000\u0000\u0000\u00f2\u00f3"+
		"\u0001\u0000\u0000\u0000\u00f3\u00f4\u0001\u0000\u0000\u0000\u00f4\u00f5"+
		"\u0005\u0003\u0000\u0000\u00f5)\u0001\u0000\u0000\u0000\u00f6\u0104\u0005"+
		"\u0006\u0000\u0000\u00f7\u00f8\u0003,\u0016\u0000\u00f8\u00f9\u0005\r"+
		"\u0000\u0000\u00f9\u0101\u0003.\u0017\u0000\u00fa\u00fb\u0005\t\u0000"+
		"\u0000\u00fb\u00fc\u0003,\u0016\u0000\u00fc\u00fd\u0005\r\u0000\u0000"+
		"\u00fd\u00fe\u0003.\u0017\u0000\u00fe\u0100\u0001\u0000\u0000\u0000\u00ff"+
		"\u00fa\u0001\u0000\u0000\u0000\u0100\u0103\u0001\u0000\u0000\u0000\u0101"+
		"\u00ff\u0001\u0000\u0000\u0000\u0101\u0102\u0001\u0000\u0000\u0000\u0102"+
		"\u0105\u0001\u0000\u0000\u0000\u0103\u0101\u0001\u0000\u0000\u0000\u0104"+
		"\u00f7\u0001\u0000\u0000\u0000\u0104\u0105\u0001\u0000\u0000\u0000\u0105"+
		"\u0106\u0001\u0000\u0000\u0000\u0106\u0107\u0005\u0007\u0000\u0000\u0107"+
		"+\u0001\u0000\u0000\u0000\u0108\u010e\u00054\u0000\u0000\u0109\u010a\u0005"+
		"\u0002\u0000\u0000\u010a\u010b\u0003\u0018\f\u0000\u010b\u010c\u0005\u0003"+
		"\u0000\u0000\u010c\u010e\u0001\u0000\u0000\u0000\u010d\u0108\u0001\u0000"+
		"\u0000\u0000\u010d\u0109\u0001\u0000\u0000\u0000\u010e-\u0001\u0000\u0000"+
		"\u0000\u010f\u0113\u0003\u0014\n\u0000\u0110\u0113\u0003*\u0015\u0000"+
		"\u0111\u0113\u0003(\u0014\u0000\u0112\u010f\u0001\u0000\u0000\u0000\u0112"+
		"\u0110\u0001\u0000\u0000\u0000\u0112\u0111\u0001\u0000\u0000\u0000\u0113"+
		"/\u0001\u0000\u0000\u0000\u0114\u0115\u0005\u0006\u0000\u0000\u0115\u0116"+
		"\u0003\u0002\u0001\u0000\u0116\u0117\u0005\u0007\u0000\u0000\u01171\u0001"+
		"\u0000\u0000\u0000\u0118\u0119\u0007\u0002\u0000\u0000\u01193\u0001\u0000"+
		"\u0000\u0000\u001a5<KQkpv\u0090\u00a1\u00a3\u00aa\u00af\u00b4\u00bc\u00c1"+
		"\u00cb\u00d4\u00d7\u00de\u00e1\u00ef\u00f2\u0101\u0104\u010d\u0112";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}