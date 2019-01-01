package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			fmt.Println("Name:", v.Name)
			fmt.Println("Pos:", v.Pos())
			fmt.Println(fset.Position(v.Pos()))
		}
		return true
	})
}
