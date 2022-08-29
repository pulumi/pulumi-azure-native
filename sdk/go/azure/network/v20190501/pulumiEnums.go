


package v20190501

type BackendEnabledState string

const (
	BackendEnabledStateEnabled  = BackendEnabledState("Enabled")
	BackendEnabledStateDisabled = BackendEnabledState("Disabled")
)

type DynamicCompressionEnabled string

const (
	DynamicCompressionEnabledEnabled  = DynamicCompressionEnabled("Enabled")
	DynamicCompressionEnabledDisabled = DynamicCompressionEnabled("Disabled")
)

type EnforceCertificateNameCheckEnabledState string

const (
	EnforceCertificateNameCheckEnabledStateEnabled  = EnforceCertificateNameCheckEnabledState("Enabled")
	EnforceCertificateNameCheckEnabledStateDisabled = EnforceCertificateNameCheckEnabledState("Disabled")
)

type FrontDoorEnabledState string

const (
	FrontDoorEnabledStateEnabled  = FrontDoorEnabledState("Enabled")
	FrontDoorEnabledStateDisabled = FrontDoorEnabledState("Disabled")
)

type FrontDoorForwardingProtocol string

const (
	FrontDoorForwardingProtocolHttpOnly     = FrontDoorForwardingProtocol("HttpOnly")
	FrontDoorForwardingProtocolHttpsOnly    = FrontDoorForwardingProtocol("HttpsOnly")
	FrontDoorForwardingProtocolMatchRequest = FrontDoorForwardingProtocol("MatchRequest")
)

type FrontDoorHealthProbeMethod string

const (
	FrontDoorHealthProbeMethodGET  = FrontDoorHealthProbeMethod("GET")
	FrontDoorHealthProbeMethodHEAD = FrontDoorHealthProbeMethod("HEAD")
)

type FrontDoorProtocol string

const (
	FrontDoorProtocolHttp  = FrontDoorProtocol("Http")
	FrontDoorProtocolHttps = FrontDoorProtocol("Https")
)

type FrontDoorQuery string

const (
	FrontDoorQueryStripNone = FrontDoorQuery("StripNone")
	FrontDoorQueryStripAll  = FrontDoorQuery("StripAll")
)

type FrontDoorRedirectProtocol string

const (
	FrontDoorRedirectProtocolHttpOnly     = FrontDoorRedirectProtocol("HttpOnly")
	FrontDoorRedirectProtocolHttpsOnly    = FrontDoorRedirectProtocol("HttpsOnly")
	FrontDoorRedirectProtocolMatchRequest = FrontDoorRedirectProtocol("MatchRequest")
)

type FrontDoorRedirectType string

const (
	FrontDoorRedirectTypeMoved             = FrontDoorRedirectType("Moved")
	FrontDoorRedirectTypeFound             = FrontDoorRedirectType("Found")
	FrontDoorRedirectTypeTemporaryRedirect = FrontDoorRedirectType("TemporaryRedirect")
	FrontDoorRedirectTypePermanentRedirect = FrontDoorRedirectType("PermanentRedirect")
)

type HealthProbeEnabled string

const (
	HealthProbeEnabledEnabled  = HealthProbeEnabled("Enabled")
	HealthProbeEnabledDisabled = HealthProbeEnabled("Disabled")
)

type RoutingRuleEnabledState string

const (
	RoutingRuleEnabledStateEnabled  = RoutingRuleEnabledState("Enabled")
	RoutingRuleEnabledStateDisabled = RoutingRuleEnabledState("Disabled")
)

type SessionAffinityEnabledState string

const (
	SessionAffinityEnabledStateEnabled  = SessionAffinityEnabledState("Enabled")
	SessionAffinityEnabledStateDisabled = SessionAffinityEnabledState("Disabled")
)

func init() {
}
