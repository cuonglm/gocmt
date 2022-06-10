package p

// CommentExisted
func CommentExisted() {
}

// CommentExistedWithSpace 
func CommentExistedWithSpace() {
}

// something
func CommentExistedWithWrong() {
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

/*
should't change C style comment
*/
var ValueWithExistedComment3 = 1

// existed comment
const (
	ParenValueWithExistedComment1 = 1
	// ParenValueWithExistedComment2 something
	ParenValueWithExistedComment2 = 1
)

// existed comment
type TypeWithExistedComment1 int

// existed comment with spaces     
type TypeWithExistedComment2 int

/*
should't change C style comment
*/
type TypeWithExistedComment3 int