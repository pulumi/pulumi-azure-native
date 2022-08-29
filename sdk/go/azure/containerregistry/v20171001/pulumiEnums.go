


package v20171001

type Action string

const (
	ActionAllow = Action("Allow")
)

type DefaultAction string

const (
	DefaultActionAllow = DefaultAction("Allow")
	DefaultActionDeny  = DefaultAction("Deny")
)

type SkuName string

const (
	SkuNameClassic  = SkuName("Classic")
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

type WebhookAction string

const (
	WebhookActionPush          = WebhookAction("push")
	WebhookActionDelete        = WebhookAction("delete")
	WebhookActionQuarantine    = WebhookAction("quarantine")
	WebhookAction_Chart_push   = WebhookAction("chart_push")
	WebhookAction_Chart_delete = WebhookAction("chart_delete")
)

type WebhookStatus string

const (
	WebhookStatusEnabled  = WebhookStatus("enabled")
	WebhookStatusDisabled = WebhookStatus("disabled")
)

func init() {
}
