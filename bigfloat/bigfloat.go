// Package bigfloat implements some helper functions for big.Float
package bigfloat

import (
	"math/big"
)

var zero = big.NewFloat(0)

const (
	cmpLt  int = -1
	cmpEq  int = 0
	cmpGt  int = 1
	cmpNil int = -42
)

// Add returns the result of adding the values of params x and y.
// Notes:
// - If x == nil && y == nil, returns nil.
// - If x != nil && y == nil, returns x.
// - If x == nil && y != nil, returns y.
func Add(x, y *big.Float) *big.Float {
	switch {
	case x == nil && y == nil:
		return nil
	case x == nil && y != nil:
		return y
	case x != nil && y == nil:
		return x
	}

	return newBigFloat().Add(x, y)
}

// Sub returns the result of subtracting the value of param y from the value of param x.
// Notes:
// - If x == nil, returns nil.
// - If x != nil && y == nil, returns x.
func Sub(x, y *big.Float) *big.Float {
	switch {
	case x == nil:
		return nil
	case x != nil && y == nil:
		return x
	}

	return newBigFloat().Sub(x, y)
}

// Mul returns the result of multiplying the values of params x and y.
// Notes:
// - If x == nil, returns nil.
// - If x != nil && y == nil, returns x.
func Mul(x, y *big.Float) *big.Float {
	switch {
	case x == nil:
		return nil
	case x != nil && y == nil:
		return x
	}

	return newBigFloat().Mul(x, y)
}

// Div returns the result of dividing the value of x by the value of y.
// Notes:
// - If x == nil, returns nil.
// - If x != nil && y == nil, returns x.
func Div(x, y *big.Float) *big.Float {
	switch {
	case x == nil:
		return nil
	case x != nil && y == nil:
		return x
	}

	return newBigFloat().Quo(x, y)
}

// Cmp wraps the Int.Cmp() function, returning the same values as the function it wraps:
// - -1 if x < y
// - 0 if x == y
// - 1 if x > y
// Note: returns -2 if either x or y == nil.
func Cmp(x, y *big.Float) int {
	if x == nil || y == nil {
		return cmpNil
	}

	return x.Cmp(y)
}

// Eq returns true if x == y.
// Returns false if either x or y == nil.
func Eq(x, y *big.Float) bool {
	return Cmp(x, y) == cmpEq
}

// Lt returns true if x < y.
// Returns false if either x or y == nil.
func Lt(x, y *big.Float) bool {
	return Cmp(x, y) == cmpLt
}

// Lte returns true if x <= y.
// Returns false if either x or y == nil.
func Lte(x, y *big.Float) bool {
	return Lt(x, y) || Eq(x, y)
}

// Gt returns true if x > y.
// Returns false if either x or y == nil.
func Gt(x, y *big.Float) bool {
	return Cmp(x, y) == cmpGt
}

// Gte returns true if x >= y.
// Returns false if either x or y == nil.
func Gte(x, y *big.Float) bool {
	return Gt(x, y) || Eq(x, y)
}

// IsZero returns true if n == 0.
// Returns nil if n == nil.
func IsZero(n *big.Float) bool {
	return Eq(zero, n)
}

// FromUint64 returns a new *big.Int with its value set to n.
func FromUint64(n uint64) *big.Float {
	return newBigFloat().SetUint64(n)
}

// FromBigInt returns a new *big.Float with is value set to the value of n.
func FromBigInt(n *big.Int) *big.Float {
	return newBigFloat().SetInt(n)
}

// FromFloat32 returns a new *big.Float with is value set to f.
func FromFloat32(f float32) *big.Float {
	return FromFloat64(float64(f))
}

// FromFloat64 returns a new *big.Float with is value set to f.
func FromFloat64(f float64) *big.Float {
	return big.NewFloat(f)
}

func newBigFloat() *big.Float {
	return new(big.Float)
}
