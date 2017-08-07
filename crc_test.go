package main

import (
	"encoding/hex"
	"testing"
)

/*
func BenchmarkDammSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DammDigit("123")
	}
}

func BenchmarkDammLarge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DammDigit("00123014764700968325")
	}
}

func BenchmarkDammSmallParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DammDigit("123")
		}
	})
}

func BenchmarkDammLargeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			DammDigit("00123014764700968325")
		}
	})
}*/
/*
func TestCrc32String(t *testing.T) {
	input := []byte("Lorem ipsum")
	crc := Crc32(input)

	if crc != 0xA8 {
		t.Errorf("Invalid CRC32. expected 0xa8 got 0x" + hex.EncodeToString([]byte{byte(crc)}))
	}
}

func TestCrc32LiteralBytes(t *testing.T) {
	input := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39}
	crc := Crc32(input)

	if crc != 0xF4 {
		t.Errorf("Invalid CRC32. expected 0xf4 got 0x" + hex.EncodeToString([]byte{byte(crc)}))
	}
}*/

func TestCrc32EmptyVector(t *testing.T) {
	input := []byte{}
	crc := Crc32(input)
	i := "00000000"
	j := hex.EncodeToString(crc)

	if i != j {
		t.Errorf("Invalid CRC32. expected 0x" + i + " got 0x" + j)
	}
}

func TestCrc32Vector1(t *testing.T) {
	input := []byte("a")
	crc := Crc32(input)
	i := "3043d0c1"
	j := hex.EncodeToString(crc)

	if i != j {
		t.Errorf("Invalid CRC32. expected 0x" + i + " got 0x" + j)
	}
}
