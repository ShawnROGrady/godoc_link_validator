package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"io"
	"strings"
)

func populateHTML(dst io.Writer, fp string) error {
	fset := token.NewFileSet()
	pkgs, err := dirPkgs(fset, fp)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		pkg.writeHTML(dst)
	}

	return nil
}

type docPkg struct {
	pkg       *doc.Package
	testFiles []*ast.File
}

func dirPkgs(fset *token.FileSet, path string) ([]*docPkg, error) {
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	docPkgs := make([]*docPkg, len(pkgs))
	i := 0
	for pkgName, pkg := range pkgs {
		var (
			testFiles  []*ast.File
			otherFiles []*ast.File
		)
		for fName, f := range pkg.Files {
			if strings.HasSuffix(fName, "_test.go") {
				testFiles = append(testFiles, f)
			} else {
				otherFiles = append(otherFiles, f)
			}
		}
		docP, err := doc.NewFromFiles(fset, otherFiles, pkgName)
		if err != nil {
			return nil, err
		}
		docPkgs[i] = &docPkg{
			pkg:       docP,
			testFiles: testFiles,
		}
		i++
	}
	return docPkgs, nil
}

func (p *docPkg) writeHTML(dst io.Writer) {
	pkgHTML(dst, p.pkg)
	for _, ex := range doc.Examples(p.testFiles...) {
		exHTML(dst, ex)
	}
}

func pkgHTML(dst io.Writer, pkg *doc.Package) {
	doc.ToHTML(dst, pkg.Doc, nil)
	for _, constV := range pkg.Consts {
		valHTML(dst, constV)
	}
	for _, varV := range pkg.Vars {
		valHTML(dst, varV)
	}
	for _, typeV := range pkg.Types {
		typeHTML(dst, typeV)
	}
	for _, fn := range pkg.Funcs {
		fnHTML(dst, fn)
	}
	for _, ex := range pkg.Examples {
		exHTML(dst, ex)
	}
}

func valHTML(dst io.Writer, val *doc.Value) {
	doc.ToHTML(dst, val.Doc, nil)
}

func typeHTML(dst io.Writer, typ *doc.Type) {
	doc.ToHTML(dst, typ.Doc, nil)
	for _, constV := range typ.Consts {
		valHTML(dst, constV)
	}
	for _, varV := range typ.Vars {
		valHTML(dst, varV)
	}
	for _, fn := range typ.Funcs {
		fnHTML(dst, fn)
	}
	for _, fn := range typ.Methods {
		fnHTML(dst, fn)
	}
	for _, ex := range typ.Examples {
		exHTML(dst, ex)
	}

}

func fnHTML(dst io.Writer, fn *doc.Func) {
	doc.ToHTML(dst, fn.Doc, nil)
	for _, ex := range fn.Examples {
		exHTML(dst, ex)
	}
}

func exHTML(dst io.Writer, ex *doc.Example) {
	doc.ToHTML(dst, ex.Doc, nil)
}
