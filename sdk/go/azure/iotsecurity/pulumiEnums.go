


package iotsecurity

type MdeIntegration string

const (
	MdeIntegrationDisabled = MdeIntegration("Disabled")
	MdeIntegrationEnabled  = MdeIntegration("Enabled")
)

type OnboardingKind string

const (
	OnboardingKindDefault         = OnboardingKind("Default")
	OnboardingKindMigratedToAzure = OnboardingKind("MigratedToAzure")
	OnboardingKindEvaluation      = OnboardingKind("Evaluation")
	OnboardingKindPurchased       = OnboardingKind("Purchased")
)

type SensorType string

const (
	SensorTypeOt         = SensorType("Ot")
	SensorTypeEnterprise = SensorType("Enterprise")
)

func init() {
}
