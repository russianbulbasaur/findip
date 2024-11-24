package dns_message

import (
	"bytes"
	"encoding/binary"
	"findip/utils"
)

type Question struct {
	qName  []byte
	qType  uint16
	qClass uint16
}

func NewQuestion(domain string, qType uint16, qClass uint16) Question {
	return Question{
		utils.DomainToBytes(domain),
		qType,
		qClass,
	}
}

func (question Question) serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	response.Write(question.qName)
	buffer := make([]byte, 2)
	binary.BigEndian.PutUint16(buffer, question.qType)
	response.Write(buffer)
	binary.BigEndian.PutUint16(buffer, question.qClass)
	response.Write(buffer)
	return response.Bytes()
}
