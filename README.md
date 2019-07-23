## 32Bit Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc32)](https://goreportcard.com/report/jancajthaml-go/crc32)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation is tableless with variable 32bit polynomial.

### Performance ###

```
BenchmarkCrcSmall  60.54 MB/s  0 B/op  0 allocs/op
BenchmarkCrcLarge  26.07 MB/s  0 B/op  0 allocs/op
```

### Usage ###

```
import "github.com/jancajthaml-go/crc32"

data := []byte("abcdefgh")
poly := 0x04C11DB7
init := 0xFFFFFFFF
xorout := 0xFFFFFFFF

crc32.Checksum(data, poly, init, xorout) // 0x5024EC61
```


