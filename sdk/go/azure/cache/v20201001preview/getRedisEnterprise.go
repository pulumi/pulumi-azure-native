


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedisEnterprise(ctx *pulumi.Context, args *LookupRedisEnterpriseArgs, opts ...pulumi.InvokeOption) (*LookupRedisEnterpriseResult, error) {
	var rv LookupRedisEnterpriseResult
	err := ctx.Invoke("azure-native:cache/v20201001preview:getRedisEnterprise", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisEnterpriseArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRedisEnterpriseResult struct {
	HostName                   string                              `pulumi:"hostName"`
	Id                         string                              `pulumi:"id"`
	Location                   string                              `pulumi:"location"`
	MinimumTlsVersion          *string                             `pulumi:"minimumTlsVersion"`
	Name                       string                              `pulumi:"name"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                              `pulumi:"provisioningState"`
	RedisVersion               string                              `pulumi:"redisVersion"`
	ResourceState              string                              `pulumi:"resourceState"`
	Sku                        EnterpriseSkuResponse               `pulumi:"sku"`
	Tags                       map[string]string                   `pulumi:"tags"`
	Type                       string                              `pulumi:"type"`
	Zones                      []string                            `pulumi:"zones"`
}
