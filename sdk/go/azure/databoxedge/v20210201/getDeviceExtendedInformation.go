


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDeviceExtendedInformation(ctx *pulumi.Context, args *GetDeviceExtendedInformationArgs, opts ...pulumi.InvokeOption) (*GetDeviceExtendedInformationResult, error) {
	var rv GetDeviceExtendedInformationResult
	err := ctx.Invoke("azure-native:databoxedge/v20210201:getDeviceExtendedInformation", args, &rv, opts...)
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

func GetDeviceExtendedInformationOutput(ctx *pulumi.Context, args GetDeviceExtendedInformationOutputArgs, opts ...pulumi.InvokeOption) GetDeviceExtendedInformationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeviceExtendedInformationResult, error) {
			args := v.(GetDeviceExtendedInformationArgs)
			r, err := GetDeviceExtendedInformation(ctx, &args, opts...)
			var s GetDeviceExtendedInformationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeviceExtendedInformationResultOutput)
}

type GetDeviceExtendedInformationOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetDeviceExtendedInformationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceExtendedInformationArgs)(nil)).Elem()
}


type GetDeviceExtendedInformationResultOutput struct{ *pulumi.OutputState }

func (GetDeviceExtendedInformationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceExtendedInformationResult)(nil)).Elem()
}

func (o GetDeviceExtendedInformationResultOutput) ToGetDeviceExtendedInformationResultOutput() GetDeviceExtendedInformationResultOutput {
	return o
}

func (o GetDeviceExtendedInformationResultOutput) ToGetDeviceExtendedInformationResultOutputWithContext(ctx context.Context) GetDeviceExtendedInformationResultOutput {
	return o
}

func (o GetDeviceExtendedInformationResultOutput) ChannelIntegrityKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.ChannelIntegrityKeyName }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) ChannelIntegrityKeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.ChannelIntegrityKeyVersion }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) ClientSecretStoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.ClientSecretStoreId }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) ClientSecretStoreUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.ClientSecretStoreUrl }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) DeviceSecrets() SecretResponseMapOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) map[string]SecretResponse { return v.DeviceSecrets }).(SecretResponseMapOutput)
}

func (o GetDeviceExtendedInformationResultOutput) EncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.EncryptionKey }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) EncryptionKeyThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.EncryptionKeyThumbprint }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDeviceExtendedInformationResultOutput) KeyVaultSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) *string { return v.KeyVaultSyncStatus }).(pulumi.StringPtrOutput)
}

func (o GetDeviceExtendedInformationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDeviceExtendedInformationResultOutput) ResourceKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) string { return v.ResourceKey }).(pulumi.StringOutput)
}

func (o GetDeviceExtendedInformationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceExtendedInformationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeviceExtendedInformationResultOutput{})
}
