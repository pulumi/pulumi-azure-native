


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAssetEncryptionKey(ctx *pulumi.Context, args *GetAssetEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*GetAssetEncryptionKeyResult, error) {
	var rv GetAssetEncryptionKeyResult
	err := ctx.Invoke("azure-native:media/v20210601:getAssetEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAssetEncryptionKeyArgs struct {
	AccountName       string `pulumi:"accountName"`
	AssetName         string `pulumi:"assetName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetAssetEncryptionKeyResult struct {
	AssetFileEncryptionMetadata []AssetFileEncryptionMetadataResponse `pulumi:"assetFileEncryptionMetadata"`
	Key                         *string                               `pulumi:"key"`
}

func GetAssetEncryptionKeyOutput(ctx *pulumi.Context, args GetAssetEncryptionKeyOutputArgs, opts ...pulumi.InvokeOption) GetAssetEncryptionKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAssetEncryptionKeyResult, error) {
			args := v.(GetAssetEncryptionKeyArgs)
			r, err := GetAssetEncryptionKey(ctx, &args, opts...)
			var s GetAssetEncryptionKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAssetEncryptionKeyResultOutput)
}

type GetAssetEncryptionKeyOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	AssetName         pulumi.StringInput `pulumi:"assetName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetAssetEncryptionKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAssetEncryptionKeyArgs)(nil)).Elem()
}


type GetAssetEncryptionKeyResultOutput struct{ *pulumi.OutputState }

func (GetAssetEncryptionKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAssetEncryptionKeyResult)(nil)).Elem()
}

func (o GetAssetEncryptionKeyResultOutput) ToGetAssetEncryptionKeyResultOutput() GetAssetEncryptionKeyResultOutput {
	return o
}

func (o GetAssetEncryptionKeyResultOutput) ToGetAssetEncryptionKeyResultOutputWithContext(ctx context.Context) GetAssetEncryptionKeyResultOutput {
	return o
}

func (o GetAssetEncryptionKeyResultOutput) AssetFileEncryptionMetadata() AssetFileEncryptionMetadataResponseArrayOutput {
	return o.ApplyT(func(v GetAssetEncryptionKeyResult) []AssetFileEncryptionMetadataResponse {
		return v.AssetFileEncryptionMetadata
	}).(AssetFileEncryptionMetadataResponseArrayOutput)
}

func (o GetAssetEncryptionKeyResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAssetEncryptionKeyResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAssetEncryptionKeyResultOutput{})
}
