package conversion

import "encoding/binary"

// Uint64ToBytes converts a uint64 to a byte slice.
func Uint64ToBytes(n uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	return buf
}

// ByteToUint64 converts a byte slice to a uint64.
func ByteToUint64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}
