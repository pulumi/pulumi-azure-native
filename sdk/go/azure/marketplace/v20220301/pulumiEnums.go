


package v20220301

type Accessibility string

const (
	AccessibilityUnknown                    = Accessibility("Unknown")
	AccessibilityPublic                     = Accessibility("Public")
	AccessibilityPrivateTenantOnLevel       = Accessibility("PrivateTenantOnLevel")
	AccessibilityPrivateSubscriptionOnLevel = Accessibility("PrivateSubscriptionOnLevel")
)

func init() {
}
