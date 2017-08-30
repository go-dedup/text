package text_test

import (
	"fmt"

	"github.com/go-dedup/text"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleTextCleanser() {
	s := "Hello, playground"

	var fn text.TextCleanser = text.Ident
	fmt.Println(fn(s))

	var fn2 = text.ToLower(fn)
	fmt.Println(fn2(s))

	var fn3 text.TextCleanser = text.Ident
	fn3 = text.ToAppend(" -GOLANG")(text.ToLower(text.ToPrepend("DECORATED: ")(fn3)))
	fmt.Println(fn3(s))

	// dec is now a text.TextCleanserDecorator, to use it, you still need to
	// pass it the function of type text.TextCleanser that you want to decorate.
	dec := text.Decorators(
		text.ToAppend(" -GOLANG"),
		text.ToLower,
		text.ToPrepend("DECORATED: "),
	)

	fn4 := dec(text.Ident)
	fmt.Println(fn4(s))

	// Output:
	// Hello, playground
	// hello, playground
	// DECORATED: hello, playground -golang
	// DECORATED: hello, playground -golang
}
