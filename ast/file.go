package ast

import "go/ast"

type Field struct {
	Annotations
}

type Type struct {
	Annotations
	Fields []Field
}

type File struct {
	ans   Annotations
	types []Type
}

type typeVisitor struct {
	ans    Annotations
	fields []Field
}

func (t *typeVisitor) Get() Type {
	return Type{
		Annotations: t.ans,
		Fields:      t.fields,
	}
}

type fileVisitor struct {
	ans   Annotations
	types []typeVisitor
}

func (f *fileVisitor) Get() File {
	types := make([]Type, 0, len(f.types))
	for _, t := range f.types {
		types = append(types, t.Get())
	}
	return File{
		ans:   f.ans,
		types: types,
	}
}

//func (f *fileVisitor) Visit(node ast.Node) (w ast.Visitor) {
//	typ, ok := node.(*ast.TypeSpec)
//	if ok {
//		w := &typeVisitor{
//			ans:    NewAnnotations(typ, typ.Doc),
//			fields: make([]Field, 0, 0),
//		}
//		f.types = append(f.types, w)
//	}
//	return f
//}

type SingleFileEntryVisitor struct {
	file *fileVisitor
}

func (s *SingleFileEntryVisitor) Get() File {
	if s.file != nil {
		return s.file.Get()
	}
	return File{}
}

func (s *SingleFileEntryVisitor) Visit(node ast.Node) ast.Visitor {
	file, ok := node.(*ast.File)
	if ok {
		s.file = &fileVisitor{
			ans: NewAnnotations(file, file.Doc),
		}
		return s.file
	}
	return s
}
