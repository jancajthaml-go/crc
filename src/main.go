package main

import (
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

	// FIXME read from input file
	// FIXME read polynomial from command line args

	crc := CRC32([]byte(os.Args[1]), 0x04C11DB7)
	os.Stdout.Write(crc)

	os.Exit(0)
}
