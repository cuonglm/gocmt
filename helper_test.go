package main

import (
	"os"
	"testing"
)

func Test_isGoFile(t *testing.T) {
	isGoFileTest := []struct {
		path     string
		expected bool
	}{
		{"main.go", true},
		{"README.md", false},
		{".travis.yml", false},
	}

	for _, tt := range isGoFileTest {
		fi, err := os.Stat(tt.path)
		if err != nil {
			t.Fatal(err)
		}

		if got := isGoFile(fi); got != tt.expected {
			t.Fatalf("isGoFile(%+v): expected %v, got %v", fi, tt.expected, got)
		}
	}
}
