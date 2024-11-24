package dns_message

import (
	"bytes"
	"encoding/binary"
)

type Answer struct {
	records []RR
}

func AnswerBuilder() Answer {
	return Answer{
		make([]RR, 0),
	}
}

func (ans Answer) addRR(record RR) Answer {
	ans.records = append(ans.records, record)
	return ans
}

func (ans Answer) serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	for _, answer := range ans.records {
		response.Write(answer.name)
		buffer16 := make([]byte, 2)
		buffer32 := make([]byte, 4)
		binary.BigEndian.PutUint16(buffer16, answer.recordType)
		response.Write(buffer16)
		binary.BigEndian.PutUint16(buffer16, answer.class)
		response.Write(buffer16)
		binary.BigEndian.PutUint32(buffer32, answer.ttl)
		response.Write(buffer32)
		binary.BigEndian.PutUint16(buffer16, answer.rdLength)
		response.Write(buffer16)
		response.Write(answer.rData)
	}
	return response.Bytes()
}
