package dns_message

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	id                     uint16 //16 bits
	qr                     uint16
	opCode                 uint16
	aa                     uint16
	tc                     uint16
	rd                     uint16
	ra                     uint16
	z                      uint16
	Rcode                  uint16
	qrOpcodeAaTcRdRaZRcode uint16
	qdCount                uint16
	anCount                uint16
	nsCount                uint16
	arCount                uint16
}

func NewHeader(id uint16, qr uint16, opCode uint16, aa uint16, tc uint16, rd uint16, ra uint16, z uint16, rCode uint16,
	qdCount uint16, anCount uint16, nsCount uint16, arCount uint16) Header {
	var qrOpcodeAaTcRdRaZRcode uint16 = 0
	qrOpcodeAaTcRdRaZRcode |= qr << 15
	qrOpcodeAaTcRdRaZRcode |= opCode << 11
	qrOpcodeAaTcRdRaZRcode |= aa << 10
	qrOpcodeAaTcRdRaZRcode |= tc << 9
	qrOpcodeAaTcRdRaZRcode |= rd << 8
	qrOpcodeAaTcRdRaZRcode |= ra << 7
	qrOpcodeAaTcRdRaZRcode |= z << 4
	qrOpcodeAaTcRdRaZRcode |= rCode
	return Header{
		id,
		qr, opCode, aa, tc, rd, ra, z, rCode,
		qrOpcodeAaTcRdRaZRcode,
		qdCount,
		anCount,
		nsCount,
		arCount,
	}
}

func ParseHeader(message []byte) Header {
	id := binary.BigEndian.Uint16(message[0:2])
	qrOpcodeAaTcRdRaZRcode := binary.BigEndian.Uint16(message[2:4])
	qdCount := binary.BigEndian.Uint16(message[4:6])
	anCount := binary.BigEndian.Uint16(message[6:8])
	nsCount := binary.BigEndian.Uint16(message[8:10])
	arCount := binary.BigEndian.Uint16(message[10:12])
	qr := qrOpcodeAaTcRdRaZRcode >> 15
	opCode := (qrOpcodeAaTcRdRaZRcode << 1) >> 12
	aa := (qrOpcodeAaTcRdRaZRcode << 5) >> 15
	tc := (qrOpcodeAaTcRdRaZRcode << 6) >> 15
	rd := (qrOpcodeAaTcRdRaZRcode << 7) >> 15
	ra := (qrOpcodeAaTcRdRaZRcode << 8) >> 15
	z := (qrOpcodeAaTcRdRaZRcode << 9) >> 13
	rCode := (qrOpcodeAaTcRdRaZRcode << 12) >> 12
	return NewHeader(id, qr, opCode, aa, tc, rd, ra, z, rCode, qdCount, anCount, nsCount, arCount)
}

func (header Header) serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, header.id)
	response.Write(buffer)
	binary.BigEndian.PutUint16(buffer, header.qrOpcodeAaTcRdRaZRcode)
	response.Write(buffer)
	binary.BigEndian.PutUint16(buffer, header.qdCount)
	response.Write(buffer)
	binary.BigEndian.PutUint16(buffer, header.anCount)
	response.Write(buffer)
	binary.BigEndian.PutUint16(buffer, header.nsCount)
	response.Write(buffer)
	binary.BigEndian.PutUint16(buffer, header.arCount)
	response.Write(buffer)
	return response.Bytes()
}
