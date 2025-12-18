// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"math"
	"math/bits"
	"unsafe"
)

func Mul[T Integer](x, y T) (T, bool) {
	isUnsigned := T(0)-1 > 0
	switch unsafe.Sizeof(T(0)) {
	case unsafe.Sizeof(int8(0)):
		if isUnsigned {
			z, ok := mulUint8(uint8(x), uint8(y))
			return T(z), ok
		} else {
			z, ok := mulInt8(int8(x), int8(y))
			return T(z), ok
		}
	case unsafe.Sizeof(int16(0)):
		if isUnsigned {
			z, ok := mulUint16(uint16(x), uint16(y))
			return T(z), ok
		} else {
			z, ok := mulInt16(int16(x), int16(y))
			return T(z), ok
		}
	case unsafe.Sizeof(int32(0)):
		if isUnsigned {
			z, ok := mulUint32(uint32(x), uint32(y))
			return T(z), ok
		} else {
			z, ok := mulInt32(int32(x), int32(y))
			return T(z), ok
		}
	case unsafe.Sizeof(int64(0)):
		if isUnsigned {
			z, ok := mulUint64(uint64(x), uint64(y))
			return T(z), ok
		} else {
			z, ok := mulInt64(int64(x), int64(y))
			return T(z), ok
		}
	default:
		panic("unreachable")
	}
}

func mulUint8(x, y uint8) (uint8, bool) {
	z := uint16(x) * uint16(y)
	if z <= math.MaxUint8 {
		return uint8(z), true
	} else {
		return 0, false
	}
}

func mulInt8(x, y int8) (int8, bool) {
	z := int16(x) * int16(y)
	if math.MinInt8 <= z && z <= math.MaxInt8 {
		return int8(z), true
	} else {
		return 0, false
	}
}

func mulUint16(x, y uint16) (uint16, bool) {
	z := uint32(x) * uint32(y)
	if z <= math.MaxUint16 {
		return uint16(z), true
	} else {
		return 0, false
	}
}

func mulInt16(x, y int16) (int16, bool) {
	z := int32(x) * int32(y)
	if math.MinInt16 <= z && z <= math.MaxInt16 {
		return int16(z), true
	} else {
		return 0, false
	}
}

func mulUint32(x, y uint32) (uint32, bool) {
	z := uint64(x) * uint64(y)
	if z <= math.MaxInt32 {
		return uint32(z), true
	} else {
		return 0, false
	}
}

func mulInt32(x, y int32) (int32, bool) {
	z := int64(x) * int64(y)
	if math.MinInt32 <= z && z <= math.MaxInt32 {
		return int32(z), true
	} else {
		return 0, false
	}
}

func mulUint64(x, y uint64) (uint64, bool) {
	hi, lo := bits.Mul64(uint64(x), uint64(y))
	if hi == 0 {
		return lo, true
	} else {
		return 0, false
	}
}

func mulInt64(x, y int64) (int64, bool) {
	neg := (x < 0) != (y < 0)
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	hi, lo := bits.Mul64(uint64(x), uint64(y)) //#nosec G115
	if hi != 0 {
		return 0, false
	} else if lo > math.MaxInt64 {
		if neg && lo == -math.MinInt64 {
			return math.MinInt64, true
		}
		return 0, false
	}
	z := int64(lo)
	if neg {
		z = -z
	}
	return z, true
}
