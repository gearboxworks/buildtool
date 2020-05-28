package pkgreflect

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	notypes    bool
	nofuncs    bool
	novars     bool
	noconsts   bool
	unexported bool
	norecurs   bool
	stdout     bool
	gofile     string
	notests    bool
)

func Main() error {
	var err error

	flag.BoolVar(&notypes, "notypes", false, "Don't list package types")
	flag.BoolVar(&nofuncs, "nofuncs", false, "Don't list package functions")
	flag.BoolVar(&novars, "novars", false, "Don't list package variables")
	flag.BoolVar(&noconsts, "noconsts", false, "Don't list package consts")
	flag.BoolVar(&unexported, "unexported", false, "Also list unexported names")
	flag.BoolVar(&norecurs, "norecurs", false, "Don't parse sub-directories resursively")
	flag.StringVar(&gofile, "gofile", "pkgreflect.go", "Name of the generated .go file")
	flag.BoolVar(&stdout, "stdout", false, "Write to stdout.")
	flag.BoolVar(&notests, "notests", false, "Don't list test related code")
	flag.Parse()

	if len(flag.Args()) > 0 {
		for _, dir := range flag.Args() {
			parseDir(dir)
		}
	} else {
		parseDir(".")
	}

	return err
}

func parseDir(dir string) {
	dirFile, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer dirFile.Close()
	info, err := dirFile.Stat()
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		panic("Path is not a directory: " + dir)
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), dir, filter, 0)
	if err != nil {
		panic(err)
	}
	for _, pkg := range pkgs {
		var buf bytes.Buffer

		fmt.Fprintln(&buf, "// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.\n")
		fmt.Fprintln(&buf, "package", pkg.Name)
		fmt.Fprintln(&buf, "")
		fmt.Fprintln(&buf, `import "reflect"`)
		fmt.Fprintln(&buf, "")

		// Types
		if !notypes {
			fmt.Fprintln(&buf, "var Types = map[string]reflect.Type{")
			print(&buf, pkg, ast.Typ, "\t\"%s\": reflect.TypeOf((*%s)(nil)).Elem(),\n")
			fmt.Fprintln(&buf, "}")
			fmt.Fprintln(&buf, "")
		}

		// Functions
		if !nofuncs {
			fmt.Fprintln(&buf, "var Functions = map[string]reflect.Value{")
			print(&buf, pkg, ast.Fun, "\t\"%s\": reflect.ValueOf(%s),\n")
			fmt.Fprintln(&buf, "}")
			fmt.Fprintln(&buf, "")
		}

		if !novars {
			// Addresses of variables
			fmt.Fprintln(&buf, "var Variables = map[string]reflect.Value{")
			print(&buf, pkg, ast.Var, "\t\"%s\": reflect.ValueOf(&%s),\n")
			fmt.Fprintln(&buf, "}")
			fmt.Fprintln(&buf, "")
		}

		if !noconsts {
			// Addresses of consts
			fmt.Fprintln(&buf, "var Consts = map[string]reflect.Value{")
			print(&buf, pkg, ast.Con, "\t\"%s\": reflect.ValueOf(%s),\n")
			fmt.Fprintln(&buf, "}")
			fmt.Fprintln(&buf, "")
		}

		if stdout {
			io.Copy(os.Stdout, &buf)
		} else {
			filename := filepath.Join(dir, gofile)
			newFileData := buf.Bytes()
			oldFileData, _ := ioutil.ReadFile(filename)
			if !bytes.Equal(newFileData, oldFileData) {
				err = ioutil.WriteFile(filename, newFileData, 0660)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	if !norecurs {
		dirs, err := dirFile.Readdir(-1)
		if err != nil {
			panic(err)
		}
		for _, info := range dirs {
			if info.IsDir() {
				parseDir(filepath.Join(dir, info.Name()))
			}
		}
	}
}

func print(w io.Writer, pkg *ast.Package, kind ast.ObjKind, format string) {
	names := []string{}
	for _, f := range pkg.Files {
		for name, object := range f.Scope.Objects {
			if object.Kind == kind && (unexported || ast.IsExported(name)) {
				names = append(names, name)
			}
		}
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Fprintf(w, format, name, name)
	}
}

func filter(info os.FileInfo) bool {

	name := info.Name()

	if info.IsDir() {
		return false
	}

	if name == gofile {
		return false
	}

	if filepath.Ext(name) != ".go" {
		return  false
	}

	if strings.HasSuffix(name, "_test.go") && notests {
		return false
	}

	return true

}
