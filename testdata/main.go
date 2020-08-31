package p

var i = 0

var I = 1

var c = "constant un-exported"

const C = "constant exported"

type t struct{}

type T struct{}

func main() {
}

func unexport(s string) {
}
func Export(s string) {
}

func ExportWithComment(s string) {
}

// ExistedComment
func ExistedComment() {}
