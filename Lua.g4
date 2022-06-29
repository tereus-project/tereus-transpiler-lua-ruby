grammar Lua;

translation: chunk;

chunk: statement*;

statement:
	variableDeclaration
	| ifStatement
	| whileStatement
	| numericForStatement;

variableDeclaration: Identifier Equal expression;

expression:
	(FloatConstant | IntegerConstant)		# ExpressionConstant
	| expression binaryOperator expression	# ExpressionBinary;

binaryOperator: Add | Subtract | Multiply | Divide;

ifStatement:
	If expression Then chunk (elseifStatement)* (elseStatement)? End;
elseifStatement: ElseIf expression chunk;
elseStatement: Else chunk;

whileStatement: While expression Do chunk End;

numericForStatement:
	For Identifier Equal expression Comma expression (
		Comma expression
	)? Do chunk End;

Equal: '=';
Add: '+';
Subtract: '-';
Multiply: '*';
Divide: '/';
Comma: ',';

And: 'and';
Break: 'break';
Do: 'do';
Else: 'else';
ElseIf: 'elseif';
End: 'end';
False: 'false';
For: 'for';
Function: 'function';
If: 'if';
In: 'in';
Local: 'local';
Nil: 'nil';
Not: 'not';
Or: 'or';
Repeat: 'repeat';
Return: 'return';
Then: 'then';
True: 'true';
Until: 'until';
While: 'while';

Identifier: [a-zA-Z_] [a-zA-Z0-9_]*;

FloatConstant: [0-9]+ '.' [0-9]+;
IntegerConstant: [1-9] [0-9]*;

BlockComment: '--[[' .*? '--]]' -> skip;
LineComment: '--' ~[\r\n]* -> skip;

Whitespace: [ \t]+ -> skip;
Newline: ( '\r' '\n'? | '\n') -> skip;