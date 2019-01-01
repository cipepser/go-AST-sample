# go-AST-sample

## 文字列として受け取ってASTを表示する

```go
package main

import (
    "go/ast"
    "go/parser"
)

func main() {
    expr, _ := parser.ParseExpr("A + 1")

    ast.Print(nil, expr)
}
```

結果

```sh
❯ go run main.go
     0  *ast.BinaryExpr {
     1  .  X: *ast.Ident {
     2  .  .  NamePos: 1
     3  .  .  Name: "A"
     4  .  .  Obj: *ast.Object {
     5  .  .  .  Kind: bad
     6  .  .  .  Name: ""
     7  .  .  }
     8  .  }
     9  .  OpPos: 3
    10  .  Op: +
    11  .  Y: *ast.BasicLit {
    12  .  .  ValuePos: 5
    13  .  .  Kind: INT
    14  .  .  Value: "1"
    15  .  }
    16  }
```

一瞬`A`を`string`としての`A`かと思ったが、よく考えると変数。
また、ちゃんと`1`も`INT`としてパースできていることがわかる。
`ParseExpr`のシグネチャは以下のようになっており、`string`を受け取る。

```go
func ParseExpr(x string) (ast.Expr, error) {
	return ParseExprFrom(token.NewFileSet(), "", []byte(x), 0)
}
```


## References
* [Go言語の golang/go パッケージで初めての構文解析](https://qiita.com/po3rin/items/a19d96d29284108ad442)
