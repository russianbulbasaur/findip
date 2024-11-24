package dns_message

import (
	"bytes"
	"encoding/binary"
)

type Answer struct {
	name       []byte
	answerType uint16
	class      uint16
	ttl        uint32
	rdLength   uint16
	rData      []byte
}

func (answer Answer) serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	response.Write(answer.name)
	buffer16 := make([]byte, 2)
	buffer32 := make([]byte, 4)
	binary.BigEndian.PutUint16(buffer16, answer.answerType)
	response.Write(buffer16)
	binary.BigEndian.PutUint16(buffer16, answer.class)
	response.Write(buffer16)
	binary.BigEndian.PutUint32(buffer32, answer.ttl)
	response.Write(buffer32)
	binary.BigEndian.PutUint16(buffer16, answer.rdLength)
	response.Write(buffer16)
	response.Write(answer.rData)
	return response.Bytes()
}
