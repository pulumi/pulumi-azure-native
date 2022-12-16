


package v20180801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInterfaceEndpoint(ctx *pulumi.Context, args *LookupInterfaceEndpointArgs, opts ...pulumi.InvokeOption) (*LookupInterfaceEndpointResult, error) {
	var rv LookupInterfaceEndpointResult
	err := ctx.Invoke("azure-native:network/v20180801:getInterfaceEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInterfaceEndpointArgs struct {
	Expand                *string `pulumi:"expand"`
	InterfaceEndpointName string  `pulumi:"interfaceEndpointName"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
}


type LookupInterfaceEndpointResult struct {
	EndpointService   *EndpointServiceResponse   `pulumi:"endpointService"`
	Etag              *string                    `pulumi:"etag"`
	Fqdn              *string                    `pulumi:"fqdn"`
	Id                *string                    `pulumi:"id"`
	Location          *string                    `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	NetworkInterfaces []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	Owner             string                     `pulumi:"owner"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Subnet            *SubnetResponse            `pulumi:"subnet"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}

func LookupInterfaceEndpointOutput(ctx *pulumi.Context, args LookupInterfaceEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupInterfaceEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInterfaceEndpointResult, error) {
			args := v.(LookupInterfaceEndpointArgs)
			r, err := LookupInterfaceEndpoint(ctx, &args, opts...)
			var s LookupInterfaceEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInterfaceEndpointResultOutput)
}

type LookupInterfaceEndpointOutputArgs struct {
	Expand                pulumi.StringPtrInput `pulumi:"expand"`
	InterfaceEndpointName pulumi.StringInput    `pulumi:"interfaceEndpointName"`
	ResourceGroupName     pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupInterfaceEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterfaceEndpointArgs)(nil)).Elem()
}


type LookupInterfaceEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupInterfaceEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInterfaceEndpointResult)(nil)).Elem()
}

func (o LookupInterfaceEndpointResultOutput) ToLookupInterfaceEndpointResultOutput() LookupInterfaceEndpointResultOutput {
	return o
}

func (o LookupInterfaceEndpointResultOutput) ToLookupInterfaceEndpointResultOutputWithContext(ctx context.Context) LookupInterfaceEndpointResultOutput {
	return o
}

func (o LookupInterfaceEndpointResultOutput) EndpointService() EndpointServiceResponsePtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *EndpointServiceResponse { return v.EndpointService }).(EndpointServiceResponsePtrOutput)
}

func (o LookupInterfaceEndpointResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupInterfaceEndpointResultOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o LookupInterfaceEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInterfaceEndpointResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupInterfaceEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInterfaceEndpointResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupInterfaceEndpointResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupInterfaceEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupInterfaceEndpointResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

func (o LookupInterfaceEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupInterfaceEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInterfaceEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInterfaceEndpointResultOutput{})
}
