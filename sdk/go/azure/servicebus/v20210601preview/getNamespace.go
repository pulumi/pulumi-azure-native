


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:servicebus/v20210601preview:getNamespace", args, &rv, opts...)
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
	CreatedAt                  string                              `pulumi:"createdAt"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	Encryption                 *EncryptionResponse                 `pulumi:"encryption"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityResponse                   `pulumi:"identity"`
	Location                   string                              `pulumi:"location"`
	MetricId                   string                              `pulumi:"metricId"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	ServiceBusEndpoint         string                              `pulumi:"serviceBusEndpoint"`
	Sku                        *SBSkuResponse                      `pulumi:"sku"`
	Status                     string                              `pulumi:"status"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	UpdatedAt                  string                              `pulumi:"updatedAt"`
	ZoneRedundant              *bool                               `pulumi:"zoneRedundant"`
}


func (val *LookupNamespaceResult) Defaults() *LookupNamespaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Encryption = tmp.Encryption.Defaults()

	return &tmp
}
