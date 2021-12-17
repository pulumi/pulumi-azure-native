


package saas

type PaymentChannelType string

const (
	PaymentChannelTypeSubscriptionDelegated = PaymentChannelType("SubscriptionDelegated")
	PaymentChannelTypeCustomerDelegated     = PaymentChannelType("CustomerDelegated")
)

func init() {
}
