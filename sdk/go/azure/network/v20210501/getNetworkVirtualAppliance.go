


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkVirtualAppliance(ctx *pulumi.Context, args *LookupNetworkVirtualApplianceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkVirtualApplianceResult, error) {
	var rv LookupNetworkVirtualApplianceResult
	err := ctx.Invoke("azure-native:network/v20210501:getNetworkVirtualAppliance", args, &rv, opts...)
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
	AddressPrefix               string                                  `pulumi:"addressPrefix"`
	BootStrapConfigurationBlobs []string                                `pulumi:"bootStrapConfigurationBlobs"`
	CloudInitConfiguration      *string                                 `pulumi:"cloudInitConfiguration"`
	CloudInitConfigurationBlobs []string                                `pulumi:"cloudInitConfigurationBlobs"`
	Etag                        string                                  `pulumi:"etag"`
	Id                          *string                                 `pulumi:"id"`
	Identity                    *ManagedServiceIdentityResponse         `pulumi:"identity"`
	InboundSecurityRules        []SubResourceResponse                   `pulumi:"inboundSecurityRules"`
	Location                    *string                                 `pulumi:"location"`
	Name                        string                                  `pulumi:"name"`
	NvaSku                      *VirtualApplianceSkuPropertiesResponse  `pulumi:"nvaSku"`
	ProvisioningState           string                                  `pulumi:"provisioningState"`
	SshPublicKey                *string                                 `pulumi:"sshPublicKey"`
	Tags                        map[string]string                       `pulumi:"tags"`
	Type                        string                                  `pulumi:"type"`
	VirtualApplianceAsn         *float64                                `pulumi:"virtualApplianceAsn"`
	VirtualApplianceNics        []VirtualApplianceNicPropertiesResponse `pulumi:"virtualApplianceNics"`
	VirtualApplianceSites       []SubResourceResponse                   `pulumi:"virtualApplianceSites"`
	VirtualHub                  *SubResourceResponse                    `pulumi:"virtualHub"`
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

func (o LookupNetworkVirtualApplianceResultOutput) AddressPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.AddressPrefix }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) BootStrapConfigurationBlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []string { return v.BootStrapConfigurationBlobs }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) CloudInitConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.CloudInitConfiguration }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) CloudInitConfigurationBlobs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []string { return v.CloudInitConfigurationBlobs }).(pulumi.StringArrayOutput)
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

func (o LookupNetworkVirtualApplianceResultOutput) InboundSecurityRules() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []SubResourceResponse { return v.InboundSecurityRules }).(SubResourceResponseArrayOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) NvaSku() VirtualApplianceSkuPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *VirtualApplianceSkuPropertiesResponse { return v.NvaSku }).(VirtualApplianceSkuPropertiesResponsePtrOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) SshPublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *string { return v.SshPublicKey }).(pulumi.StringPtrOutput)
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

func (o LookupNetworkVirtualApplianceResultOutput) VirtualApplianceSites() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) []SubResourceResponse { return v.VirtualApplianceSites }).(SubResourceResponseArrayOutput)
}

func (o LookupNetworkVirtualApplianceResultOutput) VirtualHub() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupNetworkVirtualApplianceResult) *SubResourceResponse { return v.VirtualHub }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkVirtualApplianceResultOutput{})
}
