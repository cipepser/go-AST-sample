# go-AST-sample

##  [Go言語の golang/go パッケージで初めての構文解析](https://qiita.com/po3rin/items/a19d96d29284108ad442)

### 文字列として受け取ってASTを表示する

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

### ファイルからAST

以下の方法で`example.go`からASTが得られる。

```go
func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
	}
}
```

`example.go`は以下。

```go
package example

import "log"

func add(n, m int) {
	log.Println(n + m)
}
```


結果

```sh
❯ go run main.go
     0  *ast.GenDecl {
     1  .  TokPos: ./example/example.go:3:1
     2  .  Tok: import
     3  .  Lparen: -
     4  .  Specs: []ast.Spec (len = 1) {
     5  .  .  0: *ast.ImportSpec {
     6  .  .  .  Path: *ast.BasicLit {
     7  .  .  .  .  ValuePos: ./example/example.go:3:8
     8  .  .  .  .  Kind: STRING
     9  .  .  .  .  Value: "\"log\""
    10  .  .  .  }
    11  .  .  .  EndPos: -
    12  .  .  }
    13  .  }
    14  .  Rparen: -
    15  }
     0  *ast.FuncDecl {
     1  .  Name: *ast.Ident {
     2  .  .  NamePos: ./example/example.go:5:6
     3  .  .  Name: "add"
     4  .  .  Obj: *ast.Object {
     5  .  .  .  Kind: func
     6  .  .  .  Name: "add"
     7  .  .  .  Decl: *(obj @ 0)
     8  .  .  }
     9  .  }
    10  .  Type: *ast.FuncType {
    11  .  .  Func: ./example/example.go:5:1
    12  .  .  Params: *ast.FieldList {
    13  .  .  .  Opening: ./example/example.go:5:9
    14  .  .  .  List: []*ast.Field (len = 1) {
    15  .  .  .  .  0: *ast.Field {
    16  .  .  .  .  .  Names: []*ast.Ident (len = 2) {
    17  .  .  .  .  .  .  0: *ast.Ident {
    18  .  .  .  .  .  .  .  NamePos: ./example/example.go:5:10
    19  .  .  .  .  .  .  .  Name: "n"
    20  .  .  .  .  .  .  .  Obj: *ast.Object {
    21  .  .  .  .  .  .  .  .  Kind: var
    22  .  .  .  .  .  .  .  .  Name: "n"
    23  .  .  .  .  .  .  .  .  Decl: *(obj @ 15)
    24  .  .  .  .  .  .  .  }
    25  .  .  .  .  .  .  }
    26  .  .  .  .  .  .  1: *ast.Ident {
    27  .  .  .  .  .  .  .  NamePos: ./example/example.go:5:13
    28  .  .  .  .  .  .  .  Name: "m"
    29  .  .  .  .  .  .  .  Obj: *ast.Object {
    30  .  .  .  .  .  .  .  .  Kind: var
    31  .  .  .  .  .  .  .  .  Name: "m"
    32  .  .  .  .  .  .  .  .  Decl: *(obj @ 15)
    33  .  .  .  .  .  .  .  }
    34  .  .  .  .  .  .  }
    35  .  .  .  .  .  }
    36  .  .  .  .  .  Type: *ast.Ident {
    37  .  .  .  .  .  .  NamePos: ./example/example.go:5:15
    38  .  .  .  .  .  .  Name: "int"
    39  .  .  .  .  .  }
    40  .  .  .  .  }
    41  .  .  .  }
    42  .  .  .  Closing: ./example/example.go:5:18
    43  .  .  }
    44  .  }
    45  .  Body: *ast.BlockStmt {
    46  .  .  Lbrace: ./example/example.go:5:20
    47  .  .  List: []ast.Stmt (len = 1) {
    48  .  .  .  0: *ast.ExprStmt {
    49  .  .  .  .  X: *ast.CallExpr {
    50  .  .  .  .  .  Fun: *ast.SelectorExpr {
    51  .  .  .  .  .  .  X: *ast.Ident {
    52  .  .  .  .  .  .  .  NamePos: ./example/example.go:6:2
    53  .  .  .  .  .  .  .  Name: "log"
    54  .  .  .  .  .  .  }
    55  .  .  .  .  .  .  Sel: *ast.Ident {
    56  .  .  .  .  .  .  .  NamePos: ./example/example.go:6:6
    57  .  .  .  .  .  .  .  Name: "Println"
    58  .  .  .  .  .  .  }
    59  .  .  .  .  .  }
    60  .  .  .  .  .  Lparen: ./example/example.go:6:13
    61  .  .  .  .  .  Args: []ast.Expr (len = 1) {
    62  .  .  .  .  .  .  0: *ast.BinaryExpr {
    63  .  .  .  .  .  .  .  X: *ast.Ident {
    64  .  .  .  .  .  .  .  .  NamePos: ./example/example.go:6:14
    65  .  .  .  .  .  .  .  .  Name: "n"
    66  .  .  .  .  .  .  .  .  Obj: *(obj @ 20)
    67  .  .  .  .  .  .  .  }
    68  .  .  .  .  .  .  .  OpPos: ./example/example.go:6:16
    69  .  .  .  .  .  .  .  Op: +
    70  .  .  .  .  .  .  .  Y: *ast.Ident {
    71  .  .  .  .  .  .  .  .  NamePos: ./example/example.go:6:18
    72  .  .  .  .  .  .  .  .  Name: "m"
    73  .  .  .  .  .  .  .  .  Obj: *(obj @ 29)
    74  .  .  .  .  .  .  .  }
    75  .  .  .  .  .  .  }
    76  .  .  .  .  .  }
    77  .  .  .  .  .  Ellipsis: -
    78  .  .  .  .  .  Rparen: ./example/example.go:6:19
    79  .  .  .  .  }
    80  .  .  .  }
    81  .  .  }
    82  .  .  Rbrace: ./example/example.go:7:1
    83  .  }
    84  }
```

