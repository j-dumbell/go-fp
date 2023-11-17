package bytesfp

import (
	"bytes"
	"unicode"

	"github.com/j-dumbell/go-fp"
)

// Compare is a curried version of bytes.Compare.
var Compare func(a []byte) func(b []byte) int = fp.Curry2(bytes.Compare)

// Count is a curried version of bytes.Count.
var Count func(sep []byte) func(s []byte) int = fp.CurryR2(bytes.Count)

// Contains is a curried version of bytes.Contains.
var Contains func(subslice []byte) func(b []byte) bool = fp.CurryR2(bytes.Contains)

// ContainsAny is a curried version of bytes.ContainsAny.
var ContainsAny func(chars string) func(b []byte) bool = fp.CurryR2(bytes.ContainsAny)

// ContainsRune is a curried version of bytes.ContainsRune.
var ContainsRune func(r rune) func(b []byte) bool = fp.CurryR2(bytes.ContainsRune)

// EqualFold is a curried version of bytes.EqualFold.
var EqualFold func(s []byte) func(t []byte) bool = fp.Curry2(bytes.EqualFold)

// FieldsFunc is a curried version of bytes.FieldsFunc.
var FieldsFunc func(f func(rune) bool) func(s []byte) [][]byte = fp.CurryR2(bytes.FieldsFunc)

// HasPrefix is a curried version of bytes.HasPrefix.
var HasPrefix func(prefix []byte) func(s []byte) bool = fp.CurryR2(bytes.HasPrefix)

// HasSuffix is a curried version of bytes.HasSuffix.
var HasSuffix func(suffix []byte) func(s []byte) bool = fp.CurryR2(bytes.HasSuffix)

// Index is a curried version of bytes.Index.
var Index func(sep []byte) func(s []byte) int = fp.CurryR2(bytes.Index)

// IndexAny is a curried version of bytes.IndexAny.
var IndexAny func(chars string) func(s []byte) int = fp.CurryR2(bytes.IndexAny)

// IndexByte is a curried version of bytes.IndexByte.
var IndexByte func(c byte) func(b []byte) int = fp.CurryR2(bytes.IndexByte)

// IndexFunc is a curried version of bytes.IndexFunc.
var IndexFunc func(f func(rune) bool) func(s []byte) int = fp.CurryR2(bytes.IndexFunc)

// IndexRune is a curried version of bytes.IndexRune.
var IndexRune func(r rune) func(s []byte) int = fp.CurryR2(bytes.IndexRune)

// Join is a curried version of bytes.Join.
var Join func(sep []byte) func(s [][]byte) []byte = fp.CurryR2(bytes.Join)

// LastIndex is a curried version of bytes.LastIndex.
var LastIndex func(sep []byte) func(s []byte) int = fp.CurryR2(bytes.LastIndex)

// LastIndexAny is a curried version of bytes.LastIndexAny.
var LastIndexAny func(chars string) func(s []byte) int = fp.CurryR2(bytes.LastIndexAny)

// LastIndexByte is a curried version of bytes.LastIndexByte.
var LastIndexByte func(c byte) func(s []byte) int = fp.CurryR2(bytes.LastIndexByte)

// LastIndexFunc is a curried version of bytes.LastIndexFunc.
var LastIndexFunc func(f func(rune) bool) func(s []byte) int = fp.CurryR2(bytes.LastIndexFunc)

// Map is a curried version of bytes.Map.
var Map func(mapping func(rune) rune) func(s []byte) []byte = fp.Curry2(bytes.Map)

// Repeat is a curried version of bytes.Repeat.
var Repeat func(count int) func(b []byte) []byte = fp.CurryR2(bytes.Repeat)

// Replace is a curried version of bytes.Replace.
func Replace(old []byte) func(new []byte) func(n int) func(s []byte) []byte {
	return func(new []byte) func(int) func([]byte) []byte {
		return func(n int) func([]byte) []byte {
			return func(s []byte) []byte {
				return bytes.Replace(s, old, new, n)
			}
		}
	}
}

// ReplaceAll is a curried version of bytes.ReplaceAll.
func ReplaceAll(old []byte) func(new []byte) func(s []byte) []byte {
	return func(new []byte) func([]byte) []byte {
		return func(s []byte) []byte {
			return bytes.ReplaceAll(s, old, new)
		}
	}
}

// Split is a curried version of bytes.Split.
var Split func(sep []byte) func(s []byte) [][]byte = fp.CurryR2(bytes.Split)

// SplitAfter is a curried version of bytes.SplitAfter.
var SplitAfter func(sep []byte) func(s []byte) [][]byte = fp.CurryR2(bytes.SplitAfter)

// SplitAfterN is a curried version of bytes.SplitAfterN.
func SplitAfterN(sep []byte) func(n int) func(s []byte) [][]byte {
	return func(n int) func([]byte) [][]byte {
		return func(s []byte) [][]byte {
			return bytes.SplitAfterN(s, sep, n)
		}
	}
}

// SplitN is a curried version of bytes.SplitN.
func SplitN(sep []byte) func(n int) func(s []byte) [][]byte {
	return func(n int) func([]byte) [][]byte {
		return func(s []byte) [][]byte {
			return bytes.SplitN(s, sep, n)
		}
	}
}

// ToLowerSpecial is a curried version of bytes.ToLowerSpecial.
var ToLowerSpecial func(c unicode.SpecialCase) func(s []byte) []byte = fp.Curry2(bytes.ToLowerSpecial)

// ToTitleSpecial is a curried version of bytes.ToTitleSpecial.
var ToTitleSpecial func(c unicode.SpecialCase) func(s []byte) []byte = fp.Curry2(bytes.ToTitleSpecial)

// ToUpperSpecial is a curried version of bytes.ToUpperSpecial.
var ToUpperSpecial func(c unicode.SpecialCase) func(s []byte) []byte = fp.Curry2(bytes.ToUpperSpecial)

// ToValidUTF8 is a curried version of bytes.ToValidUTF8.
var ToValidUTF8 func(replacement []byte) func(s []byte) []byte = fp.CurryR2(bytes.ToValidUTF8)

// Trim is a curried version of bytes.Trim.
var Trim func(cutset string) func(s []byte) []byte = fp.CurryR2(bytes.Trim)

// TrimFunc is a curried version of bytes.TrimFunc.
var TrimFunc func(f func(rune) bool) func(s []byte) []byte = fp.CurryR2(bytes.TrimFunc)

// TrimLeft is a curried version of bytes.TrimLeft.
var TrimLeft func(cutset string) func(s []byte) []byte = fp.CurryR2(bytes.TrimLeft)

// TrimLeftFunc is a curried version of bytes.TrimLeftFunc.
var TrimLeftFunc func(f func(rune) bool) func(s []byte) []byte = fp.CurryR2(bytes.TrimLeftFunc)

// TrimPrefix is a curried version of bytes.TrimPrefix.
var TrimPrefix func(prefix []byte) func(s []byte) []byte = fp.CurryR2(bytes.TrimPrefix)

// TrimRight is a curried version of bytes.TrimRight.
var TrimRight func(cutset string) func(s []byte) []byte = fp.CurryR2(bytes.TrimRight)

// TrimRightFunc is a curried version of bytes.TrimRightFunc.
var TrimRightFunc func(f func(r rune) bool) func(s []byte) []byte = fp.CurryR2(bytes.TrimRightFunc)

// TrimSuffix is a curried version of bytes.TrimSuffix.
var TrimSuffix func(suffix []byte) func(s []byte) []byte = fp.CurryR2(bytes.TrimSuffix)
