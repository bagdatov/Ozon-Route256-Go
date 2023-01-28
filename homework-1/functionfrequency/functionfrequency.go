package functionfrequency

import (
	"go/ast"
	"go/parser"
	"go/token"
	"sort"
)

type FunctionCall struct {
	Name  string
	Count int
}

// This is to implement sort data interface

type List []FunctionCall

func (l List) Len() int {
	return len(l)
}
func (l List) Less(i, j int) bool {
	return l[i].Count < l[j].Count
}
func (l List) Swap(i, j int) {
	l[j], l[i] = l[i], l[j]
}

func FunctionFrequency(gocode []byte) []string {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", gocode, 0)
	if err != nil {
		return nil
	}

	frequencies := make(map[string]int)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {

		case *ast.CallExpr:
			switch tt := x.Fun.(type) {

			case *ast.Ident:
				name := tt.Name
				frequencies[name] = frequencies[name] + 1

			case *ast.SelectorExpr:
				switch xx := tt.X.(type) {
				case *ast.Ident:
					name := xx.Name + "." + tt.Sel.Name
					frequencies[name] = frequencies[name] + 1
				}
			}
		}
		return true
	})

	var Calls List

	for k, v := range frequencies {
		Calls = append(Calls, FunctionCall{
			Name:  k,
			Count: v,
		})
	}
	sort.Sort(sort.Reverse(Calls))

	var funcNames []string

	for i, v := range Calls {
		funcNames = append(funcNames, v.Name)

		if i > 1 {
			break
		}
	}

	return funcNames
}
