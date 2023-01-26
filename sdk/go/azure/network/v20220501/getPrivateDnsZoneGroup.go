


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateDnsZoneGroup(ctx *pulumi.Context, args *LookupPrivateDnsZoneGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDnsZoneGroupResult, error) {
	var rv LookupPrivateDnsZoneGroupResult
	err := ctx.Invoke("azure-native:network/v20220501:getPrivateDnsZoneGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateDnsZoneGroupArgs struct {
	PrivateDnsZoneGroupName string `pulumi:"privateDnsZoneGroupName"`
	PrivateEndpointName     string `pulumi:"privateEndpointName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupPrivateDnsZoneGroupResult struct {
	Etag                  string                         `pulumi:"etag"`
	Id                    *string                        `pulumi:"id"`
	Name                  *string                        `pulumi:"name"`
	PrivateDnsZoneConfigs []PrivateDnsZoneConfigResponse `pulumi:"privateDnsZoneConfigs"`
	ProvisioningState     string                         `pulumi:"provisioningState"`
}

func LookupPrivateDnsZoneGroupOutput(ctx *pulumi.Context, args LookupPrivateDnsZoneGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDnsZoneGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateDnsZoneGroupResult, error) {
			args := v.(LookupPrivateDnsZoneGroupArgs)
			r, err := LookupPrivateDnsZoneGroup(ctx, &args, opts...)
			var s LookupPrivateDnsZoneGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateDnsZoneGroupResultOutput)
}

type LookupPrivateDnsZoneGroupOutputArgs struct {
	PrivateDnsZoneGroupName pulumi.StringInput `pulumi:"privateDnsZoneGroupName"`
	PrivateEndpointName     pulumi.StringInput `pulumi:"privateEndpointName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPrivateDnsZoneGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDnsZoneGroupArgs)(nil)).Elem()
}


type LookupPrivateDnsZoneGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDnsZoneGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDnsZoneGroupResult)(nil)).Elem()
}

func (o LookupPrivateDnsZoneGroupResultOutput) ToLookupPrivateDnsZoneGroupResultOutput() LookupPrivateDnsZoneGroupResultOutput {
	return o
}

func (o LookupPrivateDnsZoneGroupResultOutput) ToLookupPrivateDnsZoneGroupResultOutputWithContext(ctx context.Context) LookupPrivateDnsZoneGroupResultOutput {
	return o
}

func (o LookupPrivateDnsZoneGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPrivateDnsZoneGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateDnsZoneGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateDnsZoneGroupResultOutput) PrivateDnsZoneConfigs() PrivateDnsZoneConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) []PrivateDnsZoneConfigResponse { return v.PrivateDnsZoneConfigs }).(PrivateDnsZoneConfigResponseArrayOutput)
}

func (o LookupPrivateDnsZoneGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDnsZoneGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDnsZoneGroupResultOutput{})
}
