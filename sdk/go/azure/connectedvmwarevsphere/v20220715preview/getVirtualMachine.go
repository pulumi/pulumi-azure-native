


package v20220715preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere/v20220715preview:getVirtualMachine", args, &rv, opts...)
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
	CustomResourceName string                     `pulumi:"customResourceName"`
	ExtendedLocation   *ExtendedLocationResponse  `pulumi:"extendedLocation"`
	FirmwareType       *string                    `pulumi:"firmwareType"`
	FolderPath         string                     `pulumi:"folderPath"`
	GuestAgentProfile  *GuestAgentProfileResponse `pulumi:"guestAgentProfile"`
	HardwareProfile    *HardwareProfileResponse   `pulumi:"hardwareProfile"`
	Id                 string                     `pulumi:"id"`
	Identity           *IdentityResponse          `pulumi:"identity"`
	InstanceUuid       string                     `pulumi:"instanceUuid"`
	InventoryItemId    *string                    `pulumi:"inventoryItemId"`
	Kind               *string                    `pulumi:"kind"`
	Location           string                     `pulumi:"location"`
	MoName             string                     `pulumi:"moName"`
	MoRefId            *string                    `pulumi:"moRefId"`
	Name               string                     `pulumi:"name"`
	NetworkProfile     *NetworkProfileResponse    `pulumi:"networkProfile"`
	OsProfile          *OsProfileResponse         `pulumi:"osProfile"`
	PlacementProfile   *PlacementProfileResponse  `pulumi:"placementProfile"`
	PowerState         string                     `pulumi:"powerState"`
	ProvisioningState  string                     `pulumi:"provisioningState"`
	ResourcePoolId     *string                    `pulumi:"resourcePoolId"`
	SecurityProfile    *SecurityProfileResponse   `pulumi:"securityProfile"`
	SmbiosUuid         *string                    `pulumi:"smbiosUuid"`
	Statuses           []ResourceStatusResponse   `pulumi:"statuses"`
	StorageProfile     *StorageProfileResponse    `pulumi:"storageProfile"`
	SystemData         SystemDataResponse         `pulumi:"systemData"`
	Tags               map[string]string          `pulumi:"tags"`
	TemplateId         *string                    `pulumi:"templateId"`
	Type               string                     `pulumi:"type"`
	Uuid               string                     `pulumi:"uuid"`
	VCenterId          *string                    `pulumi:"vCenterId"`
	VmId               string                     `pulumi:"vmId"`
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

func (o LookupVirtualMachineResultOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) FirmwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.FirmwareType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.FolderPath }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *GuestAgentProfileResponse { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) InstanceUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.InstanceUuid }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.MoName }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.MoRefId }).(pulumi.StringPtrOutput)
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

func (o LookupVirtualMachineResultOutput) PlacementProfile() PlacementProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *PlacementProfileResponse { return v.PlacementProfile }).(PlacementProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.PowerState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ResourcePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.ResourcePoolId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) SmbiosUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.SmbiosUuid }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []ResourceStatusResponse { return v.Statuses }).(ResourceStatusResponseArrayOutput)
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

func (o LookupVirtualMachineResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
