// Generated from c://Users//chillTime//Documents//datamorf//src//DataMorf.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class DataMorfParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, OpenSquare=2, CloseSquare=3, OpenRound=4, CloseRound=5, OpenCurly=6, 
		CloseCurly=7, SemiColon=8, Comma=9, Assign=10, QuestionMark=11, QuestionMarkDot=12, 
		Colon=13, Dot=14, Plus=15, Minus=16, Not=17, Multiply=18, Divide=19, Modulus=20, 
		Power=21, LessThan=22, MoreThan=23, LessThanEquals=24, GreaterThanEquals=25, 
		Equals=26, NotEquals=27, And=28, Or=29, NullLiteral=30, BooleanLiteral=31, 
		DecimalLiteral=32, Break=33, Do=34, Of=35, Case=36, Else=37, Var=38, Return=39, 
		Continue=40, For=41, Switch=42, While=43, Function=44, Default=45, If=46, 
		Elif=47, Delete=48, Identifier=49, StringLiteral=50, WhiteSpaces=51, LineTerminator=52;
	public static final int
		RULE_program = 0, RULE_statements = 1, RULE_statement = 2, RULE_variableStatement = 3, 
		RULE_assignStatement = 4, RULE_functionStatement = 5, RULE_ifStatement = 6, 
		RULE_switchStatement = 7, RULE_deleteStatement = 8, RULE_unitStatement = 9, 
		RULE_valueStatement = 10, RULE_forStatement = 11, RULE_classicForCondition = 12, 
		RULE_iteratorForCondition = 13, RULE_unit = 14, RULE_reference = 15, RULE_accessor = 16, 
		RULE_accessProperty = 17, RULE_functionCall = 18, RULE_signatureParams = 19, 
		RULE_sendingParams = 20, RULE_arrowFunction = 21, RULE_arrayLiteral = 22, 
		RULE_objectLiteral = 23, RULE_key = 24, RULE_value = 25, RULE_keyValue = 26, 
		RULE_objectItem = 27, RULE_spread = 28, RULE_block = 29, RULE_constant = 30;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statements", "statement", "variableStatement", "assignStatement", 
			"functionStatement", "ifStatement", "switchStatement", "deleteStatement", 
			"unitStatement", "valueStatement", "forStatement", "classicForCondition", 
			"iteratorForCondition", "unit", "reference", "accessor", "accessProperty", 
			"functionCall", "signatureParams", "sendingParams", "arrowFunction", 
			"arrayLiteral", "objectLiteral", "key", "value", "keyValue", "objectItem", 
			"spread", "block", "constant"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'=>'", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','", 
			"'='", "'?'", "'?.'", "':'", "'.'", "'+'", "'-'", "'!'", "'*'", "'/'", 
			"'%'", "'**'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'", 
			"'||'", "'null'", null, null, "'break'", "'do'", "'of'", "'case'", "'else'", 
			"'var'", "'return'", "'continue'", "'for'", "'switch'", "'while'", "'function'", 
			"'default'", "'if'", "'elif'", "'delete'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "OpenSquare", "CloseSquare", "OpenRound", "CloseRound", "OpenCurly", 
			"CloseCurly", "SemiColon", "Comma", "Assign", "QuestionMark", "QuestionMarkDot", 
			"Colon", "Dot", "Plus", "Minus", "Not", "Multiply", "Divide", "Modulus", 
			"Power", "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals", 
			"Equals", "NotEquals", "And", "Or", "NullLiteral", "BooleanLiteral", 
			"DecimalLiteral", "Break", "Do", "Of", "Case", "Else", "Var", "Return", 
			"Continue", "For", "Switch", "While", "Function", "Default", "If", "Elif", 
			"Delete", "Identifier", "StringLiteral", "WhiteSpaces", "LineTerminator"
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
	public String getGrammarFileName() { return "DataMorf.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public DataMorfParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(DataMorfParser.EOF, 0); }
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
			setState(63);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2065165231194196L) != 0)) {
				{
				setState(62);
				statements();
				}
			}

			setState(65);
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
			setState(68); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(67);
				statement();
				}
				}
				setState(70); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 2065165231194196L) != 0) );
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
		public TerminalNode SemiColon() { return getToken(DataMorfParser.SemiColon, 0); }
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
		public ValueStatementContext valueStatement() {
			return getRuleContext(ValueStatementContext.class,0);
		}
		public DeleteStatementContext deleteStatement() {
			return getRuleContext(DeleteStatementContext.class,0);
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
			setState(89);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				variableStatement();
				setState(73);
				match(SemiColon);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(75);
				assignStatement();
				setState(76);
				match(SemiColon);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(78);
				functionStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(79);
				ifStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(80);
				forStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(81);
				switchStatement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(82);
				valueStatement();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(83);
				deleteStatement();
				setState(84);
				match(SemiColon);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(86);
				unitStatement();
				setState(87);
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
		public TerminalNode Var() { return getToken(DataMorfParser.Var, 0); }
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public TerminalNode Assign() { return getToken(DataMorfParser.Assign, 0); }
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
			setState(91);
			match(Var);
			setState(92);
			match(Identifier);
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Assign) {
				{
				setState(93);
				match(Assign);
				setState(94);
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
		public ReferenceContext reference() {
			return getRuleContext(ReferenceContext.class,0);
		}
		public TerminalNode Assign() { return getToken(DataMorfParser.Assign, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
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
			setState(97);
			reference();
			setState(98);
			match(Assign);
			setState(99);
			value();
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
		public TerminalNode Function() { return getToken(DataMorfParser.Function, 0); }
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public TerminalNode OpenRound() { return getToken(DataMorfParser.OpenRound, 0); }
		public SignatureParamsContext signatureParams() {
			return getRuleContext(SignatureParamsContext.class,0);
		}
		public TerminalNode CloseRound() { return getToken(DataMorfParser.CloseRound, 0); }
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
			setState(101);
			match(Function);
			setState(102);
			match(Identifier);
			setState(103);
			match(OpenRound);
			setState(104);
			signatureParams();
			setState(105);
			match(CloseRound);
			setState(106);
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
		public List<TerminalNode> If() { return getTokens(DataMorfParser.If); }
		public TerminalNode If(int i) {
			return getToken(DataMorfParser.If, i);
		}
		public List<TerminalNode> OpenRound() { return getTokens(DataMorfParser.OpenRound); }
		public TerminalNode OpenRound(int i) {
			return getToken(DataMorfParser.OpenRound, i);
		}
		public List<UnitContext> unit() {
			return getRuleContexts(UnitContext.class);
		}
		public UnitContext unit(int i) {
			return getRuleContext(UnitContext.class,i);
		}
		public List<TerminalNode> CloseRound() { return getTokens(DataMorfParser.CloseRound); }
		public TerminalNode CloseRound(int i) {
			return getToken(DataMorfParser.CloseRound, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> Else() { return getTokens(DataMorfParser.Else); }
		public TerminalNode Else(int i) {
			return getToken(DataMorfParser.Else, i);
		}
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
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(If);
			setState(109);
			match(OpenRound);
			setState(110);
			unit(0);
			setState(111);
			match(CloseRound);
			setState(112);
			block();
			setState(122);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(113);
					match(Else);
					setState(114);
					match(If);
					setState(115);
					match(OpenRound);
					setState(116);
					unit(0);
					setState(117);
					match(CloseRound);
					setState(118);
					block();
					}
					} 
				}
				setState(124);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Else) {
				{
				setState(125);
				match(Else);
				setState(126);
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
	public static class SwitchStatementContext extends ParserRuleContext {
		public TerminalNode Switch() { return getToken(DataMorfParser.Switch, 0); }
		public TerminalNode OpenRound() { return getToken(DataMorfParser.OpenRound, 0); }
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public TerminalNode CloseRound() { return getToken(DataMorfParser.CloseRound, 0); }
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
		enterRule(_localctx, 14, RULE_switchStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(Switch);
			setState(130);
			match(OpenRound);
			setState(131);
			unit(0);
			setState(132);
			match(CloseRound);
			setState(133);
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
	public static class DeleteStatementContext extends ParserRuleContext {
		public TerminalNode Delete() { return getToken(DataMorfParser.Delete, 0); }
		public List<TerminalNode> Identifier() { return getTokens(DataMorfParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(DataMorfParser.Identifier, i);
		}
		public List<TerminalNode> Dot() { return getTokens(DataMorfParser.Dot); }
		public TerminalNode Dot(int i) {
			return getToken(DataMorfParser.Dot, i);
		}
		public List<AccessPropertyContext> accessProperty() {
			return getRuleContexts(AccessPropertyContext.class);
		}
		public AccessPropertyContext accessProperty(int i) {
			return getRuleContext(AccessPropertyContext.class,i);
		}
		public DeleteStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_deleteStatement; }
	}

	public final DeleteStatementContext deleteStatement() throws RecognitionException {
		DeleteStatementContext _localctx = new DeleteStatementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_deleteStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(Delete);
			{
			setState(136);
			match(Identifier);
			setState(144);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Dot) {
				{
				{
				setState(137);
				match(Dot);
				setState(140);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case OpenSquare:
					{
					setState(138);
					accessProperty();
					}
					break;
				case Identifier:
					{
					setState(139);
					match(Identifier);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				}
				setState(146);
				_errHandler.sync(this);
				_la = _input.LA(1);
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
			setState(147);
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
	public static class ValueStatementContext extends ParserRuleContext {
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(DataMorfParser.SemiColon, 0); }
		public ValueStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valueStatement; }
	}

	public final ValueStatementContext valueStatement() throws RecognitionException {
		ValueStatementContext _localctx = new ValueStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_valueStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			value();
			setState(150);
			match(SemiColon);
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
		public TerminalNode For() { return getToken(DataMorfParser.For, 0); }
		public TerminalNode OpenRound() { return getToken(DataMorfParser.OpenRound, 0); }
		public TerminalNode CloseRound() { return getToken(DataMorfParser.CloseRound, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ClassicForConditionContext classicForCondition() {
			return getRuleContext(ClassicForConditionContext.class,0);
		}
		public IteratorForConditionContext iteratorForCondition() {
			return getRuleContext(IteratorForConditionContext.class,0);
		}
		public ForStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStatement; }
	}

	public final ForStatementContext forStatement() throws RecognitionException {
		ForStatementContext _localctx = new ForStatementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_forStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			match(For);
			setState(153);
			match(OpenRound);
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(154);
				classicForCondition();
				}
				break;
			case 2:
				{
				setState(155);
				iteratorForCondition();
				}
				break;
			}
			setState(158);
			match(CloseRound);
			setState(159);
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
	public static class ClassicForConditionContext extends ParserRuleContext {
		public VariableStatementContext init;
		public UnitStatementContext condition;
		public AssignStatementContext postOp;
		public List<TerminalNode> SemiColon() { return getTokens(DataMorfParser.SemiColon); }
		public TerminalNode SemiColon(int i) {
			return getToken(DataMorfParser.SemiColon, i);
		}
		public VariableStatementContext variableStatement() {
			return getRuleContext(VariableStatementContext.class,0);
		}
		public UnitStatementContext unitStatement() {
			return getRuleContext(UnitStatementContext.class,0);
		}
		public AssignStatementContext assignStatement() {
			return getRuleContext(AssignStatementContext.class,0);
		}
		public ClassicForConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classicForCondition; }
	}

	public final ClassicForConditionContext classicForCondition() throws RecognitionException {
		ClassicForConditionContext _localctx = new ClassicForConditionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_classicForCondition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			((ClassicForConditionContext)_localctx).init = variableStatement();
			setState(162);
			match(SemiColon);
			setState(163);
			((ClassicForConditionContext)_localctx).condition = unitStatement();
			setState(164);
			match(SemiColon);
			setState(165);
			((ClassicForConditionContext)_localctx).postOp = assignStatement();
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
	public static class IteratorForConditionContext extends ParserRuleContext {
		public TerminalNode Var() { return getToken(DataMorfParser.Var, 0); }
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public TerminalNode Of() { return getToken(DataMorfParser.Of, 0); }
		public ReferenceContext reference() {
			return getRuleContext(ReferenceContext.class,0);
		}
		public IteratorForConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iteratorForCondition; }
	}

	public final IteratorForConditionContext iteratorForCondition() throws RecognitionException {
		IteratorForConditionContext _localctx = new IteratorForConditionContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_iteratorForCondition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			match(Var);
			setState(168);
			match(Identifier);
			setState(169);
			match(Of);
			setState(170);
			reference();
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
		public Token operator;
		public UnitContext right;
		public TerminalNode OpenRound() { return getToken(DataMorfParser.OpenRound, 0); }
		public TerminalNode CloseRound() { return getToken(DataMorfParser.CloseRound, 0); }
		public List<UnitContext> unit() {
			return getRuleContexts(UnitContext.class);
		}
		public UnitContext unit(int i) {
			return getRuleContext(UnitContext.class,i);
		}
		public TerminalNode Not() { return getToken(DataMorfParser.Not, 0); }
		public ReferenceContext reference() {
			return getRuleContext(ReferenceContext.class,0);
		}
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public TerminalNode Divide() { return getToken(DataMorfParser.Divide, 0); }
		public TerminalNode Multiply() { return getToken(DataMorfParser.Multiply, 0); }
		public TerminalNode Plus() { return getToken(DataMorfParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(DataMorfParser.Minus, 0); }
		public TerminalNode LessThan() { return getToken(DataMorfParser.LessThan, 0); }
		public TerminalNode MoreThan() { return getToken(DataMorfParser.MoreThan, 0); }
		public TerminalNode LessThanEquals() { return getToken(DataMorfParser.LessThanEquals, 0); }
		public TerminalNode GreaterThanEquals() { return getToken(DataMorfParser.GreaterThanEquals, 0); }
		public TerminalNode Equals() { return getToken(DataMorfParser.Equals, 0); }
		public TerminalNode NotEquals() { return getToken(DataMorfParser.NotEquals, 0); }
		public TerminalNode And() { return getToken(DataMorfParser.And, 0); }
		public TerminalNode Or() { return getToken(DataMorfParser.Or, 0); }
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
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_unit, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenRound:
				{
				setState(173);
				match(OpenRound);
				setState(174);
				((UnitContext)_localctx).single = unit(0);
				setState(175);
				match(CloseRound);
				}
				break;
			case Not:
				{
				setState(177);
				match(Not);
				setState(178);
				((UnitContext)_localctx).single = unit(3);
				}
				break;
			case Identifier:
				{
				setState(179);
				reference();
				}
				break;
			case NullLiteral:
			case BooleanLiteral:
			case DecimalLiteral:
			case StringLiteral:
				{
				setState(180);
				constant();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(221);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(219);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(183);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(184);
						((UnitContext)_localctx).operator = match(Divide);
						setState(185);
						((UnitContext)_localctx).right = unit(16);
						}
						break;
					case 2:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(186);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(187);
						((UnitContext)_localctx).operator = match(Multiply);
						setState(188);
						((UnitContext)_localctx).right = unit(15);
						}
						break;
					case 3:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(189);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(190);
						((UnitContext)_localctx).operator = match(Plus);
						setState(191);
						((UnitContext)_localctx).right = unit(14);
						}
						break;
					case 4:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(192);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(193);
						((UnitContext)_localctx).operator = match(Minus);
						setState(194);
						((UnitContext)_localctx).right = unit(13);
						}
						break;
					case 5:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(195);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(196);
						((UnitContext)_localctx).operator = match(LessThan);
						setState(197);
						((UnitContext)_localctx).right = unit(12);
						}
						break;
					case 6:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(198);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(199);
						((UnitContext)_localctx).operator = match(MoreThan);
						setState(200);
						((UnitContext)_localctx).right = unit(11);
						}
						break;
					case 7:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(201);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(202);
						((UnitContext)_localctx).operator = match(LessThanEquals);
						setState(203);
						((UnitContext)_localctx).right = unit(10);
						}
						break;
					case 8:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(204);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(205);
						((UnitContext)_localctx).operator = match(GreaterThanEquals);
						setState(206);
						((UnitContext)_localctx).right = unit(9);
						}
						break;
					case 9:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(207);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(208);
						((UnitContext)_localctx).operator = match(Equals);
						setState(209);
						((UnitContext)_localctx).right = unit(8);
						}
						break;
					case 10:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(210);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(211);
						((UnitContext)_localctx).operator = match(NotEquals);
						setState(212);
						((UnitContext)_localctx).right = unit(7);
						}
						break;
					case 11:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(213);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(214);
						((UnitContext)_localctx).operator = match(And);
						setState(215);
						((UnitContext)_localctx).right = unit(6);
						}
						break;
					case 12:
						{
						_localctx = new UnitContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_unit);
						setState(216);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(217);
						((UnitContext)_localctx).operator = match(Or);
						setState(218);
						((UnitContext)_localctx).right = unit(5);
						}
						break;
					}
					} 
				}
				setState(223);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
	public static class ReferenceContext extends ParserRuleContext {
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public List<AccessorContext> accessor() {
			return getRuleContexts(AccessorContext.class);
		}
		public AccessorContext accessor(int i) {
			return getRuleContext(AccessorContext.class,i);
		}
		public ReferenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_reference; }
	}

	public final ReferenceContext reference() throws RecognitionException {
		ReferenceContext _localctx = new ReferenceContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_reference);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(224);
				functionCall();
				}
				break;
			case 2:
				{
				setState(225);
				match(Identifier);
				}
				break;
			}
			setState(231);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(228);
					accessor();
					}
					} 
				}
				setState(233);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
	public static class AccessorContext extends ParserRuleContext {
		public AccessPropertyContext accessProperty() {
			return getRuleContext(AccessPropertyContext.class,0);
		}
		public TerminalNode Dot() { return getToken(DataMorfParser.Dot, 0); }
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public AccessorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessor; }
	}

	public final AccessorContext accessor() throws RecognitionException {
		AccessorContext _localctx = new AccessorContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_accessor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(234);
				accessProperty();
				}
				break;
			case 2:
				{
				{
				setState(235);
				match(Dot);
				setState(236);
				match(Identifier);
				}
				}
				break;
			case 3:
				{
				{
				setState(237);
				match(Dot);
				setState(238);
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
	public static class AccessPropertyContext extends ParserRuleContext {
		public TerminalNode OpenSquare() { return getToken(DataMorfParser.OpenSquare, 0); }
		public TerminalNode CloseSquare() { return getToken(DataMorfParser.CloseSquare, 0); }
		public TerminalNode StringLiteral() { return getToken(DataMorfParser.StringLiteral, 0); }
		public TerminalNode DecimalLiteral() { return getToken(DataMorfParser.DecimalLiteral, 0); }
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public AccessPropertyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessProperty; }
	}

	public final AccessPropertyContext accessProperty() throws RecognitionException {
		AccessPropertyContext _localctx = new AccessPropertyContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_accessProperty);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			match(OpenSquare);
			setState(242);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688854155231232L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(243);
			match(CloseSquare);
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
		public TerminalNode Identifier() { return getToken(DataMorfParser.Identifier, 0); }
		public TerminalNode OpenRound() { return getToken(DataMorfParser.OpenRound, 0); }
		public TerminalNode CloseRound() { return getToken(DataMorfParser.CloseRound, 0); }
		public SendingParamsContext sendingParams() {
			return getRuleContext(SendingParamsContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_functionCall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			match(Identifier);
			setState(246);
			match(OpenRound);
			{
			setState(247);
			sendingParams();
			}
			setState(248);
			match(CloseRound);
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
	public static class SignatureParamsContext extends ParserRuleContext {
		public List<TerminalNode> Identifier() { return getTokens(DataMorfParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(DataMorfParser.Identifier, i);
		}
		public List<TerminalNode> Comma() { return getTokens(DataMorfParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(DataMorfParser.Comma, i);
		}
		public SignatureParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_signatureParams; }
	}

	public final SignatureParamsContext signatureParams() throws RecognitionException {
		SignatureParamsContext _localctx = new SignatureParamsContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_signatureParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(250);
				match(Identifier);
				setState(255);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(251);
					match(Comma);
					setState(252);
					match(Identifier);
					}
					}
					setState(257);
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
	public static class SendingParamsContext extends ParserRuleContext {
		public List<UnitContext> unit() {
			return getRuleContexts(UnitContext.class);
		}
		public UnitContext unit(int i) {
			return getRuleContext(UnitContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(DataMorfParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(DataMorfParser.Comma, i);
		}
		public SendingParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sendingParams; }
	}

	public final SendingParamsContext sendingParams() throws RecognitionException {
		SendingParamsContext _localctx = new SendingParamsContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_sendingParams);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688857376587792L) != 0)) {
				{
				setState(260);
				unit(0);
				setState(265);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(261);
					match(Comma);
					setState(262);
					unit(0);
					}
					}
					setState(267);
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
		public TerminalNode OpenRound() { return getToken(DataMorfParser.OpenRound, 0); }
		public SignatureParamsContext signatureParams() {
			return getRuleContext(SignatureParamsContext.class,0);
		}
		public TerminalNode CloseRound() { return getToken(DataMorfParser.CloseRound, 0); }
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
		enterRule(_localctx, 42, RULE_arrowFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			match(OpenRound);
			setState(271);
			signatureParams();
			setState(272);
			match(CloseRound);
			setState(273);
			match(T__0);
			setState(274);
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
		public TerminalNode OpenSquare() { return getToken(DataMorfParser.OpenSquare, 0); }
		public TerminalNode CloseSquare() { return getToken(DataMorfParser.CloseSquare, 0); }
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(DataMorfParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(DataMorfParser.Comma, i);
		}
		public ArrayLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteral; }
	}

	public final ArrayLiteralContext arrayLiteral() throws RecognitionException {
		ArrayLiteralContext _localctx = new ArrayLiteralContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(OpenSquare);
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1688857376587860L) != 0)) {
				{
				setState(277);
				value();
				setState(282);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(278);
					match(Comma);
					setState(279);
					value();
					}
					}
					setState(284);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(287);
			match(CloseSquare);
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
		public TerminalNode OpenCurly() { return getToken(DataMorfParser.OpenCurly, 0); }
		public TerminalNode CloseCurly() { return getToken(DataMorfParser.CloseCurly, 0); }
		public List<ObjectItemContext> objectItem() {
			return getRuleContexts(ObjectItemContext.class);
		}
		public ObjectItemContext objectItem(int i) {
			return getRuleContext(ObjectItemContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(DataMorfParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(DataMorfParser.Comma, i);
		}
		public ObjectLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectLiteral; }
	}

	public final ObjectLiteralContext objectLiteral() throws RecognitionException {
		ObjectLiteralContext _localctx = new ObjectLiteralContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_objectLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			match(OpenCurly);
			setState(298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1125899906859012L) != 0)) {
				{
				setState(290);
				objectItem();
				setState(295);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(291);
					match(Comma);
					{
					setState(292);
					objectItem();
					}
					}
					}
					setState(297);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(300);
			match(CloseCurly);
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
		public TerminalNode StringLiteral() { return getToken(DataMorfParser.StringLiteral, 0); }
		public TerminalNode OpenSquare() { return getToken(DataMorfParser.OpenSquare, 0); }
		public ReferenceContext reference() {
			return getRuleContext(ReferenceContext.class,0);
		}
		public TerminalNode CloseSquare() { return getToken(DataMorfParser.CloseSquare, 0); }
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case StringLiteral:
				{
				setState(302);
				match(StringLiteral);
				}
				break;
			case OpenSquare:
				{
				setState(303);
				match(OpenSquare);
				setState(304);
				reference();
				setState(305);
				match(CloseSquare);
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
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
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
		enterRule(_localctx, 50, RULE_value);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				{
				setState(309);
				functionCall();
				}
				break;
			case 2:
				{
				setState(310);
				unit(0);
				}
				break;
			case 3:
				{
				setState(311);
				objectLiteral();
				}
				break;
			case 4:
				{
				setState(312);
				arrayLiteral();
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
	public static class KeyValueContext extends ParserRuleContext {
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public TerminalNode Colon() { return getToken(DataMorfParser.Colon, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public KeyValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyValue; }
	}

	public final KeyValueContext keyValue() throws RecognitionException {
		KeyValueContext _localctx = new KeyValueContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_keyValue);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(315);
			key();
			setState(316);
			match(Colon);
			setState(317);
			value();
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
	public static class ObjectItemContext extends ParserRuleContext {
		public KeyValueContext keyValue() {
			return getRuleContext(KeyValueContext.class,0);
		}
		public SpreadContext spread() {
			return getRuleContext(SpreadContext.class,0);
		}
		public ObjectItemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectItem; }
	}

	public final ObjectItemContext objectItem() throws RecognitionException {
		ObjectItemContext _localctx = new ObjectItemContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_objectItem);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenSquare:
			case StringLiteral:
				{
				setState(319);
				keyValue();
				}
				break;
			case Dot:
				{
				setState(320);
				spread();
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
	public static class SpreadContext extends ParserRuleContext {
		public List<TerminalNode> Dot() { return getTokens(DataMorfParser.Dot); }
		public TerminalNode Dot(int i) {
			return getToken(DataMorfParser.Dot, i);
		}
		public ReferenceContext reference() {
			return getRuleContext(ReferenceContext.class,0);
		}
		public SpreadContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_spread; }
	}

	public final SpreadContext spread() throws RecognitionException {
		SpreadContext _localctx = new SpreadContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_spread);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(323);
			match(Dot);
			setState(324);
			match(Dot);
			setState(325);
			match(Dot);
			setState(326);
			reference();
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
		public TerminalNode OpenCurly() { return getToken(DataMorfParser.OpenCurly, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode CloseCurly() { return getToken(DataMorfParser.CloseCurly, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			match(OpenCurly);
			setState(329);
			statements();
			setState(330);
			match(CloseCurly);
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
		public TerminalNode NullLiteral() { return getToken(DataMorfParser.NullLiteral, 0); }
		public TerminalNode BooleanLiteral() { return getToken(DataMorfParser.BooleanLiteral, 0); }
		public TerminalNode DecimalLiteral() { return getToken(DataMorfParser.DecimalLiteral, 0); }
		public TerminalNode StringLiteral() { return getToken(DataMorfParser.StringLiteral, 0); }
		public ConstantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant; }
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_constant);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(332);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1125907423035392L) != 0)) ) {
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
		case 14:
			return unit_sempred((UnitContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean unit_sempred(UnitContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		case 7:
			return precpred(_ctx, 8);
		case 8:
			return precpred(_ctx, 7);
		case 9:
			return precpred(_ctx, 6);
		case 10:
			return precpred(_ctx, 5);
		case 11:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u00014\u014f\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0001\u0000\u0003\u0000@\b\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0004\u0001E\b\u0001\u000b\u0001\f\u0001F\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002Z\b\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0003\u0003`\b\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006y\b\u0006\n\u0006"+
		"\f\u0006|\t\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0080\b\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u008d\b\b\u0005\b\u008f\b\b"+
		"\n\b\f\b\u0092\t\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u009d\b\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0003\u000e\u00b6\b\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u00dc\b\u000e\n\u000e\f\u000e"+
		"\u00df\t\u000e\u0001\u000f\u0001\u000f\u0003\u000f\u00e3\b\u000f\u0001"+
		"\u000f\u0005\u000f\u00e6\b\u000f\n\u000f\f\u000f\u00e9\t\u000f\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u00f0\b\u0010"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0005\u0013\u00fe\b\u0013\n\u0013\f\u0013\u0101\t\u0013\u0003\u0013\u0103"+
		"\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u0108\b\u0014"+
		"\n\u0014\f\u0014\u010b\t\u0014\u0003\u0014\u010d\b\u0014\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u0119\b\u0016\n\u0016\f\u0016"+
		"\u011c\t\u0016\u0003\u0016\u011e\b\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u0126\b\u0017\n"+
		"\u0017\f\u0017\u0129\t\u0017\u0003\u0017\u012b\b\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003"+
		"\u0018\u0134\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003"+
		"\u0019\u013a\b\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001b\u0001\u001b\u0003\u001b\u0142\b\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0000\u0001\u001c\u001f\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.02468:<\u0000\u0002\u0002\u0000  12\u0002\u0000\u001e "+
		"22\u015f\u0000?\u0001\u0000\u0000\u0000\u0002D\u0001\u0000\u0000\u0000"+
		"\u0004Y\u0001\u0000\u0000\u0000\u0006[\u0001\u0000\u0000\u0000\ba\u0001"+
		"\u0000\u0000\u0000\ne\u0001\u0000\u0000\u0000\fl\u0001\u0000\u0000\u0000"+
		"\u000e\u0081\u0001\u0000\u0000\u0000\u0010\u0087\u0001\u0000\u0000\u0000"+
		"\u0012\u0093\u0001\u0000\u0000\u0000\u0014\u0095\u0001\u0000\u0000\u0000"+
		"\u0016\u0098\u0001\u0000\u0000\u0000\u0018\u00a1\u0001\u0000\u0000\u0000"+
		"\u001a\u00a7\u0001\u0000\u0000\u0000\u001c\u00b5\u0001\u0000\u0000\u0000"+
		"\u001e\u00e2\u0001\u0000\u0000\u0000 \u00ef\u0001\u0000\u0000\u0000\""+
		"\u00f1\u0001\u0000\u0000\u0000$\u00f5\u0001\u0000\u0000\u0000&\u0102\u0001"+
		"\u0000\u0000\u0000(\u010c\u0001\u0000\u0000\u0000*\u010e\u0001\u0000\u0000"+
		"\u0000,\u0114\u0001\u0000\u0000\u0000.\u0121\u0001\u0000\u0000\u00000"+
		"\u0133\u0001\u0000\u0000\u00002\u0139\u0001\u0000\u0000\u00004\u013b\u0001"+
		"\u0000\u0000\u00006\u0141\u0001\u0000\u0000\u00008\u0143\u0001\u0000\u0000"+
		"\u0000:\u0148\u0001\u0000\u0000\u0000<\u014c\u0001\u0000\u0000\u0000>"+
		"@\u0003\u0002\u0001\u0000?>\u0001\u0000\u0000\u0000?@\u0001\u0000\u0000"+
		"\u0000@A\u0001\u0000\u0000\u0000AB\u0005\u0000\u0000\u0001B\u0001\u0001"+
		"\u0000\u0000\u0000CE\u0003\u0004\u0002\u0000DC\u0001\u0000\u0000\u0000"+
		"EF\u0001\u0000\u0000\u0000FD\u0001\u0000\u0000\u0000FG\u0001\u0000\u0000"+
		"\u0000G\u0003\u0001\u0000\u0000\u0000HI\u0003\u0006\u0003\u0000IJ\u0005"+
		"\b\u0000\u0000JZ\u0001\u0000\u0000\u0000KL\u0003\b\u0004\u0000LM\u0005"+
		"\b\u0000\u0000MZ\u0001\u0000\u0000\u0000NZ\u0003\n\u0005\u0000OZ\u0003"+
		"\f\u0006\u0000PZ\u0003\u0016\u000b\u0000QZ\u0003\u000e\u0007\u0000RZ\u0003"+
		"\u0014\n\u0000ST\u0003\u0010\b\u0000TU\u0005\b\u0000\u0000UZ\u0001\u0000"+
		"\u0000\u0000VW\u0003\u0012\t\u0000WX\u0005\b\u0000\u0000XZ\u0001\u0000"+
		"\u0000\u0000YH\u0001\u0000\u0000\u0000YK\u0001\u0000\u0000\u0000YN\u0001"+
		"\u0000\u0000\u0000YO\u0001\u0000\u0000\u0000YP\u0001\u0000\u0000\u0000"+
		"YQ\u0001\u0000\u0000\u0000YR\u0001\u0000\u0000\u0000YS\u0001\u0000\u0000"+
		"\u0000YV\u0001\u0000\u0000\u0000Z\u0005\u0001\u0000\u0000\u0000[\\\u0005"+
		"&\u0000\u0000\\_\u00051\u0000\u0000]^\u0005\n\u0000\u0000^`\u00032\u0019"+
		"\u0000_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`\u0007\u0001"+
		"\u0000\u0000\u0000ab\u0003\u001e\u000f\u0000bc\u0005\n\u0000\u0000cd\u0003"+
		"2\u0019\u0000d\t\u0001\u0000\u0000\u0000ef\u0005,\u0000\u0000fg\u0005"+
		"1\u0000\u0000gh\u0005\u0004\u0000\u0000hi\u0003&\u0013\u0000ij\u0005\u0005"+
		"\u0000\u0000jk\u0003:\u001d\u0000k\u000b\u0001\u0000\u0000\u0000lm\u0005"+
		".\u0000\u0000mn\u0005\u0004\u0000\u0000no\u0003\u001c\u000e\u0000op\u0005"+
		"\u0005\u0000\u0000pz\u0003:\u001d\u0000qr\u0005%\u0000\u0000rs\u0005."+
		"\u0000\u0000st\u0005\u0004\u0000\u0000tu\u0003\u001c\u000e\u0000uv\u0005"+
		"\u0005\u0000\u0000vw\u0003:\u001d\u0000wy\u0001\u0000\u0000\u0000xq\u0001"+
		"\u0000\u0000\u0000y|\u0001\u0000\u0000\u0000zx\u0001\u0000\u0000\u0000"+
		"z{\u0001\u0000\u0000\u0000{\u007f\u0001\u0000\u0000\u0000|z\u0001\u0000"+
		"\u0000\u0000}~\u0005%\u0000\u0000~\u0080\u0003:\u001d\u0000\u007f}\u0001"+
		"\u0000\u0000\u0000\u007f\u0080\u0001\u0000\u0000\u0000\u0080\r\u0001\u0000"+
		"\u0000\u0000\u0081\u0082\u0005*\u0000\u0000\u0082\u0083\u0005\u0004\u0000"+
		"\u0000\u0083\u0084\u0003\u001c\u000e\u0000\u0084\u0085\u0005\u0005\u0000"+
		"\u0000\u0085\u0086\u0003:\u001d\u0000\u0086\u000f\u0001\u0000\u0000\u0000"+
		"\u0087\u0088\u00050\u0000\u0000\u0088\u0090\u00051\u0000\u0000\u0089\u008c"+
		"\u0005\u000e\u0000\u0000\u008a\u008d\u0003\"\u0011\u0000\u008b\u008d\u0005"+
		"1\u0000\u0000\u008c\u008a\u0001\u0000\u0000\u0000\u008c\u008b\u0001\u0000"+
		"\u0000\u0000\u008d\u008f\u0001\u0000\u0000\u0000\u008e\u0089\u0001\u0000"+
		"\u0000\u0000\u008f\u0092\u0001\u0000\u0000\u0000\u0090\u008e\u0001\u0000"+
		"\u0000\u0000\u0090\u0091\u0001\u0000\u0000\u0000\u0091\u0011\u0001\u0000"+
		"\u0000\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0093\u0094\u0003\u001c"+
		"\u000e\u0000\u0094\u0013\u0001\u0000\u0000\u0000\u0095\u0096\u00032\u0019"+
		"\u0000\u0096\u0097\u0005\b\u0000\u0000\u0097\u0015\u0001\u0000\u0000\u0000"+
		"\u0098\u0099\u0005)\u0000\u0000\u0099\u009c\u0005\u0004\u0000\u0000\u009a"+
		"\u009d\u0003\u0018\f\u0000\u009b\u009d\u0003\u001a\r\u0000\u009c\u009a"+
		"\u0001\u0000\u0000\u0000\u009c\u009b\u0001\u0000\u0000\u0000\u009d\u009e"+
		"\u0001\u0000\u0000\u0000\u009e\u009f\u0005\u0005\u0000\u0000\u009f\u00a0"+
		"\u0003:\u001d\u0000\u00a0\u0017\u0001\u0000\u0000\u0000\u00a1\u00a2\u0003"+
		"\u0006\u0003\u0000\u00a2\u00a3\u0005\b\u0000\u0000\u00a3\u00a4\u0003\u0012"+
		"\t\u0000\u00a4\u00a5\u0005\b\u0000\u0000\u00a5\u00a6\u0003\b\u0004\u0000"+
		"\u00a6\u0019\u0001\u0000\u0000\u0000\u00a7\u00a8\u0005&\u0000\u0000\u00a8"+
		"\u00a9\u00051\u0000\u0000\u00a9\u00aa\u0005#\u0000\u0000\u00aa\u00ab\u0003"+
		"\u001e\u000f\u0000\u00ab\u001b\u0001\u0000\u0000\u0000\u00ac\u00ad\u0006"+
		"\u000e\uffff\uffff\u0000\u00ad\u00ae\u0005\u0004\u0000\u0000\u00ae\u00af"+
		"\u0003\u001c\u000e\u0000\u00af\u00b0\u0005\u0005\u0000\u0000\u00b0\u00b6"+
		"\u0001\u0000\u0000\u0000\u00b1\u00b2\u0005\u0011\u0000\u0000\u00b2\u00b6"+
		"\u0003\u001c\u000e\u0003\u00b3\u00b6\u0003\u001e\u000f\u0000\u00b4\u00b6"+
		"\u0003<\u001e\u0000\u00b5\u00ac\u0001\u0000\u0000\u0000\u00b5\u00b1\u0001"+
		"\u0000\u0000\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b5\u00b4\u0001"+
		"\u0000\u0000\u0000\u00b6\u00dd\u0001\u0000\u0000\u0000\u00b7\u00b8\n\u000f"+
		"\u0000\u0000\u00b8\u00b9\u0005\u0013\u0000\u0000\u00b9\u00dc\u0003\u001c"+
		"\u000e\u0010\u00ba\u00bb\n\u000e\u0000\u0000\u00bb\u00bc\u0005\u0012\u0000"+
		"\u0000\u00bc\u00dc\u0003\u001c\u000e\u000f\u00bd\u00be\n\r\u0000\u0000"+
		"\u00be\u00bf\u0005\u000f\u0000\u0000\u00bf\u00dc\u0003\u001c\u000e\u000e"+
		"\u00c0\u00c1\n\f\u0000\u0000\u00c1\u00c2\u0005\u0010\u0000\u0000\u00c2"+
		"\u00dc\u0003\u001c\u000e\r\u00c3\u00c4\n\u000b\u0000\u0000\u00c4\u00c5"+
		"\u0005\u0016\u0000\u0000\u00c5\u00dc\u0003\u001c\u000e\f\u00c6\u00c7\n"+
		"\n\u0000\u0000\u00c7\u00c8\u0005\u0017\u0000\u0000\u00c8\u00dc\u0003\u001c"+
		"\u000e\u000b\u00c9\u00ca\n\t\u0000\u0000\u00ca\u00cb\u0005\u0018\u0000"+
		"\u0000\u00cb\u00dc\u0003\u001c\u000e\n\u00cc\u00cd\n\b\u0000\u0000\u00cd"+
		"\u00ce\u0005\u0019\u0000\u0000\u00ce\u00dc\u0003\u001c\u000e\t\u00cf\u00d0"+
		"\n\u0007\u0000\u0000\u00d0\u00d1\u0005\u001a\u0000\u0000\u00d1\u00dc\u0003"+
		"\u001c\u000e\b\u00d2\u00d3\n\u0006\u0000\u0000\u00d3\u00d4\u0005\u001b"+
		"\u0000\u0000\u00d4\u00dc\u0003\u001c\u000e\u0007\u00d5\u00d6\n\u0005\u0000"+
		"\u0000\u00d6\u00d7\u0005\u001c\u0000\u0000\u00d7\u00dc\u0003\u001c\u000e"+
		"\u0006\u00d8\u00d9\n\u0004\u0000\u0000\u00d9\u00da\u0005\u001d\u0000\u0000"+
		"\u00da\u00dc\u0003\u001c\u000e\u0005\u00db\u00b7\u0001\u0000\u0000\u0000"+
		"\u00db\u00ba\u0001\u0000\u0000\u0000\u00db\u00bd\u0001\u0000\u0000\u0000"+
		"\u00db\u00c0\u0001\u0000\u0000\u0000\u00db\u00c3\u0001\u0000\u0000\u0000"+
		"\u00db\u00c6\u0001\u0000\u0000\u0000\u00db\u00c9\u0001\u0000\u0000\u0000"+
		"\u00db\u00cc\u0001\u0000\u0000\u0000\u00db\u00cf\u0001\u0000\u0000\u0000"+
		"\u00db\u00d2\u0001\u0000\u0000\u0000\u00db\u00d5\u0001\u0000\u0000\u0000"+
		"\u00db\u00d8\u0001\u0000\u0000\u0000\u00dc\u00df\u0001\u0000\u0000\u0000"+
		"\u00dd\u00db\u0001\u0000\u0000\u0000\u00dd\u00de\u0001\u0000\u0000\u0000"+
		"\u00de\u001d\u0001\u0000\u0000\u0000\u00df\u00dd\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e3\u0003$\u0012\u0000\u00e1\u00e3\u00051\u0000\u0000\u00e2\u00e0"+
		"\u0001\u0000\u0000\u0000\u00e2\u00e1\u0001\u0000\u0000\u0000\u00e3\u00e7"+
		"\u0001\u0000\u0000\u0000\u00e4\u00e6\u0003 \u0010\u0000\u00e5\u00e4\u0001"+
		"\u0000\u0000\u0000\u00e6\u00e9\u0001\u0000\u0000\u0000\u00e7\u00e5\u0001"+
		"\u0000\u0000\u0000\u00e7\u00e8\u0001\u0000\u0000\u0000\u00e8\u001f\u0001"+
		"\u0000\u0000\u0000\u00e9\u00e7\u0001\u0000\u0000\u0000\u00ea\u00f0\u0003"+
		"\"\u0011\u0000\u00eb\u00ec\u0005\u000e\u0000\u0000\u00ec\u00f0\u00051"+
		"\u0000\u0000\u00ed\u00ee\u0005\u000e\u0000\u0000\u00ee\u00f0\u0003$\u0012"+
		"\u0000\u00ef\u00ea\u0001\u0000\u0000\u0000\u00ef\u00eb\u0001\u0000\u0000"+
		"\u0000\u00ef\u00ed\u0001\u0000\u0000\u0000\u00f0!\u0001\u0000\u0000\u0000"+
		"\u00f1\u00f2\u0005\u0002\u0000\u0000\u00f2\u00f3\u0007\u0000\u0000\u0000"+
		"\u00f3\u00f4\u0005\u0003\u0000\u0000\u00f4#\u0001\u0000\u0000\u0000\u00f5"+
		"\u00f6\u00051\u0000\u0000\u00f6\u00f7\u0005\u0004\u0000\u0000\u00f7\u00f8"+
		"\u0003(\u0014\u0000\u00f8\u00f9\u0005\u0005\u0000\u0000\u00f9%\u0001\u0000"+
		"\u0000\u0000\u00fa\u00ff\u00051\u0000\u0000\u00fb\u00fc\u0005\t\u0000"+
		"\u0000\u00fc\u00fe\u00051\u0000\u0000\u00fd\u00fb\u0001\u0000\u0000\u0000"+
		"\u00fe\u0101\u0001\u0000\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000\u0000"+
		"\u00ff\u0100\u0001\u0000\u0000\u0000\u0100\u0103\u0001\u0000\u0000\u0000"+
		"\u0101\u00ff\u0001\u0000\u0000\u0000\u0102\u00fa\u0001\u0000\u0000\u0000"+
		"\u0102\u0103\u0001\u0000\u0000\u0000\u0103\'\u0001\u0000\u0000\u0000\u0104"+
		"\u0109\u0003\u001c\u000e\u0000\u0105\u0106\u0005\t\u0000\u0000\u0106\u0108"+
		"\u0003\u001c\u000e\u0000\u0107\u0105\u0001\u0000\u0000\u0000\u0108\u010b"+
		"\u0001\u0000\u0000\u0000\u0109\u0107\u0001\u0000\u0000\u0000\u0109\u010a"+
		"\u0001\u0000\u0000\u0000\u010a\u010d\u0001\u0000\u0000\u0000\u010b\u0109"+
		"\u0001\u0000\u0000\u0000\u010c\u0104\u0001\u0000\u0000\u0000\u010c\u010d"+
		"\u0001\u0000\u0000\u0000\u010d)\u0001\u0000\u0000\u0000\u010e\u010f\u0005"+
		"\u0004\u0000\u0000\u010f\u0110\u0003&\u0013\u0000\u0110\u0111\u0005\u0005"+
		"\u0000\u0000\u0111\u0112\u0005\u0001\u0000\u0000\u0112\u0113\u0003:\u001d"+
		"\u0000\u0113+\u0001\u0000\u0000\u0000\u0114\u011d\u0005\u0002\u0000\u0000"+
		"\u0115\u011a\u00032\u0019\u0000\u0116\u0117\u0005\t\u0000\u0000\u0117"+
		"\u0119\u00032\u0019\u0000\u0118\u0116\u0001\u0000\u0000\u0000\u0119\u011c"+
		"\u0001\u0000\u0000\u0000\u011a\u0118\u0001\u0000\u0000\u0000\u011a\u011b"+
		"\u0001\u0000\u0000\u0000\u011b\u011e\u0001\u0000\u0000\u0000\u011c\u011a"+
		"\u0001\u0000\u0000\u0000\u011d\u0115\u0001\u0000\u0000\u0000\u011d\u011e"+
		"\u0001\u0000\u0000\u0000\u011e\u011f\u0001\u0000\u0000\u0000\u011f\u0120"+
		"\u0005\u0003\u0000\u0000\u0120-\u0001\u0000\u0000\u0000\u0121\u012a\u0005"+
		"\u0006\u0000\u0000\u0122\u0127\u00036\u001b\u0000\u0123\u0124\u0005\t"+
		"\u0000\u0000\u0124\u0126\u00036\u001b\u0000\u0125\u0123\u0001\u0000\u0000"+
		"\u0000\u0126\u0129\u0001\u0000\u0000\u0000\u0127\u0125\u0001\u0000\u0000"+
		"\u0000\u0127\u0128\u0001\u0000\u0000\u0000\u0128\u012b\u0001\u0000\u0000"+
		"\u0000\u0129\u0127\u0001\u0000\u0000\u0000\u012a\u0122\u0001\u0000\u0000"+
		"\u0000\u012a\u012b\u0001\u0000\u0000\u0000\u012b\u012c\u0001\u0000\u0000"+
		"\u0000\u012c\u012d\u0005\u0007\u0000\u0000\u012d/\u0001\u0000\u0000\u0000"+
		"\u012e\u0134\u00052\u0000\u0000\u012f\u0130\u0005\u0002\u0000\u0000\u0130"+
		"\u0131\u0003\u001e\u000f\u0000\u0131\u0132\u0005\u0003\u0000\u0000\u0132"+
		"\u0134\u0001\u0000\u0000\u0000\u0133\u012e\u0001\u0000\u0000\u0000\u0133"+
		"\u012f\u0001\u0000\u0000\u0000\u01341\u0001\u0000\u0000\u0000\u0135\u013a"+
		"\u0003$\u0012\u0000\u0136\u013a\u0003\u001c\u000e\u0000\u0137\u013a\u0003"+
		".\u0017\u0000\u0138\u013a\u0003,\u0016\u0000\u0139\u0135\u0001\u0000\u0000"+
		"\u0000\u0139\u0136\u0001\u0000\u0000\u0000\u0139\u0137\u0001\u0000\u0000"+
		"\u0000\u0139\u0138\u0001\u0000\u0000\u0000\u013a3\u0001\u0000\u0000\u0000"+
		"\u013b\u013c\u00030\u0018\u0000\u013c\u013d\u0005\r\u0000\u0000\u013d"+
		"\u013e\u00032\u0019\u0000\u013e5\u0001\u0000\u0000\u0000\u013f\u0142\u0003"+
		"4\u001a\u0000\u0140\u0142\u00038\u001c\u0000\u0141\u013f\u0001\u0000\u0000"+
		"\u0000\u0141\u0140\u0001\u0000\u0000\u0000\u01427\u0001\u0000\u0000\u0000"+
		"\u0143\u0144\u0005\u000e\u0000\u0000\u0144\u0145\u0005\u000e\u0000\u0000"+
		"\u0145\u0146\u0005\u000e\u0000\u0000\u0146\u0147\u0003\u001e\u000f\u0000"+
		"\u01479\u0001\u0000\u0000\u0000\u0148\u0149\u0005\u0006\u0000\u0000\u0149"+
		"\u014a\u0003\u0002\u0001\u0000\u014a\u014b\u0005\u0007\u0000\u0000\u014b"+
		";\u0001\u0000\u0000\u0000\u014c\u014d\u0007\u0001\u0000\u0000\u014d=\u0001"+
		"\u0000\u0000\u0000\u001a?FY_z\u007f\u008c\u0090\u009c\u00b5\u00db\u00dd"+
		"\u00e2\u00e7\u00ef\u00ff\u0102\u0109\u010c\u011a\u011d\u0127\u012a\u0133"+
		"\u0139\u0141";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}