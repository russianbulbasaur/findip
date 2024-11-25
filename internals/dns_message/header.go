package dns_message

import (
	"bytes"
	"encoding/binary"
	"log"
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

func ParseHeader(message []byte) Header {
	var id uint16
	err := binary.Read(bytes.NewBuffer(message[0:2]), binary.BigEndian, id)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	var qrOpcodeAaTcRdRaZRcode uint16
	err = binary.Read(bytes.NewBuffer(message[2:4]), binary.LittleEndian, qrOpcodeAaTcRdRaZRcode)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	var qdCount uint16
	err = binary.Read(bytes.NewBuffer(message[4:6]), binary.BigEndian, qdCount)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	var anCount uint16
	err = binary.Read(bytes.NewBuffer(message[6:8]), binary.BigEndian, anCount)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	var nsCount uint16
	err = binary.Read(bytes.NewBuffer(message[8:10]), binary.BigEndian, nsCount)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	var arCount uint16
	err = binary.Read(bytes.NewBuffer(message[10:12]), binary.BigEndian, arCount)
	if err != nil {
		log.Println(err)
		panic(err)
	}
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
