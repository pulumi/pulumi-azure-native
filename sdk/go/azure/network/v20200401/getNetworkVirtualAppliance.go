


package v20200401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkVirtualAppliance(ctx *pulumi.Context, args *LookupNetworkVirtualApplianceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkVirtualApplianceResult, error) {
	var rv LookupNetworkVirtualApplianceResult
	err := ctx.Invoke("azure-native:network/v20200401:getNetworkVirtualAppliance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkVirtualApplianceArgs struct {
	Expand                      *string `pulumi:"expand"`
	NetworkVirtualApplianceName string  `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           string  `pulumi:"resourceGroupName"`
}


type LookupNetworkVirtualApplianceResult struct {
	BootStrapConfigurationBlob []string                                `pulumi:"bootStrapConfigurationBlob"`
	CloudInitConfigurationBlob []string                                `pulumi:"cloudInitConfigurationBlob"`
	Etag                       string                                  `pulumi:"etag"`
	Id                         *string                                 `pulumi:"id"`
	Identity                   *ManagedServiceIdentityResponse         `pulumi:"identity"`
	Location                   *string                                 `pulumi:"location"`
	Name                       string                                  `pulumi:"name"`
	ProvisioningState          string                                  `pulumi:"provisioningState"`
	Sku                        *VirtualApplianceSkuPropertiesResponse  `pulumi:"sku"`
	Tags                       map[string]string                       `pulumi:"tags"`
	Type                       string                                  `pulumi:"type"`
	VirtualApplianceAsn        *float64                                `pulumi:"virtualApplianceAsn"`
	VirtualApplianceNics       []VirtualApplianceNicPropertiesResponse `pulumi:"virtualApplianceNics"`
	VirtualHub                 *SubResourceResponse                    `pulumi:"virtualHub"`
}

func LookupNetworkVirtualApplianceOutput(ctx *pulumi.Context, args LookupNetworkVirtualApplianceOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkVirtualApplianceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkVirtualApplianceResult, error) {
			args := v.(LookupNetworkVirtualApplianceArgs)
			r, err := LookupNetworkVirtualAppliance(ctx, &args, opts...)
			var s LookupNetworkVirtualApplianceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkVirtualApplianceResultOutput)
}

type LookupNetworkVirtualApplianceOutputArgs struct {
	Expand                      pulumi.StringPtrInput `pulumi:"expand"`
	NetworkVirtualApplianceName pulumi.StringInput    `pulumi:"networkVirtualApplianceName"`
	ResourceGroupName           pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupNetworkVirtualApplianceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkVirtualApplianceArgs)(nil)).Elem()
}


type LookupNetworkVirtualApplianceResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkVirtualApplianceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkVirtualApplianceResult)(nil)).Elem()
}

func (o LookupNetworkVirtualApplianceResultOutput) ToLookupNetworkVirtualApplianceResultOutput() LookupNetworkVirtualApplianceResultOutput {
	return o
}

func (o LookupNetworkVirtualApplianceResultOutput) ToLookupNetworkVirtualApplianceResultOutputWithContext(ctx context.Context) LookupNetworkVirtualApplianceResultOutput {
	return o
}

func (o LookupNetworkVirtualApplianceResultOutput) BootStrapConfigurationBlob() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []string { return v.BootStrapConfigurationBlob }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) CloudInitConfigurationBlob() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []string { return v.CloudInitConfigurationBlob }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Sku() VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *VirtualApplianceSkuPropertiesResponse { return v.Sku }).(VirtualApplianceSkuPropertiesResponsePtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) VirtualApplianceAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *float64 { return v.VirtualApplianceAsn }).(pulumi.Float64PtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) VirtualApplianceNics() VirtualApplianceNicPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []VirtualApplianceNicPropertiesResponse {
		return v.VirtualApplianceNics
	}).(VirtualApplianceNicPropertiesResponseArrayOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkVirtualApplianceResultOutput{})
}
