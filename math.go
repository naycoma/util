package util

import (
	"math"

	"golang.org/x/exp/constraints"
)

// DivMod returns the quotient (q) and remainder (r) of the division n/d.
// The division truncates toward zero, and the remainder has the same sign as the dividend (n).
// For example:
// DivMod(5, 3)  returns (1, 2)
// DivMod(-5, 3) returns (-1, -2)
// DivMod(5, -3) returns (-1, 2)
// DivMod(-5, -3) returns (1, -2)
func DivMod[I constraints.Integer](n, d I) (q, r I) {
	return n / d, n % d
}

// PositiveMod returns the result of the modulo operation, ensuring the result is always non-negative.
// For example, PositiveMod(-5, 3) returns 1, while -5 % 3 returns -2.
func PositiveMod[N constraints.Integer](n, d N) N {
	return (n%d + d) % d
}

// SubUnsigned performs subtraction for unsigned integers, preventing underflow.
// If a is less than b, it returns 0 instead of a negative result that would wrap around to a large positive number.
func SubUnsigned[T constraints.Unsigned](a, b T) T {
	if a > b {
		return a - b
	}
	return 0
}

// Floor returns the greatest integer value less than or equal to x.
// It supports generic Integer and Float types.
//
// Special cases are:
//
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor[I constraints.Integer, F constraints.Float](x F) I {
	return I(math.Floor(float64(x)))
}

// Ceil returns the smallest integer value greater than or equal to x.
// It supports generic Integer and Float types.
//
// Special cases are:
//
//	Ceil(±0) = ±0
//	Ceil(±Inf) = ±Inf
//	Ceil(NaN) = NaN
func Ceil[I constraints.Integer, F constraints.Float](x F) I {
	return I(math.Ceil(float64(x)))
}

// Round returns the nearest integer value to x, rounding half away from zero.
// It supports generic Integer and Float types.
//
// Special cases are:
//
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round[I constraints.Integer, F constraints.Float](x F) I {
	return I(math.Round(float64(x)))
}

// Abs returns the absolute value of x.
// It supports generic Integer and Float types.
func Abs[T constraints.Signed | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Trunc returns the integer value of x.
// It supports generic Integer and Float types.
//
// Special cases are:
//
//	Trunc(±0) = ±0
//	Trunc(±Inf) = ±Inf
//	Trunc(NaN) = NaN
func Trunc[I constraints.Integer, F constraints.Float](x F) I {
	return I(math.Trunc(float64(x)))
}

// RoundToEven returns the nearest integer, rounding ties to even.
// It supports generic Integer and Float types.
//
// Special cases are:
//
//	RoundToEven(±0) = ±0
//	RoundToEven(±Inf) = ±Inf
//	RoundToEven(NaN) = NaN
func RoundToEven[I constraints.Integer, F constraints.Float](x F) I {
	return I(math.RoundToEven(float64(x)))
}
