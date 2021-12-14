


package v20180915

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("azure-native:devtestlab/v20180915:getVirtualMachine", args, &rv, opts...)
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
