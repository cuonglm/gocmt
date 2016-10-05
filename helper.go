package main

import (
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
