package text

import "testing"

func BenchmarkSplitCamelCase(b *testing.B) {
	bw := "CamelCaseWords FooBar3Baz GNU PYTHON Standard"
	for i := 0; i < b.N; i++ {
		GetWords(bw, SplitCamelCase)
	}
}

func BenchmarkSplitCamelCaseUnicode(b *testing.B) {
	bw := "CamelCaseWords FooBar3Baz GNU PYTHON Standard"
	for i := 0; i < b.N; i++ {
		GetWords(bw, SplitCamelCaseUnicode)
	}
}

/*

$ go test -v -bench=.

goos: linux
goarch: amd64
BenchmarkSplitCamelCase-2                  20000             70751 ns/op
BenchmarkSplitCamelCaseUnicode-2          500000              2871 ns/op

*/
