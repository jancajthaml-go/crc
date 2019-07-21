package main

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"strings"
	"testing"
	"unsafe"
)

var largeText = []byte(strings.Repeat("a", 50000))
var smallText = []byte(strings.Repeat("a", 5))

const IEEE = 0xedb88320

func naive(data []byte, polynomial uint32) []byte {
	var crc uint32 = 0 ^ 0xFFFFFFFF

	for _, current := range data {
		crc ^= uint32(current)
		for j := 0; j < 8; j++ {
			crc = (crc >> 1) ^ (crc&1)*polynomial
		}
	}

	return ((*[4]byte)(unsafe.Pointer(&crc)))[:]
}

func standard(data []byte) []byte {
	result := make([]byte, 4)
	binary.LittleEndian.PutUint32(result, crc32.ChecksumIEEE(data)^0xFFFFFFFF)
	return result
}

func BenchmarkNaiveSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		naive(smallText, IEEE)
	}
}

func BenchmarkNaiveLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		naive(largeText, IEEE)
	}
}

func BenchmarkStandardSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		standard(smallText)
	}
}

func BenchmarkStandardLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		standard(largeText)
	}
}

func BenchmarkCrcSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CRC32(smallText, IEEE)
	}
}

func BenchmarkCrcLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CRC32(largeText, IEEE)
	}
}

func TestCrc32EmptyVector(t *testing.T) {
	input := []byte{}
	libResult := CRC32(input, IEEE)
	naiveResult := naive(input, IEEE)
	standardResult := standard(input)

	a := hex.EncodeToString(libResult)
	b := hex.EncodeToString(naiveResult)
	c := hex.EncodeToString(standardResult)

	if a != b || b != c {
		t.Errorf("Invalid CRC32. standard 0x" + c + " naive 0x" + b + " library 0x" + a)
	}
}

func TestCrc32Vector1(t *testing.T) {
	input := []byte("a")
	libResult := CRC32(input, IEEE)
	naiveResult := naive(input, IEEE)
	standardResult := standard(input)

	a := hex.EncodeToString(libResult)
	b := hex.EncodeToString(naiveResult)
	c := hex.EncodeToString(standardResult)

	if a != b || b != c {
		t.Errorf("Invalid CRC32. standard 0x" + c + " naive 0x" + b + " library 0x" + a)
	}
}

func TestCrc32Vector2(t *testing.T) {
	input := []byte(largeText)
	libResult := CRC32(input, IEEE)
	naiveResult := naive(input, IEEE)
	standardResult := standard(input)

	a := hex.EncodeToString(libResult)
	b := hex.EncodeToString(naiveResult)
	c := hex.EncodeToString(standardResult)

	if a != b || b != c {
		t.Errorf("Invalid CRC32. standard 0x" + c + " naive 0x" + b + " library 0x" + a)
	}
}
