package main

import (
	"go/doc"
	"go/parser"
	"go/token"
	"io"
)

func populateHTML(dst io.Writer, fp string) error {
	fset := token.NewFileSet()
	pkgs, err := dirPkgs(fset, fp)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		pkgHTML(dst, pkg)
	}

	return nil
}

func dirPkgs(fset *token.FileSet, path string) (map[string]*doc.Package, error) {
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	docPkgs := make(map[string]*doc.Package, len(pkgs))
	for pkgName, pkg := range pkgs {
		docPkgs[pkgName] = doc.New(pkg, pkg.Name, 0)
	}
	return docPkgs, nil
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
