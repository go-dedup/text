////////////////////////////////////////////////////////////////////////////
// Program: text
// Purpose: Common text handling for go-dedup
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: mkopriva https://stackoverflow.com/questions/45944781/
////////////////////////////////////////////////////////////////////////////

package text

import (
	"strings"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// TextCleanser defines the function type for text cleansing
type TextCleanser = func(string) string

// TextCleanserDecorator is the text cleansing function Decorator
type TextCleanserDecorator = func(TextCleanser) TextCleanser

////////////////////////////////////////////////////////////////////////////
// Function definitions

// ToLower cleanse the text to lower case
func ToLower(c TextCleanser) TextCleanser {
	return func(s string) string {
		lower := strings.ToLower(s)
		return c(lower)
	}
}

// ToAppend manipulates the text by appending a suffix
func ToAppend(suffix string) TextCleanserDecorator {
	return func(c TextCleanser) TextCleanser {
		return func(s string) string {
			return c(s + suffix)
		}
	}
}

// ToPrepend manipulates the text by pre-pending with a prefix
func ToPrepend(prefix string) TextCleanserDecorator {
	return func(c TextCleanser) TextCleanser {
		return func(s string) string {
			return c(prefix + s)
		}
	}
}

// Ident -- "identity" just return the same string
func Ident(s string) string {
	return s
}

// Decorators "merges" the passed in decorators and returns a singe decorator.
func Decorators(ds ...TextCleanserDecorator) TextCleanserDecorator {
	return func(c TextCleanser) TextCleanser {
		for ii := range ds {
			c = ds[len(ds)-ii-1](c)
		}
		return c
	}
}
