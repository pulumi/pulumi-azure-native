


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDnsResourceReferenceByTarResources(ctx *pulumi.Context, args *GetDnsResourceReferenceByTarResourcesArgs, opts ...pulumi.InvokeOption) (*GetDnsResourceReferenceByTarResourcesResult, error) {
	var rv GetDnsResourceReferenceByTarResourcesResult
	err := ctx.Invoke("azure-native:network:getDnsResourceReferenceByTarResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDnsResourceReferenceByTarResourcesArgs struct {
	TargetResources []SubResource `pulumi:"targetResources"`
}


type GetDnsResourceReferenceByTarResourcesResult struct {
	DnsResourceReferences []DnsResourceReferenceResponse `pulumi:"dnsResourceReferences"`
}

func GetDnsResourceReferenceByTarResourcesOutput(ctx *pulumi.Context, args GetDnsResourceReferenceByTarResourcesOutputArgs, opts ...pulumi.InvokeOption) GetDnsResourceReferenceByTarResourcesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsResourceReferenceByTarResourcesResult, error) {
			args := v.(GetDnsResourceReferenceByTarResourcesArgs)
			r, err := GetDnsResourceReferenceByTarResources(ctx, &args, opts...)
			var s GetDnsResourceReferenceByTarResourcesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDnsResourceReferenceByTarResourcesResultOutput)
}

type GetDnsResourceReferenceByTarResourcesOutputArgs struct {
	TargetResources SubResourceArrayInput `pulumi:"targetResources"`
}

func (GetDnsResourceReferenceByTarResourcesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsResourceReferenceByTarResourcesArgs)(nil)).Elem()
}


type GetDnsResourceReferenceByTarResourcesResultOutput struct{ *pulumi.OutputState }

func (GetDnsResourceReferenceByTarResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsResourceReferenceByTarResourcesResult)(nil)).Elem()
}

func (o GetDnsResourceReferenceByTarResourcesResultOutput) ToGetDnsResourceReferenceByTarResourcesResultOutput() GetDnsResourceReferenceByTarResourcesResultOutput {
	return o
}

func (o GetDnsResourceReferenceByTarResourcesResultOutput) ToGetDnsResourceReferenceByTarResourcesResultOutputWithContext(ctx context.Context) GetDnsResourceReferenceByTarResourcesResultOutput {
	return o
}

func (o GetDnsResourceReferenceByTarResourcesResultOutput) DnsResourceReferences() DnsResourceReferenceResponseArrayOutput {
	return o.ApplyT(func(v GetDnsResourceReferenceByTarResourcesResult) []DnsResourceReferenceResponse {
		return v.DnsResourceReferences
	}).(DnsResourceReferenceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsResourceReferenceByTarResourcesResultOutput{})
}