`f.Decls`のところは、`Imports`や`Comments`などもある。

### ASTのトラバース

```go
func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			fmt.Println(v.Name)
		}
		return true
	})
}
```

結果

```sh
❯ go run main.go
add
```

もう少し突っ込んでソースコードで`add`の位置を取得してみる。

```go
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
```

結果

```sh
❯ go run main.go
Name: add
Pos: 32
./example/example.go:5:1
```

### ASTの書き換え

```go
func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {
			
			v.Name = &ast.Ident{
				Name: "plus",
			}
		}
		return true
	})

	for _, d := range f.Decls {
		ast.Print(fset, d)
	}
}
```

これで`add`を`plus`に書き換えられる。
ただし、差分がASTの`Name`だけが書き換わったわけではない。
具体的には以下の通り。

```sh
     0  *ast.FuncDecl {
     1  .  Name: *ast.Ident {
     2  .  .  NamePos: ./example/example.go:5:6
     3  .  .  Name: "add"
     4  .  .  Obj: *ast.Object {
     5  .  .  .  Kind: func
     6  .  .  .  Name: "add"
     7  .  .  .  Decl: *(obj @ 0)
     8  .  .  }
     9  .  }
```

だったところが以下のようになる。（他にも実は微妙に差分がある）

```sh
     0  *ast.FuncDecl {
     1  .  Name: *ast.Ident {
     2  .  .  NamePos: -
     3  .  .  Name: "plus"
     4  .  }
```

