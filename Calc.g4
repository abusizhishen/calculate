grammar Calc;

NUMBER:[0-9]'.'?[0-9]+ |[0-9];

Add:'+';
Sub:'-';
CHENG:'*';
CHU:'/';

WS: [ \t\n\r]+ ->skip;

expr
    :expr op = (CHENG|CHU) expr #CHENGCHU
    |expr op = (Add|Sub) expr #AddSub
    |NUMBER #NUMBER
    ;

start: expr EOF;