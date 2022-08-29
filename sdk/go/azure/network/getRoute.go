


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoute(ctx *pulumi.Context, args *LookupRouteArgs, opts ...pulumi.InvokeOption) (*LookupRouteResult, error) {
	var rv LookupRouteResult
	err := ctx.Invoke("azure-native:network:getRoute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RouteName         string `pulumi:"routeName"`
	RouteTableName    string `pulumi:"routeTableName"`
}


type LookupRouteResult struct {
	AddressPrefix     *string `pulumi:"addressPrefix"`
	Etag              string  `pulumi:"etag"`
	HasBgpOverride    *bool   `pulumi:"hasBgpOverride"`
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	NextHopIpAddress  *string `pulumi:"nextHopIpAddress"`
	NextHopType       string  `pulumi:"nextHopType"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Type              *string `pulumi:"type"`
}

func LookupRouteOutput(ctx *pulumi.Context, args LookupRouteOutputArgs, opts ...pulumi.InvokeOption) LookupRouteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteResult, error) {
			args := v.(LookupRouteArgs)
			r, err := LookupRoute(ctx, &args, opts...)
			var s LookupRouteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteResultOutput)
}

type LookupRouteOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RouteName         pulumi.StringInput `pulumi:"routeName"`
	RouteTableName    pulumi.StringInput `pulumi:"routeTableName"`
}

func (LookupRouteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteArgs)(nil)).Elem()
}


type LookupRouteResultOutput struct{ *pulumi.OutputState }

func (LookupRouteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteResult)(nil)).Elem()
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutput() LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) ToLookupRouteResultOutputWithContext(ctx context.Context) LookupRouteResultOutput {
	return o
}

func (o LookupRouteResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) HasBgpOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *bool { return v.HasBgpOverride }).(pulumi.BoolPtrOutput)
}

func (o LookupRouteResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) NextHopIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.NextHopIpAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRouteResultOutput) NextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.NextHopType }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRouteResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteResultOutput{})
}
