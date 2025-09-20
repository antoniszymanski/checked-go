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

func Benchmark(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(Add(a, b))
		}
	})
	b.Run("bits.Add64", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(bits.Add64(a, b, 0))
		}
	})
	b.Run("Sub", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(Sub(a, b))
		}
	})
	b.Run("bits.Sub64", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(bits.Sub64(a, b, 0))
		}
	})
	b.Run("Mul", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(Mul(a, b))
		}
	})
	b.Run("mulUint64", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(mulUint64(a, b))
		}
	})
	b.Run("bits.Mul64", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(bits.Mul64(a, b))
		}
	})
	b.Run("Div", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(Div(a, b))
		}
	})
	b.Run("bits.Div64", func(b *testing.B) {
		for a, b := range cartesianIndices[uint64](1, 10_000) {
			keepAlive(bits.Div64(0, a, b))
		}
	})
}

func keepAlive(args ...any) {
	runtime.KeepAlive(args)
}
