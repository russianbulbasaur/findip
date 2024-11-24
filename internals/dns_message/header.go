package dns_message

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	id                     uint16 //16 bits
	qrOpcodeAaTcRdRaZRcode uint16
	qdCount                uint16
	anCount                uint16
	nsCount                uint16
	arCount                uint16
}

func NewHeader(id uint16, qr uint16, opcode uint16, aa uint16, tc uint16, rd uint16, ra uint16, z uint16, rCode uint16,
	qdCount uint16, anCount uint16, nsCount uint16, arCount uint16) Header {
	var qrOpcodeAaTcRdRaZRcode uint16 = 0
	qrOpcodeAaTcRdRaZRcode |= qr << 15
	qrOpcodeAaTcRdRaZRcode |= opcode << 11
	qrOpcodeAaTcRdRaZRcode |= aa << 10
	qrOpcodeAaTcRdRaZRcode |= tc << 9
	qrOpcodeAaTcRdRaZRcode |= rd << 8
	qrOpcodeAaTcRdRaZRcode |= ra << 7
	qrOpcodeAaTcRdRaZRcode |= z << 4
	qrOpcodeAaTcRdRaZRcode |= rCode
	return Header{
		id,
		qrOpcodeAaTcRdRaZRcode,
		qdCount,
		anCount,
		nsCount,
		arCount,
	}
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
