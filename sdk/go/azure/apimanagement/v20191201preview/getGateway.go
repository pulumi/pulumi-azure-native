


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	var rv LookupGatewayResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayArgs struct {
	GatewayId         string `pulumi:"gatewayId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupGatewayResult struct {
	Description  *string                               `pulumi:"description"`
	Id           string                                `pulumi:"id"`
	LocationData *ResourceLocationDataContractResponse `pulumi:"locationData"`
	Name         string                                `pulumi:"name"`
	Type         string                                `pulumi:"type"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

type LookupGatewayOutputArgs struct {
	GatewayId         pulumi.StringInput `pulumi:"gatewayId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}


type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) LocationData() ResourceLocationDataContractResponsePtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *ResourceLocationDataContractResponse { return v.LocationData }).(ResourceLocationDataContractResponsePtrOutput)
}

func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
