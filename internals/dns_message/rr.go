package dns_message

import (
	"findip/utils"
)

type RR struct {
	name       []byte
	recordType uint16
	class      uint16
	ttl        uint32
	rdLength   uint16
	rData      []byte
}

func NewRR(name string, recordType uint16, class uint16, ttl uint32, rdLength uint16, ipAddress string) RR {
	return RR{
		utils.DomainToBytes(name),
		recordType,
		class,
		ttl,
		rdLength,
		utils.IPTo4Bytes(ipAddress),
	}
}

func NewRRBytes(name []byte, recordType uint16, class uint16, ttl uint32, rdLength uint16, ipAddress string) RR {
	return RR{
		name,
		recordType,
		class,
		ttl,
		rdLength,
		utils.IPTo4Bytes(ipAddress),
	}
}
