package main

import (
	"go/ast"
	"go/scanner"
	"os"
	"strings"
)

func isGoFile(f os.FileInfo) bool {
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go")
}

func printError(err error) {
	scanner.PrintError(os.Stderr, err)
}

func walkFunc(path string, fi os.FileInfo, err error) error {
	if err == nil && isGoFile(fi) {
		err = processFile(path, *template, *inPlace)
	}

	if err != nil {
		return err
	}

	return nil
}

func isLineComment(group *ast.CommentGroup) bool {
	if group == nil {
		return false
	}
	if len(group.List) == 0 {
		return false
	}
	head := group.List[0].Text
	head = strings.TrimSpace(head)
	return strings.HasPrefix(head, "//")
}
