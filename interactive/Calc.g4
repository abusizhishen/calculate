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

pow
    : nu=(Id|NUMBER) POW <assoc=right> fang=(Id|NUMBER)
    ;

expr
    :expr op = (CHENG|CHU) expr #CHENGCHU
    |expr op = (Add|Sub) expr #ADDSUB
    |NUMBER #NUMBER
    |Id #ID
    |'(' expr ')' #PARENTHES
    |pow #PPOW
    ;

expression
    :Id EQ NUMBER #SetVal
    |Id     #OUTID
    |NUMBER #OUTNUMBER
    |expr #EXPR
    ;

row :expression NL?;
rows:row*;