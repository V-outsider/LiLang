grammar LiLang;

VAR         : 'var' ;
FUNCTION    : 'function' ;
RETURN      : 'return' ;
IF          : 'if' ;
ELSE        : 'else' ;
SWITCH      : 'switch' ;
CASE        : 'case' ;
DEFAULT     : 'default' ;
INT         : 'int' ;
LIST        : 'list' ;
ADD         : 'add' ;
REMOVE      : 'remove' ;
GET         : 'get' ;
SIZE        : 'size' ;
PRINT       : 'print' ;

WS          : [ \t\r\n]+ -> skip ;
COMMENT     : '//' ~[\r\n]* -> skip ;
NUMBER      : [0-9]+ ;
IDENTIFIER  : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING      : '"' .*? '"' ;
PLUS        : '+' ;
MINUS       : '-' ;
MULT        : '*' ;
DIV         : '/' ;
ASSIGN      : '=' ;
LPAREN      : '(' ;
RPAREN      : ')' ;
LBRACE      : '{' ;
RBRACE      : '}' ;
COLON       : ':' ;
SEMICOLON   : ';' ;
COMMA       : ',' ;
AMPERSAND   : '&' ;
SLBRACE     : '[' ;
LLBRACE     : ']' ;
LT          : '<' ;
GT          : '>' ;

program         : statement* ;

statement       : variableDeclaration
                | functionDeclaration
                | assignment
                | expressionStatement
                | block
                | ifStatement
                | switchStatement
                | returnStatement
                | listOperation
                ;

variableDeclaration : VAR IDENTIFIER COLON type SEMICOLON ;

functionDeclaration : FUNCTION IDENTIFIER LPAREN paramList? RPAREN type block ;

printCall : PRINT LPAREN IDENTIFIER RPAREN;

paramList        : param (COMMA param)* ;
param            : IDENTIFIER COLON type (AMPERSAND)? ;

type            : INT
                | LIST LT type GT  // Определение типа для списков
                ;

assignment       : IDENTIFIER ASSIGN expression SEMICOLON ;

expressionStatement : expression SEMICOLON ;

block            : LBRACE statement* RBRACE ;

ifStatement      : IF LPAREN expression RPAREN block (ELSE block)? ;

switchStatement  : SWITCH expression LBRACE caseStatement* RBRACE ;
caseStatement    : CASE expression COLON statement* ;

listOperation    : IDENTIFIER '.' method LPAREN expression? RPAREN SEMICOLON ;

method          : ADD
                | REMOVE
                | GET
                | SIZE
                ;

returnStatement  : RETURN expression SEMICOLON ;

expression       : expression PLUS expression
                | expression MINUS expression
                | expression MULT expression
                | expression DIV expression
                | NUMBER
                | IDENTIFIER
                | LPAREN expression RPAREN ;
