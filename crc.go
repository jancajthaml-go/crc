package crc32

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
