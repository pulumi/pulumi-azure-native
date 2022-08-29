


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRedisEnterprise(ctx *pulumi.Context, args *LookupRedisEnterpriseArgs, opts ...pulumi.InvokeOption) (*LookupRedisEnterpriseResult, error) {
	var rv LookupRedisEnterpriseResult
	err := ctx.Invoke("azure-native:cache/v20210301:getRedisEnterprise", args, &rv, opts...)
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

func LookupRedisEnterpriseOutput(ctx *pulumi.Context, args LookupRedisEnterpriseOutputArgs, opts ...pulumi.InvokeOption) LookupRedisEnterpriseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisEnterpriseResult, error) {
			args := v.(LookupRedisEnterpriseArgs)
			r, err := LookupRedisEnterprise(ctx, &args, opts...)
			var s LookupRedisEnterpriseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisEnterpriseResultOutput)
}

type LookupRedisEnterpriseOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupRedisEnterpriseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisEnterpriseArgs)(nil)).Elem()
}


type LookupRedisEnterpriseResultOutput struct{ *pulumi.OutputState }

func (LookupRedisEnterpriseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisEnterpriseResult)(nil)).Elem()
}

func (o LookupRedisEnterpriseResultOutput) ToLookupRedisEnterpriseResultOutput() LookupRedisEnterpriseResultOutput {
	return o
}

func (o LookupRedisEnterpriseResultOutput) ToLookupRedisEnterpriseResultOutputWithContext(ctx context.Context) LookupRedisEnterpriseResultOutput {
	return o
}

func (o LookupRedisEnterpriseResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.HostName }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) MinimumTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) *string { return v.MinimumTlsVersion }).(pulumi.StringPtrOutput)
}

func (o LookupRedisEnterpriseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o LookupRedisEnterpriseResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) RedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.RedisVersion }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) Sku() EnterpriseSkuResponseOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) EnterpriseSkuResponse { return v.Sku }).(EnterpriseSkuResponseOutput)
}

func (o LookupRedisEnterpriseResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRedisEnterpriseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupRedisEnterpriseResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRedisEnterpriseResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisEnterpriseResultOutput{})
}
