


package v20190301preview

type ConnectorBillingModel string

const (
	ConnectorBillingModelTrial       = ConnectorBillingModel("trial")
	ConnectorBillingModelAutoUpgrade = ConnectorBillingModel("autoUpgrade")
	ConnectorBillingModelPremium     = ConnectorBillingModel("premium")
	ConnectorBillingModelExpired     = ConnectorBillingModel("expired")
)

func init() {
}
