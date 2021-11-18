// gocmt adds missing comments on exported identifiers in Go source files.
//
// Usage:
//
//  gocmt [-i] [-p] [-t "comment template"] [-d dir] [-e exclude_dirs]
//
// This tools exists because I have to work with some existed code base, which
// is lack of documentation for many exported identifiers. Iterating over them
// is time consuming and maybe not suitable at a stage of the project. So I
// wrote this tool to quickly bypassing CI system. Once thing is settle, we can
// lookback and fix missing comments.
//
// You SHOULD always write documentation for all your exported identifiers.
package main
