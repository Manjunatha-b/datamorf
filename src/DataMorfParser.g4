parser grammar DataMorfParser;

options {
	tokenVocab = DataMorfLexer;
	superClass = DataMorfParserBase;
}

program: sourceElements? EOF;

sourceElement: statement;

statement:
	block
	| variableStatement
	| emptyStatement_
	| functionDeclaration
	| expressionStatement
	| ifStatement
	| iterationStatement
	| continueStatement
	| breakStatement
	| returnStatement
	| labelledStatement
	| switchStatement;

block: '{' statementList? '}';

statementList: statement+;

declaration: variableStatement | functionDeclaration;

variableStatement: variableDeclarationList eos;

variableDeclarationList:
	varModifier variableDeclaration (',' variableDeclaration)*;

variableDeclaration:
	assignable ('=' singleExpression)?; // ECMAScript 6: Array & Object Matching

emptyStatement_: SemiColon;

expressionStatement:
	{this.notOpenBraceAndNotFunction()}? expressionSequence eos;

ifStatement:
	If '(' expressionSequence ')' statement (Else statement)?;

iterationStatement:
	Do statement While '(' expressionSequence ')' eos	# DoStatement
	| While '(' expressionSequence ')' statement		# WhileStatement
	| For '(' (expressionSequence | variableDeclarationList)? ';' expressionSequence? ';'
		expressionSequence? ')' statement														# ForStatement
	| For '(' (singleExpression | variableDeclarationList) In expressionSequence ')' statement	#
		ForInStatement;

varModifier: Var; // let, const - ECMAScript 6

continueStatement:
	Continue ({this.notLineTerminator()}? identifier)? eos;

breakStatement:
	Break ({this.notLineTerminator()}? identifier)? eos;

returnStatement:
	Return ({this.notLineTerminator()}? expressionSequence)? eos;

switchStatement: Switch '(' expressionSequence ')' caseBlock;
caseBlock: '{' caseClauses? (defaultClause caseClauses?)? '}';
caseClauses: caseClause+;
caseClause: Case expressionSequence ':' statementList?;
defaultClause: Default ':' statementList?;

labelledStatement: identifier ':' statement;

functionDeclaration:
	Function_ '*'? identifier '(' formalParameterList? ')' functionBody;

methodDefinition: ({this.notLineTerminator()}?)? propertyName '(' formalParameterList? ')'
		functionBody
	| getter '(' ')' functionBody
	| setter '(' formalParameterList? ')' functionBody;

formalParameterList:
	formalParameterArg (',' formalParameterArg)* (
		',' lastFormalParameterArg
	)?
	| lastFormalParameterArg;

formalParameterArg:
	assignable ('=' singleExpression)?; // ECMAScript 6: Initialization

lastFormalParameterArg:
	Ellipsis singleExpression; // ECMAScript 6: Rest Parameter

functionBody: '{' sourceElements? '}';

sourceElements: sourceElement+;

arrayLiteral: ('[' elementList ']');

// JavaScript supports arrasys like [,,1,2,,].
elementList:
	','* arrayElement? (','+ arrayElement)* ','*; // Yes, everything is optional

arrayElement: Ellipsis? singleExpression;

propertyAssignment:
	propertyName ':' singleExpression					# PropertyExpressionAssignment
	| '[' singleExpression ']' ':' singleExpression		# ComputedPropertyExpressionAssignment
	| getter '(' ')' functionBody						# PropertyGetter
	| setter '(' formalParameterArg ')' functionBody	# PropertySetter
	| Ellipsis? singleExpression						# PropertyShorthand;

propertyName:
	identifierName
	| StringLiteral
	| numericLiteral
	| '[' singleExpression ']';

arguments: '(' (argument (',' argument)* ','?)? ')';

argument: Ellipsis? (singleExpression | identifier);

expressionSequence: singleExpression (',' singleExpression)*;

