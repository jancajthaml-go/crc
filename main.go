package main

import (
	"encoding/hex"
	"os"
)

func main() {
	defer func() {
		if recover() != nil {
			os.Exit(1)
		}
	}()

	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("Usage           : ./crc <input>\nValid Example   : ./crc 123; echo \"$?\""))
		return
	}

	crc := Crc32([]byte(os.Args[1]))
	os.Stdout.Write([]byte(hex.EncodeToString(crc)))

	os.Exit(0)
}
