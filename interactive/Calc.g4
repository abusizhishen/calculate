grammar Calc;

Id: [_a-zA-Z]+;
NUMBER:[0-9];

Add:'+';
Sub:'-';
CHENG:'*';
CHU:'/';
EQ: '=';
POW: '^';

WS: [ \t]+ ->skip;
NL:('\r'?'\n' | EOF);

expr
    :expr op = (CHENG|CHU) expr #CHENGCHU
    |expr op = (Add|Sub) expr #ADDSUB
    |NUMBER #NUMBER
    |Id #ID
    |'(' expr ')' #PARENTHES
    |expr '^'<assoc=right> expr #POW
    ;

expression
    :Id EQ expr #SetVal
    |Id     #OUTID
    |NUMBER #OUTNUMBER
    |expr #EXPR
    ;

row :expression NL?;
rows:row*;