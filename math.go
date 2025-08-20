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

// PositiveMod computes the true mathematical modulo of x/d, which is always non-negative.
// It handles both integer and float types efficiently.
// For example, PositiveMod(-5, 3) returns 1.
func PositiveMod[R constraints.Integer | constraints.Float](x, d R) R {
	// Use a type switch on a zero value of type R to determine if we are working with floats or integers.
	var zero R
	switch any(zero).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return R((int64(x)%int64(d) + int64(d)) % int64(d))
	default:
		x_f64 := float64(x)
		d_f64 := float64(d)
		res := math.Mod(x_f64, d_f64)
		if res < 0 {
			res += d_f64
		}
		return R(res)
	}
}

// SubUnsigned performs subtraction for unsigned integers, preventing underflow.
// If a is less than b, it returns 0 instead of a negative result that would wrap around to a large positive number.
func SubUnsigned[N constraints.Unsigned](a, b N) N {
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
func Abs[R constraints.Signed | constraints.Float](x R) R {
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

// Repeat attempts to wrap the value x into the range [start, end).
//
// It computes the result based on `math.Mod(x - start, end - start) + start`.
// Due to the behavior of `math.Mod`, if `x` is already less than `start`,
// the function may return a value outside the desired range.
// For a behavior that correctly wraps all values into the range, see `Wrap`.
//
// Examples:
//  Repeat(12, 0, 10) returns 2
//  Repeat(5, 0, 10) returns 5
//  Repeat(3, 5, 10) returns 3 (note: not wrapped into [5, 10))
func Repeat[R constraints.Integer | constraints.Float](x, start, end R) R {
	return R(math.Mod(float64(x)-float64(start), float64(end)-float64(start)) + float64(start))
}
