


package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineResource(ctx *pulumi.Context, args *LookupVirtualMachineResourceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResourceResult, error) {
	var rv LookupVirtualMachineResourceResult
	err := ctx.Invoke("azure-native:devtestlab/v20150521preview:getVirtualMachineResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineResourceArgs struct {
	LabName           string `pulumi:"labName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVirtualMachineResourceResult struct {
	ArtifactDeploymentStatus   *ArtifactDeploymentStatusPropertiesResponse `pulumi:"artifactDeploymentStatus"`
	Artifacts                  []ArtifactInstallPropertiesResponse         `pulumi:"artifacts"`
	ComputeId                  *string                                     `pulumi:"computeId"`
	CreatedByUser              *string                                     `pulumi:"createdByUser"`
	CreatedByUserId            *string                                     `pulumi:"createdByUserId"`
	CustomImageId              *string                                     `pulumi:"customImageId"`
	DisallowPublicIpAddress    *bool                                       `pulumi:"disallowPublicIpAddress"`
	Fqdn                       *string                                     `pulumi:"fqdn"`
	GalleryImageReference      *GalleryImageReferenceResponse              `pulumi:"galleryImageReference"`
	Id                         *string                                     `pulumi:"id"`
	IsAuthenticationWithSshKey *bool                                       `pulumi:"isAuthenticationWithSshKey"`
	LabSubnetName              *string                                     `pulumi:"labSubnetName"`
	LabVirtualNetworkId        *string                                     `pulumi:"labVirtualNetworkId"`
	Location                   *string                                     `pulumi:"location"`
	Name                       *string                                     `pulumi:"name"`
	Notes                      *string                                     `pulumi:"notes"`
	OsType                     *string                                     `pulumi:"osType"`
	OwnerObjectId              *string                                     `pulumi:"ownerObjectId"`
	Password                   *string                                     `pulumi:"password"`
	ProvisioningState          *string                                     `pulumi:"provisioningState"`
	Size                       *string                                     `pulumi:"size"`
	SshKey                     *string                                     `pulumi:"sshKey"`
	Tags                       map[string]string                           `pulumi:"tags"`
	Type                       *string                                     `pulumi:"type"`
	UserName                   *string                                     `pulumi:"userName"`
}
