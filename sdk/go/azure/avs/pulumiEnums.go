


package avs

type AddonType string

const (
	AddonTypeSRM = AddonType("SRM")
	AddonTypeVR  = AddonType("VR")
)

type AffinityType string

const (
	AffinityTypeAffinity     = AffinityType("Affinity")
	AffinityTypeAntiAffinity = AffinityType("AntiAffinity")
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

type PlacementPolicyStateEnum string

const (
	PlacementPolicyStateEnumEnabled  = PlacementPolicyStateEnum("Enabled")
	PlacementPolicyStateEnumDisabled = PlacementPolicyStateEnum("Disabled")
)

type PlacementPolicyType string

const (
	PlacementPolicyTypeVmVm   = PlacementPolicyType("VmVm")
	PlacementPolicyTypeVmHost = PlacementPolicyType("VmHost")
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
