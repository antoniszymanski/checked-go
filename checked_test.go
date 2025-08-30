// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"iter"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	var errors int
	for a, b := range numbers(Min[int8](), Max[int8]()) {
		if errors >= 10 {
			break
		}
		actual, actualOk := Add(a, b)
		expected := int64(a) + int64(b)
		expectedOk := inRange(&expected)
		if actualOk != expectedOk {
			t.Errorf("%v + %v should equal (%v, %v), not (%v, %v)", a, b, expected, expectedOk, actual, actualOk)
			errors++
		}
	}
}

func TestSub(t *testing.T) {
	var errors int
	for a, b := range numbers(Min[int8](), Max[int8]()) {
		if errors >= 10 {
			break
		}
		actual, actualOk := Sub(a, b)
		expected := int64(a) - int64(b)
		expectedOk := inRange(&expected)
		if actualOk != expectedOk {
			t.Errorf("%v - %v should equal (%v, %v), not (%v, %v)", a, b, expected, expectedOk, actual, actualOk)
			errors++
		}
	}
}

func TestMul(t *testing.T) {
	var errors int
	for a, b := range numbers(Min[int8](), Max[int8]()) {
		if errors >= 10 {
			break
		}
		actual, actualOk := Mul(a, b)
		expected := int64(a) * int64(b)
		expectedOk := inRange(&expected)
		if actualOk != expectedOk {
			t.Errorf("%v * %v should equal (%v, %v), not (%v, %v)", a, b, expected, expectedOk, actual, actualOk)
			errors++
		}
	}

	type TestCase struct {
		A, B, C int64
		OK      bool
	}
	for _, tc := range []TestCase{
		{
			A:  math.MinInt64 / 2,
			B:  2,
			C:  math.MinInt64,
			OK: true,
		},
		{
			A:  math.MinInt64,
			B:  1,
			C:  math.MinInt64,
			OK: true,
		},
		{
			A:  math.MinInt64,
			B:  -1,
			OK: false,
		},
	} {
		c, ok := Mul(tc.A, tc.B)
		if ok != tc.OK {
			t.Errorf("%v * %v should equal (%v, %v), not (%v, %v)", tc.A, tc.B, tc.C, tc.OK, c, ok)
		}
	}
}

func TestDiv(t *testing.T) {
	var errors int
	for a, b := range numbers(Min[int8](), Max[int8]()) {
		if errors >= 10 {
			break
		}
		actual, actualOk := Div(a, b)
		var expected int64
		var expectedOk bool
		if b != 0 {
			expected = int64(a) / int64(b)
			expectedOk = inRange(&expected)
		}
		if actualOk != expectedOk {
			t.Errorf("%v / %v should equal (%v, %v), not (%v, %v)", a, b, expected, expectedOk, actual, actualOk)
			errors++
		}
	}
}

// func TestQuotient(t *testing.T) {
// 	q, r, ok := Quotient(100, 3)
// 	if r != 1 || q != 33 || !ok {
// 		t.Errorf("expected 100/3 => 33, r=1")
// 	}
// 	if _, _, ok = Quotient(1, 0); ok {
// 		t.Error("unexpected lack of failure")
// 	}
// }

func inRange(x *int64) bool {
	ok := math.MinInt8 <= *x && *x <= math.MaxInt8
	if !ok {
		*x = 0
	}
	return ok
}

func numbers[T Integer](min, max T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for a := min; a < max; a++ {
			for b := min; b < max; b++ {
				if !yield(a, b) {
					return
				}
			}
		}
	}
}
