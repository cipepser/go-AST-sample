package main

import (
	"fmt"
	"go/ast"
	"log"

	"go/parser"
)

func main() {
	expr, err := parser.ParseExpr("1+1")

	if err != nil {
		log.Fatalln("Error:", err)
	}

	ast.Inspect(expr, func(n ast.Node) bool {
		fmt.Printf("%[1]T %[1]v\n", n)
		return true
	})
}
