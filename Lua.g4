grammar Lua;

translation: topLevelStatement+;

topLevelStatement: variableDeclaration;

variableDeclaration: Identifier Equal expression;

expression:
	(FloatConstant | IntegerConstant)		# ExpressionConstant
	| expression binaryOperator expression	# ExpressionBinary;

binaryOperator: Add | Subtract | Multiply | Divide;

Equal: '=';
Add: '+';
Subtract: '-';
Multiply: '*';
Divide: '/';

Identifier: [a-zA-Z_] [a-zA-Z0-9_]*;

FloatConstant: [0-9]+ '.' [0-9]+;
IntegerConstant: [1-9] [0-9]*;

Whitespace: [ \t]+ -> skip;
Newline: ( '\r' '\n'? | '\n') -> skip;
