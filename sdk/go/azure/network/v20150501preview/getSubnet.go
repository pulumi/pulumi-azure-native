


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SubnetName         string `pulumi:"subnetName"`
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}


type LookupSubnetResult struct {
	AddressPrefix        string                `pulumi:"addressPrefix"`
	Etag                 *string               `pulumi:"etag"`
	Id                   *string               `pulumi:"id"`
	IpConfigurations     []SubResourceResponse `pulumi:"ipConfigurations"`
	Name                 *string               `pulumi:"name"`
	NetworkSecurityGroup *SubResourceResponse  `pulumi:"networkSecurityGroup"`
	ProvisioningState    *string               `pulumi:"provisioningState"`
	RouteTable           *SubResourceResponse  `pulumi:"routeTable"`
}

func LookupSubnetOutput(ctx *pulumi.Context, args LookupSubnetOutputArgs, opts ...pulumi.InvokeOption) LookupSubnetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubnetResult, error) {
			args := v.(LookupSubnetArgs)
			r, err := LookupSubnet(ctx, &args, opts...)
			var s LookupSubnetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubnetResultOutput)
}

type LookupSubnetOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	SubnetName         pulumi.StringInput `pulumi:"subnetName"`
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (LookupSubnetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetArgs)(nil)).Elem()
}


type LookupSubnetResultOutput struct{ *pulumi.OutputState }

func (LookupSubnetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubnetResult)(nil)).Elem()
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutput() LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) ToLookupSubnetResultOutputWithContext(ctx context.Context) LookupSubnetResultOutput {
	return o
}

func (o LookupSubnetResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubnetResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

func (o LookupSubnetResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) IpConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupSubnetResult) []SubResourceResponse { return v.IpConfigurations }).(SubResourceResponseArrayOutput)
}

func (o LookupSubnetResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) NetworkSecurityGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *SubResourceResponse { return v.NetworkSecurityGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupSubnetResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupSubnetResultOutput) RouteTable() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupSubnetResult) *SubResourceResponse { return v.RouteTable }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubnetResultOutput{})
}
