package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAST(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go",
		`
package ast

import (
	"fmt"
	"go/ast"
	"reflect"
)

type printVisitor struct {
}

func (t *printVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		fmt.Println("nil")
		return t
	}

	val := reflect.ValueOf(node)
	typ := reflect.TypeOf(node)

	for typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	fmt.Printf("val: %+v, type: %s \n", val.Interface(), typ.Name())
	return t
}
`, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	ast.Walk(&printVisitor{}, f)
}
