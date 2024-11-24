package dns_message

import "bytes"

type Authority struct {
}

func (authority Authority) serialize() []byte {
	var response *bytes.Buffer = bytes.NewBuffer(make([]byte, 0))
	return response.Bytes()
}
