package p

import "embed"

// CommentExisted
func CommentExisted() {
}

// CommentExistedWithSpace 
func CommentExistedWithSpace() {
}

// something
func CommentExistedWithWrong() {
}

// multi-line comments
// something
func CommentExistedWithWrong2() {
}

/*
something
*/
func CommentWrongByDontChange() {
}

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

//go:embed dont_modify_this_comment.txt
//go:embed image/*
var Embed embed.FS

// something
//go:embed dont_modify_this_comment.txt
var Embed embed.FS

//go:generate goyacc -o gopher.go -p parser gopher.y
func Generate() {
}
