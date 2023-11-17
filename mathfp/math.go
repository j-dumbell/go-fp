package mathfp

import (
	"math"

	"github.com/j-dumbell/go-fp"
	"golang.org/x/exp/constraints"
)

// CopySign is a curried version of math.CopySign.
var CopySign func(sign float64) func(f float64) float64 = fp.CurryR2(math.Copysign)

// IsInf is a curried version of math.IsInf.
var IsInf func(sign int) func(f float64) bool = fp.CurryR2(math.IsInf)

// Min is a curried version of math.Min, however also accepts float32.
func Min[T constraints.Float](x T) func(y T) T {
	return func(y T) T {
		return T(math.Min(float64(x), float64(y)))
	}
}

// Max is a curried version of math.Max, however also accepts float32.
func Max[T constraints.Float](x T) func(y T) T {
	return func(y T) T {
		return T(math.Max(float64(x), float64(y)))
	}
}

// Mod is a curried version of math.Mod, however also accepts float32.
func Mod[T constraints.Float](x T) func(y T) T {
	return func(y T) T {
		return T(math.Mod(float64(x), float64(y)))
	}
}

// NextAfter is a curried version of math.NextAfter, however also accepts float32.
func NextAfter[T constraints.Float](x T) func(y T) T {
	return func(y T) T {
		return T(math.Nextafter(float64(x), float64(y)))
	}
}

// Pow is a curried version of math.Pow, however also accepts float32.
func Pow[T constraints.Float](y T) func(x T) T {
	return func(x T) T {
		return T(math.Pow(float64(x), float64(y)))
	}
}
