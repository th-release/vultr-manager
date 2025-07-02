package firewall

type IpType string

const (
	IpV4 IpType = "v4"
	IpV6 IpType = "v6"
)

func (i IpType) String() string {
	return string(i)
}

func (i IpType) IsValid() bool {
	return i == IpV4 || i == IpV6
}

type Protocol string

const (
	ProtocolICMP Protocol = "ICMP"
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolGRE  Protocol = "GRE"
	ProtocolESP  Protocol = "ESP"
	ProtocolAH   Protocol = "AH"
)

func (p Protocol) String() string {
	return string(p)
}

func (p Protocol) IsValid() bool {
	return p == ProtocolTCP || p == ProtocolICMP || p == ProtocolUDP || p == ProtocolGRE || p == ProtocolESP || p == ProtocolAH
}
