package pkg_test_utils

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

func exampleRiffData() []byte {
	data := make([]uint8, 256)
	for i := range data {
		data[i] = uint8(255 * (1.0 + math.Sin(2.0*math.Pi*float64(i)/64.0)))
	}

	riffHeader := []byte("RIFF")
	riffSize := uint32(256)
	waveHeader := []byte("WAVE")

	byteBuffer := bytes.NewBuffer([]byte{})
	byteBuffer.Write(riffHeader)
	binary.Write(byteBuffer, binary.LittleEndian, riffSize)
	byteBuffer.Write(waveHeader)

	// 'fmt ' sub-chunk (format)
	fmtHeader := []byte("fmt ")
	fmtSize := uint32(16)    // Size of fmt chunk
	audioFormat := uint16(1) // PCM
	numChannels := uint16(1) // Mono
	sampleRate := uint32(44100)
	byteRate := uint32(sampleRate * uint32(numChannels) * 2)
	blockAlign := uint16(numChannels * 2)
	bitsPerSample := uint16(16)

	// Write 'fmt ' sub-chunk
	byteBuffer.Write(fmtHeader)
	binary.Write(byteBuffer, binary.LittleEndian, fmtSize)
	binary.Write(byteBuffer, binary.LittleEndian, audioFormat)
	binary.Write(byteBuffer, binary.LittleEndian, numChannels)
	binary.Write(byteBuffer, binary.LittleEndian, sampleRate)
	binary.Write(byteBuffer, binary.LittleEndian, byteRate)
	binary.Write(byteBuffer, binary.LittleEndian, blockAlign)
	binary.Write(byteBuffer, binary.LittleEndian, bitsPerSample)

	// 'data' sub-chunk
	dataHeader := []byte("data")
	dataSize := uint32(256) // Size of data

	// Write 'data' sub-chunk
	byteBuffer.Write(dataHeader)
	binary.Write(byteBuffer, binary.LittleEndian, dataSize)
	byteBuffer.Write(data)
	return byteBuffer.Bytes()
}

type MockRiffReader struct {
	data         []byte
	numBytesRead int64
}

func (m MockRiffReader) Read(dst []byte) (int, error) {
	numCopied, err := m.ReadAt(dst, m.numBytesRead)
	m.numBytesRead += int64(numCopied)
	return numCopied, err
}

func (m MockRiffReader) ReadAt(dst []byte, offset int64) (int, error) {
	if offset >= int64(len(m.data)) {
		return 0, io.EOF
	}
	numCopied := copy(dst, m.data[offset:])
	return numCopied, nil
}

func NewMockRiffReader() MockRiffReader {
	return MockRiffReader{
		data: exampleRiffData(),
	}
}
