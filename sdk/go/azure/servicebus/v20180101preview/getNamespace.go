


package v20180101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus/v20180101preview:getNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNamespaceArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNamespaceResult struct {
	CreatedAt          string              `pulumi:"createdAt"`
	Encryption         *EncryptionResponse `pulumi:"encryption"`
	Id                 string              `pulumi:"id"`
	Identity           *IdentityResponse   `pulumi:"identity"`
	Location           string              `pulumi:"location"`
	MetricId           string              `pulumi:"metricId"`
	Name               string              `pulumi:"name"`
	ProvisioningState  string              `pulumi:"provisioningState"`
	ServiceBusEndpoint string              `pulumi:"serviceBusEndpoint"`
	Sku                *SBSkuResponse      `pulumi:"sku"`
	Status             string              `pulumi:"status"`
	Tags               map[string]string   `pulumi:"tags"`
	Type               string              `pulumi:"type"`
	UpdatedAt          string              `pulumi:"updatedAt"`
	ZoneRedundant      *bool               `pulumi:"zoneRedundant"`
}


func (val *LookupNamespaceResult) Defaults() *LookupNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	tmp.Identity = tmp.Identity.Defaults()

	return &tmp
}
