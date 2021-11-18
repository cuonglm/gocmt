package main

import (
	"fmt"
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
		// exclude path
		for _, e := range excludeDirs {
			if strings.HasPrefix(path, e) {
				fmt.Fprintf(os.Stdout, "ignore %s\n", path)
				return nil
			}
		}

		err = processFile(path, *template, *inPlace)
	}

	if err != nil {
		return err
	}

	return nil
}
