package stringsfp

import (
	"strings"
	"unicode"

	"github.com/j-dumbell/go-fp"
)

// Count is a curried version of strings.Count.
var Count func(substr string) func(s string) int = fp.CurryR2(strings.Count)

// Compare is a curried version of strings.Compare.
var Compare func(a string) func(b string) int = fp.Curry2(strings.Compare)

// Contains is a curried version of strings.Contains.
var Contains func(substr string) func(s string) bool = fp.CurryR2(strings.Contains)

// ContainsAny is a curried version of strings.ContainsAny.
var ContainsAny func(chars string) func(s string) bool = fp.CurryR2(strings.ContainsAny)

// ContainsRune is a curried version of strings.ContainsRune.
var ContainsRune func(r rune) func(s string) bool = fp.CurryR2(strings.ContainsRune)

// EqualFold is a curried version of strings.EqualFold.
var EqualFold func(s string) func(t string) bool = fp.Curry2(strings.EqualFold)

// FieldsFunc is a curried version of strings.FieldsFunc.
var FieldsFunc func(f func(rune) bool) func(s string) []string = fp.CurryR2(strings.FieldsFunc)

// HasPrefix is a curried version of strings.HasPrefix.
var HasPrefix func(prefix string) func(s string) bool = fp.CurryR2(strings.HasPrefix)

// HasSuffix is a curried version of strings.HasSuffix.
var HasSuffix func(suffix string) func(s string) bool = fp.CurryR2(strings.HasSuffix)

// Index is a curried version of strings.Index.
var Index func(substr string) func(s string) int = fp.CurryR2(strings.Index)

// IndexAny is a curried version of strings.IndexAny.
var IndexAny func(chars string) func(s string) int = fp.CurryR2(strings.IndexAny)

// IndexByte is a curried version of strings.IndexByte.
var IndexByte func(c byte) func(s string) int = fp.CurryR2(strings.IndexByte)

// IndexFunc is a curried version of strings.IndexFunc.
var IndexFunc func(f func(rune) bool) func(s string) int = fp.CurryR2(strings.IndexFunc)

// IndexRune is a curried version of strings.IndexRune.
var IndexRune func(r rune) func(s string) int = fp.CurryR2(strings.IndexRune)

// Join is a curried version of strings.Join.
var Join func(sep string) func(elems []string) string = fp.CurryR2(strings.Join)

// LastIndex is a curried version of strings.LastIndex.
var LastIndex func(substr string) func(s string) int = fp.CurryR2(strings.LastIndex)

// LastIndexAny is a curried version of strings.LastIndexAny.
var LastIndexAny func(chars string) func(s string) int = fp.CurryR2(strings.LastIndexAny)

// LastIndexByte is a curried version of strings.LastIndexByte.
var LastIndexByte func(c byte) func(s string) int = fp.CurryR2(strings.LastIndexByte)

// LastIndexFunc is a curried version of strings.LastIndexFunc.
var LastIndexFunc func(f func(rune) bool) func(s string) int = fp.CurryR2(strings.LastIndexFunc)

// Map is a curried version of strings.Map.
var Map func(mapping func(rune) rune) func(s string) string = fp.Curry2(strings.Map)

// Repeat is a curried version of strings.Repeat.
var Repeat func(count int) func(s string) string = fp.CurryR2(strings.Repeat)

// Replace is a curried version of strings.Replace.
func Replace(old string) func(new string) func(n int) func(s string) string {
	return func(new string) func(int) func(string) string {
		return func(n int) func(string) string {
			return func(s string) string {
				return strings.Replace(s, old, new, n)
			}
		}
	}
}

// ReplaceAll is a curried version of strings.ReplaceAll.
func ReplaceAll(old string) func(new string) func(s string) string {
	return func(new string) func(string) string {
		return func(s string) string {
			return strings.ReplaceAll(s, old, new)
		}
	}
}

// Split is a curried version of strings.Split.
var Split func(sep string) func(s string) []string = fp.CurryR2(strings.Split)

// SplitAfter is a curried version of strings.SplitAfter.
var SplitAfter func(sep string) func(s string) []string = fp.CurryR2(strings.SplitAfter)

// SplitAfterN is a curried version of strings.SplitAfterN.
func SplitAfterN(sep string) func(n int) func(s string) []string {
	return func(n int) func(string) []string {
		return func(s string) []string {
			return strings.SplitAfterN(s, sep, n)
		}
	}
}

// SplitN is a curried version of strings.SplitN.
func SplitN(sep string) func(n int) func(s string) []string {
	return func(n int) func(string) []string {
		return func(s string) []string {
			return strings.SplitN(s, sep, n)
		}
	}
}

// ToLowerSpecial is a curried version of strings.ToLowerSpecial.
var ToLowerSpecial func(c unicode.SpecialCase) func(s string) string = fp.Curry2(strings.ToLowerSpecial)

// ToTitleSpecial is a curried version of strings.ToTitleSpecial.
var ToTitleSpecial func(c unicode.SpecialCase) func(s string) string = fp.Curry2(strings.ToTitleSpecial)

// ToUpperSpecial is a curried version of strings.ToUpperSpecial.
var ToUpperSpecial func(c unicode.SpecialCase) func(s string) string = fp.Curry2(strings.ToUpperSpecial)

// ToValidUTF8 is a curried version of strings.ToValidUTF8.
var ToValidUTF8 func(replacement string) func(s string) string = fp.CurryR2(strings.ToValidUTF8)

// Trim is a curried version of strings.Trim.
var Trim func(cutset string) func(s string) string = fp.CurryR2(strings.Trim)

// TrimFunc is a curried version of strings.TrimFunc.
var TrimFunc func(f func(rune) bool) func(s string) string = fp.CurryR2(strings.TrimFunc)

// TrimLeft is a curried version of strings.TrimLeft.
var TrimLeft func(cutset string) func(s string) string = fp.CurryR2(strings.TrimLeft)

// TrimLeftFunc is a curried version of strings.TrimLeftFunc.
var TrimLeftFunc func(f func(rune) bool) func(s string) string = fp.CurryR2(strings.TrimLeftFunc)

// TrimPrefix is a curried version of strings.TrimPrefix.
var TrimPrefix func(prefix string) func(s string) string = fp.CurryR2(strings.TrimPrefix)

// TrimRight is a curried version of strings.TrimRight.
var TrimRight func(cutset string) func(s string) string = fp.CurryR2(strings.TrimRight)

// TrimRightFunc is a curried version of strings.TrimRightFunc.
var TrimRightFunc func(f func(rune) bool) func(s string) string = fp.CurryR2(strings.TrimRightFunc)

// TrimSuffix is a curried version of strings.TrimSuffix.
var TrimSuffix func(suffix string) func(s string) string = fp.CurryR2(strings.TrimSuffix)
