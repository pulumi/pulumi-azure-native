


package v20211201

type AddonType string

const (
	AddonTypeSRM = AddonType("SRM")
	AddonTypeVR  = AddonType("VR")
	AddonTypeHCX = AddonType("HCX")
)

type AffinityType string

const (
	AffinityTypeAffinity     = AffinityType("Affinity")
	AffinityTypeAntiAffinity = AffinityType("AntiAffinity")
)

type AvailabilityStrategy string

const (
	AvailabilityStrategySingleZone = AvailabilityStrategy("SingleZone")
	AvailabilityStrategyDualZone   = AvailabilityStrategy("DualZone")
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

type EncryptionState string

const (
	EncryptionStateEnabled  = EncryptionState("Enabled")
	EncryptionStateDisabled = EncryptionState("Disabled")
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
	PortMirroringDirectionEnumINGRESS       = PortMirroringDirectionEnum("INGRESS")
	PortMirroringDirectionEnumEGRESS        = PortMirroringDirectionEnum("EGRESS")
	PortMirroringDirectionEnumBIDIRECTIONAL = PortMirroringDirectionEnum("BIDIRECTIONAL")
)

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
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
