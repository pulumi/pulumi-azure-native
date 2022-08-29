


package v20210601

type AddonType string

const (
	AddonTypeSRM = AddonType("SRM")
	AddonTypeVR  = AddonType("VR")
	AddonTypeHCX = AddonType("HCX")
)

type DhcpTypeEnum string

const (
	DhcpTypeEnumSERVER = DhcpTypeEnum("SERVER")
	DhcpTypeEnumRELAY  = DhcpTypeEnum("RELAY")
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

type MountOptionEnum string

const (
	MountOptionEnumMOUNT  = MountOptionEnum("MOUNT")
	MountOptionEnumATTACH = MountOptionEnum("ATTACH")
)

type PortMirroringDirectionEnum string

const (
	PortMirroringDirectionEnum_INGRESS_EGRESS_BIDIRECTIONAL = PortMirroringDirectionEnum("INGRESS, EGRESS, BIDIRECTIONAL")
)

type ScriptExecutionParameterType string

const (
	ScriptExecutionParameterTypeValue       = ScriptExecutionParameterType("Value")
	ScriptExecutionParameterTypeSecureValue = ScriptExecutionParameterType("SecureValue")
	ScriptExecutionParameterTypeCredential  = ScriptExecutionParameterType("Credential")
)

type SslEnum string

const (
	SslEnumEnabled  = SslEnum("Enabled")
	SslEnumDisabled = SslEnum("Disabled")
)

func init() {
}
