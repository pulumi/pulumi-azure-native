


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNodeType(ctx *pulumi.Context, args *LookupNodeTypeArgs, opts ...pulumi.InvokeOption) (*LookupNodeTypeResult, error) {
	var rv LookupNodeTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20220801preview:getNodeType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNodeTypeArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	NodeTypeName      string `pulumi:"nodeTypeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNodeTypeResult struct {
	AdditionalDataDisks          []VmssDataDiskResponse            `pulumi:"additionalDataDisks"`
	ApplicationPorts             *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	Capacities                   map[string]string                 `pulumi:"capacities"`
	DataDiskLetter               *string                           `pulumi:"dataDiskLetter"`
	DataDiskSizeGB               *int                              `pulumi:"dataDiskSizeGB"`
	DataDiskType                 *string                           `pulumi:"dataDiskType"`
	EnableAcceleratedNetworking  *bool                             `pulumi:"enableAcceleratedNetworking"`
	EnableEncryptionAtHost       *bool                             `pulumi:"enableEncryptionAtHost"`
	EnableOverProvisioning       *bool                             `pulumi:"enableOverProvisioning"`
	EphemeralPorts               *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	EvictionPolicy               *string                           `pulumi:"evictionPolicy"`
	FrontendConfigurations       []FrontendConfigurationResponse   `pulumi:"frontendConfigurations"`
	HostGroupId                  *string                           `pulumi:"hostGroupId"`
	Id                           string                            `pulumi:"id"`
	IsPrimary                    bool                              `pulumi:"isPrimary"`
	IsSpotVM                     *bool                             `pulumi:"isSpotVM"`
	IsStateless                  *bool                             `pulumi:"isStateless"`
	MultiplePlacementGroups      *bool                             `pulumi:"multiplePlacementGroups"`
	Name                         string                            `pulumi:"name"`
	NetworkSecurityRules         []NetworkSecurityRuleResponse     `pulumi:"networkSecurityRules"`
	PlacementProperties          map[string]string                 `pulumi:"placementProperties"`
	ProvisioningState            string                            `pulumi:"provisioningState"`
	Sku                          *NodeTypeSkuResponse              `pulumi:"sku"`
	SpotRestoreTimeout           *string                           `pulumi:"spotRestoreTimeout"`
	SystemData                   SystemDataResponse                `pulumi:"systemData"`
	Tags                         map[string]string                 `pulumi:"tags"`
	Type                         string                            `pulumi:"type"`
	UseDefaultPublicLoadBalancer *bool                             `pulumi:"useDefaultPublicLoadBalancer"`
	UseEphemeralOSDisk           *bool                             `pulumi:"useEphemeralOSDisk"`
	UseTempDataDisk              *bool                             `pulumi:"useTempDataDisk"`
	VmExtensions                 []VMSSExtensionResponse           `pulumi:"vmExtensions"`
	VmImageOffer                 *string                           `pulumi:"vmImageOffer"`
	VmImagePublisher             *string                           `pulumi:"vmImagePublisher"`
	VmImageResourceId            *string                           `pulumi:"vmImageResourceId"`
	VmImageSku                   *string                           `pulumi:"vmImageSku"`
	VmImageVersion               *string                           `pulumi:"vmImageVersion"`
	VmInstanceCount              int                               `pulumi:"vmInstanceCount"`
	VmManagedIdentity            *VmManagedIdentityResponse        `pulumi:"vmManagedIdentity"`
	VmSecrets                    []VaultSecretGroupResponse        `pulumi:"vmSecrets"`
	VmSize                       *string                           `pulumi:"vmSize"`
	Zones                        []string                          `pulumi:"zones"`
}


func (val *LookupNodeTypeResult) Defaults() *LookupNodeTypeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableEncryptionAtHost) {
		enableEncryptionAtHost_ := false
		tmp.EnableEncryptionAtHost = &enableEncryptionAtHost_
	}
	if isZero(tmp.IsStateless) {
		isStateless_ := false
		tmp.IsStateless = &isStateless_
	}
	if isZero(tmp.MultiplePlacementGroups) {
		multiplePlacementGroups_ := false
		tmp.MultiplePlacementGroups = &multiplePlacementGroups_
	}
	return &tmp
}

func LookupNodeTypeOutput(ctx *pulumi.Context, args LookupNodeTypeOutputArgs, opts ...pulumi.InvokeOption) LookupNodeTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNodeTypeResult, error) {
			args := v.(LookupNodeTypeArgs)
			r, err := LookupNodeType(ctx, &args, opts...)
			var s LookupNodeTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNodeTypeResultOutput)
}

type LookupNodeTypeOutputArgs struct {
	ClusterName       pulumi.StringInput `pulumi:"clusterName"`
	NodeTypeName      pulumi.StringInput `pulumi:"nodeTypeName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNodeTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeTypeArgs)(nil)).Elem()
}


