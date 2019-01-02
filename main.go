package main

import (
	"fmt"
	"log"

	"golang.org/x/tools/go/ast/astutil"

	"go/parser"
)

func main() {
	expr, err := parser.ParseExpr(`func(x, y int){}(10, 20)`)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	n := astutil.Apply(expr, func(cr *astutil.Cursor) bool {
		if cr.Name() == "Args" {
			fmt.Println(cr.Name(), cr.Index())
		}
		return true
	}, nil)

	_ = n
}
