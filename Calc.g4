grammar Calc;

Id: [_a-zA-Z]+;
NUMBER:[0-9];

Add:'+';
Sub:'-';
CHENG:'*';
CHU:'/';
EQ: '=';

WS: [ \t]+ ->skip;

singleValue
    :Id     #ID
    |NUMBER #OUT
    ;

//expr
//    :expr op = (CHENG|CHU) expr #CHENGCHU
//    |expr op = (Add|Sub) expr #ADDSUB
//    |NUMBER #NUMBER
//    |'(' expr ')' #PARENTHES
//    ;

setVal
    :id=Id EQ value=NUMBER;

jia: (Id|NUMBER) Add (Id|NUMBER);

row:(setVal|singleValue|jia) ('\r'?'\n'|EOF);

start: row*;