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
		if err := docHTML(dst, pkg); err != nil {
			return err
		}
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

func docHTML(dst io.Writer, pkg *doc.Package) error {
	docWriters := []func(w io.Writer) error{
		func(w io.Writer) error {
			doc.ToHTML(dst, pkg.Doc, nil)
			return nil
		},
		func(w io.Writer) error {
			for _, constV := range pkg.Consts {
				if err := valHTML(dst, constV); err != nil {
					return err
				}
			}
			return nil
		},
		func(w io.Writer) error {
			for _, varV := range pkg.Vars {
				if err := valHTML(dst, varV); err != nil {
					return err
				}
			}
			return nil
		},
		func(w io.Writer) error {
			for _, typeV := range pkg.Types {
				if err := typeHTML(dst, typeV); err != nil {
					return err
				}
			}
			return nil
		},
		func(w io.Writer) error {
			for _, fn := range pkg.Funcs {
				if err := fnHTML(dst, fn); err != nil {
					return err
				}
			}
			return nil
		},
		func(w io.Writer) error {
			for _, ex := range pkg.Examples {
				if err := exHTML(dst, ex); err != nil {
					return err
				}
			}
			return nil
		},
	}

	for _, w := range docWriters {
		if err := w(dst); err != nil {
			return err
		}
	}
	return nil
}

func valHTML(dst io.Writer, val *doc.Value) error {
	doc.ToHTML(dst, val.Doc, nil)
	return nil
}

func typeHTML(dst io.Writer, typ *doc.Type) error {
	doc.ToHTML(dst, typ.Doc, nil)
	for _, constV := range typ.Consts {
		if err := valHTML(dst, constV); err != nil {
			return err
		}
	}
	for _, varV := range typ.Vars {
		if err := valHTML(dst, varV); err != nil {
			return err
		}
	}
	for _, fn := range typ.Funcs {
		if err := fnHTML(dst, fn); err != nil {
			return err
		}
	}
	for _, fn := range typ.Methods {
		if err := fnHTML(dst, fn); err != nil {
			return err
		}
	}
	for _, ex := range typ.Examples {
		if err := exHTML(dst, ex); err != nil {
			return err
		}
	}

	return nil
}

func fnHTML(dst io.Writer, fn *doc.Func) error {
	doc.ToHTML(dst, fn.Doc, nil)
	for _, ex := range fn.Examples {
		if err := exHTML(dst, ex); err != nil {
			return err
		}
	}
	return nil
}

func exHTML(dst io.Writer, ex *doc.Example) error {
	doc.ToHTML(dst, ex.Doc, nil)
	return nil
}
