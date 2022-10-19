


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:devtestlab:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVirtualMachineArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupVirtualMachineResult struct {
	AllowClaim                   *bool                                      `pulumi:"allowClaim"`
	ApplicableSchedule           ApplicableScheduleResponse                 `pulumi:"applicableSchedule"`
	ArtifactDeploymentStatus     ArtifactDeploymentStatusPropertiesResponse `pulumi:"artifactDeploymentStatus"`
	Artifacts                    []ArtifactInstallPropertiesResponse        `pulumi:"artifacts"`
	ComputeId                    string                                     `pulumi:"computeId"`
	ComputeVm                    ComputeVmPropertiesResponse                `pulumi:"computeVm"`
	CreatedByUser                string                                     `pulumi:"createdByUser"`
	CreatedByUserId              string                                     `pulumi:"createdByUserId"`
	CreatedDate                  *string                                    `pulumi:"createdDate"`
	CustomImageId                *string                                    `pulumi:"customImageId"`
	DataDiskParameters           []DataDiskPropertiesResponse               `pulumi:"dataDiskParameters"`
	DisallowPublicIpAddress      *bool                                      `pulumi:"disallowPublicIpAddress"`
	EnvironmentId                *string                                    `pulumi:"environmentId"`
	ExpirationDate               *string                                    `pulumi:"expirationDate"`
	Fqdn                         string                                     `pulumi:"fqdn"`
	GalleryImageReference        *GalleryImageReferenceResponse             `pulumi:"galleryImageReference"`
	Id                           string                                     `pulumi:"id"`
	IsAuthenticationWithSshKey   *bool                                      `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName                *string                                    `pulumi:"labSubnetName"`
	LabVirtualNetworkId          *string                                    `pulumi:"labVirtualNetworkId"`
	LastKnownPowerState          string                                     `pulumi:"lastKnownPowerState"`
	Location                     *string                                    `pulumi:"location"`
	Name                         string                                     `pulumi:"name"`
	NetworkInterface             *NetworkInterfacePropertiesResponse        `pulumi:"networkInterface"`
	Notes                        *string                                    `pulumi:"notes"`
	OsType                       string                                     `pulumi:"osType"`
	OwnerObjectId                *string                                    `pulumi:"ownerObjectId"`
	OwnerUserPrincipalName       *string                                    `pulumi:"ownerUserPrincipalName"`
	Password                     *string                                    `pulumi:"password"`
	PlanId                       *string                                    `pulumi:"planId"`
	ProvisioningState            string                                     `pulumi:"provisioningState"`
	ScheduleParameters           []ScheduleCreationParameterResponse        `pulumi:"scheduleParameters"`
	Size                         *string                                    `pulumi:"size"`
	SshKey                       *string                                    `pulumi:"sshKey"`
	StorageType                  *string                                    `pulumi:"storageType"`
	Tags                         map[string]string                          `pulumi:"tags"`
	Type                         string                                     `pulumi:"type"`
	UniqueIdentifier             string                                     `pulumi:"uniqueIdentifier"`
	UserName                     *string                                    `pulumi:"userName"`
	VirtualMachineCreationSource string                                     `pulumi:"virtualMachineCreationSource"`
}


func (val *LookupVirtualMachineResult) Defaults() *LookupVirtualMachineResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowClaim) {
		allowClaim_ := false
		tmp.AllowClaim = &allowClaim_
	}
	tmp.ApplicableSchedule = *tmp.ApplicableSchedule.Defaults()

	if isZero(tmp.DisallowPublicIpAddress) {
		disallowPublicIpAddress_ := false
		tmp.DisallowPublicIpAddress = &disallowPublicIpAddress_
	}
	if isZero(tmp.OwnerObjectId) {
		ownerObjectId_ := "dynamicValue"
		tmp.OwnerObjectId = &ownerObjectId_
	}
	if isZero(tmp.StorageType) {
		storageType_ := "labStorageType"
		tmp.StorageType = &storageType_
	}
	return &tmp
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
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
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

func (o LookupVirtualMachineResultOutput) AllowClaim() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.AllowClaim }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ApplicableSchedule() ApplicableScheduleResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) ApplicableScheduleResponse { return v.ApplicableSchedule }).(ApplicableScheduleResponseOutput)
}

func (o LookupVirtualMachineResultOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) ArtifactDeploymentStatusPropertiesResponse {
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesResponseOutput)
}

func (o LookupVirtualMachineResultOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []ArtifactInstallPropertiesResponse { return v.Artifacts }).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) ComputeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ComputeId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ComputeVm() ComputeVmPropertiesResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) ComputeVmPropertiesResponse { return v.ComputeVm }).(ComputeVmPropertiesResponseOutput)
}

func (o LookupVirtualMachineResultOutput) CreatedByUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.CreatedByUser }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) CreatedByUserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.CreatedByUserId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) DataDiskParameters() DataDiskPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []DataDiskPropertiesResponse { return v.DataDiskParameters }).(DataDiskPropertiesResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) EnvironmentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.EnvironmentId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *GalleryImageReferenceResponse { return v.GalleryImageReference }).(GalleryImageReferenceResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) LastKnownPowerState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.LastKnownPowerState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) NetworkInterface() NetworkInterfacePropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *NetworkInterfacePropertiesResponse { return v.NetworkInterface }).(NetworkInterfacePropertiesResponsePtrOutput)
}

func (o LookupVirtualMachineResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.OsType }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) OwnerUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.OwnerUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ScheduleParameters() ScheduleCreationParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []ScheduleCreationParameterResponse { return v.ScheduleParameters }).(ScheduleCreationParameterResponseArrayOutput)
}

func (o LookupVirtualMachineResultOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) StorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.StorageType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VirtualMachineCreationSource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.VirtualMachineCreationSource }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
