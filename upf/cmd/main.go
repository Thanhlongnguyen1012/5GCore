package main

import (
	"fmt"
	"net"
)

const (
	// PFCP ports & types
	PFCPPort                 = 8805
	SessionEstablishmentResp = 0x33
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8805")
	if err != nil {
		fmt.Println("ResolveUDPAddr:", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("ListenUDP:", err)
	}
	defer conn.Close()
	fmt.Println("PFCP server listening on", addr)

	buf := make([]byte, 1500)
	for {
		n, cli, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP:", err)
			continue
		}
		fmt.Printf("Got %d bytes from %s\n", n, cli)

		// Build Response
		resp := buildPFCPResponse(buf[:n])
		_, err = conn.WriteToUDP(resp, cli)
		if err != nil {
			fmt.Println("WriteToUDP:", err)
		} else {
			fmt.Printf("Sent PFCP Session Establishment Response (%d bytes)\n", len(resp))
		}
	}
}

// buildPFCPResponse xây dựng đúng header + SEID + seq + priority + Cause IE
func buildPFCPResponse(req []byte) []byte {
	// Byte1: version=1 (0x20) + S=1 (0x10) = 0x30
	version := byte(0x20 | 0x10)
	msgType := byte(SessionEstablishmentResp)

	// Lấy lại SEID và Sequence từ request nếu có, hoặc giả SEID=0
	// (ở request S=0, không có SEID → server tự cấp SEID)
	seid := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	// Sequence số 5–7 trong req header
	seq := []byte{req[5], req[6], req[7]}
	priority := byte(0x00) // spare + MP=0

	// IEs: Cause (type=19,len=1,value=1), NodeID (type=60,len=5,IPv4=127.0.0.1)
	causeIE := []byte{0x13, 0x00, 0x01, 0x01}
	nodeID := []byte{0x3C, 0x00, 0x05, 0x00, 127, 0, 0, 1}
	payload := append(causeIE, nodeID...)

	// Length = len(SEID)+len(seq)+1(priority)+len(payload)
	length := len(seid) + len(seq) + 1 + len(payload)
	lengthHi := byte((length >> 8) & 0xFF)
	lengthLo := byte(length & 0xFF)

	// Build header
	hdr := []byte{version, msgType, lengthHi, lengthLo}
	hdr = append(hdr, seid...)
	hdr = append(hdr, seq...)
	hdr = append(hdr, priority)

	return append(hdr, payload...)
}
