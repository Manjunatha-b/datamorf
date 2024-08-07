/*
 * The MIT License (MIT)
 
 Copyright (c) 2014 by Bart Kiers (original author) and Alexandre
 * Vitorelli (contributor -> ported to CSharp)
 Copyright (c) 2017-2020 by Ivan Kochurkin (Positive
 * Technologies):
 added ECMAScript 6 support, cleared and transformed to the universal grammar.
 * Copyright (c) 2018 by Juan Alvarez (contributor -> ported to Go)
 Copyright (c) 2019 by Student
 * Main (contributor -> ES2020)
 
 Permission is hereby granted, free of charge, to any person
 * obtaining a copy of this software and associated documentation
 files (the "Software"), to deal
 * in the Software without
 restriction, including without limitation the rights to use,
 copy,
 * modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to
 * permit persons to whom the
 Software is furnished to do so, subject to the following
 * conditions:
 
 The above copyright notice and this permission notice shall be
 included in all
 * copies or substantial portions of the Software.
 
 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT
 * WARRANTY OF ANY KIND,
 EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
 OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT
 HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 WHETHER IN
 * AN
 ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE
 OR THE USE OR
 OTHER DEALINGS IN THE SOFTWARE.
 */

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar DataMorfLexer;

channels {
    ERROR
}

options {
    superClass = DataMorfLexerBase;
}

MultiLineComment: '/*' .*? '*/' -> channel(HIDDEN);

SingleLineComment: '//' ~[\r\n\u2028\u2029]* -> channel(HIDDEN);

RegularExpressionLiteral:
    '/' RegularExpressionFirstChar RegularExpressionChar* {this.IsRegexPossible()}? '/' IdentifierPart*
;

OpenBracket: '[';

CloseBracket: ']';

OpenParen: '(';

CloseParen: ')';

OpenBrace: '{' {this.ProcessOpenBrace();};

TemplateCloseBrace:
    {this.IsInTemplateString()}? '}'
    // Break lines here to ensure proper transformation by Go/transformGrammar.py
    {this.ProcessTemplateCloseBrace();} -> popMode
;

CloseBrace: '}' {this.ProcessCloseBrace();};

SemiColon: ';';

Comma: ',';

Assign: '=';

QuestionMark: '?';

QuestionMarkDot: '?.';

Colon: ':';

Ellipsis: '...';

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

RightShiftArithmetic: '>>';

LeftShiftArithmetic: '<<';

RightShiftLogical: '>>>';

LessThan: '<';

MoreThan: '>';

LessThanEquals: '<=';

GreaterThanEquals: '>=';

Equals_: '==';

NotEquals: '!=';

BitAnd: '&';

BitXOr: '^';

BitOr: '|';

And: '&&';

Or: '||';

MultiplyAssign: '*=';

DivideAssign: '/=';

ModulusAssign: '%=';

PlusAssign: '+=';

MinusAssign: '-=';

LeftShiftArithmeticAssign: '<<=';

RightShiftArithmeticAssign: '>>=';

RightShiftLogicalAssign: '>>>=';

BitAndAssign: '&=';

BitXorAssign: '^=';

BitOrAssign: '|=';

PowerAssign: '**=';

NullishCoalescingAssign: '??=';

ARROW: '=>';

NullLiteral: 'null';

BooleanLiteral: 'true' | 'false';

DecimalLiteral:
    DecimalIntegerLiteral '.' [0-9] [0-9_]* ExponentPart?
    | '.' [0-9] [0-9_]* ExponentPart?
    | DecimalIntegerLiteral ExponentPart?
;

BigDecimalIntegerLiteral: DecimalIntegerLiteral 'n';

StringLiteral: ('"' DoubleStringCharacter* '"' | '\'' SingleStringCharacter* '\'') {this.ProcessStringLiteral();}
;

/// Keywords
Break: 'break';

Do: 'do';

Typeof: 'typeof';

Case: 'case';

Else: 'else';

New: 'new';

Var: 'var';

Return: 'return';

Continue: 'continue';

For: 'for';

Switch: 'switch';

While: 'while';

Function_: 'function';

Default: 'default';

If: 'if';

InstanceOf: 'instanceof';

Delete: 'delete';

In: 'in';

As: 'as';

From: 'from';

Of: 'of';

Identifier: IdentifierStart IdentifierPart*;

WhiteSpaces: [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN);

LineTerminator: [\r\n\u2028\u2029] -> channel(HIDDEN);

TemplateStringStartExpression: '${' {this.ProcessTemplateOpenBrace();} -> pushMode(DEFAULT_MODE);

TemplateStringAtom: ~[`];

fragment DoubleStringCharacter: ~["\\\r\n] | '\\' EscapeSequence | LineContinuation;

fragment SingleStringCharacter: ~['\\\r\n] | '\\' EscapeSequence | LineContinuation;

fragment EscapeSequence:
    CharacterEscapeSequence
    | '0' // no digit ahead! TODO
    | HexEscapeSequence
    | UnicodeEscapeSequence
    | ExtendedUnicodeEscapeSequence
;

fragment CharacterEscapeSequence: SingleEscapeCharacter | NonEscapeCharacter;

fragment HexEscapeSequence: 'x' HexDigit HexDigit;

fragment UnicodeEscapeSequence:
    'u' HexDigit HexDigit HexDigit HexDigit
    | 'u' '{' HexDigit HexDigit+ '}'
;

fragment ExtendedUnicodeEscapeSequence: 'u' '{' HexDigit+ '}';

fragment SingleEscapeCharacter: ['"\\bfnrtv];

fragment NonEscapeCharacter: ~['"\\bfnrtv0-9xu\r\n];

fragment EscapeCharacter: SingleEscapeCharacter | [0-9] | [xu];

fragment LineContinuation: '\\' [\r\n\u2028\u2029]+;

fragment HexDigit: [_0-9a-fA-F];

fragment DecimalIntegerLiteral: '0' | [1-9] [0-9_]*;

fragment ExponentPart: [eE] [+-]? [0-9_]+;

fragment IdentifierPart: IdentifierStart | [\p{Mn}] | [\p{Nd}] | [\p{Pc}] | '\u200C' | '\u200D';

fragment IdentifierStart: [\p{L}] | [$_] | '\\' UnicodeEscapeSequence;

fragment RegularExpressionFirstChar:
    ~[*\r\n\u2028\u2029\\/[]
    | RegularExpressionBackslashSequence
    | '[' RegularExpressionClassChar* ']'
;

fragment RegularExpressionChar:
    ~[\r\n\u2028\u2029\\/[]
    | RegularExpressionBackslashSequence
    | '[' RegularExpressionClassChar* ']'
;

fragment RegularExpressionClassChar: ~[\r\n\u2028\u2029\]\\] | RegularExpressionBackslashSequence;

fragment RegularExpressionBackslashSequence: '\\' ~[\r\n\u2028\u2029];