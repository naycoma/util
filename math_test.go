package util_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/naycoma/util"
)

func TestDivMod(t *testing.T) {
	a := assert.New(t)

	q, r := util.DivMod(5, 3)
	a.Equal(1, q, "5/3 quotient")
	a.Equal(2, r, "5/3 remainder")

	q, r = util.DivMod(-5, 3)
	a.Equal(-1, q, "-5/3 quotient")
	a.Equal(-2, r, "-5/3 remainder")

	q, r = util.DivMod(5, -3)
	a.Equal(-1, q, "5/-3 quotient")
	a.Equal(2, r, "5/-3 remainder")

	q, r = util.DivMod(-5, -3)
	a.Equal(1, q, "-5/-3 quotient")
	a.Equal(-2, r, "-5/-3 remainder")
}

func TestPositiveMod(t *testing.T) {
	a := assert.New(t)
	// --- Integer tests ---
	a.Equal(1, util.PositiveMod(1, 3), "1 mod 3")
	a.Equal(0, util.PositiveMod(3, 3), "3 mod 3")
	a.Equal(2, util.PositiveMod(5, 3), "5 mod 3")
	a.Equal(1, util.PositiveMod(-5, 3), "-5 mod 3")
	a.Equal(0, util.PositiveMod(-3, 3), "-3 mod 3")
	a.Equal(2, util.PositiveMod(-1, 3), "-1 mod 3")

	a.Equal(0, util.PositiveMod(0, 5), "0 mod 5")
	a.Equal(4, util.PositiveMod(-1, 5), "ｰ1 mod 5")
	a.Equal(3, util.PositiveMod(-2, 5), "ｰ2 mod 5")

	// --- Float tests ---
	a.InDelta(2.5, util.PositiveMod(7.5, 5.0), 1e-9, "float: 7.5 mod 5.0")
	a.InDelta(2.5, util.PositiveMod(-7.5, 5.0), 1e-9, "float: -7.5 mod 5.0")
	a.InDelta(0.0, util.PositiveMod(10.0, 5.0), 1e-9, "float: 10.0 mod 5.0")
	a.InDelta(0.0, util.PositiveMod(-10.0, 5.0), 1e-9, "float: -10.0 mod 5.0")
	a.InDelta(1.0, util.PositiveMod(-5.0, 3.0), 1e-9, "float: -5.0 mod 3.0")
}

func TestFloor(t *testing.T) {
	a := assert.New(t)

	// Test with positive float
	a.Equal(int(3), util.Floor[int](3.14), "Floor(3.14)")
	a.Equal(int(3), util.Floor[int](3.0), "Floor(3.0)")

	// Test with negative float
	a.Equal(int(-4), util.Floor[int](-3.14), "Floor(-3.14)")
	a.Equal(int(-3), util.Floor[int](-3.0), "Floor(-3.0)")

	// Test with different integer types
	a.Equal(int64(3), util.Floor[int64](3.14), "Floor(3.14) to int64")
	a.Equal(int32(-4), util.Floor[int32](-3.14), "Floor(-3.14) to int32")

	// Test with different float types
	a.Equal(int(3), util.Floor[int](float32(3.14)), "Floor(float32(3.14))")
}

func TestCeil(t *testing.T) {
	a := assert.New(t)

	// Test with positive float
	a.Equal(int(4), util.Ceil[int](3.14), "Ceil(3.14)")
	a.Equal(int(3), util.Ceil[int](3.0), "Ceil(3.0)")

	// Test with negative float
	a.Equal(int(-3), util.Ceil[int](-3.14), "Ceil(-3.14)")
	a.Equal(int(-3), util.Ceil[int](-3.0), "Ceil(-3.0)")

	// Test with different integer types
	a.Equal(int64(4), util.Ceil[int64](3.14), "Ceil(3.14) to int64")
	a.Equal(int32(-3), util.Ceil[int32](-3.14), "Ceil(-3.14) to int32")

	// Test with different float types
	a.Equal(int(4), util.Ceil[int](float32(3.14)), "Ceil(float32(3.14))")
}

func TestRound(t *testing.T) {
	a := assert.New(t)

	// Test with positive float
	a.Equal(int(3), util.Round[int](3.14), "Round(3.14)")
	a.Equal(int(3), util.Round[int](3.0), "Round(3.0)")
	a.Equal(int(4), util.Round[int](3.5), "Round(3.5)")
	a.Equal(int(4), util.Round[int](3.6), "Round(3.6)")

	// Test with negative float
	a.Equal(int(-3), util.Round[int](-3.14), "Round(-3.14)")
	a.Equal(int(-3), util.Round[int](-3.0), "Round(-3.0)")
	a.Equal(int(-4), util.Round[int](-3.5), "Round(-3.5)")
	a.Equal(int(-4), util.Round[int](-3.6), "Round(-3.6)")

	// Test with different integer types
	a.Equal(int64(3), util.Round[int64](3.14), "Round(3.14) to int64")
	a.Equal(int32(-4), util.Round[int32](-3.5), "Round(-3.5) to int32")

	// Test with different float types
	a.Equal(int(4), util.Round[int](float32(3.5)), "Round(float32(3.5))")
}

