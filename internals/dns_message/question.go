package dns_message

import (
	"bytes"
	"encoding/binary"
)

type Question struct {
	qName  []byte
	qType  uint16
	qClass uint16
}

func NewQuestion(qName []byte, qType uint16, qClass uint16) Question {
	return Question{
		qName,
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
