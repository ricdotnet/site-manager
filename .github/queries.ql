import javascript

from ExprStmt es
where es.getExpr().(CallExpr).getCallee().toString() = "console.log"
select es