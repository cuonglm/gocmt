package main

import (
	"bytes"
	"go/token"
	"io/ioutil"
	"testing"
)

func Test_parseFile(t *testing.T) {
	parseFileTests := []struct {
		path    string
		wantErr bool
	}{
		{"testdata/main.go", false},
		{"testdata/parenthesis.go", false},
		{"testdata/invalid_file.go", true},
	}

	for _, tt := range parseFileTests {
		fset := token.NewFileSet()
		_, err := parseFile(fset, tt.path, "...")

		if tt.wantErr && err == nil {
			t.Fatalf("Parsing %s want error, got nil", tt.path)
		}

		if !tt.wantErr && err != nil {
			t.Fatalf("Parsing %s error: %s", tt.path, err)
		}
	}
}

func TestSkipVendor(t *testing.T) {
	filePath := "testdata/vendor/main.go"
	origBuf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if err := processFile(filePath, "...", true); err != nil {
		t.Fatal(err)
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf, origBuf) {
		t.Fatal("file in vendor/ directory was edited")
	}
}
