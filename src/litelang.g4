grammar LiteLang;

program: statements? EOF;
statements: statement+;

statement:
	variableStatement SemiColon
	| assignStatement SemiColon
	| functionStatement
	| ifStatement
	| forStatement
	| switchStatement
	| valueStatement
	| deleteStatement SemiColon
	| unitStatement SemiColon;

variableStatement: Var Identifier ( Assign value)?;
assignStatement: accessLhs Assign value;
functionStatement:
	Function Identifier OpenParen paramMaker CloseParen block;
ifStatement:
	If OpenParen unit CloseParen block (
		Else If OpenParen unit CloseParen block
	)*  (Else block)?;
forStatement:
	For OpenParen (classicForParam | iteratorForParam) CloseParen block;
classicForParam: (variableStatement | assignStatement) SemiColon condition = unitStatement SemiColon
		postOp = assignStatement;
iteratorForParam: Var Identifier Of accessRhs;
switchStatement: Switch OpenParen unit CloseParen block;
deleteStatement:
	Delete (Identifier (Dot (accessProperty | Identifier))*);
unitStatement: unit;
valueStatement : value SemiColon;

unit:
	OpenParen single = unit CloseParen
	| left = unit ( Divide) right = unit
	| left = unit ( Multiply) right = unit
	| left = unit ( Plus) right = unit
	| left = unit ( Minus) right = unit
	| left = unit (
		LessThan
		| MoreThan
		| LessThanEquals
		| GreaterThanEquals
		| Equals
		| NotEquals
		| And
		| Or
	) right = unit
	| Not single = unit
	| accessRhs
	| constant;

accessLhs: Identifier (accessorLhs)*;
accessRhs: (functionCall | Identifier) (accessorRhs)*;
accessorRhs: (
		accessProperty
		| (Dot Identifier)
		| (Dot functionCall)
	);
accessorLhs: accessProperty | (Dot Identifier);
accessProperty:
	OpenBracket (StringLiteral | DecimalLiteral | Identifier) CloseBracket;
functionCall: Identifier OpenParen (params) CloseParen;
paramMaker: (Identifier (Comma Identifier)*)?;
params: (unit (Comma unit)*)?;
arrowFunction: OpenParen params CloseParen '=>' block;

arrayLiteral: OpenBracket (value (Comma value)*)? CloseBracket;
objectLiteral:
	OpenBrace (objectItem (Comma (objectItem))*)? CloseBrace;
key: ( StringLiteral | OpenBracket accessRhs CloseBracket);
value: (functionCall | unit | objectLiteral | arrayLiteral);
keyValue: key Colon value;
spread: Dot Dot Dot accessRhs;
objectItem: (keyValue | spread);

block: OpenBrace statements CloseBrace;
constant:
	NullLiteral
	| BooleanLiteral
	| DecimalLiteral
	| StringLiteral;

OpenBracket: '[';
CloseBracket: ']';
OpenParen: '(';
CloseParen: ')';
OpenBrace: '{';
CloseBrace: '}';
SemiColon: ';';
Comma: ',';
Assign: '=';
QuestionMark: '?';
QuestionMarkDot: '?.';
Colon: ':';
Dot: '.';
PlusPlus: '++';
MinusMinus: '--';
Plus: '+';
Minus: '-';
Not: '!';
Multiply: '*';
Divide: '/';
Modulus: '%';
Power: '**';
LessThan: '<';
MoreThan: '>';
LessThanEquals: '<=';
GreaterThanEquals: '>=';
Equals: '==';
NotEquals: '!=';
And: '&&';
Or: '||';

/// Null Literals
NullLiteral: 'null';

/// Boolean Literals
BooleanLiteral: 'true' | 'false';

/// Numeric Literals
DecimalLiteral:
	DecimalIntegerLiteral '.' [0-9]*
	| DecimalIntegerLiteral;

/// Keywords
Break: 'break';
Do: 'do';
Of: 'of';
Case: 'case';
Else: 'else';
New: 'new';
Var: 'var';
Return: 'return';
Continue: 'continue';
For: 'for';
Switch: 'switch';
While: 'while';
Function: 'function';
Default: 'default';
If: 'if';
Elif: 'elif';
Delete: 'delete';

Identifier: [a-z][a-zA-Z_0-9]*;
StringLiteral: '"' DoubleStringCharacter* '"';

WhiteSpaces: [\t ]+ -> channel(HIDDEN);
LineTerminator: [\r\n ] -> channel(HIDDEN);

// Fragment rules
fragment DoubleStringCharacter: ~["\\\r\n] | LineContinuation;

fragment LineContinuation: '\\' [\r\n\u2028\u2029];

fragment DecimalIntegerLiteral: '0' | [1-9] [0-9_]*;