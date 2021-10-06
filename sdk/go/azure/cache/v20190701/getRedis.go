


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedis(ctx *pulumi.Context, args *LookupRedisArgs, opts ...pulumi.InvokeOption) (*LookupRedisResult, error) {
	var rv LookupRedisResult
	err := ctx.Invoke("azure-native:cache/v20190701:getRedis", args, &rv, opts...)
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
	AccessKeys         RedisAccessKeysResponse        `pulumi:"accessKeys"`
	EnableNonSslPort   *bool                          `pulumi:"enableNonSslPort"`
	HostName           string                         `pulumi:"hostName"`
	Id                 string                         `pulumi:"id"`
	Instances          []RedisInstanceDetailsResponse `pulumi:"instances"`
	LinkedServers      []RedisLinkedServerResponse    `pulumi:"linkedServers"`
	Location           string                         `pulumi:"location"`
	MinimumTlsVersion  *string                        `pulumi:"minimumTlsVersion"`
	Name               string                         `pulumi:"name"`
	Port               int                            `pulumi:"port"`
	ProvisioningState  string                         `pulumi:"provisioningState"`
	RedisConfiguration map[string]string              `pulumi:"redisConfiguration"`
	RedisVersion       string                         `pulumi:"redisVersion"`
	ReplicasPerMaster  *int                           `pulumi:"replicasPerMaster"`
	ShardCount         *int                           `pulumi:"shardCount"`
	Sku                SkuResponse                    `pulumi:"sku"`
	SslPort            int                            `pulumi:"sslPort"`
	StaticIP           *string                        `pulumi:"staticIP"`
	SubnetId           *string                        `pulumi:"subnetId"`
	Tags               map[string]string              `pulumi:"tags"`
	TenantSettings     map[string]string              `pulumi:"tenantSettings"`
	Type               string                         `pulumi:"type"`
	Zones              []string                       `pulumi:"zones"`
}
