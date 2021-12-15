


package v20211101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNamespace(ctx *pulumi.Context, args *LookupNamespaceArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceResult, error) {
	var rv LookupNamespaceResult
	err := ctx.Invoke("azure-native:eventhub/v20211101:getNamespace", args, &rv, opts...)
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
	AlternateName              *string                             `pulumi:"alternateName"`
	ClusterArmId               *string                             `pulumi:"clusterArmId"`
	CreatedAt                  string                              `pulumi:"createdAt"`
	DisableLocalAuth           *bool                               `pulumi:"disableLocalAuth"`
	Encryption                 *EncryptionResponse                 `pulumi:"encryption"`
	Id                         string                              `pulumi:"id"`
	Identity                   *IdentityResponse                   `pulumi:"identity"`
	IsAutoInflateEnabled       *bool                               `pulumi:"isAutoInflateEnabled"`
	KafkaEnabled               *bool                               `pulumi:"kafkaEnabled"`
	Location                   *string                             `pulumi:"location"`
	MaximumThroughputUnits     *int                                `pulumi:"maximumThroughputUnits"`
	MetricId                   string                              `pulumi:"metricId"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	ServiceBusEndpoint         string                              `pulumi:"serviceBusEndpoint"`
	Sku                        *SkuResponse                        `pulumi:"sku"`
	Status                     string                              `pulumi:"status"`
	SystemData                 SystemDataResponse                  `pulumi:"systemData"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	UpdatedAt                  string                              `pulumi:"updatedAt"`
	ZoneRedundant              *bool                               `pulumi:"zoneRedundant"`
}
