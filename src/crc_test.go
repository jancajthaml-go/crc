package main

import (
	"encoding/hex"
	"testing"
	"unsafe"
)

func naive(data []byte, polynomial uint) []byte {
	var crc uint = 0 ^ 0xFFFFFFFF

	for _, current := range data {
		crc ^= uint(current)
		for j := 0; j < 8; j++ {
			crc = (crc >> 1) ^ (crc&1)*polynomial
		}
	}

	r := crc ^ 0xFFFFFFFF
	return ((*[4]byte)(unsafe.Pointer(&r)))[:]
}

func BenchmarkNaiveSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		naive([]byte("aaaaaaa"), 0xEDB88320)
	}
}

func BenchmarkNaiveLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		naive([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), 0xEDB88320)
	}
}

func BenchmarkNaiveSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			naive([]byte("aaaaaaa"), 0xEDB88320)
		}
	})
}

func BenchmarkNaiveLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			naive([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), 0xEDB88320)
		}
	})
}

func BenchmarkCrcSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CRC32([]byte("aaaaaaa"), 0xEDB88320)
	}
}

func BenchmarkCrcLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CRC32([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), 0xEDB88320)
	}
}

func BenchmarkCrcSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CRC32([]byte("aaaaaaa"), 0xEDB88320)
		}
	})
}

func BenchmarkCrcLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			CRC32([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), 0xEDB88320)
		}
	})
}

func TestCrc32EmptyVector(t *testing.T) {
	input := []byte{}
	crc := CRC32(input, 0xEDB88320)
	ref := naive(input, 0xEDB88320)
	i := hex.EncodeToString(ref)
	j := hex.EncodeToString(crc)

	if i != j {
		t.Errorf("Invalid CRC32. expected 0x" + i + " got 0x" + j)
	}
}

func TestCrc32Vector1(t *testing.T) {
	input := []byte("a")
	crc := CRC32(input, 0x04C11DB7)
	ref := naive(input, 0x04C11DB7)
	i := hex.EncodeToString(ref)
	j := hex.EncodeToString(crc)

	if i != j {
		t.Errorf("Invalid CRC32. expected naive : 0x" + i + " got Crc32 : 0x" + j)
	}
}

func TestCrc32Vector2(t *testing.T) {
	input := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	crc := CRC32(input, 0xEDB88320)
	ref := naive(input, 0xEDB88320)
	i := hex.EncodeToString(ref)
	j := hex.EncodeToString(crc)

	if i != j {
		t.Errorf("Invalid CRC32. expected naive : 0x" + i + " got Crc32 : 0x" + j)
	}
}
