


package v20180101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:relay/v20180101preview:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNamespaceArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceResult struct {
	CreatedAt          string            `pulumi:"createdAt"`
	Id                 string            `pulumi:"id"`
	Location           string            `pulumi:"location"`
	MetricId           string            `pulumi:"metricId"`
	Name               string            `pulumi:"name"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	ServiceBusEndpoint string            `pulumi:"serviceBusEndpoint"`
	Sku                *SkuResponse      `pulumi:"sku"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
	UpdatedAt          string            `pulumi:"updatedAt"`
}
