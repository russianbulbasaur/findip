package dns_message

import (
	"bytes"
)

type DNSResponse struct {
	header   Header
	question Question
	answer   Answer
	//authority Authority
	//additional space
}

type DNSRequest struct {
	header   Header
	question Question
}

func NewDNSResponse() DNSResponse {
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
	//authority := Authority{}
	return DNSResponse{
		header,
		question,
		answer,
		//authority,
	}
}

func ParseDNSRequest(message []byte) DNSRequest {
	if len(message) < 12 {
		panic("Invalid dns request")
	}
	header := ParseHeader(message[0:12])
	question := ParseQuestion(message[12:])
	return DNSRequest{
		header,
		question,
	}
}

func (message DNSRequest) GetResponse() DNSResponse {
	header := NewHeader(message.header.id,
		1,
		message.header.opCode,
		0, 0,
		message.header.rd, 1,
		0, 0,
		1,
		1,
		0,
		0)
	question := NewQuestion("codecrafters.io", 1, 1)
	answerBuilder := AnswerBuilder()
	answer := answerBuilder.addRR(NewRR("codecrafters.io",
		1,
		1, 60,
		4,
		"8.8.8.8"))
	return DNSResponse{
		header,
		question,
		answer,
	}
}

func (message DNSResponse) Serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	response.Write(message.header.serialize())
	response.Write(message.question.serialize())
	response.Write(message.answer.serialize())
	//response.Write(message.authority.serialize())
	return response.Bytes()
}
