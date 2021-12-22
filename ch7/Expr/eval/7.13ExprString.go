package eval

import (
	"fmt"
	"strconv"
)

// 练习 7.13： 为Expr增加一个String方法来打印美观的语法树。当再一次解析的时候，检查它的结果是否生成相同的语法树。
// TODO: 没有处理表达式有括号的情况
//!-env

//!+Eval1

func (v Var) String(env Env) string {
	return strconv.FormatFloat(env[v], 'g', -1, 64)
}

func (l literal) String(_ Env) string {
	return  strconv.FormatFloat(float64(l), 'g', -1, 64)
}

func (u unary) String(env Env) string {
	switch u.op {
	case '+':
		return "+" + u.x.String(env)
	case '-':
		return "-" + u.x.String(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String(env Env) string {
	switch b.op {
	case '+':
		return b.x.String(env) + "+" + b.y.String(env)
	case '-':
		return b.x.String(env) + "-" + b.y.String(env)
	case '*':
		return b.x.String(env) + "*" + b.y.String(env)
	case '/':
		return b.x.String(env) + "/" + b.y.String(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) String(env Env) string {
	switch c.fn {
	case "pow":
		return "Pow(" + c.args[0].String(env) +  ", " + c.args[1].String(env) + ")"
	case "sin":
		return "Sin(" + c.args[0].String(env) + ")"
	case "sqrt":
		return "Sqrt(" + c.args[0].String(env) + ")"
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}