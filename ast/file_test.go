package ast

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestFileVisitor_Get(t *testing.T) {
	testCases := []struct {
		src  string
		want File
	}{
		{
			src: `// annotation go through the source code and extra the annotation
// @author Deng Ming
// @date 2022/04/02
`,
			want: File{
				Annotations: Annotations{
					Ans: []Annotation{
						{
							Key:   "author",
							Value: "Deng Ming",
						},
						{
							Key:   "date",
							Value: "2022/04/02",
						},
					},
				},
				Types: []Type{
					{
						Annotations: Annotations{
							Ans: []Annotation{
								{
									Key:   "author",
									Value: "Deng Ming",
								},
								{
									Key:   "date",
									Value: "2022/04/02",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, "src.go", tc.src, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}
		tv := &SingleFileEntryVisitor{}
		ast.Walk(tv, f)
		file := tv.Get()
		assertAnnotations(t, tc.want.Annotations, file.Annotations)
		if len(tc.want.Types) != len(file.Types) {
			t.Fatal()
		}
		for i, typ := range tc.want.Types {
			wantTyp := tc.want.Types[i]
			assertAnnotations(t, wantTyp.Annotations, typ.Annotations)
			if len(wantTyp.Fields) != len(typ.Fields) {
				t.Fatal()
			}
			for j, fd := range wantTyp.Fields {
				wannFd := wantTyp.Fields[j]
				assertAnnotations(t, wannFd.Annotations, fd.Annotations)
			}
		}
	}
}

func assertAnnotations(t *testing.T, wantAns, dstAns Annotations) {
	want := wantAns.Ans
	if len(want) != len(dstAns.Ans) {
		t.Fatal()
	}
	for i, an := range want {
		val := dstAns.Ans[i]
		assert.Equal(t, an.Value, val.Value)
	}
}