singleExpression:
	anonymousFunction									# FunctionExpression
	| singleExpression '?.' singleExpression			# OptionalChainExpression
	| singleExpression '?.'? '[' expressionSequence ']'	# MemberIndexExpression
	| singleExpression '?'? '.'? identifierName			# MemberDotExpression
	// Split to try `new Date()` first, then `new Date`.
	| New identifier arguments												# NewExpression
	| New singleExpression arguments										# NewExpression
	| New singleExpression													# NewExpression
	| singleExpression arguments											# ArgumentsExpression
	| New '.' identifier													# MetaExpression // new.target
	| singleExpression {this.notLineTerminator()}? '++'						# PostIncrementExpression
	| singleExpression {this.notLineTerminator()}? '--'						# PostDecreaseExpression
	| Delete singleExpression												# DeleteExpression
	| Typeof singleExpression												# TypeofExpression
	| '++' singleExpression													# PreIncrementExpression
	| '--' singleExpression													# PreDecreaseExpression
	| '+' singleExpression													# UnaryPlusExpression
	| '-' singleExpression													# UnaryMinusExpression
	| '!' singleExpression													# NotExpression
	| <assoc = right> singleExpression '**' singleExpression				# PowerExpression
	| singleExpression ('*' | '/' | '%') singleExpression					# MultiplicativeExpression
	| singleExpression ('+' | '-') singleExpression							# AdditiveExpression
	| singleExpression ('<<' | '>>' | '>>>') singleExpression				# BitShiftExpression
	| singleExpression ('<' | '>' | '<=' | '>=') singleExpression			# RelationalExpression
	| singleExpression InstanceOf singleExpression							# InstanceofExpression
	| singleExpression In singleExpression									# InExpression
	| singleExpression ('==' | '!=') singleExpression						# EqualityExpression
	| singleExpression '&' singleExpression									# BitAndExpression
	| singleExpression '^' singleExpression									# BitXOrExpression
	| singleExpression '|' singleExpression									# BitOrExpression
	| singleExpression '&&' singleExpression								# LogicalAndExpression
	| singleExpression '||' singleExpression								# LogicalOrExpression
	| singleExpression '?' singleExpression ':' singleExpression			# TernaryExpression
	| <assoc = right> singleExpression '=' singleExpression					# AssignmentExpression
	| <assoc = right> singleExpression assignmentOperator singleExpression	#
		AssignmentOperatorExpression
	| identifier					# IdentifierExpression
	| literal						# LiteralExpression
	| arrayLiteral					# ArrayLiteralExpression
	| objectLiteral					# ObjectLiteralExpression
	| '(' expressionSequence ')'	# ParenthesizedExpression;

assignable: identifier | keyword | arrayLiteral | objectLiteral;

objectLiteral:
	'{' (propertyAssignment (',' propertyAssignment)* ','?)? '}';

anonymousFunction: functionDeclaration # NamedFunction;

arrowFunctionParameters:
	propertyName
	| '(' formalParameterList? ')';

arrowFunctionBody: singleExpression | functionBody;

assignmentOperator:
	'*='
	| '/='
	| '%='
	| '+='
	| '-='
	| '<<='
	| '>>='
	| '>>>='
	| '&='
	| '^='
	| '|='
	| '**='
	| '??=';

literal:
	NullLiteral
	| BooleanLiteral
	| StringLiteral
	| RegularExpressionLiteral
	| numericLiteral
	| bigintLiteral;

templateStringAtom:
	TemplateStringAtom
	| TemplateStringStartExpression singleExpression TemplateCloseBrace;

numericLiteral: DecimalLiteral;

bigintLiteral: BigDecimalIntegerLiteral;

getter: {this.n("get")}? identifier propertyName;

setter: {this.n("set")}? identifier propertyName;

identifierName: identifier | reservedWord;

identifier: Identifier | As | From | Of;

reservedWord: keyword | NullLiteral | BooleanLiteral;

keyword:
	Break
	| Do
	| Typeof
	| Case
	| Else
	| New
	| Var
	| Return
	| Continue
	| For
	| Switch
	| While
	| Function_
	| Default
	| If
	| Delete
	| In
	| From
	| As
	| Of;

eos:
	SemiColon
	| EOF
	| {this.lineTerminatorAhead()}?
	| {this.closeBrace()}?;

// options { superClass = JavaScriptLexerBase; }