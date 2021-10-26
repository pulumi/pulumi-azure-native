


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeviceExtendedInformation(ctx *pulumi.Context, args *GetDeviceExtendedInformationArgs, opts ...pulumi.InvokeOption) (*GetDeviceExtendedInformationResult, error) {
	var rv GetDeviceExtendedInformationResult
	err := ctx.Invoke("azure-native:databoxedge/v20210601preview:getDeviceExtendedInformation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDeviceExtendedInformationArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetDeviceExtendedInformationResult struct {
	ChannelIntegrityKeyName    *string                   `pulumi:"channelIntegrityKeyName"`
	ChannelIntegrityKeyVersion *string                   `pulumi:"channelIntegrityKeyVersion"`
	ClientSecretStoreId        *string                   `pulumi:"clientSecretStoreId"`
	ClientSecretStoreUrl       *string                   `pulumi:"clientSecretStoreUrl"`
	DeviceSecrets              map[string]SecretResponse `pulumi:"deviceSecrets"`
	EncryptionKey              *string                   `pulumi:"encryptionKey"`
	EncryptionKeyThumbprint    *string                   `pulumi:"encryptionKeyThumbprint"`
	Id                         string                    `pulumi:"id"`
	KeyVaultSyncStatus         *string                   `pulumi:"keyVaultSyncStatus"`
	Name                       string                    `pulumi:"name"`
	ResourceKey                string                    `pulumi:"resourceKey"`
	Type                       string                    `pulumi:"type"`
}
