


package v20181101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSwiftVirtualNetworkConnection(ctx *pulumi.Context, args *LookupWebAppSwiftVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSwiftVirtualNetworkConnectionResult, error) {
	var rv LookupWebAppSwiftVirtualNetworkConnectionResult
	err := ctx.Invoke("azure-native:web/v20181101:getWebAppSwiftVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSwiftVirtualNetworkConnectionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppSwiftVirtualNetworkConnectionResult struct {
	Id               string  `pulumi:"id"`
	Kind             *string `pulumi:"kind"`
	Name             string  `pulumi:"name"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	SwiftSupported   *bool   `pulumi:"swiftSupported"`
	Type             string  `pulumi:"type"`
}

func LookupWebAppSwiftVirtualNetworkConnectionOutput(ctx *pulumi.Context, args LookupWebAppSwiftVirtualNetworkConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSwiftVirtualNetworkConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppSwiftVirtualNetworkConnectionResult, error) {
			args := v.(LookupWebAppSwiftVirtualNetworkConnectionArgs)
			r, err := LookupWebAppSwiftVirtualNetworkConnection(ctx, &args, opts...)
			var s LookupWebAppSwiftVirtualNetworkConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppSwiftVirtualNetworkConnectionResultOutput)
}

type LookupWebAppSwiftVirtualNetworkConnectionOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppSwiftVirtualNetworkConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSwiftVirtualNetworkConnectionArgs)(nil)).Elem()
}


type LookupWebAppSwiftVirtualNetworkConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSwiftVirtualNetworkConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSwiftVirtualNetworkConnectionResult)(nil)).Elem()
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) ToLookupWebAppSwiftVirtualNetworkConnectionResultOutput() LookupWebAppSwiftVirtualNetworkConnectionResultOutput {
	return o
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) ToLookupWebAppSwiftVirtualNetworkConnectionResultOutputWithContext(ctx context.Context) LookupWebAppSwiftVirtualNetworkConnectionResultOutput {
	return o
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionResult) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) SwiftSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionResult) *bool { return v.SwiftSupported }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppSwiftVirtualNetworkConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSwiftVirtualNetworkConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSwiftVirtualNetworkConnectionResultOutput{})
}
