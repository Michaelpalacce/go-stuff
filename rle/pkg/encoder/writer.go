package encoder

import (
	"io"
)

const MAX_BYTE_COUNT = 255

// RleEncoder is a struct that encodes a byte stream based on Run Length.
// aka AAAABAAABBBBBEEECCDD becomes 4A3B3E2C2D
// Managing the Writer and the Reader is up to the caller.\
// NOTE: Make sure to close both Reader and Writer!
type RleEncoder struct {
	Writer io.Writer
}

// Write converst the byte stream from the reader into a RLE encoded byte stream.
func (e *RleEncoder) Write(reader io.Reader) error {
	buffer := make([]byte, 255)
	var lastByte byte
	lastByte = 0

	var lastByteCount uint8

	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		for i := 0; i < bytesRead; i++ {
			currentByte := buffer[i]
			if currentByte == lastByte && lastByteCount < MAX_BYTE_COUNT {
				lastByteCount++
				continue
			}

			if lastByteCount > 0 {
				err := e.flush(lastByte, lastByteCount)
				if err != nil {
					return err
				}
			}

			lastByte = currentByte
			lastByteCount = 1
		}
	}
	// Write the final byte after finishing the read loop
	if lastByteCount > 0 {
		err := e.flush(lastByte, lastByteCount)
		if err != nil {
			return err
		}
	}

	return nil
}

// flush writes the last byte and its count to the writer.
func (e *RleEncoder) flush(lastByte byte, lastByteCount uint8) error {
	_, err := e.Writer.Write([]byte{lastByte, lastByteCount})
	if err != nil {
		return err
	}

	return nil
}
