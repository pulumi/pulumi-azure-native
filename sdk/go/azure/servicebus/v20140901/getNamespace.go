


package v20140901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus/v20140901:getNamespace", args, &rv, opts...)
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
	CreateACSNamespace *bool             `pulumi:"createACSNamespace"`
	CreatedAt          string            `pulumi:"createdAt"`
	Enabled            *bool             `pulumi:"enabled"`
	Id                 string            `pulumi:"id"`
	Location           string            `pulumi:"location"`
	Name               string            `pulumi:"name"`
	ProvisioningState  string            `pulumi:"provisioningState"`
	ServiceBusEndpoint string            `pulumi:"serviceBusEndpoint"`
	Sku                *SkuResponse      `pulumi:"sku"`
	Status             *string           `pulumi:"status"`
	Tags               map[string]string `pulumi:"tags"`
	Type               string            `pulumi:"type"`
	UpdatedAt          string            `pulumi:"updatedAt"`
}
