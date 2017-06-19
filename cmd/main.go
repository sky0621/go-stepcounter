package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"regexp"

	"fmt"

	"github.com/Sirupsen/logrus"
)

func main() {
	target := flag.String("target", "../sampleproject", "Parse Target")
	flag.Parse()

	err := filepath.Walk(*target, Apply)
	if err != nil {
		logrus.Error(err)
	}
}

func Apply(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !filter(path, info) {
		return nil
	}

	// FIXME AST形式では要素ごとにカウントアップしてしまうため、シンプルにファイルを行リードにする

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	stepCount := 0
	commentCount := 0
	ast.Inspect(f, func(node ast.Node) bool {
		_, ok := node.(*ast.CommentGroup)
		if ok {
			commentCount = commentCount + 1
			return true
		}
		_, ok = node.(*ast.Comment)
		if ok {
			commentCount = commentCount + 1
			return true
		}
		stepCount = stepCount + 1
		return true
	})

	fmt.Printf("[%v] step: %d, comment: %d\n", path, stepCount, commentCount)
	return nil
}

func filter(path string, info os.FileInfo) bool {
	if info.IsDir() {
		return false
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}

	outDirExp, err := regexp.Compile("vendor")
	if err != nil {
		return false
	}
	if outDirExp.MatchString(absPath) {
		return false
	}

	outFileExp, err := regexp.Compile(".*test.*")
	if err != nil {
		return false
	}
	if outFileExp.MatchString(path) {
		return false
	}

	inFileExp, err := regexp.Compile(".*.go")
	if err != nil {
		return false
	}
	if !inFileExp.MatchString(path) {
		return false
	}

	return true
}
