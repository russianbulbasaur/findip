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
	header := NewHeader(
		1234,
		1,
		2,
		0, 0,
		0, 0,
		0, 0,
		0,
		1,
		0,
		0)
	answerBuilder := AnswerBuilder()
	answer := answerBuilder.addRR(NewRR("codecrafters.io",
		1,
		1, 60,
		4,
		"8.8.8.8"))
	question := NewQuestion("google.com",
		1,
		1)
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
