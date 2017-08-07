package main

import "C"
import "unsafe"

// FIXME return []byte
func Crc32(data []byte) []byte {

	var (
		i       int = 0
		current byte
		crc     uint = 0 ^ 0xFFFFFFFF
		length  int  = len(data)
	)

outer:
	if length == 0 {
		goto eos
	}
	current = data[i]
	crc ^= uint(current)
	i++

	for j := 0; j < 8; j++ {
		crc = (crc >> 1) ^ (crc&1)*0x82F63B78
	}

	length--
	goto outer
eos:

	r := crc ^ 0xFFFFFFFF
	return ((*[4]byte)(unsafe.Pointer(&r)))[:]
}
