package main

import (
	"go/token"
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
