


package scvmm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:scvmm:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualMachineName string `pulumi:"virtualMachineName"`
}


type LookupVirtualMachineResult struct {
	AvailabilitySets  []VirtualMachinePropertiesResponseAvailabilitySets `pulumi:"availabilitySets"`
	CheckpointType    *string                                            `pulumi:"checkpointType"`
	Checkpoints       []CheckpointResponse                               `pulumi:"checkpoints"`
	CloudId           *string                                            `pulumi:"cloudId"`
	ExtendedLocation  ExtendedLocationResponse                           `pulumi:"extendedLocation"`
	Generation        *int                                               `pulumi:"generation"`
	HardwareProfile   *HardwareProfileResponse                           `pulumi:"hardwareProfile"`
	Id                string                                             `pulumi:"id"`
	InventoryItemId   *string                                            `pulumi:"inventoryItemId"`
	Location          string                                             `pulumi:"location"`
	Name              string                                             `pulumi:"name"`
	NetworkProfile    *NetworkProfileResponse                            `pulumi:"networkProfile"`
	OsProfile         *OsProfileResponse                                 `pulumi:"osProfile"`
	PowerState        string                                             `pulumi:"powerState"`
	ProvisioningState string                                             `pulumi:"provisioningState"`
	StorageProfile    *StorageProfileResponse                            `pulumi:"storageProfile"`
	SystemData        SystemDataResponse                                 `pulumi:"systemData"`
	Tags              map[string]string                                  `pulumi:"tags"`
	TemplateId        *string                                            `pulumi:"templateId"`
	Type              string                                             `pulumi:"type"`
	Uuid              *string                                            `pulumi:"uuid"`
	VmName            *string                                            `pulumi:"vmName"`
	VmmServerId       *string                                            `pulumi:"vmmServerId"`
}

func LookupVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineArgs)
			r, err := LookupVirtualMachine(ctx, &args, opts...)
			var s LookupVirtualMachineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineResultOutput)
}

type LookupVirtualMachineOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualMachineName pulumi.StringInput `pulumi:"virtualMachineName"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}


type LookupVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutput() LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) AvailabilitySets() VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualMachinePropertiesResponseAvailabilitySets {
		return v.AvailabilitySets
	}).(VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput)
}

func (o LookupVirtualMachineResultOutput) CheckpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CheckpointType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Checkpoints() CheckpointResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []CheckpointResponse { return v.Checkpoints }).(CheckpointResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CloudId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o LookupVirtualMachineResultOutput) Generation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.Generation }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *OsProfileResponse { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.PowerState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineResultOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.TemplateId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VmName }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
