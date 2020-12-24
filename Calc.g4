grammar Calc;

Id: [a-zA-Z_][A-Za-z0-9_]*;
Int:[0-9]'.'?[0-9]+ |[0-9];

AddSub:'+'|'-';
EQ: '=';
CHENGCHU:'*'|'/';

WS: ' '+ ->skip;
NL: '\n';

expr
    :expr CHENGCHU expr
    |expr AddSub expr
    |'(' expr ')'
    |Id
    |Int
    |Id EQ expr
    ;

row
    : expr NL;
last
    : expr EOF;
file
    : row* last;