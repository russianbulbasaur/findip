package utils

import (
	"bytes"
	"strings"
)

func DomainToBytes(domain string) []byte {
	splitted := strings.Split(domain, ".")
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	for _, part := range splitted {
		response.WriteByte(uint8(len(part)))
		response.Write([]byte(part))
	}
	response.WriteByte(0x00)
	return response.Bytes()
}
