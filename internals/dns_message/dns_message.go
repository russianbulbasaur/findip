package dns_message

import "bytes"

type DNSMessage struct {
	header    Header
	question  Question
	answer    Answer
	authority Authority
	//additional space
}

func NewDNSMessage() DNSMessage {
	header := NewHeader(1234, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	answer := Answer{}
	question := Question{}
	authority := Authority{}
	return DNSMessage{
		header,
		question,
		answer,
		authority,
	}
}

func (message DNSMessage) Serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	response.Write(message.header.serialize())
	response.Write(message.question.serialize())
	response.Write(message.answer.serialize())
	response.Write(message.authority.serialize())
	return response.Bytes()
}
