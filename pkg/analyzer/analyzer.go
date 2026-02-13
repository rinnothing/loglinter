package analyzer

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "gologlinter",
	Doc:  "Checks that log calls are formatted correctly",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		funcCall := node.(*ast.CallExpr)

		if isLogCall(pass, funcCall) {
			fmt.Println("is logging")
		} else {
			fmt.Println("is not logging")
		}
	})

	return nil, nil
}

func isLogCall(pass *analysis.Pass, funcCall *ast.CallExpr) bool {
	// fmt.Println(reflect.TypeOf(funcCall.Fun).Elem())
	selector, ok := funcCall.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	fmt.Println(selector.Sel.Name)

	// fmt.Println(selector.X)
	// fmt.Println(reflect.TypeOf(selector.X).Elem())

	if onCall, ok := selector.X.(*ast.Ident); ok {
		onCallObj := pass.TypesInfo.ObjectOf(onCall)

		if onCallPkgName, ok := onCallObj.(*types.PkgName); ok {
			pkgPath := onCallPkgName.Imported().Path()
			if pkgPath == "log/slog" {
				return isLogMethod(selector.Sel.Name)
			}
		}

		if onCallVar, ok := onCallObj.(*types.Var); ok {
			// fmt.Println(onCallVar.Type().String())

			varType := onCallVar.Type().String()
			if varType == "*go.uber.org/zap.Logger" ||
				varType == "*go.uber.org/zap.SugaredLogger" ||
				varType == "*log/slog.Logger" {
				return isLogMethod(selector.Sel.Name)
			}

			// if structType, ok := onCallVar.Type().(*types.Pointer); ok {
			// 	fmt.Println(structType.Elem().String())
			// 	fmt.Println(reflect.TypeOf(structType.Elem()).Elem())
			// 	if namedType, ok := structType.Elem().(*types.Named); ok {
			// 		fmt.Println(namedType.)
			// 	}
			// }
		}

	}
	if onFuncCall, ok := selector.X.(*ast.CallExpr); ok {
		if onFuncSel, ok := onFuncCall.Fun.(*ast.SelectorExpr); ok {
			if onFuncSelIdent, ok := onFuncSel.X.(*ast.Ident); ok {
				onFuncSelObj := pass.TypesInfo.ObjectOf(onFuncSelIdent)

				if onFuncSelPkgName, ok := onFuncSelObj.(*types.PkgName); ok {
					pkgPath := onFuncSelPkgName.Imported().Path()
					methodName := onFuncSel.Sel.Name
					if pkgPath == "go.uber.org/zap" && (methodName == "L" || methodName == "S") {
						return isLogMethod(selector.Sel.Name)
					}
				}
			}
		}
	}

	// fmt.Println(onCall.Name)

	return false
}

var logMethods = map[string]struct{}{
	"Debug": {},
	"Info":  {},
	"Warn":  {},
	"Error": {},
	"Fatal": {},
}

func isLogMethod(name string) bool {
	_, in := logMethods[name]
	return in
}
