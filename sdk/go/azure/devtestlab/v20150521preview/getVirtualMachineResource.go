


package v20150521preview

import (
	"context"
	"reflect"

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

func LookupVirtualMachineResourceOutput(ctx *pulumi.Context, args LookupVirtualMachineResourceOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResourceResult, error) {
			args := v.(LookupVirtualMachineResourceArgs)
			r, err := LookupVirtualMachineResource(ctx, &args, opts...)
			var s LookupVirtualMachineResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineResourceResultOutput)
}

type LookupVirtualMachineResourceOutputArgs struct {
	LabName           pulumi.StringInput `pulumi:"labName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVirtualMachineResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResourceArgs)(nil)).Elem()
}


type LookupVirtualMachineResourceResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResourceResult)(nil)).Elem()
}

func (o LookupVirtualMachineResourceResultOutput) ToLookupVirtualMachineResourceResultOutput() LookupVirtualMachineResourceResultOutput {
	return o
}

func (o LookupVirtualMachineResourceResultOutput) ToLookupVirtualMachineResourceResultOutputWithContext(ctx context.Context) LookupVirtualMachineResourceResultOutput {
	return o
}

func (o LookupVirtualMachineResourceResultOutput) ArtifactDeploymentStatus() ArtifactDeploymentStatusPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *ArtifactDeploymentStatusPropertiesResponse {
		return v.ArtifactDeploymentStatus
	}).(ArtifactDeploymentStatusPropertiesResponsePtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Artifacts() ArtifactInstallPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) []ArtifactInstallPropertiesResponse { return v.Artifacts }).(ArtifactInstallPropertiesResponseArrayOutput)
}

func (o LookupVirtualMachineResourceResultOutput) ComputeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.ComputeId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) CreatedByUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.CreatedByUser }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) CreatedByUserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.CreatedByUserId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) CustomImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.CustomImageId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) DisallowPublicIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *bool { return v.DisallowPublicIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) GalleryImageReference() GalleryImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *GalleryImageReferenceResponse {
		return v.GalleryImageReference
	}).(GalleryImageReferenceResponsePtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) IsAuthenticationWithSshKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *bool { return v.IsAuthenticationWithSshKey }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) LabSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.LabSubnetName }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) LabVirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.LabVirtualNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) OwnerObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.OwnerObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) SshKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.SshKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineResourceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResourceResultOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResourceResult) *string { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResourceResultOutput{})
}
