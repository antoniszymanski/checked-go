// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"iter"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	test(t, Add, "+", func(a int64, b int64) int64 { return a + b })
}

func TestSub(t *testing.T) {
	test(t, Sub, "-", func(a int64, b int64) int64 { return a - b })
}

func TestMul(t *testing.T) {
	testWithCases(t, Mul, Mul, "*", func(a int64, b int64) int64 { return a * b },
		TestCase{
			A:    math.MinInt64 / 2,
			B:    2,
			Want: T2(int64(math.MinInt64), true),
		},
		TestCase{
			A:    math.MinInt64,
			B:    1,
			Want: T2(int64(math.MinInt64), true),
		},
		TestCase{
			A: math.MinInt64,
			B: -1,
		})
}

func TestDiv(t *testing.T) {
	test(t, Div, "/", func(a int64, b int64) int64 { return a / b })
}

func TestMin(t *testing.T) {
	assertEqual(t, "Min[uint8]()", Min[uint8](), 0)
	assertEqual(t, "Min[uint16]()", Min[uint16](), 0)
	assertEqual(t, "Min[uint32]()", Min[uint32](), 0)
	assertEqual(t, "Min[uint64]()", Min[uint64](), 0)
	assertEqual(t, "Min[int8]()", Min[int8](), math.MinInt8)
	assertEqual(t, "Min[int16]()", Min[int16](), math.MinInt16)
	assertEqual(t, "Min[int32]()", Min[int32](), math.MinInt32)
	assertEqual(t, "Min[int64]()", Min[int64](), math.MinInt64)
}

func TestMax(t *testing.T) {
	assertEqual(t, "Max[uint8]()", Max[uint8](), math.MaxUint8)
	assertEqual(t, "Max[uint16]()", Max[uint16](), math.MaxUint16)
	assertEqual(t, "Max[uint32]()", Max[uint32](), math.MaxUint32)
	assertEqual(t, "Max[uint64]()", Max[uint64](), math.MaxUint64)
	assertEqual(t, "Max[int8]()", Max[int8](), math.MaxInt8)
	assertEqual(t, "Max[int16]()", Max[int16](), math.MaxInt16)
	assertEqual(t, "Max[int32]()", Max[int32](), math.MaxInt32)
	assertEqual(t, "Max[int64]()", Max[int64](), math.MaxInt64)
}

func test(
	t *testing.T,
	fn func(int8, int8) (int8, bool),
	opName string, op func(a, b int64) int64,
) {
	var errors int
	for a, b := range int8s {
		if errors >= 10 {
			break
		}
		got := T2(fn(a, b))
		want := T2(expected(a, b, op))
		if got != want {
			t.Errorf("%v %s %v: got %v want %v", a, opName, b, got, want)
			errors++
		}
	}
}

func testWithCases(
	t *testing.T,
	fn8 func(int8, int8) (int8, bool), fn64 func(int64, int64) (int64, bool),
	opName string, op func(a, b int64) int64,
	cases ...TestCase,
) {
	var errors int
	for a, b := range int8s {
		if errors >= 10 {
			break
		}
		got := T2(fn8(a, b))
		want := T2(expected(a, b, op))
		if got != want {
			t.Errorf("%v %s %v: got %v want %v", a, opName, b, got, want)
			errors++
		}
	}
	for _, tc := range cases {
		if errors >= 10 {
			break
		}
		got := T2(fn64(tc.A, tc.B))
		if got != tc.Want {
			t.Errorf("%v %s %v: got %v want %v", tc.A, opName, tc.B, got, tc.Want)
			errors++
		}
	}
}

func assertEqual[T comparable](t *testing.T, op string, want, got T) {
	if got != want {
		t.Errorf("%s should equal (%v), not (%v)", op, got, want)
	}
}

func int8s(yield func(int8, int8) bool) {
	cartesianIndices(int8(math.MinInt8), math.MaxInt8)(yield)
}

func cartesianIndices[T Integer](min, max T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for a := min; ; a++ {
			for b := min; ; b++ {
				if !yield(a, b) {
					return
				}
				if b >= max {
					break
				}
			}
			if a >= max {
				break
			}
		}
	}
}

func expected(a, b int8, op func(a, b int64) int64) (int8, bool) {
	defer func() { _ = recover() }()
	c := op(int64(a), int64(b))
	if math.MinInt8 <= c && c <= math.MaxInt8 {
		return int8(c), true
	}
	return 0, false
}

type TestCase struct {
	A, B int64
	Want Tuple2[int64, bool]
}

type Tuple2[A, B any] struct {
	A A
	B B
}

func T2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}
