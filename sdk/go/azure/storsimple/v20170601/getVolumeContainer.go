


package v20170601

import (
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