func TestAbs(t *testing.T) {
	a := assert.New(t)

	// Test with positive integers
	a.Equal(5, util.Abs(5), "Abs(5)")
	// Test with negative integers
	a.Equal(5, util.Abs(-5), "Abs(-5)")
	// Test with zero
	a.Equal(0, util.Abs(0), "Abs(0)")

	// Test with positive floats
	a.Equal(5.5, util.Abs(5.5), "Abs(5.5)")
	// Test with negative floats
	a.Equal(5.5, util.Abs(-5.5), "Abs(-5.5)")
	// Test with zero float
	a.Equal(0.0, util.Abs(0.0), "Abs(0.0)")

	// Test with different types
	a.Equal(int64(10), util.Abs(int64(-10)), "Abs(int64(-10))")
	a.Equal(float32(10.5), util.Abs(float32(-10.5)), "Abs(float32(-10.5))")
}

func TestTrunc(t *testing.T) {
	a := assert.New(t)

	// Test with positive float
	a.Equal(int(3), util.Trunc[int](3.14), "Trunc(3.14)")
	a.Equal(int(3), util.Trunc[int](3.0), "Trunc(3.0)")

	// Test with negative float
	a.Equal(int(-3), util.Trunc[int](-3.14), "Trunc(-3.14)")
	a.Equal(int(-3), util.Trunc[int](-3.0), "Trunc(-3.0)")

	// Test with different integer types
	a.Equal(int64(3), util.Trunc[int64](3.14), "Trunc(3.14) to int64")
	a.Equal(int32(-3), util.Trunc[int32](-3.14), "Trunc(-3.14) to int32")

	// Test with different float types
	a.Equal(int(3), util.Trunc[int](float32(3.14)), "Trunc(float32(3.14))")
}

func TestRoundToEven(t *testing.T) {
	a := assert.New(t)

	// Test with positive float
	a.Equal(int(4), util.RoundToEven[int](3.5), "RoundToEven(3.5)")
	a.Equal(int(4), util.RoundToEven[int](4.5), "RoundToEven(4.5)")
	a.Equal(int(3), util.RoundToEven[int](3.14), "RoundToEven(3.14)")

	// Test with negative float
	a.Equal(int(-4), util.RoundToEven[int](-3.5), "RoundToEven(-3.5)")
	a.Equal(int(-4), util.RoundToEven[int](-4.5), "RoundToEven(-4.5)")
	a.Equal(int(-3), util.RoundToEven[int](-3.14), "RoundToEven(-3.14)")

	// Test with different integer types
	a.Equal(int64(4), util.RoundToEven[int64](3.5), "RoundToEven(3.5) to int64")
	a.Equal(int32(-4), util.RoundToEven[int32](-3.5), "RoundToEven(-3.5) to int32")

	// Test with different float types
	a.Equal(int(4), util.RoundToEven[int](float32(3.5)), "RoundToEven(float32(3.5))")
}

func TestMathModBehavior(t *testing.T) {
	a := assert.New(t)

	// math.Mod(x, y) computes the remainder of x/y.
	// The sign of the result is the same as the sign of x (the dividend).

	// Positive dividend -> Positive result
	a.InDelta(2.0, math.Mod(5.0, 3.0), 1e-9, "5 mod 3")
	a.InDelta(2.0, math.Mod(5.0, -3.0), 1e-9, "5 mod -3")

	// Negative dividend -> Negative result
	a.InDelta(-2.0, math.Mod(-5.0, 3.0), 1e-9, "-5 mod 3")
	a.InDelta(-2.0, math.Mod(-5.0, -3.0), 1e-9, "-5 mod -3")

	// The case from the Repeat function discussion
	a.InDelta(-2.0, math.Mod(-2.0, 5.0), 1e-9, "-2 mod 5")
}

func TestPositiveModCustomTypes(t *testing.T) {
	a := assert.New(t)

	// Test with custom integer type
	var mi1 MyInt = 5
	var mi2 MyInt = 3
	var mi3 MyInt = -5
	var mi4 MyInt = 3
	a.Equal(MyInt(2), util.PositiveMod(mi1, mi2), "MyInt: 5 mod 3")
	a.Equal(MyInt(1), util.PositiveMod(mi3, mi4), "MyInt: -5 mod 3")

	// Test with custom float type
	var mf1 MyFloat = 7.5
	var mf2 MyFloat = 5.0
	var mf3 MyFloat = -7.5
	var mf4 MyFloat = 5.0
	a.Equal(MyFloat(2.5), util.PositiveMod(mf1, mf2), "MyFloat: 7.5 mod 5.0")
	a.Equal(MyFloat(2.5), util.PositiveMod(mf3, mf4), "MyFloat: -7.5 mod 5.0")
}
