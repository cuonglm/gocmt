// gocmt adds missing comments on exported identifiers in Go source files.
//
// Usage:
//
//  gocmt [-i false|true] [-t "comment template"] [-d dir] [-p false|true]
//
// This tools exists because I have to work with some existed code base, which
// is lack of documentation for many exported identifiers. Iterating over them
// is time consuming and maybe not suitable at a stage of the project. So I
// wrote this tool to quickly bypassing CI system. Once thing is settle, we can
// lookback and fix missing comments.
//
// You SHOULD always write documentation for all your exported identifiers.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	// ensure that the comment starts on a newline (without the \n, sometimes it starts on the previous }
	commentBase = "\n// %s "
	// if it's in an indented block, this makes sure that the indentation is correct
	commentIndentedBase = "// %s "
	fset                = token.NewFileSet()
	defaultMode         = os.FileMode(0644)
	tralingWsRegex      = regexp.MustCompile(`(?m)[\t ]+$`)
	newlinesRegex       = regexp.MustCompile(`(?m)\n{3,}`)
)

var (
	inPlace      = flag.Bool("i", false, "Make in-place editing")
	template     = flag.String("t", "...", "Comment template")
	dir          = flag.String("d", "", "Directory to process")
	parenComment = flag.Bool("p", false, "Add comments to all const inside the parens if true")
)

func main() {
	os.Exit(gocmtRun())
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: gocmt [flags] [file ...]\n")
	flag.PrintDefaults()
}

func gocmtRun() int {
	flag.Parse()

	if *dir != "" {
		if err := filepath.Walk(*dir, walkFunc); err != nil {
			printError(err)
			return 1
		}
		return 0
	}

	if flag.NArg() == 0 {
		usage()
	}

	for i := 0; i < flag.NArg(); i++ {
		path := flag.Arg(i)
		switch fi, err := os.Stat(path); {
		case err != nil:
			printError(err)
		case fi.IsDir():
			printError(fmt.Errorf("%s is a directory", path))
		default:
			if err := processFile(path, *template, *inPlace); err != nil {
				printError(err)
				return 1
			}
		}
	}

	return 0
}

func processFile(filename, template string, inPlace bool) error {
	// skip test files and files in vendor/
	if strings.HasSuffix(filename, "_test.go") || strings.Contains(filename, "/vendor/") {
		return nil
	}

	af, modified, err := parseFile(fset, filename, template)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, af); err != nil {
		panic(err)
	}

	newBuf := buf.Bytes()
	if modified {
		newBuf = tralingWsRegex.ReplaceAll(newBuf, []byte(""))
		newBuf = newlinesRegex.ReplaceAll(newBuf, []byte("\n\n"))
		if inPlace {
			return ioutil.WriteFile(filename, newBuf, defaultMode)
		}

		fmt.Fprintf(os.Stdout, "%s", newBuf)
		return nil
	}

	fmt.Fprintf(os.Stderr, "%s no changes\n", filename)

	return nil
}
