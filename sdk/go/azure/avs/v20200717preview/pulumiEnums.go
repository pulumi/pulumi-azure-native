


package v20200717preview

type AddonType string

const (
	AddonTypeSRM = AddonType("SRM")
	AddonTypeVR  = AddonType("VR")
)

type DhcpTypeEnum string

const (
	DhcpTypeEnum_SERVER_RELAY = DhcpTypeEnum("SERVER, RELAY")
)

type DnsServiceLogLevelEnum string

const (
	DnsServiceLogLevelEnumDEBUG   = DnsServiceLogLevelEnum("DEBUG")
	DnsServiceLogLevelEnumINFO    = DnsServiceLogLevelEnum("INFO")
	DnsServiceLogLevelEnumWARNING = DnsServiceLogLevelEnum("WARNING")
	DnsServiceLogLevelEnumERROR   = DnsServiceLogLevelEnum("ERROR")
	DnsServiceLogLevelEnumFATAL   = DnsServiceLogLevelEnum("FATAL")
)

type InternetEnum string

const (
	InternetEnumEnabled  = InternetEnum("Enabled")
	InternetEnumDisabled = InternetEnum("Disabled")
)

type PortMirroringDirectionEnum string

const (
	PortMirroringDirectionEnum_INGRESS_EGRESS_BIDIRECTIONAL = PortMirroringDirectionEnum("INGRESS, EGRESS, BIDIRECTIONAL")
)

type SslEnum string

const (
	SslEnumEnabled  = SslEnum("Enabled")
	SslEnumDisabled = SslEnum("Disabled")
)

func init() {
}
