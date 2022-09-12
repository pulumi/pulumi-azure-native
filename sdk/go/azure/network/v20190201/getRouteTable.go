


package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	var rv LookupRouteTableResult
	err := ctx.Invoke("azure-native:network/v20190201:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteTableArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	RouteTableName    string  `pulumi:"routeTableName"`
}


type LookupRouteTableResult struct {
	DisableBgpRoutePropagation *bool             `pulumi:"disableBgpRoutePropagation"`
	Etag                       *string           `pulumi:"etag"`
	Id                         *string           `pulumi:"id"`
	Location                   *string           `pulumi:"location"`
	Name                       string            `pulumi:"name"`
	ProvisioningState          *string           `pulumi:"provisioningState"`
	Routes                     []RouteResponse   `pulumi:"routes"`
	Subnets                    []SubnetResponse  `pulumi:"subnets"`
	Tags                       map[string]string `pulumi:"tags"`
	Type                       string            `pulumi:"type"`
}

func LookupRouteTableOutput(ctx *pulumi.Context, args LookupRouteTableOutputArgs, opts ...pulumi.InvokeOption) LookupRouteTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteTableResult, error) {
			args := v.(LookupRouteTableArgs)
			r, err := LookupRouteTable(ctx, &args, opts...)
			var s LookupRouteTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteTableResultOutput)
}

type LookupRouteTableOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	RouteTableName    pulumi.StringInput    `pulumi:"routeTableName"`
}

func (LookupRouteTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableArgs)(nil)).Elem()
}


type LookupRouteTableResultOutput struct{ *pulumi.OutputState }

func (LookupRouteTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableResult)(nil)).Elem()
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutput() LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutputWithContext(ctx context.Context) LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) DisableBgpRoutePropagation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *bool { return v.DisableBgpRoutePropagation }).(pulumi.BoolPtrOutput)
}

func (o LookupRouteTableResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupRouteTableResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRouteTableResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupRouteTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteTableResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupRouteTableResultOutput) Routes() RouteResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []RouteResponse { return v.Routes }).(RouteResponseArrayOutput)
}

func (o LookupRouteTableResultOutput) Subnets() SubnetResponseArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []SubnetResponse { return v.Subnets }).(SubnetResponseArrayOutput)
}

func (o LookupRouteTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRouteTableResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteTableResultOutput{})
}
