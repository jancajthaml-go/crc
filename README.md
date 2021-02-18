## zero-alloc 32Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc32)](https://goreportcard.com/report/jancajthaml-go/crc32)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation provides both tableless and tabular checksum functions with variable 32bit polynomial.

### Supported standards ###

- CRC-32/BZIP2

### Usage ###

```
import "github.com/jancajthaml-go/crc32"

data := []byte("abcdefgh")
poly := 0x04C11DB7
init := 0xFFFFFFFF
xorout := 0xFFFFFFFF

// for tableless
crc32.Checksum(data, poly, init, xorout) // 0x5024EC61

// for precalculated tabular
instance = crc32.New(poly, init, xorout)
instance.Checksum(data) // 0x5024EC61
```
