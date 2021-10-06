


package v20200401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEventChannel(ctx *pulumi.Context, args *LookupEventChannelArgs, opts ...pulumi.InvokeOption) (*LookupEventChannelResult, error) {
	var rv LookupEventChannelResult
	err := ctx.Invoke("azure-native:eventgrid/v20200401preview:getEventChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventChannelArgs struct {
	EventChannelName     string `pulumi:"eventChannelName"`
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupEventChannelResult struct {
	Destination                     *EventChannelDestinationResponse `pulumi:"destination"`
	ExpirationTimeIfNotActivatedUtc *string                          `pulumi:"expirationTimeIfNotActivatedUtc"`
	Filter                          *EventChannelFilterResponse      `pulumi:"filter"`
	Id                              string                           `pulumi:"id"`
	Name                            string                           `pulumi:"name"`
	PartnerTopicFriendlyDescription *string                          `pulumi:"partnerTopicFriendlyDescription"`
	PartnerTopicReadinessState      string                           `pulumi:"partnerTopicReadinessState"`
	ProvisioningState               string                           `pulumi:"provisioningState"`
	Source                          *EventChannelSourceResponse      `pulumi:"source"`
	SystemData                      SystemDataResponse               `pulumi:"systemData"`
	Type                            string                           `pulumi:"type"`
}
