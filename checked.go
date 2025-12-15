// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"unsafe"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Add[T Integer](x, y T) (T, bool) {
	z := x + y
	if (z > x) == (y > 0) {
		return z, true
	}
	return 0, false
}

func Sub[T Integer](x, y T) (T, bool) {
	z := x - y
	if (z < x) == (y > 0) {
		return z, true
	}
	return 0, false
}

func Div[T Integer](x, y T) (T, bool) {
	if y == 0 {
		return 0, false
	}
	minusOne := T(0) - 1
	isSigned := minusOne < 0
	if isSigned && x == minusOne<<(unsafe.Sizeof(minusOne)*8-1) && y == minusOne {
		return 0, false
	}
	return x / y, true
}

func DivMod[T Integer](x, y T) (T, T, bool) {
	if y == 0 {
		return 0, 0, false
	}
	minusOne := T(0) - 1
	isSigned := minusOne < 0
	if isSigned && x == minusOne<<(unsafe.Sizeof(minusOne)*8-1) && y == minusOne {
		return 0, 0, false
	}
	return x / y, x % y, true
}

func Cast[Y, X Integer](x X) (Y, bool) {
	y := Y(x)
	hasSameType := unsafe.Sizeof(x) == unsafe.Sizeof(y) && (X(0)-1 < 0) == (Y(0)-1 < 0)
	if X(y) == x && (hasSameType || (x < 0) == (y < 0)) {
		return y, true
	}
	return 0, false
}

func Min[T Integer]() T {
	if minusOne := T(0) - 1; minusOne < 0 { // signed integer
		return minusOne << (unsafe.Sizeof(minusOne)*8 - 1)
	}
	return T(0)
}

func Max[T Integer]() T {
	if minusOne := T(0) - 1; minusOne < 0 { // signed integer
		return 1<<(unsafe.Sizeof(minusOne)*8-1) - 1
	}
	return ^T(0)
}
