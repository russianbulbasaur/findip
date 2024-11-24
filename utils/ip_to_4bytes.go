package utils

import (
	"log"
	"strconv"
	"strings"
)

func IPTo4Bytes(ip string) []byte {
	buffer := make([]byte, 0)
	split := strings.Split(ip, ".")
	for _, part := range split {
		ipPart, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalln(err, "Invalid ip address")
		}
		buffer = append(buffer, uint8(ipPart))
	}
	return buffer
}
