package p

// global comments should never be deleted
// something

import "embed"

// global comment 1

// global comment 2
// global comment 3

// ============= function =============

// FuncWithExistedComment1
func FuncWithExistedComment1() {
}

func FuncWithExistedComment2() {
	// this comments should never be deleted
}

// something
func FuncWithExistedComment3() {
}

// multi-line comments
// something
func FuncWithExistedComment4() {
}

/*
something
*/
func FuncWithExistedComment5() {
}

// ============= value =============

// existed comment
var ValueWithExistedComment1 = 1

// existed comment with spaces     
var ValueWithExistedComment2 = 1

// multi-line comments
// something
var ValueWithExistedComment3 = 1

/*
should't change C style comment
*/
var ValueWithExistedComment4 = 1

// ============= paren value =============

// existed comment
const (
	ParenValueWithExistedComment1 = 1
	// ParenValueWithExistedComment2 something
	ParenValueWithExistedComment2 = 1
)

// multi-line comments
// something
const (
	ParenValueWithExistedComment3 = 1
	// ParenValueWithExistedComment2 something
	ParenValueWithExistedComment4 = 1
)

// ============= type =============

// existed comment
type TypeWithExistedComment1 int

// existed comment with spaces     
type TypeWithExistedComment2 int

// multi-line comments
// something
type TypeWithExistedComment3 int

/*
should't change C style comment
*/
type TypeWithExistedComment4 int

// ============= marker comment =============

//go:embed dont_modify_this_comment.txt
//go:embed image/*
var Embed embed.FS

// something
//go:embed dont_modify_this_comment.txt
var Embed embed.FS

//go:generate goyacc -o gopher.go -p parser gopher.y
func Generate() {
}

// ============= end =============
