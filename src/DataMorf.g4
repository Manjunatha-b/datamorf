grammar DataMorf;

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
assignStatement: reference Assign value;
functionStatement: Function Identifier OpenRound signatureParams CloseRound block;
ifStatement: If OpenRound unit CloseRound block (Else If OpenRound unit CloseRound block)*  (Else block)?;
switchStatement: Switch OpenRound unit CloseRound block;
deleteStatement: Delete (Identifier (Dot (accessProperty | Identifier))*);
unitStatement: unit;
valueStatement : value SemiColon;

forStatement: For OpenRound (classicForCondition | iteratorForCondition) CloseRound block;
classicForCondition: init = variableStatement SemiColon condition = unitStatement SemiColon postOp = assignStatement;
iteratorForCondition: Var Identifier Of reference;

unit:
	OpenRound single = unit CloseRound
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
	| reference
	| constant;

reference: (functionCall | Identifier) (accessor)*;
accessor: (accessProperty | (Dot Identifier) | (Dot functionCall));
accessProperty: OpenSquare (StringLiteral | DecimalLiteral | Identifier) CloseSquare;

functionCall: Identifier OpenRound (sendingParams) CloseRound;
signatureParams: (Identifier (Comma Identifier)*)?;
sendingParams: (unit (Comma unit)*)?;
arrowFunction: OpenRound signatureParams CloseRound '=>' block;

arrayLiteral: OpenSquare (value (Comma value)*)? CloseSquare;
objectLiteral: OpenCurly (objectItem (Comma (objectItem))*)? CloseCurly;
key: ( StringLiteral | OpenSquare reference CloseSquare);
value: (functionCall | unit | objectLiteral | arrayLiteral);
keyValue: key Colon value;
objectItem: (keyValue | spread);

spread: Dot Dot Dot reference;

block: OpenCurly statements CloseCurly;

constant:
	NullLiteral
	| BooleanLiteral
	| DecimalLiteral
	| StringLiteral;

OpenSquare: '[';
CloseSquare: ']';
OpenRound: '(';
CloseRound: ')';
OpenCurly: '{';
CloseCurly: '}';
SemiColon: ';';
Comma: ',';
Assign: '=';
QuestionMark: '?';
QuestionMarkDot: '?.';
Colon: ':';
Dot: '.';
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

NullLiteral: 'null';
BooleanLiteral: 'true' | 'false';

DecimalLiteral:
	DecimalIntegerLiteral '.' [0-9]*
	| DecimalIntegerLiteral;

/// Keywords
Break: 'break';
Do: 'do';
Of: 'of';
Case: 'case';
Else: 'else';
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

fragment DoubleStringCharacter: ~["\\\r\n] | LineContinuation;
fragment LineContinuation: '\\' [\r\n\u2028\u2029];
fragment DecimalIntegerLiteral: '0' | [1-9] [0-9_]*;