TODO: [astutil \- GoDoc](https://godoc.org/golang.org/x/tools/go/ast/astutil)を使うと書き換えられるらしい？

```go
func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "./example/example.go", nil, parser.Mode(0))

	ast.Inspect(f, func(n ast.Node) bool {
		if v, ok := n.(*ast.FuncDecl); ok {

			v.Name = &ast.Ident{
				Name: "plus",
			}
		}
		return true
	})

	file, err := os.OpenFile("example/resutl.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pp := &printer.Config{
		Mode:     printer.UseSpaces | printer.TabIndent,
		Tabwidth: 8,
		Indent:   0,
	}
	pp.Fprint(file, fset, f)
}
```

結果(`resutl.go`)

```go
package example

import "log"

func plus(n, m int) {
	log.Println(n + m)
}
```

## [astutil\.Applyで抽象構文木を置き換える \#golang \- Qiita](https://qiita.com/tenntenn/items/40c563d6155e9ce19896)

```go
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
```

結果

```
*ast.BinaryExpr &{0xc000080060 2 + 0xc000080080}
*ast.BasicLit &{1 INT 1}
<nil> <nil>
*ast.BasicLit &{3 INT 1}
<nil> <nil>
<nil> <nil>
```


## `astutil.Apply`の使い方

シグネチャは`func Apply(root ast.Node, pre, post ApplyFunc) (result ast.Node) `である。
子ノードを処理する前後で実行したい関数を書く。`pre`だけでよければ次にようになる。


```go
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
```

結果

```
Args 0
Args 1
```

## ASTの書き換え

```go
func main() {
	expr, err := parser.ParseExpr(`x+y`)
	if err != nil {
		log.Fatal(err)
	}

	n := astutil.Apply(expr, func(cr *astutil.Cursor) bool {
		switch cr.Name() {
		case "X":
			cr.Replace(&ast.BasicLit{
				Kind:  token.INT,
				Value: "10",
			})
		case "Y":
			cr.Replace(&ast.BasicLit{
				Kind:  token.INT,
				Value: "20",
			})
		}
		return true
	}, nil)

	if err := format.Node(os.Stdout, token.NewFileSet(), n); err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println()
}
```

結果

```
10 + 20
```

## ASTからノードを削除する

```go
func main() {
	expr, err := parser.ParseExpr(`func(x, y int){}(10, 20)`)
	if err != nil {
		log.Fatal(err)
	}

	n := astutil.Apply(expr, func(cr *astutil.Cursor) bool {
		if cr.Name() == "Args" && cr.Index() == 0 {
			cr.Delete()
		}
		return true
	}, nil)

	if err := format.Node(os.Stdout, token.NewFileSet(), n); err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println()
}
```

結果

```go
func(x, y int) {
}()
```
参考までに`cr.Delete()`をコメントアウトすると以下のようになる。（当然）

```go
func(x, y int) {
}(10, 20)
```

あれ？ `cr.Index() == 0`だから`10`だけが消えて欲しかったのでは？
元記事にも以下のように書いてある。

> ここで注意したいのは、変更は直ちに行われるという点です。

`cr.Delete()`を`once.Do(cr.Delete)`に書き換えると一度だけ実行される。
※ `var once sync.Once`も必要

こんなところでも`sync.Once`使えるのね。

結果

```go
func(x, y int) {
}(20)
```

## 引数を追加する

`InsertBefore`や`InsertAfter`で実現する。

```go
func main() {
	expr, err := parser.ParseExpr(`func(x, y int){}(10, 20)`)
	if err != nil {
		log.Fatal(err)
	}

	n := astutil.Apply(expr, func(cr *astutil.Cursor) bool {
		if cr.Name() == "Args" && cr.Index() == 0 {
			cr.InsertBefore(&ast.BasicLit{
				Kind:  token.STRING,
				Value: "hi",
			})
			cr.InsertAfter(&ast.BasicLit{
				Kind:  token.STRING,
				Value: "gopher",
			})
		}
		return true
	}, nil)

	if err := format.Node(os.Stdout, token.NewFileSet(), n); err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Println()
}
```

```go
func(x, y int) {
}(hi, 10, gopher, 20)
```



## References
* [Go言語の golang/go パッケージで初めての構文解析](https://qiita.com/po3rin/items/a19d96d29284108ad442)
* [astutil \- GoDoc](https://godoc.org/golang.org/x/tools/go/ast/astutil)
* [goパッケージで簡単に静的解析して世界を広げよう \#golang \- Qiita](https://qiita.com/tenntenn/items/868704380455c5090d4b)
* [astutil\.Applyで抽象構文木を置き換える \#golang \- Qiita](https://qiita.com/tenntenn/items/40c563d6155e9ce19896)
* [Goにおける静的解析のモジュール化について \- Mercari Engineering Blog](https://tech.mercari.com/entry/2018/12/16/150000)