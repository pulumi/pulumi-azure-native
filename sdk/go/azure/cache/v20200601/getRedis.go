


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedis(ctx *pulumi.Context, args *LookupRedisArgs, opts ...pulumi.InvokeOption) (*LookupRedisResult, error) {
	var rv LookupRedisResult
	err := ctx.Invoke("azure-native:cache/v20200601:getRedis", args, &rv, opts...)
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

func LookupRedisOutput(ctx *pulumi.Context, args LookupRedisOutputArgs, opts ...pulumi.InvokeOption) LookupRedisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisResult, error) {
			args := v.(LookupRedisArgs)
			r, err := LookupRedis(ctx, &args, opts...)
			var s LookupRedisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisResultOutput)
}

type LookupRedisOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRedisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisArgs)(nil)).Elem()
}


type LookupRedisResultOutput struct{ *pulumi.OutputState }

func (LookupRedisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisResult)(nil)).Elem()
}

func (o LookupRedisResultOutput) ToLookupRedisResultOutput() LookupRedisResultOutput {
	return o
}

func (o LookupRedisResultOutput) ToLookupRedisResultOutputWithContext(ctx context.Context) LookupRedisResultOutput {
	return o
}

func (o LookupRedisResultOutput) AccessKeys() RedisAccessKeysResponseOutput {
	return o.ApplyT(func(v LookupRedisResult) RedisAccessKeysResponse { return v.AccessKeys }).(RedisAccessKeysResponseOutput)
}

func (o LookupRedisResultOutput) EnableNonSslPort() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *bool { return v.EnableNonSslPort }).(pulumi.BoolPtrOutput)
}

func (o LookupRedisResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) Instances() RedisInstanceDetailsResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []RedisInstanceDetailsResponse { return v.Instances }).(RedisInstanceDetailsResponseArrayOutput)
}

func (o LookupRedisResultOutput) LinkedServers() RedisLinkedServerResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []RedisLinkedServerResponse { return v.LinkedServers }).(RedisLinkedServerResponseArrayOutput)
}

func (o LookupRedisResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupRedisResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRedisResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupRedisResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupRedisResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupRedisResultOutput) RedisConfiguration() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *RedisCommonPropertiesResponseRedisConfiguration {
		return v.RedisConfiguration
	}).(RedisCommonPropertiesResponseRedisConfigurationPtrOutput)
}

func (o LookupRedisResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) ReplicasPerMaster() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *int { return v.ReplicasPerMaster }).(pulumi.IntPtrOutput)
}

func (o LookupRedisResultOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *int { return v.ShardCount }).(pulumi.IntPtrOutput)
}

func (o LookupRedisResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupRedisResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupRedisResultOutput) SslPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRedisResult) int { return v.SslPort }).(pulumi.IntOutput)
}

func (o LookupRedisResultOutput) StaticIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.StaticIP }).(pulumi.StringPtrOutput)
}

func (o LookupRedisResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o LookupRedisResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRedisResultOutput) TenantSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisResult) map[string]string { return v.TenantSettings }).(pulumi.StringMapOutput)
}

func (o LookupRedisResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRedisResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisResultOutput{})
}
