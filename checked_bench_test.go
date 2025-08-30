// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"math/bits"
	"os"
	"runtime"
	"testing"
)

func init() {
	os.Args = append(os.Args, "-test.benchtime=1x")
}

func BenchmarkAddBits(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(bits.Add64(a, b, 0))
	}
}

func BenchmarkAdd(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(Add(a, b))
	}
}

func BenchmarkSubBits(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(bits.Sub64(a, b, 0))
	}
}

func BenchmarkSub(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(Sub(a, b))
	}
}

func BenchmarkMulBits(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(bits.Mul64(a, b))
	}
}

func BenchmarkMul(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(Mul(a, b))
	}
}

func BenchmarkDivBits(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(bits.Div64(0, a, b))
	}
}

func BenchmarkDiv(b *testing.B) {
	for a, b := range numbers[uint64](1, 10_000) {
		use(Div(a, b))
	}
}

func use(args ...any) {
	runtime.KeepAlive(args)
}
