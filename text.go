////////////////////////////////////////////////////////////////////////////
// Program: text
// Purpose: Common text handling for go-dedup
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: mkopriva https://stackoverflow.com/questions/45944781/
////////////////////////////////////////////////////////////////////////////

package text

import (
	"regexp"
	"strings"

	"github.com/go-dedup/megophone"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// TextCleanser defines the function type for text cleansing
type TextCleanser func(string) string

// TextCleanserDecorator is the text cleansing function Decorator
type TextCleanserDecorator func(TextCleanser) TextCleanser

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// Decorator higher-order functions

// SplitCamelCase split each CamelCase word in the text to individual words
func SplitCamelCase(c TextCleanser) TextCleanser {
	return func(s string) string {
		sn := regexp.MustCompile(`([a-z0-9])([A-Z]+)`).ReplaceAllString(s, "$1 $2")
		return c(sn)
	}
}

// ToLower cleanse the text to lower case
func ToLower(c TextCleanser) TextCleanser {
	return func(s string) string {
		lower := strings.ToLower(s)
		return c(lower)
	}
}

// ToDoubleMetaphone transforms the text to DoubleMetaphones
func ToDoubleMetaphone(c TextCleanser) TextCleanser {
	return func(s string) string {
		p1, p2 := megophone.DoubleMetaphone(s)
		return c(p1 + " " + p2)
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

// RemovePunctuation cleanse all punctuations from the text
func RemovePunctuation(c TextCleanser) TextCleanser {

	removePunctuation := func(r rune) rune {
		if strings.ContainsRune(",:;", r) {
			return -1
		} else if strings.ContainsRune("_", r) {
			return ' '
		} else if strings.ContainsRune("$", r) {
			return r
		} else if regexp.MustCompile(`\W`).MatchString(string(r)) {
			return ' '
		} else {
			return r
		}
	}

	return func(s string) string {
		rp := strings.Map(removePunctuation, s)
		return c(rp)
	}
}

// Compact cleanse all consecutive punctuations into a single space
func Compact(c TextCleanser) TextCleanser {
	return func(s string) string {
		sn := regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
		return c(sn)
	}
}

// Trim cleanse all leading and trailing spaces
func Trim(c TextCleanser) TextCleanser {
	return func(s string) string {
		sn := strings.Trim(s, " ")
		return c(sn)
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

//==========================================================================
// Other support functions

func GetWords(document string, dc TextCleanserDecorator) []string {
	fn := dc(Ident)
	return strings.Split(fn(document), " ")
}

func GetDoubleMetaphone(document string, dc TextCleanserDecorator) []string {
	var ret []string
	fn := dc(Ident)
	words := strings.Split(fn(document), " ")
	for _, key := range words {
		p1, p2 := megophone.DoubleMetaphone(key)
		ret = append(ret, p1, p2)
	}
	return ret
}

//==========================================================================
// Curried functions for the above support functions

// Doc2Words defines the function type for doc to words
type Doc2Words func(document string) []string

func GetWordsFactory(dc TextCleanserDecorator) Doc2Words {
	return func(document string) []string {
		fn := dc(Ident)
		return strings.Split(fn(document), " ")
	}
}

func GetDoubleMetaphoneFactory(dc TextCleanserDecorator) Doc2Words {
	return func(document string) []string {
		var ret []string
		fn := dc(Ident)
		words := strings.Split(fn(document), " ")
		for _, key := range words {
			p1, p2 := megophone.DoubleMetaphone(key)
			ret = append(ret, p1, p2)
		}
		return ret
	}
}
