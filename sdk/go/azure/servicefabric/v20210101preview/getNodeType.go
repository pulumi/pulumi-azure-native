


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNodeType(ctx *pulumi.Context, args *LookupNodeTypeArgs, opts ...pulumi.InvokeOption) (*LookupNodeTypeResult, error) {
	var rv LookupNodeTypeResult
	err := ctx.Invoke("azure-native:servicefabric/v20210101preview:getNodeType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNodeTypeArgs struct {
	ClusterName       string `pulumi:"clusterName"`
	NodeTypeName      string `pulumi:"nodeTypeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNodeTypeResult struct {
	ApplicationPorts    *EndpointRangeDescriptionResponse `pulumi:"applicationPorts"`
	Capacities          map[string]string                 `pulumi:"capacities"`
	DataDiskSizeGB      int                               `pulumi:"dataDiskSizeGB"`
	EphemeralPorts      *EndpointRangeDescriptionResponse `pulumi:"ephemeralPorts"`
	Id                  string                            `pulumi:"id"`
	IsPrimary           bool                              `pulumi:"isPrimary"`
	Name                string                            `pulumi:"name"`
	PlacementProperties map[string]string                 `pulumi:"placementProperties"`
	ProvisioningState   string                            `pulumi:"provisioningState"`
	SystemData          SystemDataResponse                `pulumi:"systemData"`
	Tags                map[string]string                 `pulumi:"tags"`
	Type                string                            `pulumi:"type"`
	VmExtensions        []VMSSExtensionResponse           `pulumi:"vmExtensions"`
	VmImageOffer        *string                           `pulumi:"vmImageOffer"`
	VmImagePublisher    *string                           `pulumi:"vmImagePublisher"`
	VmImageSku          *string                           `pulumi:"vmImageSku"`
	VmImageVersion      *string                           `pulumi:"vmImageVersion"`
	VmInstanceCount     int                               `pulumi:"vmInstanceCount"`
	VmManagedIdentity   *VmManagedIdentityResponse        `pulumi:"vmManagedIdentity"`
	VmSecrets           []VaultSecretGroupResponse        `pulumi:"vmSecrets"`
	VmSize              *string                           `pulumi:"vmSize"`
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

func (o LookupNodeTypeResultOutput) ApplicationPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *EndpointRangeDescriptionResponse { return v.ApplicationPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o LookupNodeTypeResultOutput) Capacities() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.Capacities }).(pulumi.StringMapOutput)
}

func (o LookupNodeTypeResultOutput) DataDiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) int { return v.DataDiskSizeGB }).(pulumi.IntOutput)
}

func (o LookupNodeTypeResultOutput) EphemeralPorts() EndpointRangeDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *EndpointRangeDescriptionResponse { return v.EphemeralPorts }).(EndpointRangeDescriptionResponsePtrOutput)
}

func (o LookupNodeTypeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNodeTypeResultOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o LookupNodeTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNodeTypeResultOutput) PlacementProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) map[string]string { return v.PlacementProperties }).(pulumi.StringMapOutput)
}

func (o LookupNodeTypeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
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

func (o LookupNodeTypeResultOutput) VmExtensions() VMSSExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) []VMSSExtensionResponse { return v.VmExtensions }).(VMSSExtensionResponseArrayOutput)
}

func (o LookupNodeTypeResultOutput) VmImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImageOffer }).(pulumi.StringPtrOutput)
}

func (o LookupNodeTypeResultOutput) VmImagePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNodeTypeResult) *string { return v.VmImagePublisher }).(pulumi.StringPtrOutput)
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

func init() {
	pulumi.RegisterOutputType(LookupNodeTypeResultOutput{})
}
