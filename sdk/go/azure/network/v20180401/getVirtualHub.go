


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVirtualHub(ctx *pulumi.Context, args *LookupVirtualHubArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubResult, error) {
	var rv LookupVirtualHubResult
	err := ctx.Invoke("azure-native:network/v20180401:getVirtualHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubResult struct {
	AddressPrefix                *string                               `pulumi:"addressPrefix"`
	Etag                         string                                `pulumi:"etag"`
	HubVirtualNetworkConnections []HubVirtualNetworkConnectionResponse `pulumi:"hubVirtualNetworkConnections"`
	Id                           *string                               `pulumi:"id"`
	Location                     string                                `pulumi:"location"`
	Name                         string                                `pulumi:"name"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	Tags                         map[string]string                     `pulumi:"tags"`
	Type                         string                                `pulumi:"type"`
	VirtualWan                   *SubResourceResponse                  `pulumi:"virtualWan"`
}

func LookupVirtualHubOutput(ctx *pulumi.Context, args LookupVirtualHubOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHubResult, error) {
			args := v.(LookupVirtualHubArgs)
			r, err := LookupVirtualHub(ctx, &args, opts...)
			var s LookupVirtualHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHubResultOutput)
}

type LookupVirtualHubOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualHubName    pulumi.StringInput `pulumi:"virtualHubName"`
}

func (LookupVirtualHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubArgs)(nil)).Elem()
}


type LookupVirtualHubResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHubResult)(nil)).Elem()
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutput() LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) ToLookupVirtualHubResultOutputWithContext(ctx context.Context) LookupVirtualHubResultOutput {
	return o
}

func (o LookupVirtualHubResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) HubVirtualNetworkConnections() HubVirtualNetworkConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) []HubVirtualNetworkConnectionResponse {
		return v.HubVirtualNetworkConnections
	}).(HubVirtualNetworkConnectionResponseArrayOutput)
}

func (o LookupVirtualHubResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHubResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualHubResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualHubResultOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualHubResult) *SubResourceResponse { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHubResultOutput{})
}
