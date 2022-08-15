


package v20190601

type AlertRuleState string

const (
	AlertRuleStateEnabled  = AlertRuleState("Enabled")
	AlertRuleStateDisabled = AlertRuleState("Disabled")
)

type Severity string

const (
	SeveritySev0 = Severity("Sev0")
	SeveritySev1 = Severity("Sev1")
	SeveritySev2 = Severity("Sev2")
	SeveritySev3 = Severity("Sev3")
	SeveritySev4 = Severity("Sev4")
)

func init() {
}
