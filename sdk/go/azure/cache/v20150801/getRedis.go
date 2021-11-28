


package v20150801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedis(ctx *pulumi.Context, args *LookupRedisArgs, opts ...pulumi.InvokeOption) (*LookupRedisResult, error) {
	var rv LookupRedisResult
	err := ctx.Invoke("azure-native:cache/v20150801:getRedis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRedisResult struct {
	EnableNonSslPort   *bool             `pulumi:"enableNonSslPort"`
	HostName           *string           `pulumi:"hostName"`
	Id                 string            `pulumi:"id"`
	Location           string            `pulumi:"location"`
	Name               string            `pulumi:"name"`
	Port               *int              `pulumi:"port"`
	ProvisioningState  *string           `pulumi:"provisioningState"`
	RedisConfiguration map[string]string `pulumi:"redisConfiguration"`
	RedisVersion       *string           `pulumi:"redisVersion"`
	ShardCount         *int              `pulumi:"shardCount"`
	Sku                SkuResponse       `pulumi:"sku"`
	SslPort            *int              `pulumi:"sslPort"`
	StaticIP           *string           `pulumi:"staticIP"`
	Subnet             *string           `pulumi:"subnet"`
	Tags               map[string]string `pulumi:"tags"`
	TenantSettings     map[string]string `pulumi:"tenantSettings"`
	Type               string            `pulumi:"type"`
	VirtualNetwork     *string           `pulumi:"virtualNetwork"`
}
