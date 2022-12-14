


package v20171115

type AccessRightsDescription string

const (
	AccessRightsDescriptionServiceConfig           = AccessRightsDescription("ServiceConfig")
	AccessRightsDescriptionEnrollmentRead          = AccessRightsDescription("EnrollmentRead")
	AccessRightsDescriptionEnrollmentWrite         = AccessRightsDescription("EnrollmentWrite")
	AccessRightsDescriptionDeviceConnect           = AccessRightsDescription("DeviceConnect")
	AccessRightsDescriptionRegistrationStatusRead  = AccessRightsDescription("RegistrationStatusRead")
	AccessRightsDescriptionRegistrationStatusWrite = AccessRightsDescription("RegistrationStatusWrite")
)

type AllocationPolicy string

const (
	AllocationPolicyHashed     = AllocationPolicy("Hashed")
	AllocationPolicyGeoLatency = AllocationPolicy("GeoLatency")
	AllocationPolicyStatic     = AllocationPolicy("Static")
)

type IotDpsSku string

const (
	IotDpsSkuS1 = IotDpsSku("S1")
)

type State string

const (
	StateActivating       = State("Activating")
	StateActive           = State("Active")
	StateDeleting         = State("Deleting")
	StateDeleted          = State("Deleted")
	StateActivationFailed = State("ActivationFailed")
	StateDeletionFailed   = State("DeletionFailed")
	StateTransitioning    = State("Transitioning")
	StateSuspending       = State("Suspending")
	StateSuspended        = State("Suspended")
	StateResuming         = State("Resuming")
	StateFailingOver      = State("FailingOver")
	StateFailoverFailed   = State("FailoverFailed")
)

func init() {
}
