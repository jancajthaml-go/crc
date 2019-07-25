package crc32

import (
	"hash/crc32"
	"strings"
	"testing"
)

var largeText = []byte(strings.Repeat("a", 50000))
var smallText = []byte(strings.Repeat("a", 5))

const IEEE = 0xedb88320

func standard(data []byte) uint32 {
	return crc32.ChecksumIEEE(data) ^ 0xFFFFFFFF
}

func AssetEqual(t *testing.T, expected uint32, actual uint32) {
	if expected != actual {
		t.Errorf("Expected 0x%08X got 0x%08X", expected, actual)
	}
}

func TestCrc32EmptyVector(t *testing.T) {
	input := []byte{}
	libResult := Checksum(input, IEEE, 0xFFFFFFFF, 0x00000000)
	standardResult := standard(input)

	AssetEqual(t, standardResult, libResult)
}

func TestPrecalculatedNormalized(t *testing.T) {
	input := []byte("abcdefgh")

	t.Log("CRC-32/BZIP2")
	{
		c := New(0x04C11DB7, 0xFFFFFFFF, 0xFFFFFFFF)
		AssetEqual(t, 0x5024EC61, c.Checksum(input))
	}

}

func TestNormalized(t *testing.T) {

	input := []byte("abcdefgh")

	t.Log("CRC-32/BZIP2")
	{
		AssetEqual(t, 0x5024EC61, Checksum(input, 0x04C11DB7, 0xFFFFFFFF, 0xFFFFFFFF))
	}

}

func BenchmarkStandardSmall(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		standard(smallText)
	}
}

func BenchmarkStandardLarge(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		standard(largeText)
	}
}

func BenchmarkCrcSmall(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		Checksum(smallText, IEEE, 0xFFFFFFFF, 0x00000000)
	}
}

func BenchmarkCrcLarge(b *testing.B) {
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		Checksum(largeText, IEEE, 0xFFFFFFFF, 0x00000000)
	}
}

func BenchmarkPrecalculatedCrcSmall(b *testing.B) {
	c := New(IEEE, 0xFFFFFFFF, 0x00000000)
	b.ResetTimer()
	b.SetBytes(int64(len(smallText)))
	for n := 0; n < b.N; n++ {
		c.Checksum(smallText)
	}
}

func BenchmarkPrecalculatedCrcLarge(b *testing.B) {
	c := New(IEEE, 0xFFFFFFFF, 0x00000000)
	b.ResetTimer()
	b.SetBytes(int64(len(largeText)))
	for n := 0; n < b.N; n++ {
		c.Checksum(largeText)
	}
}
