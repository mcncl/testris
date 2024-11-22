package finder

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type TestInfo struct {
	Name    string
	Package string
	File    string
}

func FindTests(root string) ([]TestInfo, error) {
	var tests []TestInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if shouldSkipPath(path, info) {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if strings.HasSuffix(path, "_test.go") {
			fileTests, err := parseTestFile(path)
			if err != nil {
				return err
			}
			tests = append(tests, fileTests...)
		}
		return nil
	})

	return tests, err
}

func shouldSkipPath(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return info.Name() == "vendor" || info.Name() == ".git"
	}
	return !strings.HasSuffix(path, "_test.go")
}

func parseTestFile(path string) ([]TestInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var tests []TestInfo
	pkg := node.Name.Name

	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if isTestFunction(fn) {
				tests = append(tests, TestInfo{
					Name:    fn.Name.String(),
					Package: pkg,
					File:    path,
				})
			}
		}
		return true
	})

	return tests, nil
}

func isTestFunction(fn *ast.FuncDecl) bool {
	return strings.HasPrefix(fn.Name.String(), "Test") &&
		fn.Type.Params != nil &&
		len(fn.Type.Params.List) == 1
}
