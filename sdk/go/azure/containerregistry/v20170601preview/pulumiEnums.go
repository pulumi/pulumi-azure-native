


package v20170601preview

type SkuName string

const (
	SkuNameBasic             = SkuName("Basic")
	SkuName_Managed_Basic    = SkuName("Managed_Basic")
	SkuName_Managed_Standard = SkuName("Managed_Standard")
	SkuName_Managed_Premium  = SkuName("Managed_Premium")
)

type WebhookAction string

const (
	WebhookActionPush   = WebhookAction("push")
	WebhookActionDelete = WebhookAction("delete")
)

type WebhookStatus string

const (
	WebhookStatusEnabled  = WebhookStatus("enabled")
	WebhookStatusDisabled = WebhookStatus("disabled")
)

func init() {
}
