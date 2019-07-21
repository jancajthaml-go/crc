## Cyclic redundancy check

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/crc)](https://goreportcard.com/report/jancajthaml-go/crc)

CRC which encode messages by adding a fixed-length check value, for the purpose of error detection in communication networks, it can provide quick and reasonable assurance of the integrity of messages delivered.

However, it is not suitable for protection against intentional alteration of data.

Implementation is tableless with variable 32bit polynomial and same speed as zlib implementation.

### Usage ###

```
import "github.com/jancajthaml-go/crc32"

crc.Crc32([]byte("abcd"), 0x04C11DB7)
```

### Performance ###
> WIP
