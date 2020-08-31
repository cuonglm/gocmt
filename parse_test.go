package main

import (
	"bytes"
	"go/format"
	"go/token"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseSrc = `package p

var i = 0

// I ...
var I = 1

var c = "constant un-exported"

// C ...
const C = "constant exported"

type t struct{}

// T ...
type T struct{}

func main() {
}

func unexport(s string) {
}
// Export ...
func Export(s string) {
}

// ExportWithComment ...
func ExportWithComment(s string) {
}
`

const parenSrc = `package p

// Summon ...
type Summon string

const (
	// DarkOmega ...
	DarkOmega Summon = "celeste"
	// LightOmega best summon
	LightOmega Summon = "luminineria"

	// WindOmega ...
	WindOmega Summon = "tiamat"
)

// FireUtility ...
const FireUtility Summon = "the sun"

const (
	// Light best summon
	Light Summon = "lucifer"
)
`

func Test_parseFile(t *testing.T) {
	parseFileTests := []struct {
		path        string
		expectedSrc string
		modified    bool
		wantErr     bool
	}{
		{"testdata/main.go", baseSrc, true, false},
		{"testdata/parenthesis.go", parenSrc, true, false},
		{"testdata/invalid_file.go", "", false, true},
	}

	for _, tc := range parseFileTests {
		tc := tc
		t.Run(tc.path, func(t *testing.T) {
			t.Parallel()
			fset := token.NewFileSet()
			af, modified, err := parseFile(fset, tc.path, "...")
			assert.True(t, tc.wantErr == (err != nil))
			assert.Equal(t, tc.modified, modified)

			if tc.modified {
				buf := new(bytes.Buffer)
				assert.NoError(t, format.Node(buf, fset, af))
				newBuf := buf.Bytes()
				newBuf = tralingWsRegex.ReplaceAll(newBuf, []byte(""))
				newBuf = newlinesRegex.ReplaceAll(newBuf, []byte("\n\n"))
				assert.Equal(t, tc.expectedSrc, string(newBuf))
			}
		})

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
