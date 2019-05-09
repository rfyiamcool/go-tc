//+build linux

package tc

import (
	"bytes"
	"encoding/binary"
)

// Msg represents a Traffic Control Message
type Msg struct {
	Family  uint32
	Ifindex uint32
	Handle  uint32
	Parent  uint32
	Info    uint32
}

func tcmsgEncode(i *Msg) ([]byte, error) {
	var tcm bytes.Buffer
	err := binary.Write(&tcm, nativeEndian, *i)
	return tcm.Bytes(), err
}

func tcmsgDecode(data []byte, tc *Msg) error {
	b := bytes.NewReader(data)
	return binary.Read(b, nativeEndian, tc)
}
