


package v20180301beta

type PaymentChannelType string

const (
	PaymentChannelTypeSubscriptionDelegated = PaymentChannelType("SubscriptionDelegated")
	PaymentChannelTypeCustomerDelegated     = PaymentChannelType("CustomerDelegated")
)

func init() {
}
