// 用于快速转换Bytes到类型或类型到bytes
// 非自写算法，常规C语言写法位移
// version 1.0 beta
// by koangel
// email: jackliu100@gmail.com
// 2017/7/8
package grapeStream

import "math"

// Bytes 2 Values
func BTUint8(buf []byte) uint8 {
	return uint8(buf[0])
}

func BTInt8(buf []byte) int8 {
	return int8(buf[0])
}

func BTUint16(buf []byte) (r uint16) {
	r |= uint16(buf[0])
	r |= uint16(buf[1]) << 8
	return
}

func BTInt16(buf []byte) (r int16) {
	r |= int16(buf[0])
	r |= int16(buf[1]) << 8
	return
}

func BTUint32(buf []byte) (r uint32) {
	r |= uint32(buf[0])
	r |= uint32(buf[1]) << 8
	r |= uint32(buf[2]) << 16
	r |= uint32(buf[3]) << 24
	return
}

func BTInt32(buf []byte) (r int32) {
	r |= int32(buf[0])
	r |= int32(buf[1]) << 8
	r |= int32(buf[2]) << 16
	r |= int32(buf[3]) << 24
	return
}

func BTUint64(buf []byte) (r uint64) {
	r |= uint64(buf[0])
	r |= uint64(buf[1]) << 8
	r |= uint64(buf[2]) << 16
	r |= uint64(buf[3]) << 24
	r |= uint64(buf[4]) << 32
	r |= uint64(buf[5]) << 40
	r |= uint64(buf[6]) << 48
	r |= uint64(buf[7]) << 56
	return
}

func BTInt64(buf []byte) (r int64) {
	r |= int64(buf[0])
	r |= int64(buf[1]) << 8
	r |= int64(buf[2]) << 16
	r |= int64(buf[3]) << 24
	r |= int64(buf[4]) << 32
	r |= int64(buf[5]) << 40
	r |= int64(buf[6]) << 48
	r |= int64(buf[7]) << 56
	return
}

func BTFloat32(buf []byte) float32 {
	return math.Float32frombits(BTUint32(buf))
}

func BTFloat64(buf []byte) float64 {
	return math.Float64frombits(BTUint64(buf))
}

// Values 2 bytes
func U8TBytes(v uint8) []byte {
	buf := make([]byte, 1)
	buf[0] = byte(v)

	return buf
}

func I8TBytes(v int8) []byte {
	buf := make([]byte, 1)
	buf[0] = byte(v)
	return buf
}

func U16TBytes(v uint16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	return buf
}

func I16TBytes(v int16) []byte {
	buf := make([]byte, 2)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	return buf
}

func U32TBytes(v uint32) []byte {
	buf := make([]byte, 4)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return buf
}

func I32TBytes(v int32) []byte {
	buf := make([]byte, 4)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	return buf
}

func U64TBytes(v uint64) []byte {
	buf := make([]byte, 8)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)

	return buf
}

func I64TBytes(v int64) []byte {
	buf := make([]byte, 8)
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)

	return buf
}

func F32TBytes(v float32) []byte {
	return U32TBytes(math.Float32bits(v))
}

func F64TBytes(v float64) []byte {
	return U64TBytes(math.Float64bits(v))
}

func A62toi(s string) int {
	byteData := []byte(s)
	var ret = 0
	var minus = 1
	var nIndex = 0

	if byteData[0] == '-' {
		minus = -1
		nIndex++
	}

	for _, v := range byteData[nIndex:] {
		ret *= 62
		if '0' <= v && v <= '9' {
			ret += int(v - '0')
		} else if 'a' <= v && v <= 'z' {
			ret += int(v - 'a' + 10)
		} else if 'A' <= v && v <= 'Z' {
			ret += int(v - 'A' + 36)
		} else {
			return 0
		}
	}

	return ret * minus
}

var b62Base = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CNV10to62(v int) string {
	var src = 0
	var minus = 0
	var baselen = len(b62Base)

	var tmp = [64]int{}
	var bOut = [128]byte{}

	if v < 0 {
		minus = 1
		v *= -1
	}

	if v < baselen {
		if minus == 1 {
			return "-" + string(b62Base[v])
		} else {
			return string(b62Base[v])
		}
	}

	src = v
	var i = 0
	var nCnt = 0
	for i = 0; src >= baselen; i++ {
		tmp[i] = src % baselen
		src /= baselen
	}
	i--

	if minus == 1 {
		bOut[0] = byte('-')
		bOut[1] = byte(b62Base[src])
		nCnt += 2
		for j := 2; i >= 0; i-- {
			bOut[j] = byte(b62Base[tmp[i]])
			nCnt++
			j++
		}
	} else {
		bOut[0] = byte(b62Base[src])
		nCnt += 1
		for j := 1; i >= 0; i-- {
			bOut[j] = byte(b62Base[tmp[i]])
			nCnt++
			j++
		}
	}

	return string(bOut[:nCnt])
}
