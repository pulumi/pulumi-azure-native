


package cache

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedis(ctx *pulumi.Context, args *LookupRedisArgs, opts ...pulumi.InvokeOption) (*LookupRedisResult, error) {
	var rv LookupRedisResult
	err := ctx.Invoke("azure-native:cache:getRedis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRedisArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRedisResult struct {
	AccessKeys                 RedisAccessKeysResponse                          `pulumi:"accessKeys"`
	EnableNonSslPort           *bool                                            `pulumi:"enableNonSslPort"`
	HostName                   string                                           `pulumi:"hostName"`
	Id                         string                                           `pulumi:"id"`
	Instances                  []RedisInstanceDetailsResponse                   `pulumi:"instances"`
	LinkedServers              []RedisLinkedServerResponse                      `pulumi:"linkedServers"`
	Location                   string                                           `pulumi:"location"`
	MinimumTlsVersion          *string                                          `pulumi:"minimumTlsVersion"`
	Name                       string                                           `pulumi:"name"`
	Port                       int                                              `pulumi:"port"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse              `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                           `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                          `pulumi:"publicNetworkAccess"`
	RedisConfiguration         *RedisCommonPropertiesResponseRedisConfiguration `pulumi:"redisConfiguration"`
	RedisVersion               string                                           `pulumi:"redisVersion"`
	ReplicasPerMaster          *int                                             `pulumi:"replicasPerMaster"`
	ShardCount                 *int                                             `pulumi:"shardCount"`
	Sku                        SkuResponse                                      `pulumi:"sku"`
	SslPort                    int                                              `pulumi:"sslPort"`
	StaticIP                   *string                                          `pulumi:"staticIP"`
	SubnetId                   *string                                          `pulumi:"subnetId"`
	Tags                       map[string]string                                `pulumi:"tags"`
	TenantSettings             map[string]string                                `pulumi:"tenantSettings"`
	Type                       string                                           `pulumi:"type"`
	Zones                      []string                                         `pulumi:"zones"`
}


func (val *LookupRedisResult) Defaults() *LookupRedisResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableNonSslPort) {
		enableNonSslPort_ := false
		tmp.EnableNonSslPort = &enableNonSslPort_
	}
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}
