


package v20200101preview

type ProbeProtocol string

const (
	ProbeProtocolTcp   = ProbeProtocol("tcp")
	ProbeProtocolHttp  = ProbeProtocol("http")
	ProbeProtocolHttps = ProbeProtocol("https")
)

type Protocol string

const (
	ProtocolTcp = Protocol("tcp")
	ProtocolUdp = Protocol("udp")
)

func init() {
}
