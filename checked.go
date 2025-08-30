// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"unsafe"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Add[T Integer](a, b T) (T, bool) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, true
	}
	return 0, false
}

func Sub[T Integer](a, b T) (T, bool) {
	c := a - b
	if (c < a) == (b > 0) {
		return c, true
	}
	return 0, false
}

func Div[T Integer](a, b T) (T, bool) {
	q, _, ok := Quotient(a, b)
	return q, ok
}

func Quotient[T Integer](a, b T) (T, T, bool) {
	if b == 0 {
		return 0, 0, false
	}
	minusOne := T(0) - 1
	isSigned := minusOne < 0
	if isSigned && a == minusOne<<(unsafe.Sizeof(minusOne)*8-1) && b == minusOne {
		return 0, 0, false
	}
	return a / b, a % b, true
}

func Cast[Y, X Integer](x X) (y Y, ok bool) {
	hasSameType := unsafe.Sizeof(x) == unsafe.Sizeof(y) && (X(0)-1 < 0) == (Y(0)-1 < 0)
	y = Y(x)
	if X(y) == x && (hasSameType || (x < 0) == (y < 0)) {
		return y, true
	}
	return 0, false
}

func Min[T Integer]() T {
	if minusOne := T(0) - 1; minusOne > 0 { // signed integer
		return minusOne << (unsafe.Sizeof(minusOne)*8 - 1)
	}
	return T(0)
}

func Max[T Integer]() T {
	if minusOne := T(0) - 1; minusOne > 0 { // signed integer
		return 1<<(unsafe.Sizeof(minusOne)*8-1) - 1
	}
	return ^T(0)
}
