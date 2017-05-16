package sexpconv

import (
	"fmt"
	"go/ast"
	"go/types"
	"lisp/function"
	"sexp"
)

func (conv *Converter) call(fn *function.Type, args ...ast.Expr) *sexp.Call {
	return &sexp.Call{
		Fn:   fn,
		Args: conv.exprList(args),
	}
}

func (conv *Converter) CallExpr(node *ast.CallExpr) sexp.Form {
	// #REFS: 2.
	switch fn := node.Fun.(type) {
	case *ast.SelectorExpr: // x.sel()
		sel := conv.info.Selections[fn]
		if sel != nil {
			panic("method calls unimplemented")
		}

		pkg := fn.X.(*ast.Ident)
		if pkg.Name == "lisp" {
			return conv.intrinFuncCall(fn.Sel.Name, node.Args)
		}

		return conv.call(conv.makeFunction(fn.Sel, pkg.Name), node.Args...)

	case *ast.Ident: // f()
		switch fn.Name {
		case "make":
			return conv.makeBuiltin(node.Args)
		case "panic":
			return &sexp.Panic{ErrorData: conv.Expr(node.Args[0])}
		case "int", "string", "float64":
			return conv.Expr(node.Args[0])
		case "print":
			return conv.call(&function.Print, node.Args...)
		case "println":
			return conv.call(&function.Println, node.Args...)
		default:
			return conv.call(conv.makeFunction(fn, ""), node.Args...)
		}

	default:
		panic(fmt.Sprintf("unexpected func: %#v", node.Fun))
	}
}

func (conv *Converter) makeFunction(fn *ast.Ident, pkgName string) *function.Type {
	// #FIXME: can also be *types.Names (type cast), etc.
	sig := conv.typeOf(fn).(*types.Signature)

	if pkgName == "" {
		qualName := conv.symPrefix + fn.Name
		return function.New(qualName, sig)
	}
	qualName := "Go-" + pkgName + "." + fn.Name
	return function.New(qualName, sig)
}