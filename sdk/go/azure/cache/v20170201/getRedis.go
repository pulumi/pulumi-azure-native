


package v20170201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRedis(ctx *pulumi.Context, args *LookupRedisArgs, opts ...pulumi.InvokeOption) (*LookupRedisResult, error) {
	var rv LookupRedisResult
	err := ctx.Invoke("azure-native:cache/v20170201:getRedis", args, &rv, opts...)
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
	AccessKeys         RedisAccessKeysResponse       `pulumi:"accessKeys"`
	EnableNonSslPort   *bool                         `pulumi:"enableNonSslPort"`
	HostName           string                        `pulumi:"hostName"`
	Id                 string                        `pulumi:"id"`
	LinkedServers      RedisLinkedServerListResponse `pulumi:"linkedServers"`
	Location           string                        `pulumi:"location"`
	Name               string                        `pulumi:"name"`
	Port               int                           `pulumi:"port"`
	ProvisioningState  string                        `pulumi:"provisioningState"`
	RedisConfiguration map[string]string             `pulumi:"redisConfiguration"`
	RedisVersion       string                        `pulumi:"redisVersion"`
	ShardCount         *int                          `pulumi:"shardCount"`
	Sku                *SkuResponse                  `pulumi:"sku"`
	SslPort            int                           `pulumi:"sslPort"`
	StaticIP           *string                       `pulumi:"staticIP"`
	SubnetId           *string                       `pulumi:"subnetId"`
	Tags               map[string]string             `pulumi:"tags"`
	TenantSettings     map[string]string             `pulumi:"tenantSettings"`
	Type               string                        `pulumi:"type"`
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

func (o LookupRedisResultOutput) LinkedServers() RedisLinkedServerListResponseOutput {
	return o.ApplyT(func(v LookupRedisResult) RedisLinkedServerListResponse { return v.LinkedServers }).(RedisLinkedServerListResponseOutput)
}

func (o LookupRedisResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRedisResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o LookupRedisResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) RedisConfiguration() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisResult) map[string]string { return v.RedisConfiguration }).(pulumi.StringMapOutput)
}

func (o LookupRedisResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

func (o LookupRedisResultOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *int { return v.ShardCount }).(pulumi.IntPtrOutput)
}

func (o LookupRedisResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupRedisResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
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

func init() {
	pulumi.RegisterOutputType(LookupRedisResultOutput{})
}
