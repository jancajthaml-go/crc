package crc32

// CRC holds crc32 parameters and precalculated table
type CRC struct {
	table  []uint32
	poly   uint32
	xorout uint32
	init   uint32
}

// New returns CRC32 instance with precalculated table
func New(poly uint32, init uint32, xorout uint32) CRC {
	result := CRC{
		poly:   poly,
		xorout: xorout,
		init:   init,
		table:  createTable(poly, init, xorout),
	}

	return result
}

func createTable(poly uint32, init uint32, xorout uint32) []uint32 {
	result := make([]uint32, 256)
	var bit uint32
	for divident := 0; divident < 256; divident++ {
		var current uint32 = 0x00000000
		for j := uint16(0x0080); j != 0; j >>= 1 {
			if (uint16(divident) & j) != 0 {
				bit = (current & 0x80000000) ^ 0x80000000
			} else {
				bit = current & 0x80000000
			}
			switch bit {
			case 0:
				current <<= 1
			default:
				current = (current << 1) ^ poly
			}
		}
		result[divident] = current & 0xFFFFFFFF
	}
	return result
}

// Checksum returns CRC32 checksum of given CRC instance
func (crc *CRC) Checksum(data []byte) uint32 {
	var pos uint8
	var result = crc.init
	for _, item := range data {
		result = result ^ uint32(item)<<24
		pos = uint8((result >> 24) & 0xFF)
		result = (result << 8) & 0xFFFFFFFF
		result = (result ^ crc.table[pos]) & 0xFFFFFFFF

	}
	return result ^ crc.xorout
}

// Checksum returns CRC32 checksum for given parameters
func Checksum(data []byte, poly uint32, init uint32, xorout uint32) uint32 {
	var crc uint32 = init
	var bit uint32

	for _, item := range data {
		for j := uint16(0x0080); j != 0; j >>= 1 {
			if (uint16(item) & j) != 0 {
				bit = (crc & 0x80000000) ^ 0x80000000
			} else {
				bit = crc & 0x80000000
			}
			switch bit {
			case 0:
				crc <<= 1
			default:
				crc = (crc << 1) ^ poly
			}
		}
	}
	return crc ^ xorout
}
