package proto

import (
	"bytes"
	"encoding/binary"
	"io"
)

type FrameReadWriter interface {
	HasFrame() bool
	ReadFrame() ([]byte, error)
	WriteFrame([]byte) error
}

type lenEncodedFraming struct {
	rwc     io.ReadWriter
	readBuf []byte
	inBuf   bytes.Buffer
}

func (lef *lenEncodedFraming) HasFrame() bool {
	if lef.inBuf.Len() < 2 {
		return false
	}
	l := int(binary.BigEndian.Uint16(lef.inBuf.Bytes()[:2]) + 2)
	return lef.inBuf.Len() >= l
}

func (lef *lenEncodedFraming) ReadFrame() ([]byte, error) {
	// Read data until there's at least one full frame
	for !lef.HasFrame() {
		n, err := lef.rwc.Read(lef.readBuf)
		if err != nil {
			return nil, err
		}
		if _, err := lef.inBuf.Write(lef.readBuf[:n]); err != nil {
			return nil, err
		}
	}
	// Return the first frame
	l := int(binary.BigEndian.Uint16(lef.inBuf.Bytes()[:2]) + 2)
	frame := make([]byte, l)
	if _, err := io.ReadFull(&lef.inBuf, frame); err != nil {
		return nil, err
	}
	return frame[2:], nil
}

func (lef *lenEncodedFraming) WriteFrame(data []byte) error {
	var lenBytes [2]byte
	binary.BigEndian.PutUint16(lenBytes[:], uint16(len(data)))
	if _, err := lef.rwc.Write(lenBytes[:]); err != nil {
		return err
	}
	if _, err := lef.rwc.Write(data); err != nil {
		return err
	}
	return nil
}

func (lef *lenEncodedFraming) Write(b []byte) (int, error) {
	return lef.inBuf.Write(b)
}

func NewLenEncodedFraming(rwc io.ReadWriter) *lenEncodedFraming {
	return &lenEncodedFraming{
		rwc:     rwc,
		readBuf: make([]byte, 64<<10),
	}
}
