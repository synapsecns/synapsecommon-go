// Package bigint implements some helper functions for big.Int
package bigint

import (
	"math/big"
)

var zero = big.NewInt(0)

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
func Add(x, y *big.Int) *big.Int {
	if x == nil {
		if y != nil {
			return y
		}

		return nil
	} else if y == nil {
		if x != nil {
			return x
		}

		return nil
	}

	return newBigInt().Add(x, y)
}

// Sub returns the result of subtracting the value of param y from the value of param x.
// Notes:
// - If x == nil, returns nil.
// - If x != nil && y == nil, returns x.
func Sub(x, y *big.Int) *big.Int {
	if x == nil {
		return nil
	} else if x != nil && y == nil {
		return x
	}

	return newBigInt().Sub(x, y)
}

// Mul returns the result of multiplying the values of params x and y.
// Notes:
// - If x == nil, returns nil.
// - If x != nil && y == nil, returns x.
func Mul(x, y *big.Int) *big.Int {
	if x == nil {
		return nil
	} else if x != nil && y == nil {
		return x
	}

	return newBigInt().Mul(x, y)
}

// Div returns the result of dividing the value of x by the value of y.
// Notes:
// - If x == nil, returns nil.
// - If x != nil && y == nil, returns x.
func Div(x, y *big.Int) *big.Int {
	if x == nil {
		return nil
	} else if x != nil && y == nil {
		return x
	}

	return newBigInt().Quo(x, y)
}

// Cmp wraps the Int.Cmp() function, returning the same values as the function it wraps:
// - -1 if x < y
// - 0 if x == y
// - 1 if x > y
// Note: returns -2 if either x or y == nil.
func Cmp(x, y *big.Int) int {
	if x == nil || y == nil {
		return cmpNil
	}

	return x.Cmp(y)
}

// Eq returns true if x == y.
// Returns false if either x or y == nil.
func Eq(x, y *big.Int) bool {
	return Cmp(x, y) == cmpEq
}

// Lt returns true if x < y.
// Returns false if either x or y == nil.
func Lt(x, y *big.Int) bool {
	return Cmp(x, y) == cmpLt
}

// Lte returns true if x <= y.
// Returns false if either x or y == nil.
func Lte(x, y *big.Int) bool {
	return Lt(x, y) || Eq(x, y)
}

// Gt returns true if x > y.
// Returns false if either x or y == nil.
func Gt(x, y *big.Int) bool {
	return Cmp(x, y) == cmpGt
}

// Gte returns true if x >= y.
// Returns false if either x or y == nil.
func Gte(x, y *big.Int) bool {
	return Gt(x, y) || Eq(x, y)
}

// IsZero returns true if n == 0.
// Returns nil if n == nil.
func IsZero(n *big.Int) bool {
	return Eq(zero, n)
}

// FromInt64 returns a new *big.Int with its value set to n.
func FromInt64(n int64) *big.Int {
	return newBigInt().SetInt64(n)
}

// FromUint64 returns a new *big.Int with its value set to n.
func FromUint64(n uint64) *big.Int {
	return newBigInt().SetUint64(n)
}

func newBigInt() *big.Int {
	return new(big.Int)
}
