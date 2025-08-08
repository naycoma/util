package util_test

import (
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

	a.Equal(1, util.PositiveMod(1, 3), "1 mod 3")
	a.Equal(0, util.PositiveMod(3, 3), "3 mod 3")
	a.Equal(2, util.PositiveMod(5, 3), "5 mod 3")
	a.Equal(1, util.PositiveMod(-5, 3), "-5 mod 3")
	a.Equal(0, util.PositiveMod(-3, 3), "-3 mod 3")
	a.Equal(2, util.PositiveMod(-1, 3), "-1 mod 3")

	a.Equal(0, util.PositiveMod(0, 5), "0 mod 5")
	a.Equal(4, util.PositiveMod(-1, 5), "-1 mod 5")
	a.Equal(3, util.PositiveMod(-2, 5), "-2 mod 5")
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
