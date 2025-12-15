package checked

import (
	"fmt"
	"iter"
	"math"
	"math/big"
	"math/bits"
	"os"
	"runtime"
	"testing"
)

// #region Test{Add,Sub,Div,Mul}

func TestAdd(t *testing.T) {
	testOp(t, "Add", Add, func(x, y int64) int64 { return x + y })
}

func TestSub(t *testing.T) {
	testOp(t, "Sub", Sub, func(x, y int64) int64 { return x - y })
}

func TestDiv(t *testing.T) {
	testOp(t, "Div", Div, func(x, y int64) int64 { return x / y })
}

func TestMul(t *testing.T) {
	testOp(t, "Mul", Mul, func(x, y int64) int64 { return x * y })
	testCases := []struct {
		X, Y, WantValue int64
		WantOk          bool
	}{
		{
			X:         math.MinInt64 / 2,
			Y:         2,
			WantValue: math.MinInt64,
			WantOk:    true,
		},
		{
			X:         math.MinInt64,
			Y:         1,
			WantValue: math.MinInt64,
			WantOk:    true,
		},
		{
			X: math.MinInt64,
			Y: -1,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("testCases[%d]", i), func(t *testing.T) {
			gotValue, gotOk := Mul(tc.X, tc.Y)
			checkResult2(t, "Mul", tc.X, tc.Y, gotValue, gotOk, tc.WantValue, tc.WantOk)
		})
	}
}

func testOp(
	t *testing.T,
	opName string,
	op func(x, y int8) (z int8, ok bool),
	opQword func(x, y int64) int64,
) {
	for x, y := range allNumbers[int8] {
		gotValue, gotOk := op(x, y)
		wantValue, wantOk := expectedQword(x, y, opQword)
		checkResult2(t, opName, x, y, gotValue, gotOk, wantValue, wantOk)
	}
}

func expectedQword(x, y int8, qwordOp func(x, y int64) int64) (int8, bool) {
	defer func() { _ = recover() }()
	z := qwordOp(int64(x), int64(y))
	if math.MinInt8 <= z && z <= math.MaxInt8 {
		return int8(z), true
	}
	return 0, false
}

// #endregion

// #region Test{Min,Max}

func TestMin(t *testing.T) {
	checkResult1(t, "Min[uint8]()", Min[uint8](), 0)
	checkResult1(t, "Min[uint16]()", Min[uint16](), 0)
	checkResult1(t, "Min[uint32]()", Min[uint32](), 0)
	checkResult1(t, "Min[uint64]()", Min[uint64](), 0)
	checkResult1(t, "Min[int8]()", Min[int8](), math.MinInt8)
	checkResult1(t, "Min[int16]()", Min[int16](), math.MinInt16)
	checkResult1(t, "Min[int32]()", Min[int32](), math.MinInt32)
	checkResult1(t, "Min[int64]()", Min[int64](), math.MinInt64)
}

func TestMax(t *testing.T) {
	checkResult1(t, "Max[uint8]()", Max[uint8](), math.MaxUint8)
	checkResult1(t, "Max[uint16]()", Max[uint16](), math.MaxUint16)
	checkResult1(t, "Max[uint32]()", Max[uint32](), math.MaxUint32)
	checkResult1(t, "Max[uint64]()", Max[uint64](), math.MaxUint64)
	checkResult1(t, "Max[int8]()", Max[int8](), math.MaxInt8)
	checkResult1(t, "Max[int16]()", Max[int16](), math.MaxInt16)
	checkResult1(t, "Max[int32]()", Max[int32](), math.MaxInt32)
	checkResult1(t, "Max[int64]()", Max[int64](), math.MaxInt64)
}

// #endregion

// #region Fuzz{Add,Sub,Div,Mul}

func FuzzAdd(f *testing.F) {
	fuzzOp(f, "Add", Add, (*big.Int).Add)
}

func FuzzSub(f *testing.F) {
	fuzzOp(f, "Sub", Sub, (*big.Int).Sub)
}

func FuzzDiv(f *testing.F) {
	fuzzOp(f, "Div", Div, (*big.Int).Quo)
}

func FuzzMul(f *testing.F) {
	fuzzOp(f, "Mul", Mul, (*big.Int).Mul)
}

func fuzzOp(
	f *testing.F,
	opName string,
	op func(x, y int64) (z int64, ok bool),
	opBig func(z *big.Int, x *big.Int, y *big.Int) *big.Int,
) {
	f.Fuzz(func(t *testing.T, x, y int64) {
		gotValue, gotOk := op(x, y)
		wantValue, wantOk := expectedBig(x, y, opBig)
		checkResult2(t, opName, x, y, gotValue, gotOk, wantValue, wantOk)
	})
}

func expectedBig(x, y int64, bigOp func(z *big.Int, x *big.Int, y *big.Int) *big.Int) (int64, bool) {
	defer func() { _ = recover() }()
	bigX := big.NewInt(x)
	bigY := big.NewInt(y)
	bigZ := bigOp(bigX, bigX, bigY)
	var wantValue int64
	wantOk := bigZ.IsInt64()
	if wantOk {
		wantValue = bigZ.Int64()
	}
	return wantValue, wantOk
}

// #endregion

// #region Benchmark

func init() {
	os.Args = append(os.Args, "-test.benchtime=1x")
}

func Benchmark(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(Add(a, b))
		}
	})
	b.Run("bits.Add64", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(bits.Add64(a, b, 0))
		}
	})
	b.Run("Sub", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(Sub(a, b))
		}
	})
	b.Run("bits.Sub64", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(bits.Sub64(a, b, 0))
		}
	})
	b.Run("Mul", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(Mul(a, b))
		}
	})
	b.Run("mulUint64", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(mulUint64(a, b))
		}
	})
	b.Run("bits.Mul64", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(bits.Mul64(a, b))
		}
	})
	b.Run("Div", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(Div(a, b))
		}
	})
	b.Run("bits.Div64", func(b *testing.B) {
		for a, b := range range2[uint64](1, 10_000) {
			keepAlive(bits.Div64(0, a, b))
		}
	})
}

func keepAlive(args ...any) {
	runtime.KeepAlive(args)
}

// #endregion

// #region common

func allNumbers[T Integer](yield func(T, T) bool) {
	range2(Min[T](), Max[T]())(yield)
}

func range2[T Integer](start, end T) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for x := start; ; x++ {
			for y := start; ; y++ {
				if !yield(x, y) {
					return
				}
				if y >= end {
					break
				}
			}
			if x >= end {
				break
			}
		}
	}
}

func checkResult1[T Integer](t *testing.T, opName string, want, got T) {
	if got != want {
		t.Errorf("%s = %d, should be %d", opName, got, want)
	}
}

func checkResult2[T Integer](t *testing.T, opName string, x, y, gotValue T, gotOk bool, wantValue T, wantOk bool) {
	if gotValue != wantValue || gotOk != wantOk {
		t.Errorf("%s(%d, %d) = (%d, %t), should be (%d, %t)", opName, x, y, gotValue, gotOk, wantValue, wantOk)
	}
}

// #endregion
