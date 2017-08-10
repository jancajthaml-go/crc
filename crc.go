package main

import "encoding/binary"

func Crc32(data []byte, polynomial uint32) []byte {
	var (
		inx    int    = 0
		crc    uint32 = 0 ^ 0xFFFFFFFF
		length int    = len(data)
		result        = make([]byte, 4)
		p7     uint32 = polynomial >> 1
		p6     uint32 = polynomial >> 2
		p5     uint32 = polynomial >> 3
		p4     uint32 = polynomial >> 4
		p3     uint32 = polynomial >> 5
		p2     uint32 = (polynomial >> 6) ^ polynomial
		p1     uint32 = (polynomial >> 7) ^ p7
	)

loop:
	if length == 0 {
		binary.LittleEndian.PutUint32(result, crc^0xFFFFFFFF)
		return result
	}
	length--

	crc ^= uint32(data[inx])
	inx++

	crc = ((uint32(-int32(crc&1)) & p1) ^
		(uint32(-int32((crc>>1)&1)) & p2) ^
		(uint32(-int32((crc>>2)&1)) & p3) ^
		(uint32(-int32((crc>>3)&1)) & p4) ^
		(uint32(-int32((crc>>4)&1)) & p5) ^
		(uint32(-int32((crc>>5)&1)) & p6) ^
		(uint32(-int32((crc>>6)&1)) & p7) ^
		(uint32(-int32((crc>>7)&1)) & polynomial) ^
		(crc >> 8))

	goto loop
}
