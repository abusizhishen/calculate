grammar Cacl;

Id: [a-zA-Z_][A-Za-z0-9_]*;
Int:[0-9]+;

Add:'+';
Sub: '-';
EQ: '=';
CHENG:'*';
CHU:'/';

NL: '\n';
HUANHANG: '\r?\n';
WS: [\r\n ];

op: Add
    |Sub
    |EQ
    |CHENG
    |CHU
    ;

expr:
    |Int
    |setVal NL expr
    |addVal
    ;
addVal: (Id|Int) op (Id|Int);
setVal: Id op Int;
init:expr;