type LookupNodeTypeResultOutput struct{ *pulumi.OutputState }

func (LookupNodeTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeTypeResult)(nil)).Elem()
}

func (o LookupNodeTypeResultOutput) ToLookupNodeTypeResultOutput() LookupNodeTypeResultOutput {
	return o
}

func (o LookupNodeTypeResultOutput) ToLookupNodeTypeResultOutputWithContext(ctx context.Context) LookupNodeTypeResultOutput {
	return o
}

func (o LookupNodeTypeResultOutput) AdditionalDataDisks() VmssDataDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VmssDataDiskResponse { return v.AdditionalDataDisks }).(VmssDataDiskResponseArrayOutput)
}

func (o LookupNodeTypeResultOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *EndpointRangeDescriptionResponse { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o LookupNodeTypeResultOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o LookupNodeTypeResultOutput) DataDiskLetter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.DataDiskLetter }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) DataDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *int { return v.DataDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupNodeTypeResultOutput) DataDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.DataDiskType }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) EnableEncryptionAtHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableEncryptionAtHost }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) EnableOverProvisioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.EnableOverProvisioning }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *EndpointRangeDescriptionResponse { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o LookupNodeTypeResultOutput) EvictionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.EvictionPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) FrontendConfigurations() FrontendConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []FrontendConfigurationResponse { return v.FrontendConfigurations }).(FrontendConfigurationResponseArrayOutput)
}

func (o LookupNodeTypeResultOutput) HostGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.HostGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNodeTypeResultOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o LookupNodeTypeResultOutput) IsSpotVM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.IsSpotVM }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) IsStateless() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.IsStateless }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) MultiplePlacementGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.MultiplePlacementGroups }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNodeTypeResultOutput) NetworkSecurityRules() NetworkSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []NetworkSecurityRuleResponse { return v.NetworkSecurityRules }).(NetworkSecurityRuleResponseArrayOutput)
}

func (o LookupNodeTypeResultOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o LookupNodeTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNodeTypeResultOutput) Sku() NodeTypeSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *NodeTypeSkuResponse { return v.Sku }).(NodeTypeSkuResponsePtrOutput)
}

func (o LookupNodeTypeResultOutput) SpotRestoreTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.SpotRestoreTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNodeTypeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNodeTypeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupNodeTypeResultOutput) UseDefaultPublicLoadBalancer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.UseDefaultPublicLoadBalancer }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) UseEphemeralOSDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.UseEphemeralOSDisk }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) UseTempDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *bool { return v.UseTempDataDisk }).(pulumi.BoolPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmExtensions() VMSSExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VMSSExtensionResponse { return v.VmExtensions }).(VMSSExtensionResponseArrayOutput)
}

func (o LookupNodeTypeResultOutput) VmImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageOffer }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmImagePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImagePublisher }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmImageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageSku }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmImageVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageVersion }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmInstanceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) int { return v.VmInstanceCount }).(pulumi.IntOutput)
}

func (o LookupNodeTypeResultOutput) VmManagedIdentity() VmManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *VmManagedIdentityResponse { return v.VmManagedIdentity }).(VmManagedIdentityResponsePtrOutput)
}

func (o LookupNodeTypeResultOutput) VmSecrets() VaultSecretGroupResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VaultSecretGroupResponse { return v.VmSecrets }).(VaultSecretGroupResponseArrayOutput)
}

func (o LookupNodeTypeResultOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeTypeResultOutput{})
}
