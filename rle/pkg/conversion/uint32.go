package conversion

import "encoding/binary"

// Uint32ToBytes converts a uint32 to a byte slice.
func Uint32ToBytes(n uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, n)
	return buf
}

// ByteToUint32 converts a byte slice to a uint32.
func ByteToUint32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}
