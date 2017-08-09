package main

import "encoding/binary"

func Crc32(data []byte, polynomial uint32) []byte {
	var (
		inx    int    = 0
		crc    uint32 = 0 ^ 0xFFFFFFFF
		length int    = len(data)

		p7 uint32 = polynomial >> 1
		p6 uint32 = polynomial >> 2
		p5 uint32 = polynomial >> 3
		p4 uint32 = polynomial >> 4
		p3 uint32 = polynomial >> 5
		p2 uint32 = (polynomial >> 6) ^ polynomial
		p1 uint32 = (polynomial >> 7) ^ p7
	)

loop:
	if length == 0 {
		a := make([]byte, 4)
		binary.LittleEndian.PutUint32(a, crc^0xFFFFFFFF)
		return a
	}
	length--

	crc ^= uint32(data[inx])
	inx++

	var a = (uint32(-int32(crc&1)) & p1)
	var b = (uint32(-int32((crc>>1)&1)) & p2)
	var c = (uint32(-int32((crc>>2)&1)) & p3)
	var d = (uint32(-int32((crc>>3)&1)) & p4)
	var e = (uint32(-int32((crc>>4)&1)) & p5)
	var f = (uint32(-int32((crc>>5)&1)) & p6)
	var g = (uint32(-int32((crc>>6)&1)) & p7)
	var h = (uint32(-int32((crc>>7)&1)) & polynomial)
	var i = (crc >> 8)

	crc = a ^ b ^ c ^ d ^ e ^ f ^ g ^ h ^ i

	goto loop
}
