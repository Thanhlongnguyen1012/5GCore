package client

import (
	"fmt"
	"net"
	"os"
)

const (
	SessionEstablishmentReq = 0x32
)

func SendPFCPEstablismentrequest() {
	remote, err := net.ResolveUDPAddr("udp", "127.0.0.1:8805")
	if err != nil {
		fmt.Println("ResolveUDPAddr:", err)
	}
	conn, err := net.DialUDP("udp", nil, remote)
	if err != nil {
		fmt.Println("DialUDP:", err)
	}
	defer conn.Close()

	req := buildPFCPRequest()
	_, err = conn.Write(req)
	if err != nil {
		fmt.Println("Write:", err)
	}
	fmt.Printf("Sent PFCP Session Establishment Request (%d bytes)\n", len(req))

	buf := make([]byte, 65535)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ReadFromUDP:", err)
		os.Exit(1)
	}
	fmt.Printf("Received PFCP Response (%d bytes)\n", n)
	fmt.Printf("% x\n", buf[:n])
}

func buildPFCPRequest() []byte {
	// --- Header ---
	version := byte(0x20)
	msgType := byte(SessionEstablishmentReq)

	seq := []byte{0x00, 0x00, 0x01}
	priority := byte(0x00)

	// --- IEs ---
	// Node ID IE (Type=60, Len=5)
	nodeID := []byte{
		0x3C, 0x00, 0x05, // Type=60, Length=5
		0x00,         // IPv4 type
		127, 0, 0, 1, // 127.0.0.1
	}

	// F-SEID IE (Type=87, Len=13)
	fseid := []byte{
		0x57, 0x00, 0x0D, // Type=87, Length=13
		0x01,                                           // Flags: V4=1
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, // SEID (8 bytes)
		127, 0, 0, 1, // IPv4 address
	}

	// Create PDR IE (Example, adjust as needed)
	createPDR := []byte{
		0x01, 0x00, 0x0C, // Type=1 (Create PDR), Length=12
		0x00, 0x01, 0x00, 0x04, 0xff, 0xff, 0xff, 0xff, // PDR ID and other fields (dummy data)
		0x00, 0x02, 0x00, 0x01, 0x01, // Additional dummy fields
	}

	payload := append(nodeID, fseid...)
	payload = append(payload, createPDR...)

	// Calculate total length (seq(3) + priority(1) + payload)
	totalLen := len(seq) + 1 + len(payload)
	lengthHi := byte((totalLen >> 8) & 0xFF)
	lengthLo := byte(totalLen & 0xFF)

	header := []byte{
		version,
		msgType,
		lengthHi,
		lengthLo,
	}
	header = append(header, seq...)
	header = append(header, priority)

	return append(header, payload...)
}
