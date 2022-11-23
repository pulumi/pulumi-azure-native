


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:network/v20190801:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointArgs struct {
	Expand              *string `pulumi:"expand"`
	PrivateEndpointName string  `pulumi:"privateEndpointName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointResult struct {
	Etag                                *string                                `pulumi:"etag"`
	Id                                  *string                                `pulumi:"id"`
	Location                            *string                                `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                string                                 `pulumi:"name"`
	NetworkInterfaces                   []NetworkInterfaceResponse             `pulumi:"networkInterfaces"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnectionResponse `pulumi:"privateLinkServiceConnections"`
	ProvisioningState                   string                                 `pulumi:"provisioningState"`
	Subnet                              *SubnetResponse                        `pulumi:"subnet"`
	Tags                                map[string]string                      `pulumi:"tags"`
	Type                                string                                 `pulumi:"type"`
}

func LookupPrivateEndpointOutput(ctx *pulumi.Context, args LookupPrivateEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointResult, error) {
			args := v.(LookupPrivateEndpointArgs)
			r, err := LookupPrivateEndpoint(ctx, &args, opts...)
			var s LookupPrivateEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointResultOutput)
}

type LookupPrivateEndpointOutputArgs struct {
	Expand              pulumi.StringPtrInput `pulumi:"expand"`
	PrivateEndpointName pulumi.StringInput    `pulumi:"privateEndpointName"`
	ResourceGroupName   pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupPrivateEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointArgs)(nil)).Elem()
}


type LookupPrivateEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointResult)(nil)).Elem()
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutput() LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) ToLookupPrivateEndpointResultOutputWithContext(ctx context.Context) LookupPrivateEndpointResultOutput {
	return o
}

func (o LookupPrivateEndpointResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPrivateEndpointResultOutput) ManualPrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.ManualPrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) PrivateLinkServiceConnections() PrivateLinkServiceConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) []PrivateLinkServiceConnectionResponse {
		return v.PrivateLinkServiceConnections
	}).(PrivateLinkServiceConnectionResponseArrayOutput)
}

func (o LookupPrivateEndpointResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointResultOutput) Subnet() SubnetResponsePtrOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) *SubnetResponse { return v.Subnet }).(SubnetResponsePtrOutput)
}

func (o LookupPrivateEndpointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrivateEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointResultOutput{})
}
