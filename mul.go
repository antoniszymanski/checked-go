// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package checked

import (
	"math"
	"math/bits"
	"unsafe"
)

func Mul[T Integer](a, b T) (T, bool) {
	isUnsigned := T(0)-1 > 0
	switch unsafe.Sizeof(T(0)) {
	case unsafe.Sizeof(int8(0)):
		if isUnsigned {
			res, ok := mulUint8(uint8(a), uint8(b))
			return T(res), ok
		} else {
			res, ok := mulInt8(int8(a), int8(b))
			return T(res), ok
		}
	case unsafe.Sizeof(int16(0)):
		if isUnsigned {
			res, ok := mulUint16(uint16(a), uint16(b))
			return T(res), ok
		} else {
			res, ok := mulInt16(int16(a), int16(b))
			return T(res), ok
		}
	case unsafe.Sizeof(int32(0)):
		if isUnsigned {
			res, ok := mulUint32(uint32(a), uint32(b))
			return T(res), ok
		} else {
			res, ok := mulInt32(int32(a), int32(b))
			return T(res), ok
		}
	case unsafe.Sizeof(int64(0)):
		if isUnsigned {
			res, ok := mulUint64(uint64(a), uint64(b))
			return T(res), ok
		} else {
			res, ok := mulInt64(int64(a), int64(b))
			return T(res), ok
		}
	default:
		panic("unreachable")
	}
}

func mulUint8[T ~uint8](a, b T) (T, bool) {
	res := uint16(a) * uint16(b)
	if res <= math.MaxUint8 {
		return T(res), true
	} else {
		return 0, false
	}
}

func mulInt8[T ~int8](a, b T) (T, bool) {
	res := int16(a) * int16(b)
	if math.MinInt8 <= res && res <= math.MaxInt8 {
		return T(res), true
	} else {
		return 0, false
	}
}

func mulUint16[T ~uint16](a, b T) (T, bool) {
	res := uint32(a) * uint32(b)
	if res <= math.MaxUint16 {
		return T(res), true
	} else {
		return 0, false
	}
}

func mulInt16[T ~int16](a, b T) (T, bool) {
	res := int32(a) * int32(b)
	if math.MinInt16 <= res && res <= math.MaxInt16 {
		return T(res), true
	} else {
		return 0, false
	}
}

func mulUint32[T ~uint32](a, b T) (T, bool) {
	res := uint64(a) * uint64(b)
	if res <= math.MaxInt32 {
		return T(res), true
	} else {
		return 0, false
	}
}

func mulInt32[T ~int32](a, b T) (T, bool) {
	res := int64(a) * int64(b)
	if math.MinInt32 <= res && res <= math.MaxInt32 {
		return T(res), true
	} else {
		return 0, false
	}
}

func mulUint64[T ~uint64](a, b T) (T, bool) {
	hi, lo := bits.Mul64(uint64(a), uint64(b))
	if hi == 0 {
		return T(lo), true
	} else {
		return 0, false
	}
}

func mulInt64[T ~int64](a, b T) (T, bool) {
	neg := (a < 0) != (b < 0)
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	hi, lo := bits.Mul64(uint64(a), uint64(b)) //#nosec G115
	if hi != 0 {
		return 0, false
	} else if lo > math.MaxInt64 {
		if neg && lo == -math.MinInt64 {
			return math.MinInt64, true
		}
		return 0, false
	}
	res := int64(lo)
	if neg {
		res = -res
	}
	return T(res), true
}
