


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetvirtualmachineRetrieve(ctx *pulumi.Context, args *GetvirtualmachineRetrieveArgs, opts ...pulumi.InvokeOption) (*GetvirtualmachineRetrieveResult, error) {
	var rv GetvirtualmachineRetrieveResult
	err := ctx.Invoke("azure-native:azurestackhci/v20210701preview:getvirtualmachineRetrieve", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetvirtualmachineRetrieveArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VirtualmachinesName string `pulumi:"virtualmachinesName"`
}


type GetvirtualmachineRetrieveResult struct {
	ExtendedLocation  *ExtendedLocationResponse                         `pulumi:"extendedLocation"`
	HardwareProfile   *VirtualmachinesPropertiesResponseHardwareProfile `pulumi:"hardwareProfile"`
	Id                string                                            `pulumi:"id"`
	Location          string                                            `pulumi:"location"`
	Name              string                                            `pulumi:"name"`
	NetworkProfile    *VirtualmachinesPropertiesResponseNetworkProfile  `pulumi:"networkProfile"`
	OsProfile         *VirtualmachinesPropertiesResponseOsProfile       `pulumi:"osProfile"`
	ProvisioningState string                                            `pulumi:"provisioningState"`
	ResourceName      *string                                           `pulumi:"resourceName"`
	SecurityProfile   *VirtualmachinesPropertiesResponseSecurityProfile `pulumi:"securityProfile"`
	Status            VirtualMachineStatusResponse                      `pulumi:"status"`
	StorageProfile    *VirtualmachinesPropertiesResponseStorageProfile  `pulumi:"storageProfile"`
	SystemData        SystemDataResponse                                `pulumi:"systemData"`
	Tags              map[string]string                                 `pulumi:"tags"`
	Type              string                                            `pulumi:"type"`
}

func GetvirtualmachineRetrieveOutput(ctx *pulumi.Context, args GetvirtualmachineRetrieveOutputArgs, opts ...pulumi.InvokeOption) GetvirtualmachineRetrieveResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetvirtualmachineRetrieveResult, error) {
			args := v.(GetvirtualmachineRetrieveArgs)
			r, err := GetvirtualmachineRetrieve(ctx, &args, opts...)
			var s GetvirtualmachineRetrieveResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetvirtualmachineRetrieveResultOutput)
}

type GetvirtualmachineRetrieveOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualmachinesName pulumi.StringInput `pulumi:"virtualmachinesName"`
}

func (GetvirtualmachineRetrieveOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualmachineRetrieveArgs)(nil)).Elem()
}


type GetvirtualmachineRetrieveResultOutput struct{ *pulumi.OutputState }

func (GetvirtualmachineRetrieveResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetvirtualmachineRetrieveResult)(nil)).Elem()
}

func (o GetvirtualmachineRetrieveResultOutput) ToGetvirtualmachineRetrieveResultOutput() GetvirtualmachineRetrieveResultOutput {
	return o
}

func (o GetvirtualmachineRetrieveResultOutput) ToGetvirtualmachineRetrieveResultOutputWithContext(ctx context.Context) GetvirtualmachineRetrieveResultOutput {
	return o
}

func (o GetvirtualmachineRetrieveResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) HardwareProfile() VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseHardwareProfile {
		return v.HardwareProfile
	}).(VirtualmachinesPropertiesResponseHardwareProfilePtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) NetworkProfile() VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseNetworkProfile {
		return v.NetworkProfile
	}).(VirtualmachinesPropertiesResponseNetworkProfilePtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) OsProfile() VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseOsProfile {
		return v.OsProfile
	}).(VirtualmachinesPropertiesResponseOsProfilePtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *string { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) SecurityProfile() VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseSecurityProfile {
		return v.SecurityProfile
	}).(VirtualmachinesPropertiesResponseSecurityProfilePtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) Status() VirtualMachineStatusResponseOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) VirtualMachineStatusResponse { return v.Status }).(VirtualMachineStatusResponseOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) StorageProfile() VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) *VirtualmachinesPropertiesResponseStorageProfile {
		return v.StorageProfile
	}).(VirtualmachinesPropertiesResponseStorageProfilePtrOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetvirtualmachineRetrieveResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetvirtualmachineRetrieveResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetvirtualmachineRetrieveResultOutput{})
}
