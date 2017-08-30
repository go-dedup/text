package text_test

import (
	"fmt"

	"github.com/go-dedup/text"
)

// for standalone test, change package to `main` and the next func def to,
// func main() {
func ExampleTextCleanser() {
	s := "Hello~~, play_ground#5!"

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
		text.SplitCamelCase,
		text.ToLower,
		text.ToPrepend("DECORATED: "),
		text.RemovePunctuation,
	)

	fn4 := dec(text.Ident)
	fmt.Println(fn4(s))
	s += "\n\n%% Something extra: UpperCamelCase and someInitMethod.\n"
	fmt.Printf("\n>>>>\n'%s'\n", s)
	fmt.Printf("%#v\n", text.GetWords(s, dec))

	dec = text.Decorators(
		dec,
		text.Compact,
	)
	fmt.Printf("%#v\n", text.GetWords(s, dec))

	// Output:
	// Hello~~, play_ground#5!
	// hello~~, play_ground#5!
	// DECORATED: hello~~, play_ground#5! -golang
	// DECORATED hello   play ground 5    golang
	//
	// >>>>
	// 'Hello~~, play_ground#5!
	//
	// %% Something extra: UpperCamelCase and someInitMethod.
	// '
	// []string{"DECORATED", "hello", "", "", "play", "ground", "5", "", "", "", "", "", "", "something", "extra", "", "upper", "camel", "case", "and", "some", "init", "method", "", "", "", "", "golang"}
	// []string{"DECORATED", "hello", "play", "ground", "5", "something", "extra", "upper", "camel", "case", "and", "some", "init", "method", "golang"}
}

// to show the full code in GoDoc
type dummy struct {
}
