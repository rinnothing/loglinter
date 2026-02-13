package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

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

		if !isLogCall(pass, funcCall) {
			return
		}

		checkLowerFirst(pass, funcCall)
		checkEnglish(pass, funcCall)
		checkNoSpecialSymbols(pass, funcCall)
		checkNoSensitiveData(pass, funcCall)
	})

	return nil, nil
}

func isLogCall(pass *analysis.Pass, funcCall *ast.CallExpr) bool {
	selector, ok := funcCall.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	if onCall, ok := selector.X.(*ast.Ident); ok {
		onCallObj := pass.TypesInfo.ObjectOf(onCall)

		if onCallPkgName, ok := onCallObj.(*types.PkgName); ok {
			pkgPath := onCallPkgName.Imported().Path()
			if pkgPath == "log/slog" {
				return isLogMethod(selector.Sel.Name)
			}
		}

		if onCallVar, ok := onCallObj.(*types.Var); ok {
			varType := onCallVar.Type().String()
			if varType == "*go.uber.org/zap.Logger" ||
				varType == "*go.uber.org/zap.SugaredLogger" ||
				varType == "*log/slog.Logger" {
				return isLogMethod(selector.Sel.Name)
			}
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

func getFirstArgString(pass *analysis.Pass, funcCall *ast.CallExpr) (string, token.Pos) {
	if len(funcCall.Args) == 0 {
		return "", 0
	}

	fst := funcCall.Args[0]
	strLit, ok := fst.(*ast.BasicLit)
	if !ok {
		return "", 0
	}

	str, err := strconv.Unquote(strLit.Value)
	if err != nil {
		pass.Reportf(strLit.ValuePos, "failed to unquote the string literal: '%s'", strLit.Value)
		return "", 0
	}

	return str, strLit.ValuePos
}

func checkLowerFirst(pass *analysis.Pass, funcCall *ast.CallExpr) {
	fst, pos := getFirstArgString(pass, funcCall)
	if fst == "" {
		return
	}

	chr, _ := utf8.DecodeRuneInString(fst)
	if chr == utf8.RuneError || unicode.IsLower(chr) {
		return
	}

	pass.Reportf(pos, "log messages should start from lowercase letter, but %10q doesn't", fst)
}

func checkEnglish(pass *analysis.Pass, funcCall *ast.CallExpr) {
	fst, pos := getFirstArgString(pass, funcCall)
	if fst == "" {
		return
	}

	for i, ch := range fst {
		if !unicode.Is(unicode.Latin, ch) {
			pass.Reportf(pos+token.Pos(i), "log messages should be english only, but %d char in %10q isn't", i+1, fst)
			return
		}
	}
}

func checkNoSpecialSymbols(pass *analysis.Pass, funcCall *ast.CallExpr) {
	fst, pos := getFirstArgString(pass, funcCall)
	if fst == "" {
		return
	}

	for i, ch := range fst {
		if !(unicode.IsSpace(ch) || unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '%' || ch == '=' || ch == '-') {
			pass.Reportf(pos+token.Pos(i), "log messages shouldn't contain special symbols or emoji but %d char in %10q seem to break the rule", i+1, fst)
			return
		}
	}
}

var sensitiveData = []string{
	"password=",
	"key=",
	"token=",
}

func checkNoSensitiveData(pass *analysis.Pass, funcCall *ast.CallExpr) {
	fst, pos := getFirstArgString(pass, funcCall)
	if fst == "" {
		return
	}

	for _, sens := range sensitiveData {
		idx := strings.Index(fst, sens)
		if idx != -1 {
			pass.Reportf(pos+token.Pos(idx), "log messages shouldn't contain sensitive information but there's %q in string %10q", sens, fst)
			return
		}
	}
}
