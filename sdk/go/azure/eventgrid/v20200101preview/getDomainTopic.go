


package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainTopic(ctx *pulumi.Context, args *LookupDomainTopicArgs, opts ...pulumi.InvokeOption) (*LookupDomainTopicResult, error) {
	var rv LookupDomainTopicResult
	err := ctx.Invoke("azure-native:eventgrid/v20200101preview:getDomainTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainTopicArgs struct {
	DomainName        string `pulumi:"domainName"`
	DomainTopicName   string `pulumi:"domainTopicName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainTopicResult struct {
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	Type              string  `pulumi:"type"`
}
