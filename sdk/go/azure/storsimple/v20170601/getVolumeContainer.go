


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolumeContainer(ctx *pulumi.Context, args *LookupVolumeContainerArgs, opts ...pulumi.InvokeOption) (*LookupVolumeContainerResult, error) {
	var rv LookupVolumeContainerResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getVolumeContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeContainerArgs struct {
	DeviceName          string `pulumi:"deviceName"`
	ManagerName         string `pulumi:"managerName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	VolumeContainerName string `pulumi:"volumeContainerName"`
}


type LookupVolumeContainerResult struct {
	BandWidthRateInMbps           *int                               `pulumi:"bandWidthRateInMbps"`
	BandwidthSettingId            *string                            `pulumi:"bandwidthSettingId"`
	EncryptionKey                 *AsymmetricEncryptedSecretResponse `pulumi:"encryptionKey"`
	EncryptionStatus              string                             `pulumi:"encryptionStatus"`
	Id                            string                             `pulumi:"id"`
	Kind                          *string                            `pulumi:"kind"`
	Name                          string                             `pulumi:"name"`
	OwnerShipStatus               string                             `pulumi:"ownerShipStatus"`
	StorageAccountCredentialId    string                             `pulumi:"storageAccountCredentialId"`
	TotalCloudStorageUsageInBytes float64                            `pulumi:"totalCloudStorageUsageInBytes"`
	Type                          string                             `pulumi:"type"`
	VolumeCount                   int                                `pulumi:"volumeCount"`
}

func LookupVolumeContainerOutput(ctx *pulumi.Context, args LookupVolumeContainerOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeContainerResult, error) {
			args := v.(LookupVolumeContainerArgs)
			r, err := LookupVolumeContainer(ctx, &args, opts...)
			var s LookupVolumeContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeContainerResultOutput)
}

type LookupVolumeContainerOutputArgs struct {
	DeviceName          pulumi.StringInput `pulumi:"deviceName"`
	ManagerName         pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeContainerName pulumi.StringInput `pulumi:"volumeContainerName"`
}

func (LookupVolumeContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeContainerArgs)(nil)).Elem()
}


type LookupVolumeContainerResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeContainerResult)(nil)).Elem()
}

func (o LookupVolumeContainerResultOutput) ToLookupVolumeContainerResultOutput() LookupVolumeContainerResultOutput {
	return o
}

func (o LookupVolumeContainerResultOutput) ToLookupVolumeContainerResultOutputWithContext(ctx context.Context) LookupVolumeContainerResultOutput {
	return o
}

func (o LookupVolumeContainerResultOutput) BandWidthRateInMbps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) *int { return v.BandWidthRateInMbps }).(pulumi.IntPtrOutput)
}

func (o LookupVolumeContainerResultOutput) BandwidthSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) *string { return v.BandwidthSettingId }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeContainerResultOutput) EncryptionKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) *AsymmetricEncryptedSecretResponse { return v.EncryptionKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o LookupVolumeContainerResultOutput) EncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) string { return v.EncryptionStatus }).(pulumi.StringOutput)
}

func (o LookupVolumeContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeContainerResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeContainerResultOutput) OwnerShipStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) string { return v.OwnerShipStatus }).(pulumi.StringOutput)
}

func (o LookupVolumeContainerResultOutput) StorageAccountCredentialId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) string { return v.StorageAccountCredentialId }).(pulumi.StringOutput)
}

func (o LookupVolumeContainerResultOutput) TotalCloudStorageUsageInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeContainerResult) float64 { return v.TotalCloudStorageUsageInBytes }).(pulumi.Float64Output)
}

func (o LookupVolumeContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeContainerResultOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeContainerResult) int { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeContainerResultOutput{})
}
