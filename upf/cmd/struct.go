package main

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// PFCP Header (8 bytes)
type PfcpHeader struct {
	Version       uint8 // 3 bits
	Flags         uint8 // 5 bits
	MessageType   uint8
	MessageLength uint16
	SEID          uint64
}

// Marshal PFCP Header thành byte slice (Big-Endian)
func (h *PfcpHeader) Marshal() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Pack Version (3 bits) + Flags (5 bits) vào 1 byte
	versionAndFlags := (h.Version << 5) | (h.Flags & 0x1F)
	if err := binary.Write(buf, binary.BigEndian, versionAndFlags); err != nil {
		return nil, err
	}

	// Các trường còn lại
	if err := binary.Write(buf, binary.BigEndian, h.MessageType); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, h.MessageLength); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, h.SEID); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal byte slice thành PFCP Header
func (h *PfcpHeader) Unmarshal(data []byte) error {
	if len(data) < 12 { // 1 + 1 + 2 + 8 = 12 bytes
		return errors.New("invalid PFCP header length")
	}

	buf := bytes.NewReader(data)

	// Đọc Version và Flags
	var versionAndFlags uint8
	if err := binary.Read(buf, binary.BigEndian, &versionAndFlags); err != nil {
		return err
	}
	h.Version = versionAndFlags >> 5
	h.Flags = versionAndFlags & 0x1F

	// Các trường còn lại
	if err := binary.Read(buf, binary.BigEndian, &h.MessageType); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &h.MessageLength); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &h.SEID); err != nil {
		return err
	}

	return nil
}
