


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:compute/v20180601:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmName            string  `pulumi:"vmName"`
}


type LookupVirtualMachineResult struct {
	AdditionalCapabilities  *AdditionalCapabilitiesResponse    `pulumi:"additionalCapabilities"`
	AvailabilitySet         *SubResourceResponse               `pulumi:"availabilitySet"`
	DiagnosticsProfile      *DiagnosticsProfileResponse        `pulumi:"diagnosticsProfile"`
	HardwareProfile         *HardwareProfileResponse           `pulumi:"hardwareProfile"`
	Id                      string                             `pulumi:"id"`
	Identity                *VirtualMachineIdentityResponse    `pulumi:"identity"`
	InstanceView            VirtualMachineInstanceViewResponse `pulumi:"instanceView"`
	LicenseType             *string                            `pulumi:"licenseType"`
	Location                string                             `pulumi:"location"`
	Name                    string                             `pulumi:"name"`
	NetworkProfile          *NetworkProfileResponse            `pulumi:"networkProfile"`
	OsProfile               *OSProfileResponse                 `pulumi:"osProfile"`
	Plan                    *PlanResponse                      `pulumi:"plan"`
	ProvisioningState       string                             `pulumi:"provisioningState"`
	ProximityPlacementGroup *SubResourceResponse               `pulumi:"proximityPlacementGroup"`
	Resources               []VirtualMachineExtensionResponse  `pulumi:"resources"`
	StorageProfile          *StorageProfileResponse            `pulumi:"storageProfile"`
	Tags                    map[string]string                  `pulumi:"tags"`
	Type                    string                             `pulumi:"type"`
	VmId                    string                             `pulumi:"vmId"`
	Zones                   []string                           `pulumi:"zones"`
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
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmName            pulumi.StringInput    `pulumi:"vmName"`
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

func (o LookupVirtualMachineResultOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *AdditionalCapabilitiesResponse { return v.AdditionalCapabilities }).(AdditionalCapabilitiesResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *DiagnosticsProfileResponse { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *VirtualMachineIdentityResponse { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) InstanceView() VirtualMachineInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) VirtualMachineInstanceViewResponse { return v.InstanceView }).(VirtualMachineInstanceViewResponseOutput)
}

func (o LookupVirtualMachineResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
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

func (o LookupVirtualMachineResultOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *SubResourceResponse { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []VirtualMachineExtensionResponse { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VmId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
